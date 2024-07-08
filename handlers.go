package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func renderTemplate(res http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".html")

	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		_, _ = fmt.Fprintf(res, "Error executing template")
	}

}

func Home(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "home.page.tmpl")

}

func About(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "about.page.tmpl")

}
