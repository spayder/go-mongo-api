package main

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/spayder/go-mongo-api/db"
	"github.com/spayder/go-mongo-api/handlers"
	"github.com/spayder/go-mongo-api/services"
	"go.mongodb.org/mongo-driver/mongo"
)

type Application struct {
	models services.Models
}

func main() {
	mongoClient, err := db.ConnectToMongo()
	if err != nil {
		log.Panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)

	defer cancel()

	defer func() {
		if err = mongoClient.Disconnect(ctx); err != nil {
			log.Panic(err)
		}
	}()

	services.New(mongoClient)

	log.Fatal(http.ListenAndServe(":8080", handlers.CreateRouter()))
}
