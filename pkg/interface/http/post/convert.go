package http_post

import (
	"errors"
	"fmt"

	ep "github.com/mcustiel/go-blog/pkg/logic/entities/post"
	"github.com/mcustiel/go-blog/pkg/utils"
)

func postFromMap(m map[string]string) (ep.Post, error) {
	var slug ep.Slug
	var title ep.Title
	var text ep.Text
	var err error
	utils.Exec(func() error {
		slug, err = ep.CreateSlug(m["slug"])
		return err
	}).Bind(func() error {
		title, err = ep.CreateTitle(m["title"])
		return err
	}).Bind(func() error {
		text, err = ep.CreateText(m["text"])
		return err
	})

	if err != nil {
		return ep.NilPost(), err
	}

	return ep.CreatePost(ep.NilPostId(), slug, title, text), nil
}

func mapOrderString(order string) (string, error) {
	switch order {
	case "i":
		return "Id", nil
	case "s":
		return "Slug", nil
	case "tl":
		return "Title", nil
	case "tx":
		return "Text", nil
	case "":
		return "", nil
	}
	return "", errors.New(fmt.Sprintf("Invalid order string: %s", order))
}

