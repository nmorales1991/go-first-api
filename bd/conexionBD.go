package bd

import (
	"context"
	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
	"log"
	"os"
)

var MongoClient = ConexionBD()

func ConexionBD() *mongo.Client {
	// reviso si existe el archivo .env
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found")
	}
	// obtengo la url del env
	uri := os.Getenv("MONGODB_URI")
	if uri == "" {
		log.Fatal("You must set your 'MONGODB_URI' environmental variable.")
	}
	// conexión
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(uri))
	if err != nil {
		panic(err)
	}

	return client
}

func Ping() bool {
	if err := MongoClient.Ping(context.TODO(), readpref.Primary()); err != nil {
		return false
	}
	return true
}
