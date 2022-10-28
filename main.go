package main

import (
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"practice_go/database"
	"practice_go/router"
)

func main() {
	var err error
	var db *sqlx.DB

	if err := godotenv.Load(".env"); err != nil {
		log.Println(err.Error())
	}

	if db, err = database.ConnectDB(); err != nil {
		log.Println(err.Error())
	}

	r := router.NewRouter(db)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "this is default router")
	})

	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", r))
}
