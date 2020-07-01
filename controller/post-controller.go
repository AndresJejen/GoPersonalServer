package controller

import (
	"encoding/json"
	"net/http"

	"github.com/AndresJejen/GoPersonalServer/entities"
	errors "github.com/AndresJejen/GoPersonalServer/errors"
	service "github.com/AndresJejen/GoPersonalServer/service"
)

type controller struct{}

var (
	postService service.PostService = service.NewPostService()
)

type PostController interface {
	GetPost(resp http.ResponseWriter, req *http.Request)
	AddPost(resp http.ResponseWriter, req *http.Request)
}

func NewPostController() PostController {
	return &controller{}
}

func (*controller) GetPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	posts, err := postService.FindAll()
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error getting the posts"})
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(posts)
}

func (*controller) AddPost(resp http.ResponseWriter, req *http.Request) {
	resp.Header().Set("Content-type", "application/json")
	var post entities.Post
	err := json.NewDecoder(req.Body).Decode(&post)
	if err != nil {
		resp.WriteHeader(http.StatusInternalServerError)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error unmarshalling the request"})
		return
	}
	err1 := postService.Validate(&post)
	if err1 != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: err1.Error()})
		return
	}
	result, err2 := postService.Create(&post)
	if err2 != nil {
		resp.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(resp).Encode(errors.ServiceError{Message: "Error saving the post"})
		return
	}
	resp.WriteHeader(http.StatusOK)
	json.NewEncoder(resp).Encode(result)
}
