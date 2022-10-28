package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"practice_go/controllers"
)

func NewRouter(con *controllers.MyAppController) *mux.Router {

	r := mux.NewRouter()

	// simply return hello
	r.HandleFunc("/hello", controllers.HelloHandler).Methods(http.MethodGet)

	// Post an article
	r.HandleFunc("/article", con.PostArticle).Methods(http.MethodPost)

	// Get specific article based on the article_id
	r.HandleFunc("/article/{id:[0-9]+}", con.GetArticle).Methods(http.MethodGet)

	// List all articles
	r.HandleFunc("/article/all", con.ListArticles).Methods(http.MethodGet)

	// Add nice to the specific article
	r.HandleFunc("/article/nice", con.PostNice).Methods(http.MethodPost)

	// Add comment to the specific article
	r.HandleFunc("/article/comment", con.PostComment).Methods(http.MethodPost)

	return r
}
