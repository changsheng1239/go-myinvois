package myinvois

import (
	"io"
	"net/http"
)

func newRequest(httpMethod, endpoint string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(httpMethod, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("Accept-Language", "en")
	req.Header.Set("Content-Type", "application/json")

	return req, nil
}

func newAuthenticatedRequest(httpMethod, endpoint, token string, body io.Reader) (*http.Request, error) {
	req, err := newRequest(httpMethod, endpoint, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Authorization", "Bearer "+token)

	return req, nil
}
