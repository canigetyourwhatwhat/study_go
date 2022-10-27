package main

import (
	_ "github.com/go-sql-driver/mysql" // It is the key to connect DB, but don't use it explicitly.
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"practice_go/controllers"
	"practice_go/database"
)

func main() {
	r := mux.NewRouter()

	var err error

	if err := godotenv.Load(".env"); err != nil {
		log.Println(err.Error())
	}

	if err = database.ConnectDB(); err != nil {
		log.Println(err.Error())
	}
	// simply return hello
	r.HandleFunc("/hello", controllers.HelloHandler).Methods(http.MethodGet)

	// Post an article
	r.HandleFunc("/article", controllers.PostArticle).Methods(http.MethodPost)

	// Get specific article based on the article_id
	r.HandleFunc("/article/{id:[0-9]+}", controllers.GetArticle).Methods(http.MethodGet)

	// List all articles
	r.HandleFunc("/article/all", controllers.ListArticles).Methods(http.MethodGet)

	// Add nice to the specific article
	r.HandleFunc("/article/nice", controllers.PostNice).Methods(http.MethodPost)

	// Add comment to the specific article
	r.HandleFunc("/article/comment", controllers.PostComment).Methods(http.MethodPost)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "this is default router")
	})

	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", r))
}
