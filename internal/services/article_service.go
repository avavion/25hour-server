package services

import (
	"github.com/avavion/25hour-server/internal/models"
	"github.com/avavion/25hour-server/internal/repositories"
)

type ArticleService interface {
	GetAllArticles() ([]models.Article, error)
	GetArticleBySlug(slug string) (models.Article, error)
}

type ArticleServiceImplementation struct {
	articleRepository repositories.ArticleRepository
}

func (s *ArticleServiceImplementation) GetAllArticles() ([]models.Article, error) {
	articles, err := s.articleRepository.GetAllArticles()

	if err != nil {
		return nil, err
	}

	return articles, nil
}

func (s *ArticleServiceImplementation) GetArticleBySlug(slug string) (models.Article, error) {
	//TODO implement me
	panic("implement me")
}
