package facade

import (
	"errors"
	"fmt"
	"log"

	ec "github.com/mcustiel/go-blog/pkg/logic/entities/common"
	ep "github.com/mcustiel/go-blog/pkg/logic/entities/post"
	rp "github.com/mcustiel/go-blog/pkg/repository/post"
	"github.com/mcustiel/go-blog/pkg/utils"
)

type PostFacade struct {
	pd rp.PostDao
}

func CreatePostFacade(pd rp.PostDao) PostFacade {
	return PostFacade{pd}
}

func (f PostFacade) CreatePost(p ep.Post) (ep.Post, error) {
	if p.Id() != ep.NilPostId() {
		return p, errors.New(
			fmt.Sprintf("Trying to create a post with an already existing id: %v", p))
	}
	log.Printf("Creating post through dao")
	newPost, err := f.pd.Create(convertToDbEntity(p))
	if err != nil {
		return ep.NilPost(), err
	}
	return convertToDomainEntity(newPost)
}

func (f PostFacade) UpdatePost(p ep.Post) (ep.Post, error) {
	log.Printf("Updating post through dao")
	err := f.pd.Update(convertToDbEntity(p))
	if err != nil {
		return p, err
	}
	return p, nil
}

func (f PostFacade) GetPost(p ep.PostId) (ep.Post, error) {
	log.Printf("Fetching post through dao")
	post, err := f.pd.Get(p.String())
	log.Printf("Fetching post %v", post)

	if err != nil {
		return ep.NilPost(), err
	}
	return convertToDomainEntity(post)
}

func (f PostFacade) ListPosts(order string, offset ec.Offset, limit ec.Limit) (*PostIterator, error) {
	var err error
	ret, err := f.pd.GetMany(map[string]any{}, order, offset.AsUint(), limit.AsUint())
	if err != nil {
		return nil, err
	}
	return NewPostIterator(ret, convertToDomainEntity), nil
}

func convertToDbEntity(in ep.Post) rp.Post {
	return rp.CreatePost(
		in.Id().String(),
		in.Slug().String(),
		in.Title().String(),
		in.Text().String())
}

func convertToDomainEntity(in rp.Post) (ep.Post, error) {
	var err error
	var id ep.PostId
	var slug ep.Slug
	var title ep.Title
	var text ep.Text

	res := utils.Exec(func() error {
		id, err = ep.CreateId(in.Id)
		return err
	}).Bind(func() error {
		slug, err = ep.CreateSlug(in.Slug)
		return err
	}).Bind(func() error {
		title, err = ep.CreateTitle(in.Title)
		return err
	}).Bind(func() error {
		text, err = ep.CreateText(in.Text)
		return err
	})

	err = res.Error()
	if err != nil {
		return ep.NilPost(), err
	}

	return ep.CreatePost(
		id,
		slug,
		title,
		text), nil
}
