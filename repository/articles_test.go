package repository

import (
	"github.com/jmoiron/sqlx"
	"practice_go/database"
	"practice_go/entity"
	"reflect"
	"testing"
)

func TestAddNiceByArticle(t *testing.T) {

	expectedNiceNum := database.ArticleTestData[0].NiceNum
	initialNiceNum := expectedNiceNum
	var actualNiceNum int
	targetArticleId := 1

	type args struct {
		db        *sqlx.DB
		articleId int
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{db: database.DB, articleId: targetArticleId}, false},
		{"2", args{db: database.DB, articleId: targetArticleId}, false},
		{"3", args{db: database.DB, articleId: targetArticleId}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := AddNiceByArticle(tt.args.db, tt.args.articleId); (err != nil) != tt.wantErr {
				t.Errorf("AddNiceByArticle() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := database.DB.Get(&actualNiceNum, "select nice_num from articles where id = ?", targetArticleId); err != nil {
				t.Errorf(err.Error())
			}

			// niceNum should be added by 1
			expectedNiceNum += 1

			if expectedNiceNum != actualNiceNum {
				t.Errorf("Expected: %d,    Got: %d", expectedNiceNum, actualNiceNum)
			}

		})
	}

	t.Cleanup(func() {
		if _, err := database.DB.Exec("update articles set nice_num = ? where id = ?", initialNiceNum, targetArticleId); err != nil {
			t.Errorf(err.Error())
		}
	})

}

func TestGetArticleByArticleID(t *testing.T) {

	type args struct {
		db        *sqlx.DB
		articleID int
	}
	tests := []struct {
		name    string
		args    args
		want    *entity.Article
		wantErr bool
	}{
		{"1", args{db: database.DB, articleID: 2}, &database.ArticleTestData[1], false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := GetArticleByArticleID(tt.args.db, tt.args.articleID)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetArticleByArticleID() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			// Solve the time stamp issue
			got.UpdatedAt = tt.want.UpdatedAt
			got.CreatedAt = tt.want.CreatedAt

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetArticleByArticleID() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestInsertArticle(t *testing.T) {

	actualArticle := database.ArticleTestData[2]
	var expectedArticle entity.Article

	type args struct {
		db      *sqlx.DB
		article *entity.Article
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{"1", args{db: database.DB, article: &actualArticle}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := InsertArticle(tt.args.db, tt.args.article); (err != nil) != tt.wantErr {
				t.Errorf("InsertArticle() error = %v, wantErr %v", err, tt.wantErr)
			}

			if err := database.DB.Get(&expectedArticle, "select * from articles where id = ?", actualArticle.ID); err != nil {
				t.Errorf(err.Error())
			}

			// Solve the time stamp issue
			expectedArticle.UpdatedAt = actualArticle.UpdatedAt
			expectedArticle.CreatedAt = actualArticle.CreatedAt

			if !reflect.DeepEqual(expectedArticle, actualArticle) {
				t.Errorf("GetArticleByArticleID() got = %v, want %v", actualArticle, expectedArticle)
			}

		})
	}

	t.Cleanup(func() {
		if _, err := database.DB.Exec("delete from articles where id = ?", actualArticle.ID); err != nil {
			t.Errorf(err.Error())
		}
	})
}

func TestListArticles(t *testing.T) {

	articles := []*entity.Article{&database.ArticleTestData[0], &database.ArticleTestData[1]}

	type args struct {
		db *sqlx.DB
	}
	tests := []struct {
		name    string
		args    args
		want    []*entity.Article
		wantErr bool
	}{
		{"1", args{database.DB}, articles, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ListArticles(tt.args.db)

			// resolve time issues
			for i := range tt.want {
				got[i].CreatedAt = tt.want[i].CreatedAt
				got[i].UpdatedAt = tt.want[i].UpdatedAt
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("ListArticles() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ListArticles() got = %v, want %v", got, tt.want)
			}
		})
	}
}
