package db

import (
	"bugtracker-backend/internal/models"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateComment(t *testing.T) {
	// Setup
	err := Init()
	assert.NoError(t, err)
	defer Cleanup()

	tests := []struct {
		name       string
		bugID      string
		comment    *models.Comment
		shouldErr  bool
		errMessage string
	}{
		{
			name:  "Valid comment creation",
			bugID: "1",
			comment: &models.Comment{
				Author:  "Test User",
				Content: "Test Comment",
			},
			shouldErr: false,
		},
		{
			name:  "Invalid bug ID",
			bugID: "999",
			comment: &models.Comment{
				Author:  "Test User",
				Content: "Test Comment",
			},
			shouldErr:  true,
			errMessage: "bug not found",
		},
		{
			name:  "Invalid bug ID format",
			bugID: "invalid",
			comment: &models.Comment{
				Author:  "Test User",
				Content: "Test Comment",
			},
			shouldErr:  true,
			errMessage: "invalid bug ID format",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := CreateComment(tt.bugID, tt.comment)
			if tt.shouldErr {
				assert.Error(t, err)
				assert.Contains(t, err.Error(), tt.errMessage)
			} else {
				assert.NoError(t, err)
				assert.NotEmpty(t, tt.comment.ID)
				assert.NotEmpty(t, tt.comment.CreatedAt)
				assert.Equal(t, tt.bugID, tt.comment.BugID)
			}
		})
	}
}

func TestGetComments(t *testing.T) {
	// Setup
	err := Init()
	assert.NoError(t, err)
	defer Cleanup()

	// Create a bug first
	bug := &models.Bug{
		Title:       "Test Bug",
		Description: "Test Description",
		Status:      "Open",
		Priority:    "High",
	}
	err = CreateBug(bug)
	assert.NoError(t, err)

	// Create some test comments
	testComments := []*models.Comment{
		{Author: "User1", Content: "Comment 1"},
		{Author: "User2", Content: "Comment 2"},
		{Author: "User3", Content: "Comment 3"},
	}

	for _, comment := range testComments {
		err := CreateComment(strconv.Itoa(bug.ID), comment)
		assert.NoError(t, err)
	}

	tests := []struct {
		name        string
		bugID       string
		expectedLen int
		shouldErr   bool
	}{
		{
			name:        "Get existing comments",
			bugID:       strconv.Itoa(bug.ID),
			expectedLen: 3,
			shouldErr:   false,
		},
		{
			name:        "Get comments for non-existent bug",
			bugID:       "999",
			expectedLen: 0,
			shouldErr:   false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			comments, err := GetComments(tt.bugID)
			if tt.shouldErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Len(t, comments, tt.expectedLen)
				if tt.expectedLen > 0 {
					// Verify comment fields
					for _, comment := range comments {
						assert.NotEmpty(t, comment.ID)
						assert.NotEmpty(t, comment.CreatedAt)
						assert.Equal(t, tt.bugID, comment.BugID)
					}
				}
			}
		})
	}
}
