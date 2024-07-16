package render

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

func RenderTemplateTest(res http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/"+tmpl+".page.gohtml", "./templates/base.layout.gohtml")

	err := parsedTemplate.Execute(res, nil)
	if err != nil {
		_, _ = fmt.Fprintf(res, "Error executing template")
	}

}

var tc = make(map[string]*template.Template)

func RenderTemplate(res http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error

	//check to see if we already have template in cache
	_, inMap := tc[t]
	if !inMap {
		// create template
		log.Println("creating new template")
		err = createTemplateCache(t)
		if err != nil {
			log.Fatal("Error creating template cache")
		}
	} else {
		log.Println("using cached template")
	}
	tmpl = tc[t]
	err = tmpl.Execute(res, nil)
	if err != nil {
		log.Fatal("Error creating template cache")
	}

}

func createTemplateCache(t string) error {
	templatePath := "templates/" + t + ".page.gohtml"
	templates := []string{
		templatePath,
		"templates/base.layout.gohtml",
	}
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	tc[t] = tmpl
	return nil

}
