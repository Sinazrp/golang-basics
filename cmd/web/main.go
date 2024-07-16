package main

import (
	"fmt"
	"golang-bookingwebapp/pkg/handlers"
	"net/http"
)

const portNumber = ":8080"

func main() {

	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	http.HandleFunc("/favicon.ico", handlers.HandlerICon)

	fmt.Println(fmt.Sprintf("Listening onn port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
