package main

import (
	"github.com/bmizerany/pat"
	"golang-bookingwebapp/pkg/config"
	"golang-bookingwebapp/pkg/handlers"
	"net/http"
)

func routes(app *config.AppConfig) http.Handler {
	mux := pat.New()
	mux.Get("/home", http.HandlerFunc(handlers.Repo.Home))
	mux.Get("/about", http.HandlerFunc(handlers.Repo.About))
	return mux

}
