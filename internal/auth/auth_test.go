package auth

import (
	"net/http"
	"testing"
)

func TestGetAPIKey(t *testing.T) {
	header := http.Header{}
	header.Add("Authorization", "ApiKey testKey")
	key, err := GetAPIKey(header)
	if err != nil {
		t.Fatal("Recieved Error:", err)
	}
	if key != "testKey" {
		t.Fatal("Wrong key. Expected: testKey. Recieved:", key)
	}
}
