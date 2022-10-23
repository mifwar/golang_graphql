package database

import (
	"context"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"

	"graphql_example/graph/model"
)

type DB struct {
	client *mongo.Client
}

func Connect() *DB {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://localhost:27017"))

	if err != nil {
		log.Fatalf("error at : %v", err)
	}

	return &DB{
		client: client,
	}
}

func (db *DB) Save(input *model.NewUser) *model.User {
	collection := db.client.Database("people").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	res, err := collection.InsertOne(ctx, input)

	if err != nil {
		log.Fatalf("error at: %v", err)
	}

	return &model.User{
		ID:    res.InsertedID.(primitive.ObjectID).Hex(),
		Name:  input.Name,
		Class: input.Class,
	}
}

func (db *DB) FindById(ID string) *model.User {
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	collection := db.client.Database("people").Collection("users")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	user := model.User{}
	err = collection.FindOne(ctx, bson.M{"_id": ObjectID}).Decode(&user)

	if err != nil {
		log.Fatalf("error at: %v", err)
	}

	return &user
}

func (db *DB) GetAll() []*model.User {
	collection := db.client.Database("people").Collection("users")

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	defer cur.Close(ctx)

	var Users []*model.User

	for cur.Next(ctx) {
		var user *model.User
		err := cur.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}

		Users = append(Users, user)
	}
	return Users
}
