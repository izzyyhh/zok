package httpclient

import (
	"io"
	"net/http"
)

var httpclient = &http.Client{}

const (
	HTTP  = "http"
	HTTPS = "https"
)

func Get(url string) ([]byte, error) {
	resp, err := httpclient.Get(url)

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	body, err := io.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}

	return body, nil
}
