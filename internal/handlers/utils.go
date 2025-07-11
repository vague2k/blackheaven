package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/a-h/templ"
	"github.com/vague2k/blackheaven/views/components/toast"
)

var ErrInternal = "Internal server error"

func scanForm(r *http.Request, dst any) error {
	if err := r.ParseForm(); err != nil {
		return fmt.Errorf("%v", ErrInternal)
	}

	v := reflect.ValueOf(dst).Elem()
	t := v.Type()

	for i := range v.NumField() {
		field := v.Field(i)
		structField := t.Field(i)
		formKey := strings.ToLower(structField.Name)

		if formValues, ok := r.Form[formKey]; ok && len(formValues) > 0 {
			switch field.Kind() {
			case reflect.String:
				field.SetString(formValues[0])
			}
		}
	}

	return nil
}

func render(w http.ResponseWriter, r *http.Request, components ...templ.Component) {
	if len(components) == 0 {
		return
	}
	for _, c := range components {
		if c == nil {
			continue
		}
		c.Render(r.Context(), w)
	}
}

func showToast(variant toast.Variant, description string, w http.ResponseWriter, r *http.Request) {
	var title string
	switch variant {
	case "error":
		title = "Form Error"
	case "success":
		title = "Form Successfully submitted"
	}
	toast.Toast(toast.Props{
		Icon:        true,
		Title:       title,
		Description: description,
		Variant:     variant,
		Position:    "top-center",
		Dismissible: true,
	}).Render(r.Context(), w)
}
