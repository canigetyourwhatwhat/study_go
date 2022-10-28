package main

import (
	"github.com/joho/godotenv"
	"io"
	"log"
	"net/http"
	"practice_go/controllers"
	"practice_go/database"
	"practice_go/router"
	"practice_go/service"
)

func main() {
	var err error

	if err := godotenv.Load(".env"); err != nil {
		log.Println(err.Error())
	}

	if err = database.ConnectDB(); err != nil {
		log.Println(err.Error())
	}

	ser := service.NewMyAppService(database.DB)
	con := controllers.NewMyAppController(ser)

	r := router.NewRouter(con)

	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		_, _ = io.WriteString(writer, "this is default router")
	})

	log.Println("server started")
	log.Fatal(http.ListenAndServe(":8080", r))
}
