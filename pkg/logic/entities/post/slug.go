package entities_post

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Slug struct {
	s string
}

const SLUG_MAX_LEN = 255

func CreateSlug(s string) (Slug, error) {
	if utf8.RuneCountInString(s) > SLUG_MAX_LEN {
		return Slug{}, errors.New(fmt.Sprintf("Slug max length is %d. Got one with length %d", SLUG_MAX_LEN, utf8.RuneCountInString(s)))
	}
	return Slug{s}, nil
}

func (s Slug) String() string {
	return s.s
}
