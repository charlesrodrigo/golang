package database

import (
	"context"
	"fmt"
	"os"

	"br.com.charlesrodrigo/ms-person/helper/constants"
	"br.com.charlesrodrigo/ms-person/helper/logger"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connection *mongo.Database

func getConnectionDb() *mongo.Database {

	logger.Info("Starting connect database")

	if connection != nil {
		logger.Info("return connect database")
		return connection
	}

	ctx, _ := context.WithTimeout(context.Background(), constants.TIMEOUT_CONTEXT)

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(os.Getenv(constants.DATABASE_URI)))

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed connect database %s", err))
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed connect database %s", err))
	}

	logger.Info("Connected database with successful")

	connection = client.Database(os.Getenv(constants.DATABASE_NAME))

	return connection
}
