package repository_post

import "github.com/mcustiel/go-blog/pkg/persistence"

type PostIterator struct {
	it persistence.RowsIterator
}

func NewPostIterator(it persistence.RowsIterator) *PostIterator {
	return &PostIterator{it}
}

func (it *PostIterator) Next() bool {
	return it.it.Next()
}

func (it *PostIterator) Get() (Post, error) {
	data, err := it.it.Get()
	if err != nil {
		return Post{}, err
	}
	return data.(Post), nil
}

func (it *PostIterator) Error() error {
	return it.it.Error()
}

func (it *PostIterator) Close() error {
	return it.it.Close()
}
