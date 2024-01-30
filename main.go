package main

import (
	"context"
	"fmt"
	"log"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name  string
	Age   int
	City  string
	Email string
}

func main() {
	// Set up MongoDB client
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb://localhost:27017"))
	if err != nil {
		log.Fatal(err)
	}
	defer client.Disconnect(context.Background())

	// Access the database and collection
	database := client.Database("testdb")
	collection := database.Collection("people")

	// Create
	person := Person{Name: "SWATHA M", Age: 21, City: "Coimbatore", Email: "swatha.manoharan@bootlabstech.com"}
	insertResult, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Inserted document ID:", insertResult.InsertedID)

	// Read
	var result Person
	filter := bson.D{{"name", "SWATHA M"}}
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found document: %+v\n", result)

	// Update
	update := bson.D{{"$set", bson.D{{"age", 31}}}}
	updateResult, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Matched %v document(s) and modified %v document(s)\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Read again to see the updated document
	err = collection.FindOne(context.Background(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Updated document: %+v\n", result)

	// Delete
	deleteResult, err := collection.DeleteOne(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Deleted %v document(s)\n", deleteResult.DeletedCount)
}

