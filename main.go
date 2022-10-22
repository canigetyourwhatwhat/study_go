package main

import (
	_ "github.com/go-sql-driver/mysql" // It is the key to connect DB, but don't use it explicitly.
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"practice_go/database"
	"practice_go/handlers"
)

func main() {
	var err error
	r := mux.NewRouter()
	if err = database.ConnectDB(); err != nil {
		panic(err)
	}

	// simply return hello
	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)

	// Post an article
	r.HandleFunc("/article", handlers.PostArticle).Methods(http.MethodPost)

	// Get specific article based on the article_id
	r.HandleFunc("/article/{id:[0-9]+}", handlers.GetArticle).Methods(http.MethodGet)

	// List all articles
	r.HandleFunc("/article/all", handlers.ListArticles).Methods(http.MethodGet)

	// Add nice to the specific article
	r.HandleFunc("/article/nice", handlers.PostNice).Methods(http.MethodPost)

	// Add comment to the specific article
	//r.HandleFunc("article/nice", handlers.PostComment).Methods(http.MethodPost)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "this is default router")
	})

	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", r))
}
