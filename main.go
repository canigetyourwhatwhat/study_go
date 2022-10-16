package main

import (
	// It doesn't connect to DB.
	// It enables abstraction since we just need to change driver name and the string to connect DB.
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // It is the key to connect DB, but don't use it explicitly.
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"practice_go/handlers"
)

func main() {
	r := mux.NewRouter()

	db, err := connectDB()
	if err != nil {
		err.Error()
	}

	db.Ping()

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

func connectDB() (*sql.DB, error) {
	//dbUser := os.Getenv("USER_NAME")
	//dbPass := os.Getenv("USER_PASS")
	//dbName := os.Getenv("DATABASE")
	dbUser := "docker"
	dbPass := "docker"
	dbName := "sampledb"
	connectDbStr := fmt.Sprintf("%s:%s@tcp(127.0.0.1:3306)/%s?parseTime=true", dbUser, dbPass, dbName)

	db, err := sql.Open("mysql", connectDbStr)
	if err != nil {
		fmt.Println("Couldn't connect the Database")
		return nil, err
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		fmt.Println("Ping to DB is failed")
		return nil, err
	} else {
		log.Println("DB is healthy")
	}
	return db, nil
}
