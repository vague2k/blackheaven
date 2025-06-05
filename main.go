package main

import (
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/vague2k/blackheaven/handlers"
	"github.com/vague2k/blackheaven/internal/assets"
	"github.com/vague2k/blackheaven/internal/pages"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	handler := handlers.NewHandler()
	mux.HandleFunc("POST /inquiry", handler.InquiryEndpoint)
	mux.HandleFunc("POST /inquiry/validate/email", handler.ValidateEmailEndpoint)

	// pages
	SetupAssetsRoutes(mux)
	mux.Handle("GET /inquiry", templ.Handler(pages.Inquiry()))

	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

func SetupAssetsRoutes(mux *http.ServeMux) {
	isDevelopment := os.Getenv("GO_ENV") != "production"

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./internal/assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
}
