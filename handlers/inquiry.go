package handlers

import (
	"fmt"
	"net/http"

	"github.com/vague2k/blackheaven/ui/modules"
)

type Inquiry struct {
	Kind     string `json:"kind"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	OrderNum string `json:"order"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
}

func (h *Handler) InquiryEndpoint(w http.ResponseWriter, r *http.Request) {
	request := &Inquiry{}

	err := r.ParseForm()
	if err != nil {
		respErr(w, http.StatusBadRequest, err.Error())
		return
	}

	for k, v := range r.Form {
		switch k {
		case "kind":
			request.Kind = v[0]
		case "email":
			request.Email = v[0]
		case "name":
			request.Name = v[0]
		case "order":
			request.OrderNum = v[0]
		case "subject":
			request.Subject = v[0]
		case "content":
			request.Content = v[0]
		default:
			respErr(w, http.StatusBadRequest, "sum went wrong gang")
			return
		}
	}

	if err := isValidInquiry(request.Kind); err != nil {
		respErr(w, http.StatusBadRequest, err.Error())
		return
	} else if err := isValidEmail(request.Email); err != nil {
		respErr(w, http.StatusBadRequest, err.Error())
		return
	} else if request.Content == "" {
		respErr(w, http.StatusBadRequest, ErrInquiryContentEmpty)
		return
	}

	if request.Name == "" {
		request.Name = "No name given"
	}
	if request.OrderNum == "" {
		request.OrderNum = "No order number given"
	}
	if request.Subject == "" {
		request.Subject = "New Message"
	}

	modules.InquirySuccess().Render(r.Context(), w)
}

func isValidInquiry(v string) error {
	switch v {
	case "order":
		return nil
	case "release":
		return nil
	case "submission":
		return nil
	case "":
		return fmt.Errorf("%s", ErrInquiryKindEmpty)
	}

	return fmt.Errorf("not a valid inquiry '%s'", v)
}
