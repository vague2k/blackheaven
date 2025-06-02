package handlers

import (
	"bufio"
	"bytes"
	"embed"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/mail"
	"strings"
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
			respErr(w, http.StatusBadRequest, "sum went wrong gang")
			return
		}
	}
	// err := json.NewDecoder(r.Body).Decode(request)
	// if err != nil {
	// 	respErr(w, http.StatusBadRequest, err.Error())
	// 	return
	// }

	if err := isValidInquiry(request.Kind); err != nil {
		respErr(w, http.StatusBadRequest, err.Error())
		return
	} else if err := isValidEmail(request.Email); err != nil {
		respErr(w, http.StatusBadRequest, err.Error())
		return
	} else if request.Content == "" {
		respErr(w, http.StatusBadRequest, ErrInquiryContentEmpty)
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

	b, err := json.Marshal(request)
	if err != nil {
		respErr(w, http.StatusBadRequest, err.Error())
		return
	}

	w.Header().Set("Content-Type", "application/json")
	_, err = w.Write(b)
	if err != nil {
		log.Fatal(err)
	}
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
