package db

import (
	"context"
	"log"
	"sync"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	Client *mongo.Client
}

var client *MongoDB
var lock = &sync.Mutex{}

const uri = "mongodb://localhost:27017"

func GetMongoDBClient() *MongoDB {
	if client == nil {
		lock.Lock()
		defer lock.Unlock()
		if client == nil {
			client = client.init()
		}
	}
	return client
}

func (*MongoDB) init() *MongoDB {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(uri).SetServerAPIOptions(serverAPI)
	// Create a new client and connect to the server
	tempClient, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Fatalf("Failed to establish mongodb connection to: %v \n %v", uri, err.Error())
	}
	client = &MongoDB{Client: tempClient}
	log.Println("Connection established with grpc")
	return client
}

func (*MongoDB) Cleanup() {
	if client != nil {
		err := client.Client.Disconnect(context.TODO())
		if err != nil {
			log.Printf("Failed to close mongodb connection. %v", err.Error())
		}
		client = nil
	}
}
