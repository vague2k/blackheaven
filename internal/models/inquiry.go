package models

import (
	"bufio"
	"bytes"
	"embed"
	"fmt"
	"net"
	"net/mail"
	"os"
	"strings"
)

var (
	ErrInternal        = "Internal server error"
	ErrTopicRequired   = "Inquiry Topic is required"
	ErrEmailRequired   = "Email is required"
	ErrEmailInvalid    = "Email is invalid"
	ErrContentRequired = "Message is required"
	ErrOrderRequired   = "Order number is required"
)

type Inquiry struct {
	Topic,
	Email,
	Name,
	Order,
	Subject,
	Content string
}

func (i *Inquiry) IsValidTopic() error {
	switch i.Topic {
	case "order", "release", "submission", "general":
		return nil
	case "":
		return fmt.Errorf("%v", ErrTopicRequired)
	default:
		return fmt.Errorf("not a valid inquiry '%s'", i.Topic)
	}
}

func (i *Inquiry) IsContentEmpty() error {
	if i.Content == "" {
		return fmt.Errorf("%v", ErrContentRequired)
	}
	return nil
}

func (i *Inquiry) IsTopicOrder() error {
	if i.Order == "" && i.Topic == "order" {
		return fmt.Errorf("%v", ErrOrderRequired)
	}
	return nil
}

func (i *Inquiry) IsValidEmail() error {
	if i.Email == "" {
		return fmt.Errorf("%v", ErrEmailRequired)
	}

	invalid := fmt.Errorf("%v", ErrEmailInvalid)

	email, err := mail.ParseAddress(i.Email)
	if err != nil {
		fmt.Fprint(os.Stdout, err)
		return invalid
	}
	_, domain, _ := strings.Cut(email.Address, "@")

	_, err = net.LookupMX(domain)
	if err != nil {
		fmt.Fprint(os.Stdout, err)
		return invalid
	}

	if err := checkDisposable(domain); err != nil {
		fmt.Fprint(os.Stdout, err)
		return invalid
	}

	_, err = net.LookupTXT(domain)
	if err != nil {
		fmt.Fprint(os.Stdout, err)
		return invalid
	}

	return nil
}

//go:embed disposable_emails.txt
var f embed.FS

// TODO: the embedded list of disposable emails is over 1.5 mb big. perhaps
// reading the file and creating the hashmap could be done concurrently since
// order doesn't matter anyways?

func checkDisposable(v string) error {
	hashMap := make(map[string]bool)
	b, err := f.ReadFile("disposable_emails.txt")
	if err != nil {
		return err
	}
	scanner := bufio.NewScanner(bytes.NewReader(b))

	for scanner.Scan() {
		line := scanner.Text()
		hashMap[line] = true
	}

	if hashMap[v] {
		return fmt.Errorf("%v", ErrEmailInvalid)
	}
	return nil
}
