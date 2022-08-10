package http_post

import (
	"errors"
	"log"
	"net/http"
	"time"

	ep "github.com/mcustiel/go-blog/pkg/logic/entities/post"
	"github.com/mcustiel/gorro"
)

func (ph *PostHandler) GetPost(w http.ResponseWriter, r *gorro.Request) error {
	start := time.Now()
	id, ok := r.NamedParams["id"]
	log.Printf("Id: %s", id)
	if !ok {
		return errors.New("Missing post id parameter")
	}

	postId, err := ep.CreateId(id)
	if err != nil {
		return err
	}

	start = time.Now()
	ent, err := ph.f.GetPost(postId)
	log.Printf("Retrieving from db took %s", time.Since(start))
	if err != nil {
		if err.Error() == "Entity Not Found" {
			http.Error(w, "Post not found", http.StatusNotFound)
			return nil
		}
		return err
	}
	return writeOnePost(w, ent)
	
}
