package repository_post

import (
	gosql "database/sql"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"

	"github.com/mcustiel/go-blog/pkg/persistence"
	"github.com/mcustiel/go-blog/pkg/persistence/sql"
)

type SqlPostDao struct {
	conn *sql.SqlConnection
}

func NewSqlPostDao(conn *sql.SqlConnection) *SqlPostDao {
	return &SqlPostDao{conn}
}

func (dao *SqlPostDao) Create(post Post) (Post, error) {
	var err error
	var result gosql.Result
	var id int64
	result, err = dao.conn.Exec("INSERT INTO posts (slug, title, text) VALUES ($1, $2, $3)",
		[]any{post.Slug, post.Title, post.Text})

	if err == nil {
		id, err = result.LastInsertId()
		if err == nil {
			post.Id = fmt.Sprint(id)
		}
	}
	return post, err
}

func (dao *SqlPostDao) Update(post Post) error {
	var err error
	var id uint64

	id, err = strconv.ParseUint(post.Id, 0, 64)
	if err != nil {
		return err
	}

	_, err = dao.conn.Exec("UPDATE posts SET slug = $1, title = $2, text = $3 WHERE id = $4",
		[]any{post.Slug, post.Title, post.Text, id})

	return err
}

func (dao *SqlPostDao) Delete(id string) error {
	var intid uint64
	var err error

	intid, err = strconv.ParseUint(id, 0, 64)
	if err != nil {
		return err
	}
	_, err = dao.conn.Exec("DELETE FROM posts WHERE id = $1",
		[]any{intid})

	return err
}

func (dao *SqlPostDao) Get(id string) (Post, error) {
	var intid uint64
	var err error

	intid, err = strconv.ParseUint(id, 0, 64)
	if err != nil {
		return Post{}, err
	}

	res, err := dao.conn.QueryOne("SELECT id, slug , title, text FROM posts WHERE id = $1", []any{intid}, postEntityToPostObject)

	if err != nil {
		return Post{}, err
	}
	return res.(Post), nil
}

func buildWhere(filters map[string]any) (string, []any) {
	var filtersString string = "WHERE true"
	var filtersValues []any = make([]any, len(filters))
	var counter int = 0
	for field, value := range filters {
		filtersString += fmt.Sprintf(" AND %s = $%d", field, counter)
		filtersValues[counter] = value
		counter++
	}
	return filtersString, filtersValues
}

func orderString(order string) string {
	if order != "" {
		return fmt.Sprintf("ORDER BY %s", strings.ToLower(order))
	}
	return ""
}

func (dao *SqlPostDao) GetMany(filters map[string]any, order string, offset uint64, limit uint64) (*PostIterator, error) {
	var err error

	filtersString, filtersValues := buildWhere(filters)

	var limitStr string
	var offsetStr string
	if limit > 0 {
		limitStr = fmt.Sprintf(" LIMIT %d", limit)

		if offset > 0 {
			offsetStr = fmt.Sprintf(" OFFSET %d", offset)
		} else {
			offsetStr = ""
		}
	} else {
		limitStr = ""
		offsetStr = ""
	}

	log.Printf("Query %s, values %v",
		fmt.Sprintf(
			"SELECT id, slug, title, text FROM posts %s %s%s%s",
			filtersString,
			orderString(order),
			limitStr, offsetStr),
		filtersValues)

	res, err := dao.conn.Query(
		fmt.Sprintf(
			"SELECT id, slug, title, text FROM posts %s %s%s%s",
			filtersString,
			orderString(order),
			limitStr, offsetStr),
		filtersValues,
		postEntityToPostObject)

	if err != nil {
		return nil, err
	}
	return NewPostIterator(res), nil
}

func postEntityToPostObject(s persistence.Scanneable) (interface{}, error) {
	var p Post
	var id uint64
	err := s.Scan(&id, &p.Slug, &p.Title, &p.Text)
	if err != nil {
		if err == gosql.ErrNoRows {
			return nil, notFound()
		}
		return nil, err
	}
	p.Id = fmt.Sprint(id)
	return p, nil
}

func notFound() error {
	return errors.New("Entity Not Found")
}
