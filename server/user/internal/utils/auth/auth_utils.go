package utils

import (
	"io"
	"net/http"
	"net/url"
	"time"
)

func httpCall(method, rawURL, accessToken string) ([]byte, error) {
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
		return nil, nil
	}

	return body, nil
}
