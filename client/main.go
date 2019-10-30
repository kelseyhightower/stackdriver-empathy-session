package main

import (
	"flag"
	"log"
	"net/http"
)

var (
	frontendAddr string
)

func main() {
	flag.StringVar(&frontendAddr, "frontend", "http://127.0.0.1:8080", "The frontend address")
	flag.Parse()

	request, err := http.NewRequest("GET", frontendAddr, nil)
	if err != nil {
		log.Fatalf("error calling the frontend service: %v", err)
	}

	log.Println("calling the frontend service...")

	response, err := http.DefaultClient.Do(request)
	if err != nil {
		log.Fatalf("error calling the frontend service: %v", err)
	}

	if response.StatusCode != 200 {
		log.Fatalf("error calling the frontend service, http status: %s", response.Status)
	}

	log.Println("call to the frontend was successful.")
}
