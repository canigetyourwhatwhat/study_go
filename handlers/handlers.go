package handlers

import (
	"encoding/json"
	"fmt"
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
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

}

func ListArticles(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	var page int

	if query, ok := queryParam["page"]; ok == true && len(query) > 0 {
		num, err := strconv.Atoi(query[0])
		if err != nil {
			http.Error(w, "Invalid input for page", http.StatusBadRequest)
			return
		}
		page = num
	} else {
		page = 1
	}
	resStr := fmt.Sprintf("This is page %d", page)
	if _, err := io.WriteString(w, resStr); err != nil {
		http.Error(w, "Can't write", http.StatusInternalServerError)
		return
	}
}

func PostArticle(w http.ResponseWriter, r *http.Request) {
	var article models.Article

	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}
	log.Println("inserted article: ", article)

	_, err = repository.InsertArticle(database.DB, &article)
	if err != nil {
		http.Error(w, "Failed to post article", http.StatusInternalServerError)
	}
}
