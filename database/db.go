package database

import (
    "context"
    "log"

    "go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"
)

func ConnectDB() *mongo.Database {
    clientOptions := options.Client().ApplyURI("mongodb+srv://Aldiyar:Aldiyar@cluster0.rbc4h.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0")
    client, err := mongo.Connect(context.Background(), clientOptions)
    if err != nil {
        log.Fatal(err)
    }
    return client.Database("TestDB")
}
