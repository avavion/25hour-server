package services

import (
	"github.com/avavion/25hour-server/internal/models"
	"github.com/avavion/25hour-server/internal/repositories"
)

type ArticleServiceInterface interface {
	GetAllArticles() ([]models.Article, error)
	GetArticleBySlug(slug string) (models.Article, error)
}

type ArticleService struct {
	articleRepository repositories.ArticleRepositoryInterface
}

func NewArticleService(repository repositories.ArticleRepositoryInterface) ArticleServiceInterface {
	return &ArticleService{
		articleRepository: repository,
	}
}

func (s *ArticleService) GetAllArticles() ([]models.Article, error) {
	articles, err := s.articleRepository.GetAllArticles()

	if err == nil {
		return nil, err
	}

	return articles, nil
}

func (s *ArticleService) GetArticleBySlug(slug string) (models.Article, error) {
	//TODO implement me
	panic("implement me")
}
