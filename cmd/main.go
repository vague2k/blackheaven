package main

import (
	"log"
	"net/http"

	"github.com/vague2k/blackheaven/handlers"
)

func main() {
	mux := http.NewServeMux()

	handler := handlers.NewHandler()
	mux.HandleFunc("POST /inquiry", handler.InquiryEndpoint)

	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		log.Fatal(err)
	}
}
