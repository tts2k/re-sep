package utils

import (
	"crypto/rand"
	"encoding/base64"
	"io"
	"net/http"
	"net/url"
	"time"
)

func HTTPCall(method, rawURL, accessToken string) ([]byte, error) {
	httpClient := http.Client{
		Timeout: 30 * time.Second,
	}

	URL, _ := url.Parse(rawURL)
	queries := URL.Query()
	queries.Add("accessToken", accessToken)
	URL.RawQuery = queries.Encode()

	request := http.Request{
		Method: method,
		URL:    URL,
	}

	response, err := httpClient.Do(&request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, err
	}

	return body, nil
}

func RandString(nByte int) string {
	b := make([]byte, nByte)
	rand.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

func SetCallbackCookie(w http.ResponseWriter, r *http.Request, name, value string) {
	c := &http.Cookie{
		Name:     name,
		Value:    value,
		Secure:   r.TLS != nil,
		HttpOnly: true,
	}
	http.SetCookie(w, c)
}
