package service

import (
	"net/http/httptest"
	"testing"
)

func TestGoogleLogin(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	res := httptest.NewRecorder()

	Login(res, req)

	result := res.Result()
	cookies := result.Cookies()

	if len(cookies) != 2 {
		t.Fatalf("Invalid cookie length. Expected %d but got %d instead", 2, len(cookies))
	}

	state := cookies[0]
	location, err := result.Location()
	if err != nil {
		t.Fatal(err)
	}

	locationState := location.Query().Get("state")
	if locationState != state.Value {
		t.Fatalf(`Invalid redirect state. Expected "%s" but got "%s" instead`, state.Value, locationState)
	}
}
