package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	t.Run("Get Valid API Key", func(t *testing.T) {
		headers := http.Header{}
		expectedAPIKey := "secret_key"
		authHeaderValue := "ApiKey " + expectedAPIKey
		headers.Add("Authorization", authHeaderValue)

		actualAPIKey, err := GetAPIKey(headers)
		if err != nil {
			t.Error("expected err to be nil")
		}
		if actualAPIKey != expectedAPIKey {
			t.Errorf("expected API key: %v, got: %v", expectedAPIKey, actualAPIKey)
		}
	})
}
