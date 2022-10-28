package interfaces

import (
	"practice_go/entity"
)

type ArticleService interface {
	GetArticle(articleID int) (*entity.Article, error)
	ListArticles() ([]*entity.Article, error)
	InsertArticle(article *entity.Article) error
	PostNice(articleID int) error
}

type CommentService interface {
	InsertComment(comment *entity.Comment) error
}
