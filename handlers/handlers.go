package handlers

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"practice_go/models"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("hello handler")
	if _, err := io.WriteString(w, "hello"); err != nil {
		http.Error(w, "Can't write", http.StatusInternalServerError)
		return
	}
}

func GetOneArticle(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	categoryName := mux.Vars(r)["category"]
	if err != nil || categoryName == "" {
		http.Error(w, "invalid article ID", http.StatusBadRequest)
		return
	}
	resStr := fmt.Sprintf("This articel is category of %s and number is %d", categoryName, articleID)
	if _, err = io.WriteString(w, resStr); err != nil {
		http.Error(w, "Can't write", http.StatusInternalServerError)
		return
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
	if err := json.NewDecoder(r.Body).Decode(&article); err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}
	log.Println(article)

	if err := json.NewEncoder(w).Encode(article); err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

}
