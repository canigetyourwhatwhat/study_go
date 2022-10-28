package api

import (
	"encoding/json"
	"net/http"
	"practice_go/entity"
	"practice_go/interfaces"
)

type CommentController struct {
	service interfaces.CommentService
}

func NewCommentController(s interfaces.CommentService) *CommentController {
	return &CommentController{s}
}

func (s *CommentController) PostComment(w http.ResponseWriter, r *http.Request) {
	var comment entity.Comment

	err := json.NewDecoder(r.Body).Decode(&comment)
	if err != nil {
		http.Error(w, "Failed to decode", http.StatusBadRequest)
	}

	err = s.service.InsertComment(&comment)
	if err != nil {
		http.Error(w, "Failed to post comment", http.StatusInternalServerError)
	}
}
