package auth

import (
	"net/http"
)

const AUTH_HEADER = "Authorization"

type ApiKeyNotFound struct{}
type InvalidApiKey struct{}

func (e ApiKeyNotFound) Error() string {
	return "API key not found"
}

func (e InvalidApiKey) Error() string {
	return "Invalid API key"
}

func GetAPIKey(header http.Header) (string, error) {
	apikey := header.Get(AUTH_HEADER)

	if apikey == "" {
		return "", ApiKeyNotFound{}
	}

	return apikey, nil
}
