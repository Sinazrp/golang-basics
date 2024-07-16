package handlers

import (
	"golang-bookingwebapp/pkg/render"
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "home")

}

func About(res http.ResponseWriter, req *http.Request) {
	render.RenderTemplate(res, "about")

}
