package service

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"

	"github.com/saifuljnu/crudApp/models"
)

type UserService interface {
	Insert(data interface{}) error
	GetAll() ([]models.User, error)
	DeleteData(primitive.ObjectID) error
	UpdateData(models.User) error
}

type userService struct {
	db *mongo.Client
}

func NewUserService(db *mongo.Client) UserService {
	return &userService{
		db: db,
	}
}

// getCollection returns the MongoDB collection for "collectionUser"
func (s *userService) getCollection() *mongo.Collection {
	return s.db.Database("crudDB").Collection("collectionUser")
}

func (s *userService) Insert(data interface{}) error {
	orgCollection := s.getCollection()
	_, err := orgCollection.InsertOne(context.Background(), data)
	return err
}

func (s *userService) GetAll() ([]models.User, error) {
	orgCollection := s.getCollection()

	cur, err := orgCollection.Find(context.Background(), bson.M{})
	if err != nil {
		return nil, err
	}

	defer cur.Close(context.Background())

	var data []models.User
	for cur.Next(context.Background()) {
		var d models.User
		err := cur.Decode(&d)
		if err != nil {
			return nil, err
		}
		data = append(data, d)
	}

	if err := cur.Err(); err != nil {
		return nil, err
	}

	return data, nil
}

func (s *userService) DeleteData(id primitive.ObjectID) error {
	orgCollection := s.getCollection()
	filter := bson.D{{"_id", id}}
	_, err := orgCollection.DeleteOne(context.Background(), filter)
	return err
}

func (s *userService) UpdateData(data models.User) error {
	orgCollection := s.getCollection()
	filter := bson.D{{"_id", data.ID}}
	update := bson.D{{"$set", data}}
	_, err := orgCollection.UpdateOne(context.Background(), filter, update)
	return err
}
