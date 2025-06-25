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

	"github.com/a-h/templ"
	"github.com/vague2k/blackheaven/views/modules/form"
)

var (
	ErrInternal        = "Internal server error"
	ErrTopicRequired   = "Inquiry Topic is required"
	ErrEmailRequired   = "Email is required"
	ErrEmailInvalid    = "Email is invalid"
	ErrContentRequired = "Message is required"
	ErrOrderRequired   = "Order number is required"
)

const formID = "inquiry-form"

type InquiryForm struct {
	Topic,
	Email,
	Name,
	Order,
	Subject,
	Content string

	errMsgs []string
}

func (i *InquiryForm) ErrMsgs() []string {
	return i.errMsgs
}

func (i *InquiryForm) IsValidTopic() templ.Component {
	var hasError bool
	var desc string
	switch i.Topic {
	case "order", "release", "submission", "general":
		desc = "What kind of topic is it?"
	default:
		hasError = true
		desc = ErrTopicRequired
		i.errMsgs = append(i.errMsgs, desc)
	}
	return form.Selectbox(form.SelectboxProps{
		FormID:      formID,
		Name:        "topic",
		Class:       "w-1/2",
		Label:       "Inquiry Topic",
		HasError:    hasError,
		Required:    true,
		Description: desc,
		Value:       i.Topic,
		Placeholder: "Select a topic",
		Options:     []string{"general", "order", "submission"},
		Attributes: templ.Attributes{
			"hx-swap-oob": "outerHTML:#inquiry-form-topic-element-container",
		},
	})
}

func (i *InquiryForm) IsContentEmpty() templ.Component {
	var hasError bool
	desc := "The message box will expand as you type"
	if i.Content == "" {
		hasError = true
		desc = ErrContentRequired
		i.errMsgs = append(i.errMsgs, desc)
	}
	return form.Textarea(form.TextareaProps{
		FormID:      formID,
		Name:        "content",
		Label:       "Message",
		Required:    true,
		HasError:    hasError,
		Description: desc,
		Placeholder: "What do you have to say...",
		AutoResize:  true,
		Value:       i.Content,
		Attributes: templ.Attributes{
			"hx-swap-oob": "outerHTML:#inquiry-form-content-element-container",
		},
	})
}

func (i *InquiryForm) IsTopicOrder() templ.Component {
	var hasError, required bool
	desc := "Required if your topic is about an order"
	if i.Order == "" && i.Topic == "order" {
		hasError = true
		required = true
		desc = ErrOrderRequired
		i.errMsgs = append(i.errMsgs, desc)
	}
	return form.Input(form.InputProps{
		Class:       "w-1/2",
		FormID:      formID,
		Name:        "order",
		Label:       "Order #",
		HasError:    hasError,
		Required:    required,
		Description: desc,
		Value:       i.Order,
		Type:        "text",
		Placeholder: "Order # here",
		Attributes: templ.Attributes{
			"hx-swap-oob": "outerHTML:#inquiry-form-order-element-container",
		},
	})
}

func (i *InquiryForm) IsValidEmail() templ.Component {
	var hasError, hasNoError bool
	var desc string
	if err := i.isValidEmail(); err != nil {
		hasError = true
		desc = err.Error()
		i.errMsgs = append(i.errMsgs, desc)
	} else {
		desc = "Looks good to me!"
		hasNoError = true
	}
	return form.Input(form.InputProps{
		FormID:      formID,
		Name:        "email",
		Class:       "w-1/2",
		Label:       "Email",
		HasError:    hasError,
		HasNoError:  hasNoError,
		Required:    true,
		Description: desc,
		Value:       i.Email,
		Placeholder: "johnsmith@email.com",
		Attributes: templ.Attributes{
			"hx-swap-oob": "outerHTML:#inquiry-form-email-element-container",
		},
	})
}

func (i *InquiryForm) isValidEmail() error {
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
