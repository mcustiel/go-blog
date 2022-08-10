package entities_post

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type Text struct {
	t string
}

const TEXT_MAX_LEN = 65535

func CreateText(s string) (Text, error) {
	if utf8.RuneCountInString(s) > ID_MAX_LEN {
		return Text{}, errors.New(fmt.Sprintf("Text max length is %d. Got one with length %d", TEXT_MAX_LEN, utf8.RuneCountInString(s)))
	}
	return Text{s}, nil
}

func (t Text) String() string {
	return t.t
}
