package handlers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestServerStartup(t *testing.T) {
	// Create a new router
	router := mux.NewRouter()

	// Register routes
	RegisterRoutes(router)

	// Create a test server
	ts := httptest.NewServer(router)
	defer ts.Close()

	// Make a request to the health check endpoint
	resp, err := http.Get(ts.URL + "/api/health")
	assert.NoError(t, err)
	assert.Equal(t, http.StatusOK, resp.StatusCode)
}
