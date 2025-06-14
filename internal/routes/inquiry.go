package routes

import (
	"fmt"
	"net"
	"net/mail"
	"strings"
)

var (
	ErrInternal        = "Internal server error"
	ErrTopicRequired   = "Inquiry Topic is required"
	ErrEmailRequired   = "Email is required"
	ErrEmailInvalid    = "Email is invalid"
	ErrContentRequired = "Message is required"
	ErrOrderRequired   = "An order number is required if the topic is about an order"
)

type inquiry struct {
	Topic,
	Email,
	Name,
	Order,
	Subject,
	Content string
}

func (i *inquiry) IsValidTopic() error {
	switch i.Topic {
	case "order", "release", "submission", "general":
		return nil
	case "":
		return fmt.Errorf("%v", ErrTopicRequired)
	default:
		return fmt.Errorf("not a valid inquiry '%s'", i.Topic)
	}
}

func (i *inquiry) IsContentEmpty() error {
	if i.Content == "" {
		return fmt.Errorf("%v", ErrContentRequired)
	}
	return nil
}

func (i *inquiry) IsTopicOrder() error {
	if i.Order == "" && i.Topic == "order" {
		return fmt.Errorf("%v", ErrOrderRequired)
	}
	return nil
}

func (i *inquiry) IsValidEmail() error {
	if i.Email == "" {
		return fmt.Errorf("%v", ErrEmailRequired)
	}

	invalid := fmt.Errorf("%v", ErrEmailInvalid)

	email, err := mail.ParseAddress(i.Email)
	if err != nil {
		return invalid
	}
	_, domain, _ := strings.Cut(email.Address, "@")

	_, err = net.LookupMX(domain)
	if err != nil {
		return invalid
	}

	if err := checkDisposable(domain); err != nil {
		return invalid
	}

	_, err = net.LookupTXT(domain)
	if err != nil {
		return invalid
	}

	return nil
}
