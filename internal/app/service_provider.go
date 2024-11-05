package app

import (
	"github.com/avavion/25hour-server/internal/handlers"
	"github.com/avavion/25hour-server/internal/repositories"
	"github.com/avavion/25hour-server/internal/services"
	"github.com/avavion/25hour-server/package/drivers/databases"
)

func NewServiceProvider() ServiceProviderInterface {
	return &ServiceProvider{
		databaseRepository: nil,
		articleHandler:     nil,
		articleService:     nil,
		articleRepository:  nil,
	}
}

type ServiceProviderInterface interface {
	GetArticleHandler() handlers.ArticleHandlerInterface
	GetDatabaseRepository() databases.DatabaseRepositoryInterface
	GetArticleService() services.ArticleServiceInterface
	GetArticleRepository() repositories.ArticleRepositoryInterface
}

type ServiceProvider struct {
	databaseRepository databases.DatabaseRepositoryInterface
	articleHandler     handlers.ArticleHandlerInterface
	articleService     services.ArticleServiceInterface
	articleRepository  repositories.ArticleRepositoryInterface
}

func (p *ServiceProvider) GetArticleHandler() handlers.ArticleHandlerInterface {
	if p.articleHandler == nil {
		p.articleHandler = handlers.NewArticleHandler(p.GetArticleService())
	}

	return p.articleHandler
}

func (p *ServiceProvider) GetDatabaseRepository() databases.DatabaseRepositoryInterface {
	if p.databaseRepository == nil {
		// TODO: тянуть из конфига
		p.databaseRepository, _ = databases.NewDatabaseRepository(databases.DatabaseConnection{
			Username:     "root",
			Password:     "root",
			DatabaseName: "25_hours",
			Host:         "localhost",
			Port:         3306,
		})
	}

	return p.databaseRepository
}

func (p *ServiceProvider) GetArticleService() services.ArticleServiceInterface {
	if p.articleService == nil {
		p.articleService = services.NewArticleService(p.GetArticleRepository())
	}

	return p.articleService
}

func (p *ServiceProvider) GetArticleRepository() repositories.ArticleRepositoryInterface {
	if p.articleRepository == nil {
		p.articleRepository = repositories.NewArticleRepository(p.GetDatabaseRepository())
	}

	return p.articleRepository
}
