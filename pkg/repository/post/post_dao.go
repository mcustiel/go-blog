package repository_post

type Post struct {
	Id    string
	Slug  string
	Title string
	Text  string
}

type PostDao interface {
	Create(post Post) (Post, error)
	Get(id string) (Post, error)
	GetMany(filters map[string]any, order string, offset uint64, limit uint64) (*PostIterator, error)
	Update(post Post) error
	Delete(id string) error
}

func CreatePost(id string, slug string, title string, text string) Post {
	return Post{
		id, slug, title, text}
}
