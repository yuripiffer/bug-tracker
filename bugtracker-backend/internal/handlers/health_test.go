package handlers

import (
	"bugtracker-backend/internal/config"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthCheck(t *testing.T) {
	// Create a request to pass to our handler
	req, err := http.NewRequest("GET", "/api/health", nil)
	assert.NoError(t, err)

	// Create a ResponseRecorder to record the response
	rr := httptest.NewRecorder()
	handler := http.HandlerFunc(HealthCheck)

	// Call the handler
	handler.ServeHTTP(rr, req)

	// Check the status code
	assert.Equal(t, http.StatusOK, rr.Code)

	// Check the response body
	var response HealthResponse
	err = json.NewDecoder(rr.Body).Decode(&response)
	assert.NoError(t, err)

	// Verify the response fields
	assert.Equal(t, "ok", response.Status)
	assert.Equal(t, config.Backend_Version, response.Version)

	// Verify Content-Type header
	assert.Equal(t, "application/json", rr.Header().Get("Content-Type"))
} 