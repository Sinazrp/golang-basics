package handlers

import (
	"golang-bookingwebapp/pkg/config"
	"golang-bookingwebapp/pkg/models"
	"golang-bookingwebapp/pkg/render"
	"net/http"
)

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

func NewRepository(app *config.AppConfig) {
	Repo = &Repository{App: app}
}

//func NewHandlers(r *Repository) {
//	Repo = r
//}

func (m *Repository) Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home.page.gohtml", &models.TemplateData{})

}

func (m *Repository) About(res http.ResponseWriter, req *http.Request) {
	stringMap := make(map[string]string)
	stringMap["test"] = "1.0.0"
	render.RenderTemplate(res, "about.page.gohtml", &models.TemplateData{StringMap: stringMap})

}
func HandlerICon(w http.ResponseWriter, r *http.Request) {}
