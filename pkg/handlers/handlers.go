package handlers

import (
	"golang-bookingwebapp/pkg/render"
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home.page.gohtml")

}

func About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "about.page.gohtml")

}
func HandlerICon(w http.ResponseWriter, r *http.Request) {}
