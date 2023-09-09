package controller

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hiteshchoudhary/mongoapi/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)



const connectionString = "mongodb+srv://azad:9HmGicBGpxr8shqe@one.ybrid8w.mongodb.net/?retryWrites=true&w=majority"
const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

//connect with mongoDb

func init() {
	//client options
    clientOption := options.Client().ApplyURI(connectionString)

	//connect to mongo database
	client, err := mongo.Connect(context.TODO(), clientOption)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("MongoDB connection successful")

	collection = client.Database(dbName).Collection(colName)

	// collection instance
	fmt.Println("Collection instance is ready")
}

// MongoDB Helpers

func insertOneMovie(movie model.Netflix) {
   inserted, err := collection.InsertOne(context.Background(), movie)

   if err != nil {
	log.Fatal(err)
   }

   fmt.Println("Inserted One movie in the db with id:", inserted.InsertedID)
}

// update 1 record
func updateOneMovie(movieId string) {
	id, err := primitive.ObjectIDFromHex(movieId)
	if err != nil {
		log.Fatal(err)
	}

	filter := bson.M{"_id": id}
	update := bson.M{"$set": bson.M{"watched": true}}

	result, err := collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Modified count: ", result.ModifiedCount)
}

// delete 1 record
func deleteOneMovie(movieId string) {
	id, _ := primitive.ObjectIDFromHex(movieId)
	filter := bson.M{"_id": id}
	deleteCount, err := collection.DeleteOne(context.Background(), filter)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Movie got deleted with delete count: ", deleteCount)
}

// delete all records from mongodb
func deleteAllMovie() int64 {

	deleteResult, err := collection.DeleteMany(context.Background(), bson.D{{}}, nil)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Number of movies deleted: ", deleteResult.DeletedCount)
	return deleteResult.DeletedCount
}