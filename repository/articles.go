package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"practice_go/models"
)

func InsertArticle(db *sqlx.DB, article *models.Article) (int, error) {
	result, err := db.Exec("insert into articles (title, contents, username) values (?, ?, ?)", article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	id, _ := result.LastInsertId()

	return int(id), err
}

func GetArticleByArticleID(db *sqlx.DB, articleID int) (*models.Article, error) {
	var article models.Article
	err := db.Get(&article, `select * from articles where id = ?`, &articleID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("no article for that ID")
	}
	if err != nil {
		return nil, err
	}

	return &article, nil

}

func ListArticles(db *sqlx.DB) ([]*models.Article, error) {
	var articles []*models.Article
	err := db.Select(&articles, `select * from articles`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("no article for that ID")
	}
	if err != nil {
		return nil, err
	}

	return articles, nil
}
