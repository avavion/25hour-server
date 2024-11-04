package repositories

import "github.com/avavion/25hour-server/internal/models"

type ArticleRepository interface {
	GetAllArticles() ([]models.Article, error)
	GetArticleBySlug(slug string) (models.Article, error)
}

type ArticleRepositoryImplementation struct {
	database
}

func (a ArticleRepositoryImplementation) GetAllArticles() ([]models.Article, error) {

}

func (a ArticleRepositoryImplementation) GetArticleBySlug(slug string) (models.Article, error) {
	//TODO implement me
	panic("implement me")
}
