package repository

import (
	"context"
	"go-auth/domain/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IUserRepo interface {
	GetUserByID(id string) (model.User, error)
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