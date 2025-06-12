package handlers

import (
	"bufio"
	"bytes"
	"embed"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/mail"
	"reflect"
	"strings"
)

var (
	ErrInternal        = "Internal server error"
	ErrTopicRequired   = "Inquiry topic is required"
	ErrEmailRequired   = "Inquiry email is required"
	ErrEmailInvalid    = "Inquiry email is invalid"
	ErrContentRequired = "Inquiry content is required"
)

func ScanForm(r *http.Request, dst any) error {
	if err := r.ParseForm(); err != nil {
		return errors.New(ErrInternal)
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

func isValidEmail(v string) error {
	if v == "" {
		return errors.New(ErrEmailInvalid)
	}

	email, err := mail.ParseAddress(v)
	if err != nil {
		return fmt.Errorf("%s", ErrEmailInvalid)
	}
	_, domain, _ := strings.Cut(email.Address, "@")

	_, err = net.LookupMX(domain)
	if err != nil {
		return errors.New(ErrEmailInvalid)
	}

	if err := checkDisposable(domain); err != nil {
		return errors.New(ErrEmailInvalid)
	}

	_, err = net.LookupTXT(domain)
	if err != nil {
		return errors.New(ErrEmailInvalid)
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
		return errors.New(ErrEmailInvalid)
	}
	return nil
}
