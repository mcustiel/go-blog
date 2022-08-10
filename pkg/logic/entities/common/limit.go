package entities_common

import (
	"strconv"
)

type Limit struct {
	limit uint64
}

func LimitFromUint64(l uint64) Limit {
	return Limit{l}
}

func LimitFromString(l string) (Limit, error) {
	lim, err := strconv.ParseUint(l, 10, 64)
	if err != nil {
		return Limit{0}, err
	}
	return Limit{lim}, err
}

func (l Limit) AsUint() uint64 {
	return l.limit
}