package handlers

import (
	"github.com/avavion/25hour-server/internal/services"
	"github.com/avavion/25hour-server/package/response"
	"net/http"
)

type ArticleHandlerInterface interface {
	GetAllArticles(w http.ResponseWriter, r *http.Request)
	GetArticleByID(w http.ResponseWriter, r *http.Request)
}

type ArticleHandler struct {
	service services.ArticleServiceInterface
}

func NewArticleHandler(service services.ArticleServiceInterface) ArticleHandlerInterface {
	return &ArticleHandler{service: service}
}

func (handler *ArticleHandler) GetAllArticles(w http.ResponseWriter, r *http.Request) {
	articles, err := handler.service.GetAllArticles()

	if err == nil {
		response.ToJson(w, map[string]string{
			"error": "Not Found Articles",
		}, 404)

		return
	}

	response.ToJson(w, articles, 200)
}

func (handler *ArticleHandler) GetArticleByID(w http.ResponseWriter, r *http.Request) {
	//vars := mux.Vars(r)
	//id, _ := strconv.Atoi(vars["id"])
	//
	//article, err := handler.service.GetArticleByID(id)
	//if err != nil {
	//	response.JSON(w, http.StatusNotFound, map[string]string{"error": "Article not found"})
	//	return
	//}
	//response.JSON(w, http.StatusOK, article)
}
