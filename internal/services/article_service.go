package services

import (
	"github.com/avavion/25hour-server/internal/models"
	"github.com/avavion/25hour-server/internal/repositories"
)

type ArticleService interface {
	GetAllArticles() ([]models.Article, error)
	GetArticleBySlug(slug string) (models.Article, error)
}

type articleService struct {
	articleRepository repositories.ArticleRepository
}

func (a articleService) GetAllArticles() ([]models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleService) GetArticleBySlug(slug string) (models.Article, error) {
	//TODO implement me
	panic("implement me")
}
