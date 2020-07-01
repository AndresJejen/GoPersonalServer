package repository

import (
	"context"
	"log"

	firestore "cloud.google.com/go/firestore"
	entities "github.com/AndresJejen/GoPersonalServer/entities"
	"google.golang.org/api/iterator"
	"google.golang.org/api/option"
)

type repo struct{}

// NewFirestoreRepository Comentario
func NewFirestoreRepository() PostRepository {
	return &repo{}
}

const (
	projectID      string = "goserver-1dae4"
	collectionName string = "post"
)

func (*repo) Save(post *entities.Post) (*entities.Post, error) {
	ctx := context.Background()

	client, err := firestore.NewClient(ctx, projectID)

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

	var options []option.ClientOption
	options = append(options, option.WithCredentialsFile("d:\\go\\src\\github.com\\AndresJejen\\GoPersonalServer\\FireBaseCredentialsFile.json"))

	client, err := firestore.NewClient(ctx, projectID, options...)

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
