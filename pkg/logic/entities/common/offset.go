package entities_common

import (
	"strconv"
)

type Offset struct {
	offset uint64
}

func OffsetFromUint64(o uint64) Offset {
	return Offset{o}
}

func OffsetFromString(o string) (Offset, error) {
	off, err := strconv.ParseUint(o, 10, 64)
	if err != nil {
		return Offset{0}, err
	}
	return Offset{off}, err
}

func (o Offset) AsUint() uint64 {
	return o.offset
}
