package tests

import "testing"

func TestSetupApp(t *testing.T) {
	_, err := SetupApp()

	if err != nil {
		t.Fatalf("Error Setup App")
	}
}
