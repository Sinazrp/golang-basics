package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func RenderTemplate(res http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".html")

	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		_, _ = fmt.Fprintf(res, "Error executing template")
	}

}
