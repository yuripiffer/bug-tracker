package models

import "time"

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
