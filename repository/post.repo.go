package repository

import (
	"../entities"
)

type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}
