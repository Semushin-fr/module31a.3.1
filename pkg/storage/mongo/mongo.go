package mongo

import (
	"GoNews/pkg/storage"
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Store struct {
	client *mongo.Client
	db     *mongo.Database
}

func New(uri string, dbName string) (*Store, error) {
	clientOptions := options.Client().ApplyURI(uri)

	client, err := mongo.Connect(context.Background(), clientOptions)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	err = client.Ping(context.Background(), nil)

	if err != nil {
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database(dbName)

	return &Store{
		client: client,
		db:     db,
	}, nil
}

func (s *Store) Posts() ([]storage.Post, error) {
	collection := s.db.Collection("posts")

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return nil, fmt.Errorf("failed to find posts: %w", err)
	}

	defer cursor.Close(context.Background())

	var posts []storage.Post

	if err = cursor.All(context.Background(), &posts); err != nil {
		return nil, fmt.Errorf("failed to decode posts: %w", err)
	}

	return posts, nil
}

func (s *Store) AddPost(post storage.Post) error {
	collection := s.db.Collection("posts")

	_, err := collection.InsertOne(context.Background(), post)
	if err != nil {
		return fmt.Errorf("failed to insert post: %w", err)
	}

	return nil
}

func (s *Store) UpdatePost(post storage.Post) error {
	collection := s.db.Collection("posts")

	filter := bson.M{"id": post.ID}

	update := bson.M{
		"$set": bson.M{
			"title":      post.Title,
			"content":    post.Content,
			"author_id":  post.AuthorID,
			"created_at": post.CreatedAt,
		},
	}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return fmt.Errorf("failed to update post: %w", err)
	}

	if result.MatchedCount == 0 {
		return fmt.Errorf("post with ID %d not found", post.ID)
	}

	return nil
}

func (s *Store) DeletePost(post storage.Post) error {
	collection := s.db.Collection("posts")

	filter := bson.M{"id": post.ID}

	result, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		return fmt.Errorf("failed to delete post: %w", err)
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("post with ID %d not found", post.ID)
	}

	return nil
}
