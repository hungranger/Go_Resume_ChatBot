package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hungranger/Go_Resume_ChatBot/pkg/config"
)

type ConfigMock struct {
}

func (c ConfigMock) LoadConfig() config.IConfig {
	panic("not implemented")
}

func (c ConfigMock) GetPort() string {
	return "8080"
}

func (c ConfigMock) GetFbApi() string {
	panic("not implemented")
}

func (c ConfigMock) GetImage() string {
	panic("not implemented")
}

func TestHTTPServerHome(t *testing.T) {
	server := HTTPServer{ConfigMock{}}
	t.Run("Test HTTP server", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/", nil)
		response := httptest.NewRecorder()

		server.home(response, request)

		got := response.Body.String()
		want := "hello from home"

		if got != want {
			t.Errorf("got '%s', want '%s' ", got, want)
		}
	})
}

func TestHTTPServerMessage(t *testing.T) {
	server := HTTPServer{ConfigMock{}}
	t.Run("Test HTTP server 'message; handler", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodPost, "/", nil)
		response := httptest.NewRecorder()

		server.messages(response, request)

		got := response.Body.String()
		want := "hello from messages"

		if got != want {
			t.Errorf("got '%s', want '%s' ", got, want)
		}
	})
}
