package handlers

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"net"
	"net/mail"
	"strings"
)

var (
	ErrInternal        = "Internal server error"
	ErrTopicRequired   = "Inquiry topic is required"
	ErrEmailRequired   = "Inquiry email is required"
	ErrEmailInvalid    = "Inquiry email is invalid"
	ErrContentRequired = "Inquiry content is required"
)

func isValidEmail(v string) error {
	if v == "" {
		return fmt.Errorf("%s", ErrEmailRequired)
	}

	email, err := mail.ParseAddress(v)
	if err != nil {
		return fmt.Errorf("%s", ErrEmailInvalid)
	}
	_, domain, _ := strings.Cut(email.Address, "@")

	_, err = net.LookupMX(domain)
	if err != nil {
		return fmt.Errorf("%s", ErrEmailInvalid)
	}

	if err := checkDisposable(domain); err != nil {
		return fmt.Errorf("%s", ErrEmailInvalid)
	}

	_, err = net.LookupTXT(domain)
	if err != nil {
		return fmt.Errorf("%s", ErrEmailInvalid)
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
		return fmt.Errorf("%s", ErrEmailInvalid)
	}
	return nil
}
