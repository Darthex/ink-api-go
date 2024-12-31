package tags

import (
	"github.com/Darthex/ink-golang/types/articles"
	"github.com/Darthex/ink-golang/utils"
	"net/http"
)

func Router(router *http.ServeMux) {
	router.HandleFunc("GET /tags/",
		func(w http.ResponseWriter, r *http.Request) {
			t := []articles.Tag{
				articles.LIFESTYLE,
				articles.PROGRAMMING,
				articles.TECHNOLOGY,
				articles.BUSINESS,
				articles.ENTERTAINMENT,
				articles.EDUCATION,
				articles.ENVIRONMENT,
				articles.DESIGN,
				articles.PERSONAL,
				articles.FINANCE,
				articles.NEWS_POLITICS,
				articles.SPORTS,
			}
			_ = utils.WriteJson(w, http.StatusOK, map[string]interface{}{"tags": t})
			return
		})
}
