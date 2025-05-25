package main

import (
	"log"
	"net/http"

	"github.com/vague2k/blackheaven/handlers"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	handler := handlers.NewHandler()
	mux.HandleFunc("POST /inquiry", handler.InquiryEndpoint)

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
