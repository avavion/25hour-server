package app

import "net/http"

type Interface interface {
	Run()
}

type App struct {
	provider ServiceProviderInterface
}

func (a *App) Run() {
	handler := a.provider.GetArticleHandler()

	server := &http.ServeMux{}

	server.HandleFunc("/api/articles", handler.GetAllArticles)
	server.HandleFunc("/api/articles/:id", handler.GetAllArticles)

	_ = http.ListenAndServe(":8080", server)
}

func NewApplication() Interface {
	return &App{NewServiceProvider()}
}
