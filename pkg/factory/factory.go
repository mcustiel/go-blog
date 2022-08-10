package factory

import (
	gosql "database/sql"

	http_post "github.com/mcustiel/go-blog/pkg/interface/http/post"
	"github.com/mcustiel/go-blog/pkg/logic/facade"
	"github.com/mcustiel/go-blog/pkg/persistence"
	"github.com/mcustiel/go-blog/pkg/persistence/sql"
	repository_post "github.com/mcustiel/go-blog/pkg/repository/post"
	"github.com/mcustiel/gorro"

	_ "github.com/mattn/go-sqlite3"
)

var connCache persistence.ConnectionManager[gosql.Result] = nil

func CreateDbConnectionManager() persistence.ConnectionManager[gosql.Result] {
	if connCache == nil {
		connCache = sql.NewDefaultConnectionManager()
	}
	return connCache
}

func CreatePostDao() repository_post.PostDao {
	conn := CreateDbConnectionManager().GetConnection()
	return repository_post.NewSqlPostDao(conn.(*sql.SqlConnection))
}

func CreatePostFacade() facade.PostFacade {
	dao := CreatePostDao()
	return facade.CreatePostFacade(dao)
}

func CreatePostHandler() *http_post.PostHandler {
	f := CreatePostFacade()
	return http_post.NewPostHandler(f)
}

func CreateRouter() gorro.Router {
	return gorro.NewRouter()
}
