package http

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
)

const start_valid_status_code = 200
const end_valid_status_code = 400

func Download(url url.URL) (*http.Response, error) {
	log.Printf("Downloading %s\n", url.String())
	res, err := http.Get(url.String())
	if err != nil {
		return nil, err
	}
	if res.StatusCode < start_valid_status_code {
		return nil, fmt.Errorf("Invalid StatusCode %d", res.StatusCode)
	}
	if res.StatusCode >= end_valid_status_code {
		return nil, fmt.Errorf("Invalid StatusCode %d", res.StatusCode)
	}

	return res, nil
}
