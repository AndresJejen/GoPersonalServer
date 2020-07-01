package repository

import entities "github.com/AndresJejen/GoPersonalServer/entities"

// PostRepository Interface del repositorio
type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}
