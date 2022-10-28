package repository

import (
	"database/sql"
	"errors"
	_ "github.com/go-sql-driver/mysql" // It is the key to connect DB, but don't use it explicitly.
	"github.com/jmoiron/sqlx"
	"practice_go/entity"
)

func InsertArticle(db *sqlx.DB, article *entity.Article) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("insert into articles (title, contents, username, nice_num) values (?, ?, ?, ?)", article.Title, article.Contents, article.UserName, article.NiceNum)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	if err = tx.Commit(); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return err
}

func GetArticleByArticleID(db *sqlx.DB, articleID int) (*entity.Article, error) {
	var article entity.Article
	err := db.Get(&article, `select * from articles where id = ?`, &articleID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("no article for that ID")
	}
	if err != nil {
		return nil, err
	}

	return &article, nil
}

func ListArticles(db *sqlx.DB) ([]*entity.Article, error) {
	var articles []*entity.Article
	err := db.Select(&articles, `select * from articles`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("no article for that ID")
	}
	if err != nil {
		return nil, err
	}

	return articles, nil
}

func AddNiceByArticle(db *sqlx.DB, articleId int) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	var targetNiceNum int
	err = db.Get(&targetNiceNum, `select nice_num from articles where id = ?`, &articleId)
	if errors.Is(err, sql.ErrNoRows) {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return errors.New("no article for that ID")
	} else if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	_, err = tx.Exec(`update articles set nice_num = ? where id = ?`, targetNiceNum+1, articleId)
	if err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	// Always good to use tx since I can roll back wherever I got an error.
	if err = tx.Commit(); err != nil {
		if err = tx.Rollback(); err != nil {
			return err
		}
		return err
	}

	return nil

}
