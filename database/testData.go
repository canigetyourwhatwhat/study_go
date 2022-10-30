package database

import (
	"practice_go/entity"
	"time"
)

var tempTime, _ = time.Parse("2000-01-01 00:00:00", "2000-01-01 00:00:00")

var ArticleTestData = []entity.Article{
	{
		Title:     "firstPost",
		Contents:  "This is my first blog",
		UserName:  "john",
		NiceNum:   2,
		CreatedAt: tempTime,
		UpdatedAt: tempTime,
	},
	{
		Title:     "2nd",
		Contents:  "Second blog post",
		UserName:  "bob",
		NiceNum:   4,
		CreatedAt: tempTime,
		UpdatedAt: tempTime,
	},
	{
		Title:     "For Insertion",
		Contents:  "Third blog post",
		UserName:  "david",
		NiceNum:   3,
		CreatedAt: tempTime,
		UpdatedAt: tempTime,
	},
}
