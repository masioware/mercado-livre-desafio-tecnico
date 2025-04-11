package http

import (
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

type fakeResponse struct {
	Message string `json:"message"`
}

func TestDoRequest_UnsupportedMethod(t *testing.T) {
	err := DoRequest(RequestOptions{
		Method: "PATCH",
		URL:    "http://example.com",
		Result: &fakeResponse{},
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "unsupported HTTP method")
}

func TestDoRequest_HTTPErrorStatus(t *testing.T) {
	// servidor falso que retorna 500
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	err := DoRequest(RequestOptions{
		Method: "GET",
		URL:    server.URL,
		Result: &fakeResponse{},
	})

	assert.Error(t, err)
	assert.Contains(t, err.Error(), "API error")
}

func TestDoRequest_SuccessfulGET(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		if _, err := w.Write([]byte(`{"message":"ok"}`)); err != nil {
			log.Println("failed to write response:", err)
		}
	}))

	defer server.Close()

	var resp fakeResponse
	err := DoRequest(RequestOptions{
		Method: "GET",
		URL:    server.URL,
		Result: &resp,
	})

	assert.NoError(t, err)
	assert.Equal(t, "ok", resp.Message)
}
