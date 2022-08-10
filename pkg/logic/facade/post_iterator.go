package facade

import (
	ep "github.com/mcustiel/go-blog/pkg/logic/entities/post"
	rp "github.com/mcustiel/go-blog/pkg/repository/post"
)

type PostIterator struct {
	it        *rp.PostIterator
	converter func(rp.Post) (ep.Post, error)
}

func NewPostIterator(it *rp.PostIterator, converter func(rp.Post) (ep.Post, error)) *PostIterator {
	return &PostIterator{it, converter}
}

func (it *PostIterator) Next() bool {
	return it.it.Next()
}

func (it *PostIterator) Get() (ep.Post, error) {
	data, err := it.it.Get()
	if err != nil {
		return ep.Post{}, err
	}
	return it.converter(data)
}

func (it *PostIterator) Error() error {
	return it.it.Error()
}

func (it *PostIterator) Close() error {
	return it.it.Close()
}
