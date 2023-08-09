package Mongo

import (
	"Backend/Core/Globals"
	"context"
	"github.com/rs/zerolog/log"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"time"
)

const (
	startingToConnectDbMessage      = "Starting to connect to mongo database"
	errorConnectingToDbMessage      = "Error connecting to mongo database "
	connectionToDbSuccessfulMessage = "Connected to mongo database"
	startingToPingDbMessage         = "Starting to ping mongo database"
	errorPingingDbMessage           = "Error pinging mongo database "
	pingingDbSuccessfulMessage      = "Ping to mongo database was successful"
)

type MongoDatabase struct {
}

func (self MongoDatabase) Connect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	connectionString := Globals.EnvValues.MongoConnectionString
	log.Info().Msg(startingToConnectDbMessage)
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		log.Error().Msg(errorConnectingToDbMessage + err.Error())
		return err
	}
	log.Info().Msg(connectionToDbSuccessfulMessage)
	MongoConnectionItem.MapFromConnection(client)
	if pingError := self.ping(client); pingError != nil {
		log.Error().Msg(errorPingingDbMessage + pingError.Error())
		return pingError
	}
	return nil
}

func (self MongoDatabase) ping(client *mongo.Client) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	log.Info().Msg(startingToPingDbMessage)
	if pingError := client.Ping(context.TODO(), options.Client().ReadPreference); pingError != nil {
		client.Disconnect(ctx)
		return pingError
	}
	log.Info().Msg(pingingDbSuccessfulMessage)
	return nil
}
