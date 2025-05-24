package handlers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func endpoint(method, target string, handler http.HandlerFunc, payload io.Reader) (*httptest.ResponseRecorder, *http.Request) {
	r := httptest.NewRequest(method, target, payload)
	w := httptest.NewRecorder()
	handler(w, r)

	return w, r
}

func TestInquiryEndpoint(t *testing.T) {
	t.Run("Errors if payload is EOF/empty body", func(t *testing.T) {
		handler := NewHandler()
		w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, nil)
		resp := w.Result()
		respErr := &RespErr{}
		err := json.NewDecoder(resp.Body).Decode(respErr)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, http.StatusBadRequest, respErr.Status)
		assert.Equal(t, "EOF", respErr.ErrorMsg)
	})
	t.Run("Errors on invalid inquiry type values", func(t *testing.T) {
		testCases := []struct {
			name     string
			payload  map[string]string
			expected string
		}{
			{
				name:     "inquiry type is empty string",
				payload:  map[string]string{"type": ""},
				expected: "inquiry type can't be empty",
			},
			{
				name:     "inquiry type is missing entirely", // simulate missing inquiry type
				payload:  map[string]string{},
				expected: "inquiry type can't be empty",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, _ := json.Marshal(tc.payload)
				handler := NewHandler()
				w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, bytes.NewBuffer(b))
				resp := w.Result()
				defer resp.Body.Close()

				var respErr RespErr
				err := json.NewDecoder(resp.Body).Decode(&respErr)
				assert.NoError(t, err)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				assert.Equal(t, http.StatusBadRequest, respErr.Status)
				assert.Equal(t, tc.expected, respErr.ErrorMsg)
			})
		}
	})
	t.Run("Errors on invalid inquiry content values", func(t *testing.T) {
		payload := map[string]string{
			"type":    "test",
			"content": "",
		}
		b, _ := json.Marshal(payload)
		handler := NewHandler()
		w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, bytes.NewBuffer(b))
		resp := w.Result()
		defer resp.Body.Close()

		var respErr RespErr
		err := json.NewDecoder(resp.Body).Decode(&respErr)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, http.StatusBadRequest, respErr.Status)
		assert.Equal(t, "inquiry content can't be empty", respErr.ErrorMsg)
	})
	t.Run("Inquiry subject defaults if empty string", func(t *testing.T) {
		payload := map[string]string{
			"type":    "test",
			"content": "test",
		}
		b, _ := json.Marshal(payload)
		handler := NewHandler()
		w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, bytes.NewBuffer(b))
		resp := w.Result()
		defer resp.Body.Close()

		var inquiry Inquiry
		err := json.NewDecoder(resp.Body).Decode(&inquiry)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "test", inquiry.Type)
		assert.Equal(t, "New Message", inquiry.Subject)
		assert.Equal(t, "test", inquiry.Content)
	})
}
