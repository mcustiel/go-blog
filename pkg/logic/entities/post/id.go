package entities_post

import (
	"errors"
	"fmt"
	"unicode/utf8"
)

type PostId struct {
	i string
}

const ID_MAX_LEN = 32

func NilPostId() PostId {
	return PostId{""}
}

func CreateId(s string) (PostId, error) {
	if utf8.RuneCountInString(s) > ID_MAX_LEN {
		return PostId{}, errors.New(fmt.Sprintf("ID max length is %d. Got one with length %d", ID_MAX_LEN, utf8.RuneCountInString(s)))
	}
	return PostId{s}, nil
}

func (p PostId) String() string {
	return p.i
}
