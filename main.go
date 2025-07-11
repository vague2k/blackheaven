package main

import (
	"github.com/a-h/templ"
	"github.com/vague2k/blackheaven/internal/handlers"
	"github.com/vague2k/blackheaven/server"
	"github.com/vague2k/blackheaven/views/pages"
)

func main() {
	s := server.NewServer(":3000")
	s.SetupAssets()
	// service endpoints
	s.Router.Post("/create-inquiry", handlers.CreateInquiry)
	// pages
	s.Router.Handle("/inquiry", templ.Handler(pages.Inquiry()))
	s.Router.Handle("/inquiry/submit-successful", templ.Handler(pages.FormSubmitSuccessful()))
	s.Router.Handle("/manager", templ.Handler(pages.ManagerView()))

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
