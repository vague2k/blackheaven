package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/vague2k/blackheaven/internal/components/input"
	"github.com/vague2k/blackheaven/internal/components/toast"
	"github.com/vague2k/blackheaven/internal/modules"
)

type Inquiry struct {
	Topic    string `json:"topic"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	OrderNum string `json:"order"`
	Subject  string `json:"subject"`
	Content  string `json:"content"`
}

func (h *Handler) InquiryEndpoint(w http.ResponseWriter, r *http.Request) {
	i := &Inquiry{}

	err := r.ParseForm()
	if err != nil {
		showInquiryErrorToast(ErrInternal, fmt.Sprintf("Status code: %d", http.StatusInternalServerError), w, r)
		return
	}

	for k, v := range r.Form {
		switch k {
		case "topic":
			i.Topic = v[0]
		case "email":
			i.Email = v[0]
		case "name":
			i.Name = v[0]
		case "order":
			i.OrderNum = v[0]
		case "subject":
			i.Subject = v[0]
		case "content":
			i.Content = v[0]
		default:
			showInquiryErrorToast("Internal issue", "an internal server error has occured", w, r)
			return
		}
	}

	if err := isValidTopic(i.Topic); err != nil {
		showInquiryErrorToast("Inquiry Topic", err.Error(), w, r)
		return
	}
	if err := isValidEmail(i.Email); err != nil {
		showInquiryErrorToast("Email", err.Error(), w, r)
		modules.FormInput(modules.FormInputProps{
			FormID:      "inquiry-form",
			Name:        "email",
			Class:       "w-1/2",
			Label:       "Email",
			HasError:    true,
			Required:    true,
			Description: "Not a valid email, try again",
			InputProps: input.Props{
				Value:       i.Email,
				Placeholder: "johnsmith@email.com",
			},
			Attributes: templ.Attributes{
				"hx-swap-oob": "outerHTML:#inquiry-form-email-element-container",
			},
		}).Render(r.Context(), w)
		return
	}
	if i.Content == "" {
		showInquiryErrorToast("Content", ErrContentRequired, w, r)
		return
	}

	if i.Name == "" {
		i.Name = "No name given"
	}
	if i.OrderNum == "" {
		i.OrderNum = "No order number given"
	}
	if i.Subject == "" {
		i.Subject = "New Message"
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

func isValidTopic(v string) error {
	switch v {
	case "order", "release", "submission", "general":
		return nil
	case "":
		return fmt.Errorf("%s", ErrTopicRequired)
	default:
		return fmt.Errorf("not a valid inquiry '%s'", v)
	}
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
