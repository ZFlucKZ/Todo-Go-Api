package db

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var DB *mongo.Database

func connectToMongo() *mongo.Client {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	connectionString := os.Getenv("MONGODB_CONNECTION_STRING")

	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(connectionString).SetServerAPIOptions(serverAPI)

	ctx, cancel := context.WithTimeout(context.TODO(), 10*time.Second)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		panic(err)
	}

	defer cancel()

	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	DB = client.Database("todo-app")

	return client
}

func Init() {
	connectToMongo()
}
