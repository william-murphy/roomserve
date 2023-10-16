package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func MockPostRequest(t *testing.T, body interface{}, route string) *http.Request {
	var buf bytes.Buffer
	err := json.NewEncoder(&buf).Encode(body)
	if err != nil {
		t.Fatal("Error encoding body to json")
	}
	req := httptest.NewRequest(http.MethodPost, route, &buf)
	return req
}
