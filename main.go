package main

import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

func home(res http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(res, "Welcome to the Home!")

}
func Divide(res http.ResponseWriter, req *http.Request) {
	f, err := divideValues(100.0, 0)
	if err != nil {
		http.Error(res, err.Error(), http.StatusInternalServerError)
		return
	}
	_, _ = fmt.Fprintf(res, "%d / %d = %f", 100, 10, f)
}

func about(res http.ResponseWriter, req *http.Request) {
	sum, _ := addValues(2, 1)
	_, _ = fmt.Fprintf(res, fmt.Sprintf("some of numbers is %d", sum))

}
func addValues(x, y int) (int, error) {

	return x + y, nil
}
func divideValues(x, y float32) (float32, error) {
	if y == 0 {
		return 0, fmt.Errorf("cannot divide by zero")
	}
	result := x / y
	return result, nil
}

func main() {
	http.HandleFunc("/", home)
	http.HandleFunc("/about", about)
	http.HandleFunc("/divide", Divide)

	fmt.Println(fmt.Sprintf("Listening on port %s", portNumber))
	_ = http.ListenAndServe(portNumber, nil)

}
