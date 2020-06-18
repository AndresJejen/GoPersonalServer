package repository

import (
	"context"
	"log"

	firebase "firebase.google.com/go"
	entities "github.com/AndresJejen/GoPersonalServer/entities"
	"google.golang.org/api/iterator"
)

// PostRepository Interface del repositorio
type PostRepository interface {
	Save(post *entities.Post) (*entities.Post, error)
	FindAll() ([]entities.Post, error)
}

type repo struct{}

// NewPostRepository Comentario
func NewPostRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "GoServer"
	collectionName string = "post"
)

func (*repo) Save(post *entities.Post) (*entities.Post, error) {
	ctx := context.Background()
	// conf := &firebase.Config{ProjectID: projectID}
	// opt := option.WithCredentialsFile("github.com/AndresJejen/GoPersonalServer/FireBaseCredentialsFile.json")
	app, err := firebase.NewApp(ctx, nil)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()
	_, _, err = client.Collection(collectionName).Add(ctx, map[string]interface{}{
		"ID":    post.ID,
		"Title": post.Title,
		"Text":  post.Text,
	})

	if err != nil {
		log.Fatalf("Failed to create Data: %v", err)
		return nil, err
	}

	return post, nil

}

func (*repo) FindAll() ([]entities.Post, error) {
	ctx := context.Background()
	// conf := &firebase.Config{ProjectID: projectID}
	// opt := option.WithCredentialsFile("../FireBaseCredentialsFile.json")
	app, err := firebase.NewApp(ctx, nil)

	if err != nil {
		log.Fatalln(err)
	}

	client, err := app.Firestore(ctx)

	if err != nil {
		log.Fatalf("Failed to create a Firestore Client: %v", err)
		return nil, err
	}
	defer client.Close()

	var posts []entities.Post
	iter := client.Collection(collectionName).Documents(ctx)

	for {
		doc, err := iter.Next()
		if err == iterator.Done {
			break
		}
		if err != nil {
			log.Fatalf("Failed to iterate the list of posts: %v", err)
			return nil, err
		}
		post := entities.Post{
			ID:    doc.Data()["ID"].(int64),
			Title: doc.Data()["Title"].(string),
			Text:  doc.Data()["Text"].(string),
		}
		posts = append(posts, post)
	}
	return posts, nil
}
