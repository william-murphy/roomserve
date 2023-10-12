package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func MockPostRequest(t *testing.T, body interface{}, route string) *http.Request {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		t.Fatal("Error encoding body to json")
	}
	req, err := http.NewRequest(http.MethodPost, route, &buf)
	if err != nil {
		t.Fatal("Error creating NewRequest")
	}
	return req
}
