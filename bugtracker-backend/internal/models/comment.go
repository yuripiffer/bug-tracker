package models

import (
	"fmt"
	"time"
)

type Comment struct {
	ID        string    `json:"id"`
	BugID     string    `json:"bugId"`
	Author    string    `json:"author"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"createdAt"`
}

type CreateCommentRequest struct {
	Author  string `json:"author"`
	Content string `json:"content"`
}

func (c *Comment) Validate() error {
	if c.Author == "" {
		return fmt.Errorf("author is required")
	}
	if c.Content == "" {
		return fmt.Errorf("content is required")
	}
	if c.BugID == "" {
		return fmt.Errorf("bug ID is required")
	}
	return nil
}

func (r *CreateCommentRequest) Validate() error {
	if r.Author == "" {
		return fmt.Errorf("author is required")
	}
	if r.Content == "" {
		return fmt.Errorf("content is required")
	}
	return nil
}
