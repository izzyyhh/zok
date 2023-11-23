package httpclient

import (
	"bytes"
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

func Post(url *string, data *string) ([]byte, error) {
	bodyBytes := []byte(*data)
	resp, err := httpclient.Post(*url, http.DetectContentType(bodyBytes), bytes.NewReader(bodyBytes))

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
