package main

import (
	"io/ioutil"
	"log"
	"net/http"
)

func main() {
	MakeRequest()
}

func MakeRequest() {
	for true {
		resp, err := http.Get("http://51.250.85.125:80")
		if err != nil {
			log.Fatalln(err)
		}

		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}

		log.Println(resp.StatusCode, string(body))
	}
}
