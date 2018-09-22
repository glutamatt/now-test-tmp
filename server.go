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
		file, err := http.Get("https://raw.githubusercontent.com/glutamatt/now-test-tmp/7a46bb9a6e24b97058a17949cb2ad5cfca3dbf68/sound.mp3")
		if err != nil {
			fmt.Fprint(w, err.Error())
			return
		}
		defer file.Body.Close()
		_, err = io.Copy(w, file.Body)
		if err != nil {
			fmt.Fprint(w, err.Error())
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, content) })

	log.Fatal(http.ListenAndServe(":"+servicePort, nil))
}
