package repository

import (
	"context"
	"errors"
	"go-auth/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	GetUserByID(id string) (model.User, error)
	GetAllUsers() ([]model.User, error)
	GetAllUsersByRole(roles []string) ([]model.User, error)
	UpdateUser(user model.User) (model.User, error)
}

type userRepo struct {
	db *mongo.Database
	indexRepo *IndexRepo
}

func NewUserRepo(db *mongo.Database, indexRepo *IndexRepo) *userRepo {
	err := indexRepo.CreateIndex("users", "id", true)
	if err != nil {
		logger.Errorf("error creating index: %v\n", err)
	}
	return &userRepo{
		db: db,
		indexRepo: indexRepo,
	}
}

func (r *userRepo) GetUserByID(id string) (model.User, error) {
	collection := r.db.Collection("users")
	var user model.User
	err := collection.FindOne(context.Background(), bson.M{"id": id}).Decode(&user)
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return user, nil 
		}
		return model.User{}, err 
	}
	return user, nil
}

func (r *userRepo) GetAllUsers() ([]model.User, error) {
    var users []model.User
    cursor, err := r.db.Collection("users").Find(context.Background(), bson.M{})
    if err != nil {
        return nil, errors.New("failed to fetch users")
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var user model.User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (r *userRepo) GetAllUsersByRole(roles []string) ([]model.User, error) {
	var users []model.User
    filter := bson.M{"role": bson.M{"$in": roles}}
    cursor, err := r.db.Collection("users").Find(context.Background(), filter)
    if err != nil {
        return nil, errors.New("failed to fetch users by role")
    }
    defer cursor.Close(context.Background())

    for cursor.Next(context.Background()) {
        var user model.User
        if err := cursor.Decode(&user); err != nil {
            return nil, err
        }
        users = append(users, user)
    }

    if err := cursor.Err(); err != nil {
        return nil, err
    }

    return users, nil
}

func (r *userRepo) UpdateUser(user model.User) (model.User, error) {
	filter := bson.M{"id": user.Id}
	update := bson.M{"$set": user}


	collection := r.db.Collection("users")
	_, err := collection.UpdateOne(context.Background(), filter,update)
	if err != nil {
		return model.User{}, err
	}
	return user, nil
}