package main

import (
	"fmt"
	"io"
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

	http.HandleFunc("/sound.mp3", func(w http.ResponseWriter, r *http.Request) {
		file, err := os.Open("sound.mp3")
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err)
		}
		_, err = io.Copy(w, file)
		if err != nil {
			w.WriteHeader(500)
			fmt.Fprint(w, err)
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, content) })

	log.Fatal(http.ListenAndServe("0.0.0.0:"+servicePort, nil))
}
