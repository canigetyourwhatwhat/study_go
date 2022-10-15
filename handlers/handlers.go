package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gorilla/mux"
	"io"
	"log"
	"net/http"
	"practice_go/models"
	"strconv"
)

func HelloHandler(w http.ResponseWriter, _ *http.Request) {
	log.Println("hello handler")
	if _, err := io.WriteString(w, "hello"); err != nil {
		http.Error(w, "Can't write", http.StatusInternalServerError)
		return
	}
}

func GetOneArticle(w http.ResponseWriter, r *http.Request) {
	articleID, err := strconv.Atoi(mux.Vars(r)["id"])
	categoryName := mux.Vars(r)["category"]
	if err != nil || categoryName == "" {
		http.Error(w, "invalid article ID", http.StatusBadRequest)
		return
	}
	resStr := fmt.Sprintf("This articel is category of %s and number is %d", categoryName, articleID)
	if _, err = io.WriteString(w, resStr); err != nil {
		http.Error(w, "Can't write", http.StatusInternalServerError)
		return
	}
}

func ListArticles(w http.ResponseWriter, r *http.Request) {
	queryParam := r.URL.Query()
	var page int

	if query, ok := queryParam["page"]; ok == true && len(query) > 0 {
		num, err := strconv.Atoi(query[0])
		if err != nil {
			http.Error(w, "Invalid input for page", http.StatusBadRequest)
			return
		}
		page = num
	} else {
		page = 1
	}
	resStr := fmt.Sprintf("This is page %d", page)
	if _, err := io.WriteString(w, resStr); err != nil {
		http.Error(w, "Can't write", http.StatusInternalServerError)
		return
	}
}

func PostArticle(w http.ResponseWriter, r *http.Request) {

	// ------ Read the Body in Request Header   -----
	// Get Content-Length to make a slice
	length, err := strconv.Atoi(r.Header.Get("Content-Length"))
	if err != nil {
		http.Error(w, "Invalid value in Content-Length", http.StatusBadRequest)
		return
	}
	reqBodyBuffer := make([]byte, length)
	// Read the Body info and store in byte slice
	if _, err = r.Body.Read(reqBodyBuffer); !errors.Is(err, io.EOF) {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}
	defer r.Body.Close()
	log.Println("Retrieved body: ", reqBodyBuffer)

	// ------ Unmarshal the data  -----
	var reqArticle models.Article
	if err = json.Unmarshal(reqBodyBuffer, &reqArticle); err != nil {
		http.Error(w, "Failed to unmarshal", http.StatusBadRequest)
		return
	}
	log.Println("Unmarshalled / decoded body: ", reqArticle)

	// ------ Marshal the data and return it  -----
	jsonData, err := json.Marshal(reqArticle)
	if err != nil {
		http.Error(w, "Invalid json format", http.StatusInternalServerError)
		return
	}
	w.Write(jsonData)
	log.Println("Marshalled / encoded body: ", reqArticle)
}
