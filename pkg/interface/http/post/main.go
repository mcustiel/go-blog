package http_post

import (
	"errors"
	"log"
	"net/http"
	"time"

	ep "github.com/mcustiel/go-blog/pkg/logic/entities/post"
	"github.com/mcustiel/go-blog/pkg/logic/facade"
	"github.com/mcustiel/go-blog/pkg/utils"
	"github.com/mcustiel/gorro"
)

type PostHandler struct {
	f facade.PostFacade
}

func NewPostHandler(facade facade.PostFacade) *PostHandler {
	return &PostHandler{facade}
}

func (ph *PostHandler) UpdatePost(w http.ResponseWriter, r *gorro.Request) error {
	start := time.Now()
	id, ok := r.NamedParams["id"]
	if !ok {
		return errors.New("Missing post id parameter")
	}
	ent, err := getPostEntity(r)
	log.Printf("Validating and converting to domain entity took %s", time.Since(start))
	if err != nil {
		return err
	}
	postId, err := ep.CreateId(id)
	if err != nil {
		return err
	}

	start = time.Now()
	ent, err = ph.f.UpdatePost(ent.WithId(postId))
	log.Printf("Writing to db took %s", time.Since(start))
	if err != nil {
		return err
	}
	return writeOnePost(w, ent)
}

func (ph *PostHandler) RegisterRoutes(r gorro.Router) error {
	return utils.Exec(func() error {
		return r.Register(`^/post$`, gorro.HandlersMap{
			http.MethodPost: ph.CreatePost,
			http.MethodGet:  ph.ListPosts})
	}).Bind(func() error {
		return r.Register(`^/post/(?P<id>\d+)$`, gorro.HandlersMap{
			http.MethodPut: ph.UpdatePost,
			http.MethodGet: ph.GetPost})
	}).Bind(func() error {
		return r.Register(`^/posts$`, gorro.HandlersMap{
			http.MethodGet: ph.ListPosts})
	}).Error()
}
