package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"practice_go/database"
	"practice_go/models"
	"practice_go/repository"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("hello handler")
	if _, err := io.WriteString(w, "hello"); err != nil {
		http.Error(w, "Can't write", http.StatusInternalServerError)
		return
	}
}

func GetArticle(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid article ID", http.StatusBadRequest)
		return
	}

	var article *models.Article
	article, err = repository.GetArticleByArticleID(database.DB, articleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	if err := json.NewEncoder(w).Encode(article); err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to show on the screen", http.StatusInternalServerError)
	}

}

func ListArticles(w http.ResponseWriter, r *http.Request) {
	var articles []*models.Article
	var err error

	articles, err = repository.ListArticles(database.DB)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to list all articles", http.StatusInternalServerError)
	}

	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to show on the screen", http.StatusInternalServerError)
	}

}

func PostArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article

	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

	err = repository.InsertArticle(database.DB, &article)
	if err != nil {
		http.Error(w, "Failed to post article", http.StatusInternalServerError)
	}
}

func PostNice(w http.ResponseWriter, r *http.Request) {

	type Req struct {
		ArticleId int `json:"articleId"`
	}

	var req Req

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

	err = repository.AddNiceByArticle(database.DB, req.ArticleId)
	if err != nil {
		http.Error(w, "Failed to add nice on the article", http.StatusInternalServerError)
	}
}

func PostComment(w http.ResponseWriter, r *http.Request) {
	var comment models.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

	err = repository.InsertComment(database.DB, &comment)
	if err != nil {
		http.Error(w, "Failed to post comment", http.StatusInternalServerError)
	}
}
