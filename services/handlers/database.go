package handlers

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"strings"
)

var (
	ErrParamParseInt  = &APIError{http.StatusInternalServerError, "error parsing int from url parameter"}
	ErrFetchDBRecords = &APIError{http.StatusInternalServerError, "error fetching inquiries from database"}
	ErrMarshal        = &APIError{http.StatusInternalServerError, "error while trying to marshall into json"}
)

func (h *Handler) SelectInquiries(w http.ResponseWriter, r *http.Request) {
	limitParam := r.URL.Query().Get("limit")
	if limitParam == "" || strings.HasPrefix(limitParam, "-") {
		limitParam = "1" // default value if not set
	}

	// NOTE: at this point, the param should ALWAYS be a "positive int" string,
	// if this errors something has gone terribly wrong
	limit, err := strconv.ParseInt(limitParam, 0, 64)
	if err != nil {
		http.Error(w, ErrParamParseInt.msg, ErrParamParseInt.status)
		return
	}

	inquiries, err := h.DB.SelectInquiries(r.Context(), limit)
	if err != nil {
		log.Println(err)
		http.Error(w, ErrFetchDBRecords.msg, ErrFetchDBRecords.status)
		return
	}

	encodedInquiries, err := json.Marshal(inquiries)
	if err != nil {
		http.Error(w, ErrMarshal.msg, ErrMarshal.status)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(encodedInquiries)

	// Example response (replace with actual logic)
	fmt.Fprintf(w, "Fetching inquiries with limit: %s", limitParam)
}
