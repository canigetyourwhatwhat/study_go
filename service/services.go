package service

import (
	"github.com/jmoiron/sqlx"
	"log"
	"practice_go/customErrors"
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
	article, err := repository.GetArticleByArticleID(s.db, articleID)
	if err != nil {
		errCode := customErrors.InsertDataFailed.Wrap(err, "fail to get data")
		return nil, errCode
	}
	return article, nil
}

func (s *MyAppService) ListArticles() ([]*entity.Article, error) {
	articles, err := repository.ListArticles(s.db)
	if err != nil {
		log.Println(err.Error())
		return nil, err
	}
	return articles, nil
}

func (s *MyAppService) InsertArticle(article *entity.Article) error {
	err := repository.InsertArticle(s.db, article)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s *MyAppService) PostNice(articleID int) error {
	err := repository.AddNiceByArticle(s.db, articleID)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}

func (s *MyAppService) InsertComment(comment *entity.Comment) error {
	err := repository.InsertComment(s.db, comment)
	if err != nil {
		log.Println(err.Error())
		return err
	}
	return nil
}
