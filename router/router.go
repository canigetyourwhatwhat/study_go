package router

import (
	"github.com/gorilla/mux"
	"github.com/jmoiron/sqlx"
	"net/http"
	"practice_go/api"
	"practice_go/service"
)

func NewRouter(db *sqlx.DB) *mux.Router {

	r := mux.NewRouter()

	ser := service.NewMyAppService(db)
	articleCon := api.NewArticleController(ser)
	commentCon := api.NewCommentController(ser)

	// Post an article
	r.HandleFunc("/article", articleCon.PostArticle).Methods(http.MethodPost)

	// Get specific article based on the article_id
	r.HandleFunc("/article/{id:[0-9]+}", articleCon.GetArticle).Methods(http.MethodGet)

	// List all articles
	r.HandleFunc("/article/all", articleCon.ListArticles).Methods(http.MethodGet)

	// Add nice to the specific article
	r.HandleFunc("/article/nice", articleCon.PostNice).Methods(http.MethodPost)

	// Add comment to the specific article
	r.HandleFunc("/article/comment", commentCon.PostComment).Methods(http.MethodPost)

	return r
}
