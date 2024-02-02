package userrepo

import (
	"BackendCoursyclopedia/db"
	"BackendCoursyclopedia/model"
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepository interface {
	FindAllUsers(ctx context.Context) ([]model.User, error)
	FindUserByID(ctx context.Context, userID string) (*model.User, error)
	CreateUser(ctx context.Context, user model.User) (*model.User, error)
	DeleteUserByID(ctx context.Context, userID string) error
	UpdateUserByID(ctx context.Context, userID string, updateUser model.User) (*model.User, error)
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

func (r *UserRepository) FindUserByID(ctx context.Context, userID string) (*model.User, error) {
	collection := db.GetCollection("users")
	var user model.User

	// Convert string to ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	err = collection.FindOne(ctx, bson.M{"_id": objID}).Decode(&user)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) CreateUser(ctx context.Context, user model.User) (*model.User, error) {
	collection := db.GetCollection("users")
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return nil, err
	}

	user.ID = result.InsertedID.(primitive.ObjectID)

	return &user, nil
}

func (r *UserRepository) DeleteUserByID(ctx context.Context, userID string) error {
	collection := db.GetCollection("users")

	// Convert string to ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	_, err = collection.DeleteOne(ctx, bson.M{"_id": objID})
	if err != nil {
		return err
	}

	return nil
}

func (r *UserRepository) UpdateUserByID(ctx context.Context, userID string, updateUser model.User) (*model.User, error) {
	collection := db.GetCollection("users")

	// Convert string to ObjectID
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return nil, err
	}

	// Create an update document with the fields to be updated
	update := bson.M{
		"$set": updateUser,
	}

	result, err := collection.UpdateOne(ctx, bson.M{"_id": objID}, update)
	if err != nil {
		return nil, err
	}

	if result.MatchedCount == 0 {
		return nil, errors.New("no user found with the given ID")
	}

	return &updateUser, nil
}
