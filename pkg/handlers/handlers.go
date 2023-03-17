package handlers

import (
	"net/http"

	"github.com/Faraday1987/fgoutils/pkg/config"
	"github.com/Faraday1987/fgoutils/pkg/models"
	"github.com/Faraday1987/fgoutils/pkg/render"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepo(a *config.AppConfig) *Repository {
	return &Repository{
		App: a,
	}
}

func NewHandlers(r *Repository) {
	Repo = r
}

func (f *Repository) Home(w http.ResponseWriter, r *http.Request) {

	render.RenderTemplate(w, "home.page.tmpl", &models.TemplateData{})
}

func (f *Repository) About(w http.ResponseWriter, r *http.Request) {

	//passing params to the template
	stringMap := make(map[string]string)
	stringMap["micoolparam"] = "This is mi cool param ❤️"

	render.RenderTemplate(w, "about.page.tmpl", &models.TemplateData{
		StringMap: stringMap,
	})
}
