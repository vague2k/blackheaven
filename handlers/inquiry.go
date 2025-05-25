package handlers

import (
	"encoding/json"
	"log"
	"net/http"
)

type Inquiry struct {
	Type    string `json:"type"`
	Subject string `json:"subject"`
	Content string `json:"content"`
}

func (h *Handler) InquiryEndpoint(w http.ResponseWriter, r *http.Request) {
	reqInquiry := &Inquiry{}

	err := json.NewDecoder(r.Body).Decode(reqInquiry)
	if err != nil {
		respErr(w, http.StatusBadRequest, err.Error())
		return
	}

	if reqInquiry.Type == "" {
		respErr(w, http.StatusBadRequest, ErrInquiryTypeEmpty)
		return
	} else if reqInquiry.Content == "" {
		respErr(w, http.StatusBadRequest, ErrInquiryContentEmpty)
		return
	} else if reqInquiry.Subject == "" {
		reqInquiry.Subject = "New Message"
	}

	b, err := json.Marshal(reqInquiry)
	if err != nil {
		respErr(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err)
	}
}
