package models

import "time"

type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title"`
	Contents string `json:"contents"`
	UserName string `json:"userName"`
	NiceNum  int    `json:"niceNum"`
	//CommentList []Comment `json:"commentList"`
	CreatedAt time.Time `json:"createdAt"`
	UpdatedAt time.Time `json:"updatedAt"`
}

type Comment struct {
	CommentID int       `json:"commentID"`
	ArticleID int       `json:"articleID"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt"`
}
