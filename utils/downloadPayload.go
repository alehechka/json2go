package utils

import (
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

	return ioutil.ReadAll(resp.Body)
}
