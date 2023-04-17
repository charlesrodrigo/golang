package database

import (
	"context"

	"br.com.charlesrodrigo/ms-person/helper/constants"
	"br.com.charlesrodrigo/ms-person/helper/function"
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

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(constants.DATABASE_URI))

	function.IfErrorFatal("Failed connect databae", err)

	err = client.Ping(ctx, readpref.Primary())

	function.IfErrorFatal("Failed connect databae", err)

	logger.Info("Connected database with successful")

	connection = client.Database(constants.DATABASE_NAME)

	return connection
}
