package tests

import (
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestMainApi(t *testing.T) {
	app := SetupApp()

	req := httptest.NewRequest(http.MethodGet, "/api", nil)

	resp, err := app.Test(req)
	if err != nil {
		t.Fatalf("Failed to send request: %v", err)
	}

	if resp.StatusCode != http.StatusOK {
		t.Errorf("Expected status 200 OK, got %d", resp.StatusCode)
	}
}
