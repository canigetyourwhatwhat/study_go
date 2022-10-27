package service

import (
	"github.com/jmoiron/sqlx"
	"log"
	"practice_go/database"
	"practice_go/entity"
	"practice_go/repository"
)

type MyAppService struct {
	db *sqlx.DB
}

func NewMyAppService(db *sqlx.DB) *MyAppService {
	return &MyAppService{db: db}
}

func (s *MyAppService) GetArticle(articleID int) (*entity.Article, error) {
	article, err := repository.GetArticleByArticleID(database.DB, articleID)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return article, nil
}

func (s *MyAppService) ListArticles() ([]*entity.Article, error) {
	articles, err := repository.ListArticles(database.DB)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return articles, nil
}

func (s *MyAppService) InsertArticle(article *entity.Article) error {
	err := repository.InsertArticle(database.DB, article)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s *MyAppService) PostNice(articleID int) error {
	err := repository.AddNiceByArticle(database.DB, articleID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s *MyAppService) PostComment(comment *entity.Comment) error {
	err := repository.InsertComment(database.DB, comment)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
