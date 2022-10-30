package entity

import "time"

type Article struct {
	Title    string `json:"title" `
	Contents string `json:"contents"`
	UserName string `json:"userName"`
	// I had to add db because of sqlx doesn't recognize under bar (_) in db.
	NiceNum   int       `json:"niceNum" db:"nice_num"`
	Comments  []Comment `json:"comments"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
	UpdatedAt time.Time `json:"updatedAt" db:"updated_at"`
}

type Comment struct {
	ArticleID int       `json:"articleID" db:"article_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"createdAt" db:"created_at"`
}

type ArticleComment struct {
	ArticleID            int       `json:"article_id" db:"article_id"`
	Title                string    `json:"title"`
	Contents             string    `json:"contents"`
	UserName             string    `json:"userName"`
	NiceNum              int       `json:"niceNum" db:"nice_num"`
	ArticleCreatedAt     time.Time `json:"article_created_at" db:"article_created_at"`
	ArticleUpdatedAt     time.Time `json:"article_updated_at" db:"article_updated_at"`
	ArticleIDFromComment int       `json:"article_id_from_comment" db:"article_id_from_comment"`
	Message              string    `json:"message"`
	CommentCreatedAt     time.Time `json:"comment_create_at" db:"comment_create_at"`
}
