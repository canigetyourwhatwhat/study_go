package models

import "time"

type Article struct {
	ID       int    `json:"id"`
	Title    string `json:"title" `
	Contents string `json:"contents"`
	UserName string `json:"userName"`
	// I had to add db because of sqlx doesn't recognize under bar (_) in db.
	NiceNum int `json:"niceNum" db:"nice_num"`
	//CommentList []Comment `json:"commentList"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Comment struct {
	CommentID int       `json:"commentID"`
	ArticleID int       `json:"articleID"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}
