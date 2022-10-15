package models

import "time"

type Article struct {
	ID          int
	Title       string
	Contents    string
	UserName    string
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
