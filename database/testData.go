package database

import (
	"practice_go/entity"
	"time"
)

var tempTime, _ = time.Parse("2000-01-01 00:00:00", "2000-01-01 00:00:00")

var ArticleTestData = []entity.Article{
	{
		ID:        1,
		Title:     "firstPost",
		Contents:  "This is my first blog",
		UserName:  "john",
		NiceNum:   2,
		CreatedAt: tempTime,
		UpdatedAt: tempTime,
	},
	{
		ID:        2,
		Title:     "2nd",
		Contents:  "Second blog post",
		UserName:  "bob",
		NiceNum:   4,
		CreatedAt: tempTime,
		UpdatedAt: tempTime,
	},
	{
		ID:        3,
		Title:     "For Insertion",
		Contents:  "Third blog post",
		UserName:  "david",
		NiceNum:   3,
		CreatedAt: tempTime,
		UpdatedAt: tempTime,
	},
}
