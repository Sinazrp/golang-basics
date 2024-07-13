package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)

	fmt.Println(fmt.Sprintf("Listening onn port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
