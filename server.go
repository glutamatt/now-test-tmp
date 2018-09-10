package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	servicePort := os.Getenv("SERVER_LISTENING_PORT")
	if servicePort == "" {
		servicePort = "443"
	}

	css, err := ioutil.ReadFile("assets/theme.css")
	if err != nil {
		panic(err)
	}
	page, err := ioutil.ReadFile("assets/page.html")
	if err != nil {
		panic(err)
	}

	content := "<style type=\"text/css\">" + string(css) + "</style>" + string(page)

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, content) })

	log.Fatal(http.ListenAndServe(":"+servicePort, nil))
}
