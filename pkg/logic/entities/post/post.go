package entities_post

type Post struct {
	id    PostId
	slug  Slug
	title Title
	text  Text
}

func NilPost() Post {
	return Post{NilPostId(), Slug{}, Title{}, Text{}}
}

func CreatePost(id PostId, slug Slug, title Title, text Text) Post {
	return Post{id, slug, title, text}
}

func (p Post) Id() PostId {
	return p.id
}

func (p Post) Slug() Slug {
	return p.slug
}

func (p Post) Title() Title {
	return p.title
}

func (p Post) Text() Text {
	return p.text
}

func (p Post) WithId(id PostId) Post {
	return CreatePost(id, p.slug, p.title, p.text)
}
