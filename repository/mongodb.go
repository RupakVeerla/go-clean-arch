package repository

import (
	"context"
	"errors"
	"fmt"
	"go-demo/entity"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	collection *mongo.Collection
	ctx        context.Context
)

func init() {
	client, err := mongo.NewClient(options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	collection = client.Database("userdb").Collection("user")
	collection.DeleteMany(context.TODO(), bson.D{{}})
}

type mongoDB struct{}

func NewMongoDBRepo() Repository {
	return &mongoDB{}
}

func (*mongoDB) Post(user *entity.User) error {
	u := collection.FindOne(ctx, bson.D{{"id", user.ID}})
	if u.Err() == nil {
		return fmt.Errorf("user already exists with ID %d", user.ID)
	}
	_, err := collection.InsertOne(ctx, user)
	if err != nil {
		return err
	}
	return nil
}

func (*mongoDB) Get() (users []*entity.User, err error) {
	cur, err := collection.Find(ctx, bson.D{})
	if err != nil {
		return nil, err
	}
	err = cur.All(ctx, &users)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	if len(users) == 0 {
		return nil, errors.New("no user data")
	}
	return users, nil
}
