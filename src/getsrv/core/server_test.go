package core

import (
	"testing"
	"net/http"
	"net/http/httptest"
)

func TestRoutes(t *testing.T) {
	c := NewDummyContext()
	r := BuildRouter(c)
	request, _ := http.NewRequest("GET", "/healthz", nil)
    response := httptest.NewRecorder()
	r.ServeHTTP(response, request)
	if response.Code != 200 {
		t.Errorf("HTTP OK response is expected")
	}
}
