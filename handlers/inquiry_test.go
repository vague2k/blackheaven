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
	t.Run("Errors on invalid inquiry kind values", func(t *testing.T) {
		testCases := []struct {
			name     string
			payload  map[string]string
			expected string
		}{
			{
				name:     "inquiry kind is empty string",
				payload:  map[string]string{"kind": ""},
				expected: "inquiry kind can't be empty",
			},
			{
				name:     "inquiry kind is missing entirely", // simulate missing inquiry kind
				payload:  map[string]string{},
				expected: "inquiry kind can't be empty",
			},
			{
				name:     "inquiry kind is 'asd asda asd'",
				payload:  map[string]string{"kind": "asd asda asd"},
				expected: "not a valid inquiry 'asd asda asd'",
			},
			{
				name:     "inquiry kind is 'kindnotexist'",
				payload:  map[string]string{"kind": "kindnotexist"},
				expected: "not a valid inquiry 'kindnotexist'",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := json.Marshal(tc.payload)
				assert.NoError(t, err)
				handler := NewHandler()
				w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, bytes.NewBuffer(b))
				resp := w.Result()
				defer resp.Body.Close()

				var respErr RespErr
				err = json.NewDecoder(resp.Body).Decode(&respErr)
				assert.NoError(t, err)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				assert.Equal(t, http.StatusBadRequest, respErr.Status)
				assert.Equal(t, tc.expected, respErr.ErrorMsg)
			})
		}
	})
	t.Run("Errors on invalid inquiry email values", func(t *testing.T) {
		testCases := []struct {
			name     string
			payload  map[string]string
			expected string
		}{
			{
				name: "inquiry kind is empty string",
				payload: map[string]string{
					"kind":  "order",
					"email": "",
				},
				expected: "inquiry email can't be empty",
			},
			{
				name: "inquiry kind is empty string", // simulate missing inquiry kind
				payload: map[string]string{
					"kind": "order",
				},
				expected: "inquiry email can't be empty",
			},
			{
				name: "inquiry kind is 'test@doesntexist.fuckyou'",
				payload: map[string]string{
					"kind":  "order",
					"email": "test@doesn'texist.com",
				},
				expected: "inquiry email is invalid",
			},
			{
				name: "inquiry email is 'whatdafuqbro'",
				payload: map[string]string{
					"kind":  "order",
					"email": "whatdafuqbro",
				},
				expected: "inquiry email is invalid",
			},
			{
				name: "inquiry email is 'test@0wnd.net'",
				payload: map[string]string{
					"kind":  "order",
					"email": "test@0wnd.net",
				},
				expected: "inquiry email is invalid",
			},
			{
				name: "inquiry email is 'test@1clck2.com'",
				payload: map[string]string{
					"kind":  "order",
					"email": "test@1clck2.com",
				},
				expected: "inquiry email is invalid",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := json.Marshal(tc.payload)
				assert.NoError(t, err)
				handler := NewHandler()
				w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, bytes.NewBuffer(b))
				resp := w.Result()
				defer resp.Body.Close()

				var respErr RespErr
				err = json.NewDecoder(resp.Body).Decode(&respErr)
				assert.NoError(t, err)

				assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
				assert.Equal(t, http.StatusBadRequest, respErr.Status)
				assert.Equal(t, tc.expected, respErr.ErrorMsg)
			})
		}
	})
	t.Run("Ok on common inquiry email values", func(t *testing.T) {
		testCases := []struct {
			name     string
			payload  map[string]string
			expected string
		}{
			{
				name: "inquiry email is 'test@gmail.com'",
				payload: map[string]string{
					"kind":    "order",
					"email":   "test@gmail.com",
					"content": "test",
				},
				expected: "",
			},
			{
				name: "inquiry email is 'test@yahoo.com'",
				payload: map[string]string{
					"kind":    "order",
					"email":   "test@yahoo.com",
					"content": "test",
				},
				expected: "",
			},
			{
				name: "inquiry email is 'test@yahoo.es'",
				payload: map[string]string{
					"kind":    "order",
					"email":   "test@yahoo.es",
					"content": "test",
				},
				expected: "",
			},
			{
				name: "inquiry email is 'test@aol.com'",
				payload: map[string]string{
					"kind":    "order",
					"email":   "test@aol.com",
					"content": "test",
				},
				expected: "",
			},
			{
				name: "inquiry email is 'test@proton.me'",
				payload: map[string]string{
					"kind":    "order",
					"email":   "test@proton.me",
					"content": "test",
				},
				expected: "",
			},
		}

		for _, tc := range testCases {
			t.Run(tc.name, func(t *testing.T) {
				b, err := json.Marshal(tc.payload)
				assert.NoError(t, err)
				handler := NewHandler()
				w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, bytes.NewBuffer(b))
				resp := w.Result()
				defer resp.Body.Close()

				var actual RespErr
				err = json.NewDecoder(resp.Body).Decode(&actual)
				assert.NoError(t, err)

				assert.Equal(t, http.StatusOK, resp.StatusCode)
			})
		}
	})
	t.Run("Errors on invalid inquiry content values", func(t *testing.T) {
		payload := map[string]string{
			"kind":    "order",
			"email":   "test@gmail.com",
			"content": "",
		}
		b, err := json.Marshal(payload)
		assert.NoError(t, err)
		handler := NewHandler()
		w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, bytes.NewBuffer(b))
		resp := w.Result()
		defer resp.Body.Close()

		var respErr RespErr
		err = json.NewDecoder(resp.Body).Decode(&respErr)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusBadRequest, resp.StatusCode)
		assert.Equal(t, http.StatusBadRequest, respErr.Status)
		assert.Equal(t, "inquiry content can't be empty", respErr.ErrorMsg)
	})
	t.Run("Inquiry name defaults if empty string", func(t *testing.T) {
		payload := map[string]string{
			"kind":    "order",
			"email":   "test@gmail.com",
			"content": "test",
		}
		b, err := json.Marshal(payload)
		assert.NoError(t, err)
		handler := NewHandler()
		w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, bytes.NewBuffer(b))
		resp := w.Result()
		defer resp.Body.Close()

		var inquiry Inquiry
		err = json.NewDecoder(resp.Body).Decode(&inquiry)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "order", inquiry.Kind)
		assert.Equal(t, "test@gmail.com", inquiry.Email)
		assert.Equal(t, "No name given", inquiry.Name)
		assert.Equal(t, "New Message", inquiry.Subject)
		assert.Equal(t, "test", inquiry.Content)
	})
	t.Run("Inquiry subject defaults if empty string", func(t *testing.T) {
		payload := map[string]string{
			"kind":    "order",
			"email":   "test@gmail.com",
			"name":    "name",
			"content": "test",
		}
		b, err := json.Marshal(payload)
		assert.NoError(t, err)
		handler := NewHandler()
		w, _ := endpoint("POST", "/inquiry", handler.InquiryEndpoint, bytes.NewBuffer(b))
		resp := w.Result()
		defer resp.Body.Close()

		var inquiry Inquiry
		err = json.NewDecoder(resp.Body).Decode(&inquiry)
		assert.NoError(t, err)

		assert.Equal(t, http.StatusOK, resp.StatusCode)
		assert.Equal(t, "order", inquiry.Kind)
		assert.Equal(t, "test@gmail.com", inquiry.Email)
		assert.Equal(t, "name", inquiry.Name)
		assert.Equal(t, "New Message", inquiry.Subject)
		assert.Equal(t, "test", inquiry.Content)
	})
}
