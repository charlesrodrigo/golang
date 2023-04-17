package database

import (
	"context"
	"errors"

	"br.com.charlesrodrigo/ms-person/helper/function"
	"br.com.charlesrodrigo/ms-person/helper/logger"
	"br.com.charlesrodrigo/ms-person/internal/model"
	"br.com.charlesrodrigo/ms-person/internal/repository"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PersonRepositoryImpl struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}

func NewPersonRepositoryImpl() repository.PersonRepository {
	connection := getConnectionDb()
	collection := connection.Collection("person")

	return &PersonRepositoryImpl{Db: connection, Collection: collection}
}

// Save implements Person
func (repo *PersonRepositoryImpl) Create(ctx context.Context, person *model.Person) error {

	result, err := repo.Collection.InsertOne(ctx, person)

	if err != nil {
		logger.Error("Failed insert", err)
		return err
	}

	logger.Info("Inserted a single document: ", result.InsertedID.(primitive.ObjectID).Hex())

	return nil
}

// Update implements Person
func (repo *PersonRepositoryImpl) Update(ctx context.Context, person *model.Person) error {

	filter := bson.M{"_id": person.ID}
	update := bson.M{"$set": person}

	result, err := repo.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		logger.Error("Failed Update", err)
		return err
	}

	if result.ModifiedCount == 0 {
		return errors.New("Not found document for update")
	}

	logger.Info("update a single document: ", result.UpsertedID)

	return nil
}

// Delete implements Person
func (repo *PersonRepositoryImpl) Delete(ctx context.Context, id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.Error("id user invalid:", err.Error())
		return err
	}

	_, err = repo.Collection.DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		logger.Error("cannot delete user: %w", err)
		return err
	}

	logger.Info("deleted a single document: ", id)

	return nil
}

// FindById implements Person
func (repo *PersonRepositoryImpl) FindById(ctx context.Context, id string) (model.Person, error) {

	person := model.Person{}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.Error("Failed FindById", err)
		return model.Person{}, err
	}

	err = repo.Collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&person)

	if err != nil {
		return model.Person{}, nil
	}

	return person, nil
}

// FindAll implements Person
func (repo *PersonRepositoryImpl) FindAll(ctx context.Context) []model.Person {
	return repo.Find(ctx, &ListPersonParams{})
}

func (repo *PersonRepositoryImpl) Find(ctx context.Context, params *ListPersonParams) []model.Person {
	opts := options.Find()
	opts = withListPersonParams(opts, params)

	cur, err := repo.Collection.Find(ctx, bson.M{}, opts)

	if err != nil {
		return []model.Person{}
	}

	defer cur.Close(ctx)

	res := make([]model.Person, 0)

	for cur.Next(ctx) {
		person := model.Person{}

		err = cur.Decode(&person)

		function.IfErrorPanic("Failed Find", err)

		res = append(res, person)
	}

	return res
}

type ListPersonParams struct {
	Limit  *int64
	Offset *int64
}

func withListPersonParams(opts *options.FindOptions, params *ListPersonParams) *options.FindOptions {
	if params.Limit != nil {
		opts = opts.SetLimit(*params.Limit)
	}
	if params.Offset != nil {
		opts = opts.SetSkip(*params.Offset)
	}

	return opts
}
