package userrepo

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/model"
	"context"

	"go.mongodb.org/mongo-driver/bson"
	//"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	FindAllUsers(ctx context.Context) ([]model.User, error)
}

type UserRepository struct {
	DB *mongo.Client
}

func NewUserRepository(db *mongo.Client) IUserRepository {
	return &UserRepository{
		DB: db,
	}
}

func (r *UserRepository) FindAllUsers(ctx context.Context) ([]model.User, error) {
	collection := db.GetCollection("users")

	var users []model.User
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}