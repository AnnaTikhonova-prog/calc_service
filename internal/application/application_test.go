package application

import (
	"bytes"
	"encoding/json"
	"net/http"
	"testing"
)

func TestCalculateHandler(t *testing.T) {
	tests := []struct {
		expression string
		statusCode int
	}{
		{"2+2*2", http.StatusOK},
		{"2+2*", http.StatusUnprocessableEntity},
		{"10/0", http.StatusInternalServerError},
	}

	for _, test := range tests {
		reqBody, _ := json.Marshal(map[string]string{"expression": test.expression})
		resp, err := http.Post("http://localhost:8080/api/v1/calculate", "application/json", bytes.NewBuffer(reqBody))
		if err != nil {
			t.Fatalf("Error making request: %v", err)
		}

		if resp.StatusCode != test.statusCode {
			t.Errorf("Expected status %d, got %d", test.statusCode, resp.StatusCode)
		}
	}
}
