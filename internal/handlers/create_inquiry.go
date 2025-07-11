package handlers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/vague2k/blackheaven/internal/models"
	"github.com/vague2k/blackheaven/internal/services"
)

const formID = "inquiry-form"

func CreateInquiry(w http.ResponseWriter, r *http.Request) {
	time.Sleep(300 * time.Millisecond) // FIXME: DELETE THIS WHEN NOT IN DEV, MAKES BUTTON LOADING STATE OBVIOUS
	inquiryForm := &models.InquiryForm{}

	if err := scanForm(r, inquiryForm); err != nil {
		showToast("error", fmt.Sprintf("%s. Status code: %d", ErrInternal, http.StatusInternalServerError), w, r)
		return
	}

	topicSelectbox := inquiryForm.IsValidTopic()
	emailInput := inquiryForm.IsValidEmail()
	orderInput := inquiryForm.IsTopicOrder()
	contentTextarea := inquiryForm.IsContentEmpty()

	// NOTE: render new state ONLY if there's any erroneous states present
	// the "HX-Redirect" header needs to be the only thing returned from the resp to work properly
	errs := inquiryForm.ErrMsgs()
	if errs != nil {
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

	// TODO: service should return err and the database.Inquiry that was created
	services.CreateInquiry(inquiryForm)

	w.Header().Set("HX-Redirect", "/form/inquiry/successful")
}
