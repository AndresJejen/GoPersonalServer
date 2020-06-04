package main

import (
	"encoding/json"
	"net/http"

	"./entities"
)

var (
	posts []entities.Post
)

func init() {
	posts = []entities.Post{entities.Post{ID: 1, Title: "Title 1", Text: "Text 1"}}
}

func getPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	result, err := json.Marshal(posts)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		resp.Write([]byte(`{"error": "Error marshalling the post array"}`))
		return
	}
	resp.WriteHeader(http.StatusOK)
	resp.Write(result)
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
	post.ID = (int64)(len(posts) + 1)
	posts = append(posts, post)
	resp.WriteHeader(http.StatusOK)
	result, err := json.Marshal(post)
	resp.Write(result)
}
