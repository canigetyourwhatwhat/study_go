package repository

import (
	"fmt"
	"github.com/jmoiron/sqlx"
	"practice_go/models"
)

func InsertComment(db *sqlx.DB, comment *models.Comment) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into comments (article_id, message) values (?, ?)", comment.ArticleID, comment.Message)
	if err != nil {
		fmt.Println(err.Error())
		tx.Rollback()
		return err
	}

	if err = tx.Commit(); err != nil {
		tx.Rollback()
		return err
	}

	return nil
}
