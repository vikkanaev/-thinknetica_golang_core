package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type url struct {
	Long  string `bson:"long"`
	Short string `bson:"short"`
}

const (
	mongoConn      = "mongodb://localhost:27017/"
	databaseName   = "shortener" // имя БД
	collectionName = "urls"      // имя коллекции в БД
)

func main() {
	var u url
	mongoOpts := options.Client().ApplyURI(mongoConn)
	client, err := mongo.Connect(context.Background(), mongoOpts)
	if err != nil {
		log.Println("Err Connect")
		log.Fatal(err)
	}

	collection := client.Database(databaseName).Collection(collectionName)
	ctx := context.Background()

	shortUrl := "ekpo9u"
	filter := bson.D{primitive.E{Key: "short", Value: shortUrl}}

	err = collection.FindOne(ctx, filter).Decode(&u)
	if err != nil {
		// ErrNoDocuments means that the filter did not match any documents in
		// the collection.
		if err == mongo.ErrNoDocuments {
			return
		}
		log.Println("Err FindOne")
		log.Fatal(err)
	}

	fmt.Printf("Found the document: %+v\n", u.Short)

	count, _ := collection.EstimatedDocumentCount(ctx)
	fmt.Printf("Count documents: %+v\n", count)
}
