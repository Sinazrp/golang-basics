package main

import (
	"fmt"
	"golang-bookingwebapp/pkg/config"
	"golang-bookingwebapp/pkg/handlers"
	"golang-bookingwebapp/pkg/render"
	"log"
	"net/http"
)

const portNumber = ":8080"

func main() {
	var app config.AppConfig
	app.UseCache = false
	tc, err := render.CreateTemplateCache()
	if err != nil {
		log.Fatal("Error initializing template cache")
	}
	app.TemplateCache = tc
	repo := handlers.NewRepository(&app)
	handlers.NewHandlers(repo)
	render.NewTemplates(&app)

	http.HandleFunc("/", handlers.Repo.Home)
	http.HandleFunc("/about", handlers.Repo.About)
	http.HandleFunc("/favicon.ico", handlers.HandlerICon)

	fmt.Println(fmt.Sprintf("Listening onn port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
