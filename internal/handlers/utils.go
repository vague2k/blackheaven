package handlers

import (
	"fmt"
	"net/http"
	"reflect"
	"strings"

	"github.com/a-h/templ"
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
	if components == nil || len(components) == 0 {
		return
	}
	for _, c := range components {
		c.Render(r.Context(), w)
	}
}
