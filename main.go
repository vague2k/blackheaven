package main

import (
	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/vague2k/blackheaven/internal/handlers"
	"github.com/vague2k/blackheaven/server"
	"github.com/vague2k/blackheaven/views/pages"
)

func main() {
	s := server.NewServer(":3000")
	s.Router.Use(middleware.Logger)

	s.SetupAssets()

	// service endpoints
	s.Router.Post("/create-inquiry", handlers.CreateInquiry)

	// pages
	s.Router.Handle("/inquiry", templ.Handler(pages.Inquiry()))
	s.Router.Handle("/manager", templ.Handler(pages.ManagerView()))

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
