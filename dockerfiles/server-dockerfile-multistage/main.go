package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		//fmt.Println("Hello, World!")
		fmt.Fprintf(rw, "Hello, World!\n")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
