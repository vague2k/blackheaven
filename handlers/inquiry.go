package handlers

import (
	"fmt"
	"net/http"

	"github.com/vague2k/blackheaven/ui/components/toast"
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
			showInquiryErrorToast("Internal issue", "an internal server error has occured", w, r)
			return
		}
	}

	if err := isValidInquiry(request.Kind); err != nil {
		showInquiryErrorToast("Inquiry Topic", err.Error(), w, r)
		return
	} else if err := isValidEmail(request.Email); err != nil {
		showInquiryErrorToast("Email", err.Error(), w, r)
		return
	} else if request.Content == "" {
		showInquiryErrorToast("Content", ErrInquiryContentEmpty, w, r)
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

	// TODO: setup way to only get here through redirect and not outside source

	// w.Header().Set("HX-Redirect", "/inquiry/success")
	toast.Toast(toast.Props{
		Icon:        true,
		Title:       "Success",
		Description: "Your form has been submitted",
		Variant:     "success",
		Position:    "top-center",
		Dismissible: true,
	}).Render(r.Context(), w)
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

func showInquiryErrorToast(title, description string, w http.ResponseWriter, r *http.Request) {
	toast.Toast(toast.Props{
		Icon:        true,
		Title:       title,
		Description: description,
		Variant:     "error",
		Position:    "top-center",
		Dismissible: true,
	}).Render(r.Context(), w)
}
