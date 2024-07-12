package service

import (
	"net/http/httptest"
	"testing"

	"golang.org/x/oauth2"
)

func newTestOauthConf(url string) oauth2.Config {
	return oauth2.Config{
		ClientID:     "CLIENT_ID",
		ClientSecret: "CLIENT_SECRET",
		RedirectURL:  "REDIRECT_URL",
		Scopes:       []string{"scope1", "scope2"},
		Endpoint: oauth2.Endpoint{
			AuthURL:  url + "/auth",
			TokenURL: url + "/token",
		},
	}
}

func TestGoogleLogin(t *testing.T) {
	req := httptest.NewRequest("POST", "/", nil)
	res := httptest.NewRecorder()

	Login(res, req)

	result := res.Result()
	cookies := result.Cookies()

	for _, v := range cookies {
		println(v.Name, ":", v.Value)
	}

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
