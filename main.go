package main

import (
	"fmt"
	"net/http"
	"text/template"
)

const portNumber = ":8080"

func renderTemplate(res http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl + ".html")

	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		_, _ = fmt.Fprintf(res, "Error executing template")
	}

}

func home(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "home.page.tmpl")

}

func about(res http.ResponseWriter, req *http.Request) {
	renderTemplate(res, "about.page.tmpl")

}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)

	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
