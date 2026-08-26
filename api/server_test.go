package api

import (
	"coffeeware/service"
	"coffeeware/store"
	"net/http/httptest"
	"path/filepath"
	"testing"
)

func TestHealthEndpoint(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "db"))
	defer s.Close()
	req := httptest.NewRequest("GET", "/health", nil)
	rr := httptest.NewRecorder()
	New(service.New(s)).Routes().ServeHTTP(rr, req)
	if rr.Code != 204 {
		t.Fatal(rr.Code)
	}
}
