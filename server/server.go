package server

import (
	"net/http"
	"os"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vague2k/blackheaven/views/assets"
)

type Server struct {
	Port   string
	Router *chi.Mux

	server http.Server
}

func NewServer(port string) *Server {
	s := &Server{
		Port:   port,
		Router: chi.NewRouter(),
	}

	s.Router.Use(middleware.Logger)
	return s
}

// TODO: implement graceful shutdown
func (s *Server) Run() error {
	s.server = http.Server{
		Addr:    s.Port,
		Handler: s.Router,
	}
	return s.server.ListenAndServe()
}

func (s *Server) SetupAssets() {
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

	s.Router.Handle("GET /assets/*", http.StripPrefix("/assets/", assetHandler))
}
