package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"
)

func main() {
	port := ":8080"

	if os.Getenv("PORT") != "" {
		port = fmt.Sprintf(":%s", os.Getenv("PORT"))
	}

	http.HandleFunc("/time", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, time.Now().Format(time.RFC3339))
	})

	log.Fatal(http.ListenAndServe(port, nil))
}
