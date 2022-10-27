package controllers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"practice_go/entity"
	"practice_go/service"
	"strconv"
)

type MyAppController struct {
	service *service.MyAppService
}

func NewMyAppController(s *service.MyAppService) *MyAppController {
	return &MyAppController{service: s}
}

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("hello handler")
	if _, err := io.WriteString(w, "hello"); err != nil {
		http.Error(w, "Can't write", http.StatusInternalServerError)
		return
	}
}

func (c *MyAppController) GetArticle(w http.ResponseWriter, r *http.Request) {

	// Get the ID of the article
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		http.Error(w, "invalid article ID", http.StatusBadRequest)
		return
	}

	// Pass variable to business logic
	result, err := c.service.GetArticle(articleID)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the result by encoding
	if err := json.NewEncoder(w).Encode(result); err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to show on the screen", http.StatusInternalServerError)
	}

}

func (c *MyAppController) ListArticles(w http.ResponseWriter, _ *http.Request) {

	// Call business logic without parameter
	articles, err := c.service.ListArticles()
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to list all articles", http.StatusInternalServerError)
	}

	// Return the result by encoding
	err = json.NewEncoder(w).Encode(articles)
	if err != nil {
		log.Println(err.Error())
		http.Error(w, "Failed to show on the screen", http.StatusInternalServerError)
	}

}

func (c *MyAppController) PostArticle(w http.ResponseWriter, r *http.Request) {
	var article entity.Article

	err := json.NewDecoder(r.Body).Decode(&article)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

	err = c.service.InsertArticle(&article)
	if err != nil {
		http.Error(w, "Failed to post article", http.StatusInternalServerError)
	}
}

func (c *MyAppController) PostNice(w http.ResponseWriter, r *http.Request) {

	// Get the article ID by decoding
	type Req struct {
		ArticleId int `json:"articleId"`
	}
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

	err = c.service.PostNice(req.ArticleId)
	if err != nil {
		http.Error(w, "Failed to add nice on the article", http.StatusInternalServerError)
	}
}

func (c *MyAppController) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment entity.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

	err = c.service.PostComment(&comment)
	if err != nil {
		http.Error(w, "Failed to post comment", http.StatusInternalServerError)
	}
}
