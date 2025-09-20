package util

import (
	"fmt"
	"net/url"
)

func ExtractCodeFromURL(rawURL string) (string, error) {
	parsedURL, err := url.Parse(rawURL)
	if err != nil {
		return "", fmt.Errorf("invalid URL: %w", err)
	}

	queryParams, err := url.ParseQuery(parsedURL.RawQuery)
	if err != nil {
		return "", fmt.Errorf("invalid query parameters: %w", err)
	}

	code := queryParams.Get("code")
	if code == "" {
		return "", fmt.Errorf("code not found in URL")
	}

	return code, nil
}
