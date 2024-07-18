package handlers

import (
	"golang-bookingwebapp/pkg/config"
	"golang-bookingwebapp/pkg/render"
	"net/http"
)

var Repo *Repository

type Repository struct {
	App *config.AppConfig
}

func NewRepository(app *config.AppConfig) *Repository {
	return &Repository{App: app}
}
func NewHandlers(r *Repository) {
	Repo = r
}

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home.page.gohtml")

}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "about.page.gohtml")

}
func HandlerICon(w http.ResponseWriter, r *http.Request) {}
