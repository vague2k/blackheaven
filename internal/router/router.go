package router

import (
	"net/http"

	"github.com/vague2k/blackheaven/internal/handlers"
)

func SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /create-inquiry", handlers.CreateInquiry)
}
