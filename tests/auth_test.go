package tests

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterEndpoint(t *testing.T) {
	app := SetupApp()

	body := map[string]string{
		"name":             "Test User",
		"email":            "testuser@example.com",
		"password":         "password123",
		"confirm_password": "password123",
	}

	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/register", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}

func TestRegisterInvalidEmail(t *testing.T) {
	app := SetupApp()

	body := map[string]string{
		"name":             "Bad Email",
		"email":            "not-an-email",
		"password":         "password123",
		"confirm_password": "password123",
	}

	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/register", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %d", resp.StatusCode)
	}
}
