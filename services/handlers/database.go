package handlers

import (
	"fmt"
	"net/http"
)

// var ErrNonURLEncondedSemiColon = &APIError{http.StatusNotImplemented, "invalid URL encoded semicolon"}

// get req
func (h *Handler) SelectInquiries(w http.ResponseWriter, r *http.Request) {
	limit := r.URL.Query().Get("limit")

	// Add optional validation for "limit" (e.g., convert to int)
	if limit == "" {
		limit = "1" // default value if not set
	}

	// Example response (replace with actual logic)
	fmt.Fprintf(w, "Fetching inquiries with limit: %s", limit)
}
