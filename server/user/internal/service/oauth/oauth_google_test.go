package service

import (
	"net/http/httptest"
	"net/url"
	"testing"
)

func TestLogin(t *testing.T) {
	oAuthGoogle := newOAuthGoogle(nil)

	req := httptest.NewRequest("POST", "/", nil)
	res := httptest.NewRecorder()

	oAuthGoogle.Login(res, req)

	result := res.Result()
	cookies := result.Cookies()

	if len(cookies) != 1 {
		t.Fatalf("Invalid cookie length. Expected %d but got %d instead", 1, len(cookies))
	}

	state := cookies[0]
	location := result.Header.Get("location")
	locationURL, err := url.Parse(location)
	if err != nil {
		t.Fatal(err)
	}

	locationState := locationURL.Query().Get("state")
	if locationState != state.Value {
		t.Fatalf(`Invalid redirect state. Expected "%s" but got "%s" instead`, state.Value, locationState)
	}
}
