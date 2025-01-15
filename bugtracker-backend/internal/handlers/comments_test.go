package handlers

import (
	"bugtracker-backend/internal/db"
	"bugtracker-backend/internal/models"
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"os"
	"strconv"
	"testing"

	"bugtracker-backend/internal/testutil"

	"github.com/gorilla/mux"
	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	// Set test database path
	os.Setenv("DB_PATH", testutil.GetTestDBPath())
	defer testutil.CleanupTestDB()

	err := db.Init()
	assert.NoError(t, err)
	defer func() {
		db.CleanupTestDB()
		db.Cleanup()
	}()

	// Create a test bug first
	bug := &models.Bug{
		Title:       "Test Bug",
		Description: "Test Description",
	}
	err = db.CreateBug(bug)
	assert.NoError(t, err)

	tests := []struct {
		name           string
		bugID          string
		payload        interface{}
		expectedStatus int
		expectedError  string
	}{
		{
			name:  "Valid comment creation",
			bugID: "1",
			payload: models.CreateCommentRequest{
				Author:  "Test User",
				Content: "Test Comment",
			},
			expectedStatus: http.StatusCreated,
		},
		{
			name:  "Invalid bug ID",
			bugID: "999",
			payload: models.CreateCommentRequest{
				Author:  "Test User",
				Content: "Test Comment",
			},
			expectedStatus: http.StatusNotFound,
			expectedError:  "bug not found",
		},
		{
			name:  "Missing content",
			bugID: "1",
			payload: models.CreateCommentRequest{
				Author: "Test User",
			},
			expectedStatus: http.StatusBadRequest,
			expectedError:  "content is required",
		},
		{
			name:           "Invalid JSON",
			bugID:          "1",
			payload:        `{"invalid": json}`,
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid request body",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var body bytes.Buffer
			if str, ok := tt.payload.(string); ok {
				body.WriteString(str)
			} else {
				err := json.NewEncoder(&body).Encode(tt.payload)
				assert.NoError(t, err)
			}

			req := httptest.NewRequest("POST", "/api/bugs/"+tt.bugID+"/comments", &body)
			w := httptest.NewRecorder()

			// Set up router to handle URL parameters
			router := mux.NewRouter()
			router.HandleFunc("/api/bugs/{bugId}/comments", CreateComment)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedError != "" {
				var resp map[string]string
				err := json.NewDecoder(w.Body).Decode(&resp)
				assert.NoError(t, err)
				assert.Contains(t, resp["error"], tt.expectedError)
			} else {
				var comment models.Comment
				err := json.NewDecoder(w.Body).Decode(&comment)
				assert.NoError(t, err)
				assert.NotEmpty(t, comment.ID)
				assert.NotEmpty(t, comment.CreatedAt)
				assert.Equal(t, tt.bugID, comment.BugID)
			}
		})
	}
}

func TestGetComments(t *testing.T) {
	// Set test database path
	os.Setenv("DB_PATH", testutil.GetTestDBPath())
	defer testutil.CleanupTestDB()

	err := db.Init()
	assert.NoError(t, err)
	defer func() {
		db.CleanupTestDB()
		db.Cleanup()
	}()

	// Create a test bug
	bug := &models.Bug{
		Title:       "Test Bug",
		Description: "Test Description",
	}
	err = db.CreateBug(bug)
	assert.NoError(t, err)

	// Create some test comments
	testComments := []*models.Comment{
		{Author: "User1", Content: "Comment 1"},
		{Author: "User2", Content: "Comment 2"},
		{Author: "User3", Content: "Comment 3"},
	}

	for _, comment := range testComments {
		err := db.CreateComment(strconv.Itoa(bug.ID), comment)
		assert.NoError(t, err)
	}

	tests := []struct {
		name           string
		bugID          string
		expectedStatus int
		expectedCount  int
		expectedError  string
	}{
		{
			name:           "Get existing comments",
			bugID:          "1",
			expectedStatus: http.StatusOK,
			expectedCount:  3,
		},
		{
			name:           "Get comments for non-existent bug",
			bugID:          "999",
			expectedStatus: http.StatusNotFound,
			expectedError:  "bug not found",
		},
		{
			name:           "Invalid bug ID",
			bugID:          "invalid",
			expectedStatus: http.StatusBadRequest,
			expectedError:  "invalid bug ID",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			req := httptest.NewRequest("GET", "/api/bugs/"+tt.bugID+"/comments", nil)
			w := httptest.NewRecorder()

			router := mux.NewRouter()
			router.HandleFunc("/api/bugs/{bugId}/comments", GetComments)
			router.ServeHTTP(w, req)

			assert.Equal(t, tt.expectedStatus, w.Code)

			if tt.expectedError != "" {
				var resp map[string]string
				err := json.NewDecoder(w.Body).Decode(&resp)
				assert.NoError(t, err)
				assert.Contains(t, resp["error"], tt.expectedError)
			} else {
				var comments []models.Comment
				err := json.NewDecoder(w.Body).Decode(&comments)
				assert.NoError(t, err)
				assert.Len(t, comments, tt.expectedCount)
			}
		})
	}
}
