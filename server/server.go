package server

import (
	"embed"
	"log"
	"net/http"
	"os"
)

type Server struct {
	Port   string
	Router *Router

	logger *log.Logger
	server http.Server
}

func NewServer(port string) *Server {
	return &Server{
		Port:   port,
		Router: NewRouter(),
	}
}

func (s *Server) Run() error {
	s.server = http.Server{
		Addr:    s.Port,
		Handler: s.Router.mux,
	}
	return s.server.ListenAndServe()
}

func (s *Server) SetupAssets(assets embed.FS) {
	isDevelopment := os.Getenv("GO_ENV") != "production"

	assetHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if isDevelopment {
			w.Header().Set("Cache-Control", "no-store")
		}

		var fs http.Handler
		if isDevelopment {
			fs = http.FileServer(http.Dir("./views/assets"))
		} else {
			fs = http.FileServer(http.FS(assets))
		}

		fs.ServeHTTP(w, r)
	})

	s.Router.mux.Handle("/assets/", http.StripPrefix("/assets/", assetHandler))
}
