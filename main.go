package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		fprintf, err := fmt.Fprintf(writer, "Hello World")
		if err != nil {

			fmt.Println(err)
		}
		fmt.Println(fmt.Sprintf("number of bytes written: %d", fprintf))

	})
	_ = http.ListenAndServe(":8080", nil)

}
