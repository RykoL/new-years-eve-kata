package handler

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"net/http"
	"os"
)

type Wish struct {
	Message string `json:"message",bson:"message"`
}

type CreateWishPayload struct {
	Message string `json:"message",bson:"message"`
}

var client, mongoErr = mongo.Connect(context.TODO(), options.Client().ApplyURI(os.Getenv("MONGODB_URI")))

func handleCreateWish(w http.ResponseWriter, r *http.Request) {

	if r.ContentLength == 0 {
		w.WriteHeader(400)
		return
	}

	decoder := json.NewDecoder(r.Body)
	var createWishPayload CreateWishPayload

	err := decoder.Decode(&createWishPayload)
	log.Println(createWishPayload.Message)

	if err != nil {
		log.Println(err)
		w.WriteHeader(400)
		return
	}
	wishesCollection := client.Database("new-years-eve").Collection("wishes")

	_, err = wishesCollection.InsertOne(context.TODO(), createWishPayload)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.WriteHeader(201)
}

func handleGetWishes(w http.ResponseWriter, r *http.Request) {
	wishesCollection := client.Database("new-years-eve").Collection("wishes")

	// retrieve all the documents in a collection
	cursor, err := wishesCollection.Find(context.TODO(), bson.D{})
	// check for errors in the finding
	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	// convert the cursor result to bson
	var wishes []Wish
	// check for errors in the conversion
	if err = cursor.All(context.TODO(), &wishes); err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	for _, result := range wishes {
		log.Println(result)
	}

	payload, err := json.Marshal(wishes)

	if err != nil {
		log.Println(err)
		w.WriteHeader(500)
		return
	}

	w.Header().Add("Content-Type", "application/json")
	w.Write(payload)

}

func Router(w http.ResponseWriter, r *http.Request) {
	if mongoErr != nil {
		panic(mongoErr)
	}

	if r.Method == "GET" {
		handleGetWishes(w, r)
	} else if r.Method == "POST" {
		handleCreateWish(w, r)
	} else {
		w.WriteHeader(405)
		return
	}
}
