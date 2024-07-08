package main

import (
	"net/http"
)

func Home(res http.ResponseWriter, req *http.Request) {
	RenderTemplate(res, "home.page.tmpl")

}

func About(res http.ResponseWriter, req *http.Request) {
	RenderTemplate(res, "about.page.tmpl")

}
