package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	// Get the MongoDB URI from the environment variable
	mongoURI := os.Getenv("MONGO_URL")
	if mongoURI == "" {
		log.Fatal("MONGO_URL environment variable not set")
	}

	// Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Create a new router
	r := mux.NewRouter()

	// Define the GET API endpoint to get all people
	r.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		// Get a handle for your collection
		collection := client.Database("godb").Collection("people")

		// Find all documents in the collection
		cursor, err := collection.Find(context.TODO(), bson.D{})
		if err != nil {
			log.Fatal(err)
		}
		defer cursor.Close(context.TODO())

		// Iterate over the cursor and append results to a slice
		var people []Person
		for cursor.Next(context.TODO()) {
			var person Person
			err := cursor.Decode(&person)
			if err != nil {
				log.Fatal(err)
			}
			people = append(people, person)
		}

		// Marshal the slice into JSON and write it to the response writer
		jsonBytes, err := json.Marshal(people)
		if err != nil {
			log.Fatal(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(jsonBytes)
	}).Methods("GET")

	// Define the POST API endpoint to add a new person
	r.HandleFunc("/people", func(w http.ResponseWriter, r *http.Request) {
		// Parse the request body into a Person struct
		var person Person
		err := json.NewDecoder(r.Body).Decode(&person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Get a handle for your collection
		collection := client.Database("godb").Collection("people")

		// Insert the new document into the collection
		_, err = collection.InsertOne(context.TODO(), person)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		// Return a success message
		w.WriteHeader(http.StatusCreated)
		fmt.Fprintf(w, "Person added successfully")
	}).Methods("POST")

	// Start the HTTP server
	log.Fatal(http.ListenAndServe(":6000", r))
}
