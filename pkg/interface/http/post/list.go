package http_post

import (
	"encoding/json"
	"net/http"
	"net/url"

	ep "github.com/mcustiel/go-blog/pkg/logic/entities/post"
	ec "github.com/mcustiel/go-blog/pkg/logic/entities/common"
	"github.com/mcustiel/go-blog/pkg/logic/facade"
	"github.com/mcustiel/gorro"
)

func getLimit(queryString url.Values) (ec.Limit, error) {
	if param, ok := queryString["count"]; ok && len(param) > 0 {
		return ec.LimitFromString(param[0])
	}
	return ec.LimitFromUint64(0), nil
}

func getOffset(queryString url.Values) (ec.Offset, error) {
	if param, ok := queryString["from"]; ok && len(param) > 0 {
		return ec.OffsetFromString(param[0])
	}
	return ec.OffsetFromUint64(0), nil
}

func getOrder(queryString url.Values) (string, error) {
	var order string
	if orderParam, ok := queryString["order"]; ok && len(orderParam) > 0 {
		order = orderParam[0]
	} else {
		order = ""
	}

	return mapOrderString(order)
}

func (ph *PostHandler) ListPosts(w http.ResponseWriter, r *gorro.Request) error {
	var queryString url.Values = r.URL.Query()
	var limit ec.Limit
	var offset ec.Offset
	var order string
	var err error

	limit, err = getLimit(queryString)
	if err != nil {
		return err
	}
	offset, err = getOffset(queryString)
	if err != nil {
		return err
	}
	order, err = getOrder(queryString)
	if err != nil {
		return err
	}

	iterator, err := ph.f.ListPosts(order, offset, limit)
	if err != nil {
		return err
	}

	return writeList(w, iterator)
}

func writeList(w http.ResponseWriter, iterator *facade.PostIterator) error {
	var first bool = true
	var post ep.Post
	var err error

	w.Header().Add("Content-Type", "application/json")
	d := json.NewEncoder(w)
	w.Write([]byte("["))

	for iterator.Next() {
		post, err = iterator.Get()
		if err != nil {
			iterator.Close()
			return err
		}
		if !first {
			w.Write([]byte(","))
		} else {
			first = false
		}
		err = encodePost(post, d)
		if err != nil {
			iterator.Close()
			return err
		}
	}
	w.Write([]byte("]"))
	return nil
}
