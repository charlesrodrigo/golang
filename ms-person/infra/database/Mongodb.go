package database

import (
	"context"
	"fmt"
	"os"

	"br.com.charlesrodrigo/ms-person/helper/constants"
	"br.com.charlesrodrigo/ms-person/helper/logger"
	mongoprom "github.com/globocom/mongo-go-prometheus"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

var connection *mongo.Database

func getConnectionDb() *mongo.Database {
	databaseName := os.Getenv(constants.DATABASE_NAME)
	databaseURI := os.Getenv(constants.DATABASE_URI)

	monitor := mongoprom.NewCommandMonitor(
		mongoprom.WithInstanceName(databaseName),
		mongoprom.WithNamespace(os.Getenv(constants.METRIC_NAME)),
		mongoprom.WithDurationBuckets([]float64{.001, .005, .01}),
	)

	logger.Info("Starting connect database")

	if connection != nil {
		logger.Info("return connect database")
		return connection
	}

	ctx, _ := context.WithTimeout(context.Background(), constants.TIMEOUT_CONTEXT)

	opts := options.Client().ApplyURI(databaseURI).SetMonitor(monitor)

	client, err := mongo.Connect(ctx, opts)

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed connect database %s", err))
	}

	err = client.Ping(ctx, readpref.Primary())

	if err != nil {
		logger.Fatal(fmt.Sprintf("Failed connect database %s", err))
	}

	logger.Info("Connected database with successful")

	connection = client.Database(databaseName)

	return connection
}
