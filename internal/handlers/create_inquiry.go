package handlers

import (
	"fmt"
	"net/http"

	"github.com/vague2k/blackheaven/internal/models"
	"github.com/vague2k/blackheaven/internal/services"
	"github.com/vague2k/blackheaven/views/components/toast"
)

const formID = "inquiry-form"

func CreateInquiry(w http.ResponseWriter, r *http.Request) {
	inquiryForm := &models.InquiryForm{}

	if err := scanForm(r, inquiryForm); err != nil {
		showToast("error", fmt.Sprintf("%s. Status code: %d", ErrInternal, http.StatusInternalServerError), w, r)
		return
	}

	topicSelectbox := inquiryForm.IsValidTopic()
	emailInput := inquiryForm.IsValidEmail()
	orderInput := inquiryForm.IsTopicOrder()
	contentTextarea := inquiryForm.IsContentEmpty()

	// if form has any errors, render all components with erronous state,
	// and render a toast with the first error from the list
	errs := inquiryForm.ErrMsgs()
	if len(errs) > 0 {
		showToast("error", errs[0], w, r)
		render(w, r,
			topicSelectbox,
			emailInput,
			orderInput,
			contentTextarea,
		)
		return
	}

	if inquiryForm.Subject == "" {
		inquiryForm.Subject = "New Message"
	}

	showToast("success", "Your form has been submitted", w, r)

	// TODO: service should return err and the database.Inquiry that was created
	services.CreateInquiry(inquiryForm)

	// TODO: redirect to /inquiry/submit/{id}/success if the creation of the inquiry was successful
	// w.Header().Add("Hx-Redirect", "/contacts")
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
