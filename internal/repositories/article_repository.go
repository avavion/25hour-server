package repositories

import (
	"github.com/avavion/25hour-server/internal/models"
	"github.com/avavion/25hour-server/package/drivers/databases"
)

type ArticleRepositoryInterface interface {
	GetAllArticles() ([]models.Article, error)
	GetArticleBySlug(slug string) (models.Article, error)
}

type ArticleRepository struct {
	database databases.DatabaseRepositoryInterface
}

func NewArticleRepository(database databases.DatabaseRepositoryInterface) ArticleRepositoryInterface {
	return ArticleRepository{
		database: database,
	}
}

func (a ArticleRepository) GetAllArticles() ([]models.Article, error) {
	return nil, nil

	//return []models.Article{}, nil
}

func (a ArticleRepository) GetArticleBySlug(slug string) (models.Article, error) {
	//TODO implement me
	panic("implement me")
}
