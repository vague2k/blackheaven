package routes

import "net/http"

func SetupSwapRoutes(mux *http.ServeMux) {
	mux.HandleFunc("POST /validate/inquiry-form/all", validateInquiryForm)
}
