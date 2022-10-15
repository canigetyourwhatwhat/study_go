package models

import "time"

type Article struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Contents    string
	UserName    string `json:"userName"`
	NiceNum     int
	CommentList []Comment
	CreatedAt   time.Time
}

type Comment struct {
	CommentID int
	ArticleID int
	Message   string
	CreatedAt time.Time
}
