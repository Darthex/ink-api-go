package articles

import (
	"fmt"
	"github.com/Darthex/ink-golang/types"
	"github.com/Darthex/ink-golang/types/articles"
	"github.com/Darthex/ink-golang/utils"
	"net/http"
	"strconv"
)

func Router(router *http.ServeMux, store *Store) {

	router.HandleFunc("GET /article/",
		func(w http.ResponseWriter, r *http.Request) {
			pagination, err := types.GetPaginationParams(r)
			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid pagination params: %v", err))
				return
			}
			res, err := store.GetArticles(pagination)
			if err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to get articles: %v", err))
				return
			}
			_ = utils.WriteJson(w, http.StatusOK, res)
			return
		})

	router.HandleFunc("GET /article/{articleId}",
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("articleId")
			pid, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid article id: %v", err))
				return
			}
			a, err := store.GetArticleById(pid)
			if err != nil {
				utils.WriteError(w, http.StatusNotFound, fmt.Errorf("%v", err))
				return
			}
			_ = utils.WriteJson(w, http.StatusOK, a)
		})

	router.HandleFunc("POST /article/publish",
		func(w http.ResponseWriter, r *http.Request) {
			var payload articles.ArticlePublishPayload
			if err := utils.ParseAndValidate(w, r, &payload); err != nil {
				return
			}
			if err := store.CreateNewArticle(payload); err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to publish article: %v", err))
				return
			}
			_ = utils.WriteJson(w, http.StatusOK, nil)
			return
		})

	router.HandleFunc("PUT /article/{articleId}",
		func(w http.ResponseWriter, r *http.Request) {
			id := r.PathValue("articleId")
			pid, err := strconv.ParseInt(id, 10, 64)
			if err != nil {
				utils.WriteError(w, http.StatusBadRequest, fmt.Errorf("invalid article id: %v", err))
				return
			}
			var payload articles.ArticlePublishPayload
			if err := utils.ParseAndValidate(w, r, &payload); err != nil {
				return
			}
			if err := store.UpdateArticle(payload, pid); err != nil {
				utils.WriteError(w, http.StatusInternalServerError, fmt.Errorf("failed to update article: %v", err))
				return
			}
			_ = utils.WriteJson(w, http.StatusOK, nil)
			return
		})
}
