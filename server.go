package main

import (
	"fmt"
	"html"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello you 2, %q", html.EscapeString(r.URL.Path))
	})

	log.Fatal(http.ListenAndServe(":443", nil))
}
