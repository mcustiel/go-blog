package entities_post

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Title struct {
	t string
}

const TITLE_MAX_LEN = 255

func CreateTitle(s string) (Title, error) {
	if utf8.RuneCountInString(s) > ID_MAX_LEN {
		return Title{}, errors.New(fmt.Sprintf("Title max length is %d. Got one with length %d", ID_MAX_LEN, utf8.RuneCountInString(s)))
	}
	return Title{s}, nil
}

func (t Title) String() string {
	return t.t
}
