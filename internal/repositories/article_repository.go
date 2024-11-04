package repositories

import "github.com/avavion/25hour-server/internal/models"

type ArticleRepository interface {
	GetAllArticles() ([]models.Article, error)
	GetArticleBySlug(slug string) (models.Article, error)
}

type articleRepository struct{}

func (a articleRepository) GetAllArticles() ([]models.Article, error) {
	//TODO implement me
	panic("implement me")
}

func (a articleRepository) GetArticleBySlug(slug string) (models.Article, error) {
	//TODO implement me
	panic("implement me")
}
