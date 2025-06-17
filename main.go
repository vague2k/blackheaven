package main

import (
	"github.com/vague2k/blackheaven/internal/handlers"
	"github.com/vague2k/blackheaven/server"
	"github.com/vague2k/blackheaven/views/assets"
	"github.com/vague2k/blackheaven/views/pages"
)

func main() {
	s := server.NewServer(":3000")

	s.SetupAssets(assets.Assets)

	// service endpoints
	s.Router.Post("/create-inquiry", handlers.CreateInquiry)

	// pages
	s.Router.HandleView("/inquiry", pages.Inquiry)
	s.Router.HandleView("/manager", pages.ManagerView)

	err := s.Run()
	if err != nil {
		panic(err)
	}
}
