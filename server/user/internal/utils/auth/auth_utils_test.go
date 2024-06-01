package utils

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func newUnstartedTestServer() *httptest.Server {
	mux := http.NewServeMux()

	return httptest.NewUnstartedServer(mux)
}

func newTestServer() *httptest.Server {
	srv := newUnstartedTestServer()
	srv.Start()
	return srv
}

func TestHttpCall(t *testing.T) {
	url := "https://www.googleapis.com/oauth2/v2/userinfo"

	response, err := httpCall("GET", url, "")
	if err != nil {
		t.Fatal(err)
	}

	t.Fatal(string(response))
}
