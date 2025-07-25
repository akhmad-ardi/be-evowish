package tests

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestRegisterSuccess(t *testing.T) {
	app, _ := SetupApp()

	body := map[string]string{
		"name":             "Test User",
		"email":            "testuser@example.com",
		"password":         "password123",
		"confirm_password": "password123",
	}

	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/user/register", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req, 5000)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}

	t.Log("User registered")

	email := body["email"]
	DeleteDataUser(&email)
}

func TestRegisterInvalidEmail(t *testing.T) {
	app, _ := SetupApp()

	body := map[string]string{
		"name":             "Bad Email",
		"email":            "not-an-email",
		"password":         "password123",
		"confirm_password": "password123",
	}

	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/user/register", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %d", resp.StatusCode)
	}
}

func TestRegisterValidationError(t *testing.T) {
	app, _ := SetupApp()

	body := map[string]string{
		"name":             "",
		"email":            "",
		"password":         "password123",
		"confirm_password": "password123",
	}

	jsonBody, _ := json.Marshal(body)
	req := httptest.NewRequest("POST", "/api/user/register", bytes.NewReader(jsonBody))
	req.Header.Set("Content-Type", "application/json")

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusBadRequest {
		t.Errorf("Expected status 400 Bad Request, got %d", resp.StatusCode)
	}

	// Baca body respons
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		t.Fatalf("Failed to read response body: %v", err)
	}

	var jsonResp map[string]interface{}
	if err := json.Unmarshal(respBody, &jsonResp); err != nil {
		t.Fatalf("Failed to unmarshal response body: %v", err)
	}

	// Pastikan key validation_error ada
	if _, ok := jsonResp["validation_error"]; !ok {
		t.Errorf("Expected key 'validation_errors' in response")
	}

	// Pastikan validation_errors tidak kosong (karena name & email kosong)
	errors := jsonResp["validation_error"].(map[string]interface{})
	if len(errors) == 0 {
		t.Errorf("Expected validation_error to contain error messages, got empty object")
	}
}
