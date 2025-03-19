package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestProxyHandler(t *testing.T) {
	// Create a test server to act as our target API
	targetServer := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Echo back request details for verification
		fmt.Fprintf(w, `{"path":"%s","method":"%s","user_agent":"%s"}`,
			r.URL.Path,
			r.Method,
			r.Header.Get("User-Agent"),
		)
	}))
	defer targetServer.Close()

	tests := []struct {
		name           string
		path           string
		method         string
		targetAPIURL   string
		envTargetAPI   string
		expectedStatus int
		expectedBody   string
	}{
		{
			name:           "GET request with header target API",
			path:           "/test/path",
			method:         "GET",
			targetAPIURL:   targetServer.URL,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"path":"/test/path","method":"GET"`,
		},
		{
			name:           "POST request with header target API",
			path:           "/api/data",
			method:         "POST",
			targetAPIURL:   targetServer.URL,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"path":"/api/data","method":"POST"`,
		},
		{
			name:           "Root path with no target API",
			path:           "/",
			method:         "GET",
			expectedStatus: http.StatusOK,
			expectedBody:   `{"status":"Reliable Proxy server is running"`,
		},
		{
			name:           "Missing target API URL",
			path:           "/api/data",
			method:         "GET",
			expectedStatus: http.StatusBadRequest,
			expectedBody:   `{"error":"Missing target_api_url header or TARGET_API_URL environment variable"}`,
		},
		{
			name:           "GET request with env target API",
			path:           "/test/env",
			method:         "GET",
			envTargetAPI:   targetServer.URL,
			expectedStatus: http.StatusOK,
			expectedBody:   `{"path":"/test/env","method":"GET"`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Set environment variable if needed
			if tt.envTargetAPI != "" {
				oldEnv := os.Getenv("TARGET_API_URL")
				os.Setenv("TARGET_API_URL", tt.envTargetAPI)
				defer os.Setenv("TARGET_API_URL", oldEnv)
			} else {
				os.Unsetenv("TARGET_API_URL")
			}

			// Create request
			req, err := http.NewRequest(tt.method, tt.path, nil)
			if err != nil {
				t.Fatalf("Could not create request: %v", err)
			}

			// Add headers
			if tt.targetAPIURL != "" {
				req.Header.Set("target-api-url", tt.targetAPIURL)
			}
			req.Header.Set("User-Agent", "ProxyTest")

			// Create response recorder
			rr := httptest.NewRecorder()

			// Serve the request
			handler := http.HandlerFunc(proxyHandler)
			handler.ServeHTTP(rr, req)

			// Check status code
			if status := rr.Code; status != tt.expectedStatus {
				t.Errorf("handler returned wrong status code: got %v want %v", status, tt.expectedStatus)
			}

			// Check response body
			if body := rr.Body.String(); !strings.Contains(body, tt.expectedBody) {
				t.Errorf("handler returned unexpected body: got %v, want to contain %v", body, tt.expectedBody)
			}
		})
	}
}

func TestLoadEnvFile(t *testing.T) {
	// Create a temporary .env file
	content := []byte(`
TEST_VAR1=value1
TEST_VAR2=value2
# This is a comment
TEST_VAR3=value with spaces
`)
	tmpfile, err := ioutil.TempFile("", "test.env")
	if err != nil {
		t.Fatalf("Failed to create temp file: %v", err)
	}
	defer os.Remove(tmpfile.Name())

	if _, err := tmpfile.Write(content); err != nil {
		t.Fatalf("Failed to write to temp file: %v", err)
	}
	if err := tmpfile.Close(); err != nil {
		t.Fatalf("Failed to close temp file: %v", err)
	}

	// Temporarily rename the file to .env
	os.Rename(tmpfile.Name(), ".env")
	defer os.Rename(".env", tmpfile.Name())

	// Clear environment variables
	os.Unsetenv("TEST_VAR1")
	os.Unsetenv("TEST_VAR2")
	os.Unsetenv("TEST_VAR3")

	// Run the function
	loadEnvFile()

	// Check if environment variables were set correctly
	if val := os.Getenv("TEST_VAR1"); val != "value1" {
		t.Errorf("Expected TEST_VAR1 to be 'value1', got '%s'", val)
	}
	if val := os.Getenv("TEST_VAR2"); val != "value2" {
		t.Errorf("Expected TEST_VAR2 to be 'value2', got '%s'", val)
	}
	if val := os.Getenv("TEST_VAR3"); val != "value with spaces" {
		t.Errorf("Expected TEST_VAR3 to be 'value with spaces', got '%s'", val)
	}
}

func TestRespondWithError(t *testing.T) {
	tests := []struct {
		name       string
		message    string
		statusCode int
	}{
		{
			name:       "Bad Request Error",
			message:    "Bad Request",
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Internal Server Error",
			message:    "Internal Server Error",
			statusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create response recorder
			rr := httptest.NewRecorder()

			// Call the function
			respondWithError(rr, tt.message, tt.statusCode)

			// Check status code
			if status := rr.Code; status != tt.statusCode {
				t.Errorf("respondWithError returned wrong status code: got %v want %v", status, tt.statusCode)
			}

			// Check response body
			expected := fmt.Sprintf(`{"error":"%s"}`, tt.message)
			if body := strings.TrimSpace(rr.Body.String()); body != expected {
				t.Errorf("respondWithError returned unexpected body: got %v want %v", body, expected)
			}
		})
	}
} 