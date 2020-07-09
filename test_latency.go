package main

import (
	"context"
	"fmt"
	"log"
	"sort"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

// Movie contains fields of interest from movie documents.
type Movie struct {
	ID       primitive.ObjectID `bson:"_id"`
	Title    string             `bson:"title"`
	Plot     string             `bson:"plot"`
	Rated    string             `bson:"Rated"`
	Runtime  int                `bson:"runtime"`
	Type     string             `bson:"type"`
	Released time.Time          `bson:"released"`
	Cast     []string           `bson:"cast"`
}

const mongoURI = "mongodb+srv://veronica:<PASSWORD>@pluto.e3h8i.mongodb.net/test?retryWrites=true&w=majority"
const n = 100

func main() {
	
	var result Movie
	filter := bson.D{{"title", "Leaving Las Vegas"}}
	
	ctx := context.TODO()

	//  Set client options
	clientOptions := options.Client().ApplyURI(mongoURI)
	
	// Connect to MongoDB
	start := time.Now()
	client, err := mongo.Connect(ctx, clientOptions)
		if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)
	fmt.Println("Connected to MongoDB!")
	log.Printf("Connect and Ping took: %s", elapsed)

	defer client.Disconnect(ctx)
	
	col := client.Database("sample_mflix").Collection("movies")
	
	var durations []time.Duration
	
	for i := 0; i < n; i++ {
		findStart := time.Now()
		err = col.FindOne(ctx, filter).Decode(&result)
		if err != nil {
			log.Fatal(err)
		}
		findElapsed := time.Since(findStart)
		log.Printf("Found %+v in %s ", result.Title, findElapsed)
		durations = append(durations, findElapsed)
	}

	elapsed = time.Since(start)

	err = client.Disconnect(ctx)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Connection to MongoDB closed.")
	log.Printf("Total Duration: %s", elapsed)
	
	sort.Slice(durations, func(i, j int) bool {
		return durations[i] < (durations[j])
	})
	
	for i := 0; i < len(durations); i++ {
		log.Printf("%d : %s", i, durations[i])
	}
	

}
