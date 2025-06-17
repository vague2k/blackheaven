package server

import (
	"net/http"

	"github.com/a-h/templ"
)

type Router struct {
	mux *http.ServeMux
}

func NewRouter() *Router {
	return &Router{
		mux: http.NewServeMux(),
	}
}

func (r *Router) Get(pattern string, handler http.HandlerFunc) {
	r.mux.Handle("GET "+pattern, handler)
}

func (r *Router) Post(pattern string, handler http.HandlerFunc) {
	r.mux.Handle("POST "+pattern, handler)
}

func (r *Router) HandleView(pattern string, component func() templ.Component) {
	r.mux.Handle("GET "+pattern, templ.Handler(component()))
}
