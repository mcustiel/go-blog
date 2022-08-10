package http_post

import (
	"log"
	"net/http"
	"time"

	"github.com/mcustiel/gorro"
)

func (ph *PostHandler) CreatePost(w http.ResponseWriter, r *gorro.Request) error {
	start := time.Now()
	ent, err := getPostEntity(r)
	log.Printf("Validating and converting to domain entity took %s", time.Since(start))
	if err != nil {
		return err
	}
	start = time.Now()
	ent, err = ph.f.CreatePost(ent)
	log.Printf("Writing to db took %s", time.Since(start))
	if err != nil {
		return err
	}
	return writeOnePost(w, ent)
}
