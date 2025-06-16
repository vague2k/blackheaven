package main

import (
	"log"
	"net/http"
	"os"

	"github.com/a-h/templ"
	"github.com/vague2k/blackheaven/internal/router"
	"github.com/vague2k/blackheaven/views/assets"
	"github.com/vague2k/blackheaven/views/pages"
)

func main() {
	mux := http.NewServeMux()
	server := http.Server{
		Addr:    ":3000",
		Handler: mux,
	}

	router.SetupRoutes(mux)

	// mux.HandleFunc("GET /db/inquiry/select-inquiries", h.SelectInquiries)
	// mux.HandleFunc("GET /db/inquiry/select-inquiries", h.SelectInquiries)

	// pages
	SetupAssetsRoutes(mux)
	mux.Handle("GET /inquiry", templ.Handler(pages.Inquiry()))
	mux.Handle("GET /manager", templ.Handler(pages.ManagerView()))

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
			fs = http.FileServer(http.Dir("./views/assets"))
		} else {
			fs = http.FileServer(http.FS(assets.Assets))
		}

		fs.ServeHTTP(w, r)
	})

	mux.Handle("GET /assets/", http.StripPrefix("/assets/", assetHandler))
}
