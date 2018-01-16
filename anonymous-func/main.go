package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("starting...")

	http.ListenAndServe(":9001", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("Handled %s %v", r.Method, r.URL.Path)

		fmt.Fprint(w, "Hello World\n")
	}))
}
