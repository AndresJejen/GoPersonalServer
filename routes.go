package main

import (
	"encoding/json"
	"math/rand"
	"net/http"

	entities "github.com/AndresJejen/GoPersonalServer/entities"
	repository "github.com/AndresJejen/GoPersonalServer/repository"
)

var (
	repo repository.PostRepository = repository.NewPostRepository()
)

// func init() {
// 	posts = []entities.Post{entities.Post{ID: 1, Title: "Title 1", Text: "Text 1"}}
// }

func getPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := repo.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error getting the posts"}`))
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func addPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entities.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error unmarshalling the request"}`))
		return
	}
	post.ID = rand.Int63()
	repo.Save(&post)
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(post)
}
