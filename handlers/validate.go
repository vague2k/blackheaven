package handlers

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"net"
	"net/http"
	"net/mail"
	"strings"

	"github.com/a-h/templ"
	"github.com/vague2k/blackheaven/internal/components/form"
	"github.com/vague2k/blackheaven/internal/components/input"
	"github.com/vague2k/blackheaven/internal/modules"
)

func (h *Handler) ValidateEmailEndpoint(w http.ResponseWriter, r *http.Request) {
	var email string
	err := r.ParseForm()
	if err != nil {
		respErr(w, http.StatusBadRequest, err.Error())
		return
	}

	for k, v := range r.Form {
		if k == "email" {
			email = v[0]
			break
		}
	}

	var component templ.Component
	if err := isValidEmail(email); err != nil {
		form.Description(form.DescriptionProps{})
		component = modules.FormInput(modules.FormInputProps{
			FormID:      "inquiry-form",
			Name:        "email",
			Label:       "Email",
			Description: "Not a valid email, please try again",
			InputProps: input.Props{
				HasError:    true,
				Class:       "focus:outline-destructive",
				Value:       email,
				Placeholder: "johnsmith@email.com",
				Attributes: templ.Attributes{
					"hx-post":              "/inquiry/validate/email",
					"hx-target":            "#inquiry-form-email-element-container",
					"hx-trigger":           "keyup delay:500ms changed",
					"hx-on:htmx:afterSwap": "document.getElementById('inquiry-form-email-input').focus();",
				},
			},
			FormDescProps: form.DescriptionProps{
				Class: "text-xs mt-2 text-destructive",
			},
		})
	} else {
		component = modules.FormInput(modules.FormInputProps{
			FormID:      "inquiry-form",
			Name:        "email",
			Label:       "Email",
			Description: "Email valid!",
			InputProps: input.Props{
				Class:       "border-success",
				Value:       email,
				Placeholder: "johnsmith@email.com",
				Attributes: templ.Attributes{
					"hx-post":              "/inquiry/validate/email",
					"hx-target":            "#inquiry-form-email-element-container",
					"hx-trigger":           "keyup delay:500ms changed",
					"hx-on:htmx:afterSwap": "document.getElementById('inquiry-form-email-input').focus();",
				},
			},
			FormDescProps: form.DescriptionProps{
				Class: "text-xs mt-2 text-success",
			},
		})
	}
	component.Render(r.Context(), w)
}

func isValidEmail(v string) error {
	if v == "" {
		return fmt.Errorf("%s", ErrInquiryEmailEmpty)
	}

	email, err := mail.ParseAddress(v)
	if err != nil {
		return fmt.Errorf("%s", ErrInquiryEmailInvalid)
	}
	_, domain, _ := strings.Cut(email.Address, "@")

	_, err = net.LookupMX(domain)
	if err != nil {
		return fmt.Errorf("%s", ErrInquiryEmailInvalid)
	}

	if err := checkDisposable(domain); err != nil {
		return fmt.Errorf("%s", ErrInquiryEmailInvalid)
	}

	_, err = net.LookupTXT(domain)
	if err != nil {
		return fmt.Errorf("%s", ErrInquiryEmailInvalid)
	}

	return nil
}

//go:embed disposable.txt
var f embed.FS

// TODO: the embedded list of disposable emails is over 1.5 mb big. perhaps
// reading the file and creating the hashmap could be done concurrently since
// order doesn't matter anyways?

func checkDisposable(v string) error {
	hashMap := make(map[string]bool)
	b, err := f.ReadFile("disposable.txt")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))

	for scanner.Scan() {
		line := scanner.Text()
		hashMap[line] = true
	}

	if hashMap[v] {
		return fmt.Errorf("%s", ErrInquiryEmailInvalid)
	}
	return nil
}
