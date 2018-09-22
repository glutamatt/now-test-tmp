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

	onlineAssets := map[string]string{
		"/sound.mp3":      "https://raw.githubusercontent.com/glutamatt/now-test-tmp/fe19110e57d1c5cdf80c19f9c232ced3c7b41efd/sound.mp3",
		"/background.jpg": "https://raw.githubusercontent.com/glutamatt/now-test-tmp/fe19110e57d1c5cdf80c19f9c232ced3c7b41efd/background.jpg",
	}

	for path, dest := range onlineAssets {
		func(path, dest string) {
			http.HandleFunc(path, func(w http.ResponseWriter, r *http.Request) {
				file, err := http.Get(dest)
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
		}(path, dest)

	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) { fmt.Fprint(w, content) })

	log.Fatal(http.ListenAndServe(":"+servicePort, nil))
}
