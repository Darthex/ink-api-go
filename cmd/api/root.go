package _api

import (
	"database/sql"
	"github.com/Darthex/ink-golang/service/articles"
	"github.com/Darthex/ink-golang/service/auth"
	"github.com/Darthex/ink-golang/service/tags"
	"net/http"
)

type Root struct {
	router *http.ServeMux
	store  *sql.DB
}

func newRootRouter(router *http.ServeMux, store *sql.DB) *Root {
	return &Root{router: router, store: store}
}

func (r *Root) registerRoutes() {
	// Auth module
	authStore := auth.NewAuthStore(r.store)
	auth.Router(r.router, authStore)

	// Article module
	articleStore := articles.NewArticleStore(r.store)
	articles.Router(r.router, articleStore)

	// Tags module
	tags.Router(r.router)
}
