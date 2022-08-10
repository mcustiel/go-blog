package http_post

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"time"

	ep "github.com/mcustiel/go-blog/pkg/logic/entities/post"
	"github.com/mcustiel/gorro"
)

func getPostEntity(r *gorro.Request) (ep.Post, error) {
	var m map[string]string
	var err error

	jd := json.NewDecoder(r.Body)
	jd.DisallowUnknownFields()

	err = jd.Decode(&m)
	if err != nil {
		return ep.NilPost(), err
	}

	err = ensureValidPost(m)
	if err != nil {
		return ep.NilPost(), err
	}

	return postFromMap(m)
}

func ensureValidPost(m map[string]string) error {
	var ok bool

	_, ok = m["slug"]
	if !ok {
		return errors.New("Slug field is not present")
	}
	_, ok = m["title"]
	if !ok {
		return errors.New("Title field is not present")
	}
	_, ok = m["text"]
	if !ok {
		return errors.New("Text field is not present")
	}
	return nil
}

func writeOnePost(w http.ResponseWriter, p ep.Post) error {
	start := time.Now()
	w.Header().Add("Content-Type", "application/json")
	d := json.NewEncoder(w)
	err := encodePost(p, d)

	log.Printf("Writing the response took %s", time.Since(start))
	return err
}

func encodePost(p ep.Post, d *json.Encoder) error {
	return d.Encode(map[string]string{
		"id":    p.Id().String(),
		"slug":  p.Slug().String(),
		"title": p.Title().String(),
		"text":  p.Text().String()})
}
