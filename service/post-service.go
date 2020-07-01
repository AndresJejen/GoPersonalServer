package service

import (
	"errors"
	"math/rand"

	"github.com/AndresJejen/GoPersonalServer/entities"
	"github.com/AndresJejen/GoPersonalServer/repository"
)

type PostService interface {
	Validate(post *entities.Post) error
	Create(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}

type service struct{}

var (
	repo repository.PostRepository = repository.NewFirestoreRepository()
)

func NewPostService() PostService {
	return &service{}
}

func (*service) Validate(post *entities.Post) error {
	if post == nil {
		err := errors.New("The Post is empty")
		return err
	}
	if post.Title == "" {
		err := errors.New("The Post Titile is empty")
		return err
	}
	return nil
}

func (*service) Create(post *entities.Post) (*entities.Post, error) {
	post.ID = rand.Int63()
	return repo.Save(post)
}

func (*service) FindAll() ([]entities.Post, error) {
	return repo.FindAll()
}
