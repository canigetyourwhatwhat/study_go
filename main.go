package main

import (
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"practice_go/handlers"
)

func main() {
	r := mux.NewRouter()

	r.HandleFunc("/hello", handlers.HelloHandler).Methods(http.MethodGet)
	r.HandleFunc("/article/{category}/{id}", handlers.GetOneArticle).Methods(http.MethodGet)
	r.HandleFunc("/article/list", handlers.ListArticles).Methods(http.MethodGet)
	r.HandleFunc("/article", handlers.PostArticle).Methods(http.MethodPost)
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		io.WriteString(writer, "this is default router")
	})

	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", r))
}
