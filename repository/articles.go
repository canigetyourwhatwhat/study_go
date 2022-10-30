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
	var articleData []entity.ArticleComment
	err := db.Select(&articleData,
		`select
    		a.title, 
    		a.contents, 
    		a.username, 
    		a.nice_num, 
    		a.created_at as article_created_at, 
    		a.updated_at as article_updated_at,
    		c.message, 
    		c.created_at as comment_create_at
		from articles as a inner join comments c on a.id = c.article_id where a.id = ?`, &articleID)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("no article for that ID")
	}
	if err != nil {
		return nil, err
	}

	var comments []entity.Comment
	for _, row := range articleData {
		comment := entity.Comment{
			Message:   row.Message,
			CreatedAt: row.CommentCreatedAt}
		comments = append(comments, comment)
	}

	article := entity.Article{
		Title:     articleData[0].Title,
		Contents:  articleData[0].Contents,
		UserName:  articleData[0].UserName,
		NiceNum:   articleData[0].NiceNum,
		Comments:  comments,
		CreatedAt: articleData[0].ArticleCreatedAt,
		UpdatedAt: articleData[0].ArticleUpdatedAt,
	}

	return &article, nil
}

func ListArticles(db *sqlx.DB) ([]*entity.Article, error) {
	var articlesData []entity.ArticleComment
	err := db.Select(&articlesData,
		`select
    		a.id as article_id,
    		a.title, 
    		a.contents, 
    		a.username, 
    		a.nice_num, 
    		a.created_at as article_created_at, 
    		a.updated_at as article_updated_at,
    		c.article_id as article_id_from_comment,
    		c.message, 
    		c.created_at as comment_create_at
		from articles as a inner join comments c on a.id = c.article_id order by a.id`)
	if errors.Is(err, sql.ErrNoRows) {
		return nil, errors.New("no article for that ID")
	}
	if err != nil {
		return nil, err
	}

	// Track the current article index
	currentArticleID := 1

	articles := []*entity.Article{{
		Title:     articlesData[0].Title,
		Contents:  articlesData[0].Contents,
		UserName:  articlesData[0].UserName,
		NiceNum:   articlesData[0].NiceNum,
		Comments:  nil,
		CreatedAt: articlesData[0].ArticleCreatedAt,
		UpdatedAt: articlesData[0].ArticleUpdatedAt,
	}}

	for _, row := range articlesData {

		comment := entity.Comment{
			Message:   row.Message,
			CreatedAt: row.CommentCreatedAt}

		if currentArticleID != row.ArticleIDFromComment {
			article := &entity.Article{
				Title:     row.Title,
				Contents:  row.Contents,
				UserName:  row.UserName,
				NiceNum:   row.NiceNum,
				Comments:  []entity.Comment{comment},
				CreatedAt: row.ArticleCreatedAt,
				UpdatedAt: row.ArticleUpdatedAt,
			}
			articles = append(articles, article)
			currentArticleID++
		} else {
			articles[currentArticleID-1].Comments = append(articles[currentArticleID-1].Comments, comment)
		}

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
