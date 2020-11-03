package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
)

func main() {
	port := ":8080"
	timeServer := "localhost"

	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	if os.Getenv("TIME_SERVER") != "" {
		timeServer = os.Getenv("TIME_SERVER")
	}

	http.HandleFunc("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, fmt.Sprintf("Hello, World! %s\n", getTime(timeServer)))
	})

	log.Fatal(http.ListenAndServe(port, nil))
}

func getTime(timeServer string) string {
	resp, err := http.Get(fmt.Sprintf("http://%s:8000/time", timeServer))
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, _ := ioutil.ReadAll(resp.Body)
	return string(body)
}
