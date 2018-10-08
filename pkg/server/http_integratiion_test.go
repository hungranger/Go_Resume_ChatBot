package server

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hungranger/Go_Resume_ChatBot/pkg/config"
)

func TestHttpServer(t *testing.T) {
	iConfig := config.ProvideIniConfig()
	httpServer := ProvideServer(iConfig)

	request, _ := http.NewRequest(http.MethodGet, "/", nil)
	response := httptest.NewRecorder()

	httpServer.home(response, request)

	got := response.Body.String()
	want := "hello from Hung"

	if got != want {
		t.Errorf("got '%s', want '%s' ", got, want)
	}
}
