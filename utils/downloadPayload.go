package utils

import (
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
)

// DownloadPayload calls a URL and returns the bytes of the response body.
func DownloadPayload(uri string) ([]byte, error) {
	u, err := url.ParseRequestURI(uri)
	if err != nil {
		return []byte{}, err
	}

	resp, err := http.Get(u.String())
	if err != nil {
		return []byte{}, err
	}

	if resp.StatusCode < 200 || resp.StatusCode >= 300 {
		return []byte{}, errors.New("failed to retrieve data from url")
	}

	return ioutil.ReadAll(resp.Body)
}
