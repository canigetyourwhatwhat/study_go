package api

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"practice_go/customErrors"
	"practice_go/entity"
	"practice_go/interfaces"
	"strconv"
)

type ArticleController struct {
	service interfaces.ArticleService
}

func NewArticleController(s interfaces.ArticleService) *ArticleController {
	return &ArticleController{service: s}
}

func (c *ArticleController) GetArticle(w http.ResponseWriter, r *http.Request) {

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

func (c *ArticleController) ListArticles(w http.ResponseWriter, _ *http.Request) {

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

func (c *ArticleController) PostArticle(w http.ResponseWriter, r *http.Request) {
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

func (c *ArticleController) PostNice(w http.ResponseWriter, r *http.Request) {

	// Get the article ID by decoding
	type Req struct {
		ArticleId int `json:"articleId"`
	}
	var req Req
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		err = customErrors.ReqBodyDecodeFailed.Wrap(err, "bad request body")
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

	err = c.service.PostNice(req.ArticleId)
	if err != nil {
		http.Error(w, "Failed to add nice on the article", http.StatusInternalServerError)
	}
}
