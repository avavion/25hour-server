package handlers

import (
	"github.com/avavion/25hour-server/internal/services"
	"net/http"
	"strconv"
)

type ArticleHandler struct {
	service services.ArticleService
}

func (h *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := h.service.GetAllArticles()
	if err != nil {
		response.JSON(w, http.StatusInternalServerError, map[string]string{"error": "Unable to fetch articles"})
		return
	}
	response.JSON(w, http.StatusOK, articles)
}

func (h *ArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	article, err := h.service.GetArticleByID(id)
	if err != nil {
		response.JSON(w, http.StatusNotFound, map[string]string{"error": "Article not found"})
		return
	}
	response.JSON(w, http.StatusOK, article)
}
