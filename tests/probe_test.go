package tests

import (
	"github.com/kenjitheman/url-probe/core"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestProbeURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	}))
	defer server.Close()

	url := server.URL

	err := core.ProbeURL(url)

	if err != nil {
		t.Errorf("ProbeURL returned an error for a valid URL: %v", err)
	}
}
