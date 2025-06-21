package handlers

import (
	"fmt"
	"net/http"

	"github.com/a-h/templ"
	"github.com/vague2k/blackheaven/internal/models"
	"github.com/vague2k/blackheaven/internal/services"
	"github.com/vague2k/blackheaven/views/components/input"
	"github.com/vague2k/blackheaven/views/components/textarea"
	"github.com/vague2k/blackheaven/views/components/toast"
	"github.com/vague2k/blackheaven/views/modules"
)

const formID = "inquiry-form"

func CreateInquiry(w http.ResponseWriter, r *http.Request) {
	inquiry := &models.Inquiry{}
	var topicSelectbox, emailInput, orderInput, contentTextarea templ.Component
	var formHasError bool
	var errs []error

	if err := scanForm(r, inquiry); err != nil {
		showInquiryErrorToast(ErrInternal, fmt.Sprintf("Status code: %d", http.StatusInternalServerError), w, r)
		return
	}

	topicErr := inquiry.IsValidTopic()
	emailErr := inquiry.IsValidEmail()
	orderErr := inquiry.IsTopicOrder()
	contentErr := inquiry.IsContentEmpty()

	if topicErr != nil {
		formHasError = true
		errs = append(errs, topicErr)
		topicSelectbox = modules.FormSelectBox(modules.FormSelectBoxProps{
			FormID:      formID,
			Name:        "topic",
			Class:       "w-1/2",
			Label:       "Inquiry Topic",
			HasError:    true,
			Required:    true,
			Description: topicErr.Error(),
			Placeholder: "Select a topic",
			Options:     []string{"general", "order", "submission"},
			Attributes: templ.Attributes{
				"hx-swap-oob": "outerHTML:#inquiry-form-topic-element-container",
			},
		})
	} else {
		topicSelectbox = modules.FormSelectBox(modules.FormSelectBoxProps{
			FormID:      formID,
			Name:        "topic",
			Class:       "w-1/2",
			Label:       "Inquiry Topic",
			Required:    true,
			Value:       inquiry.Topic,
			Description: "What kind of topic is it?",
			Placeholder: "Select a topic",
			Options:     []string{"general", "order", "submission"},
			Attributes: templ.Attributes{
				"hx-swap-oob": "outerHTML:#inquiry-form-topic-element-container",
			},
		})
	}
	if emailErr != nil {
		formHasError = true
		errs = append(errs, emailErr)
		emailInput = modules.FormInput(modules.FormInputProps{
			FormID:      formID,
			Name:        "email",
			Class:       "w-1/2",
			Label:       "Email",
			HasError:    true,
			Required:    true,
			Description: emailErr.Error(),
			Value:       inquiry.Email,
			Placeholder: "johnsmith@email.com",
			Attributes: templ.Attributes{
				"hx-swap-oob": "outerHTML:#inquiry-form-email-element-container",
			},
		})
	} else {
		emailInput = modules.FormInput(modules.FormInputProps{
			FormID:      formID,
			Name:        "email",
			Class:       "w-1/2",
			Label:       "Email",
			HasNoError:  true,
			Required:    true,
			Description: "Looks good to me!",
			Value:       inquiry.Email,
			Placeholder: "johnsmith@email.com",
			Attributes: templ.Attributes{
				"hx-swap-oob": "outerHTML:#inquiry-form-email-element-container",
			},
		})
	}
	if orderErr != nil {
		formHasError = true
		errs = append(errs, orderErr)
		orderInput = modules.FormInput(modules.FormInputProps{
			Class:       "w-1/2",
			FormID:      formID,
			Name:        "order",
			Label:       "Order #",
			HasError:    true,
			Required:    true,
			Description: orderErr.Error(),
			Type:        input.TypeText,
			Placeholder: "Order # here",
			Attributes: templ.Attributes{
				"hx-swap-oob": "outerHTML:#inquiry-form-order-element-container",
			},
		})
	}
	if contentErr != nil {
		formHasError = true
		errs = append(errs, contentErr)
		contentTextarea = modules.FormTextarea(modules.FormTextareaProps{
			FormID:   formID,
			Name:     "content",
			Label:    "Message",
			Required: true,
			HasError: true,
			TextareaProps: textarea.Props{
				Placeholder: "What do you have to say...",
				AutoResize:  true,
			},
			Attributes: templ.Attributes{
				"hx-swap-oob": "outerHTML:#inquiry-form-content-element-container",
			},
		})
	}

	if formHasError {
		showInquiryErrorToast("Form error", errs[0].Error(), w, r)
		render(w, r,
			topicSelectbox,
			emailInput,
			orderInput,
			contentTextarea,
		)
		return
	}

	if inquiry.Subject == "" {
		inquiry.Subject = "New Message"
	}

	toast.Toast(toast.Props{
		Icon:        true,
		Title:       "Success",
		Description: "Your form has been submitted",
		Variant:     "success",
		Position:    "top-center",
		Dismissible: true,
	}).Render(r.Context(), w)

	services.CreateInquiry(inquiry)
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
