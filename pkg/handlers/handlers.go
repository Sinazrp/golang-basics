package handlers

import (
	"golang-bookingwebapp/pkg/config"
	"golang-bookingwebapp/pkg/render"
	"net/http"
)

type TemplateData struct {
	StringMap map[string]string
	IntMap    map[string]int
	FloatMap  map[string]float32
	Data      map[string]interface{}
	CSRFToken string
	Flash     string
	Warning   string
	Error     string
}

type Repository struct {
	App *config.AppConfig
}

var Repo *Repository

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
