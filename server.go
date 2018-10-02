package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strings"
)

func main() {
	servicePort := os.Getenv("SERVER_LISTENING_PORT")
	if servicePort == "" {
		servicePort = "443"
	}
	assetsBase := "https://raw.githubusercontent.com/glutamatt/now-test-tmp/fe19110e57d1c5cdf80c19f9c232ced3c7b41efd/"

	home := homeContent()
	handler := &RegexpHandler{}

	handler.HandleFunc("/assets", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, assetsBase+strings.TrimPrefix(r.RequestURI, "/assets/"), http.StatusPermanentRedirect)
	})
	handler.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, home)
	})

	log.Println("Server listening on : http://0.0.0.0:" + servicePort)
	log.Fatal(http.ListenAndServe(":"+servicePort, handler))
}

func homeContent() string {
	css, err := ioutil.ReadFile("assets/theme.css")
	if err != nil {
		panic(err)
	}
	page, err := ioutil.ReadFile("assets/page.html")
	if err != nil {
		panic(err)
	}

	return "<style type=\"text/css\">" + string(css) + "</style>" + string(page)
}
