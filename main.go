package main

import (
	"context"
	"fmt"
	"github.com/nmorales1991/go-first-api/bd"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func main() {
	coll := bd.MongoClient.Database("sample_mflix").Collection("movies")
	title := "Back to the Future"
	var result bson.M
	err := coll.FindOne(context.TODO(), bson.D{{"title", title}}).Decode(&result)
	if err == mongo.ErrNoDocuments {
		fmt.Printf("No document was found with the title %s\n", title)
		return
	}
	if err != nil {
		panic(err)
	}

	fmt.Println(result)
}
