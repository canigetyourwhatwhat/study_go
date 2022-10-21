package repository

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"log"
	"practice_go/models"
)

func InsertArticle(db *sql.DB, article *models.Article) (int, error) {
	result, err := db.Exec("insert into articles (title, contents, username) values (?, ?, ?)", article.Title, article.Contents, article.UserName)
	if err != nil {
		fmt.Println(err.Error())
		return 0, err
	}
	id, _ := result.LastInsertId()

	return int(id), err
}

func GetArticleByArticleID(db *sql.DB, articleID int) (*models.Article, error) {
	var article models.Article
	row := db.QueryRow(`select * from articles where id = ?`, &articleID)
	if errors.Is(row.Err(), sql.ErrNoRows) {
		return nil, errors.New("no article for that ID")
	}
	if row.Err() != nil {
		return nil, row.Err()
	}

	if err := row.Scan(&article.ID, &article.Title, &article.Contents, &article.UserName, &article.NiceNum, &article.CreatedAt, &article.UpdatedAt); err != nil {
		fmt.Println("here")
		log.Println(err.Error())
	}

	return &article, nil

}
