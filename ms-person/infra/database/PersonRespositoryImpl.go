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
	ctx        context.Context
}

func NewPersonRepositoryImpl(c context.Context) repository.PersonRepository {
	connection := getConnectionDb()
	collection := connection.Collection("person")

	return &PersonRepositoryImpl{ctx: c, Db: connection, Collection: collection}
}

// Save implements Person
func (repo *PersonRepositoryImpl) Create(person *model.Person) (err error) {

	result, err := repo.Collection.InsertOne(repo.ctx, person)

	if err != nil {
		logger.Error("Failed insert %s", err.Error())
		return
	}

	logger.Info("Inserted a single document: %s", result.InsertedID.(primitive.ObjectID).Hex())

	return
}

// Update implements Person
func (repo *PersonRepositoryImpl) Update(person *model.Person) (err error) {

	filter := bson.M{"_id": person.ID}
	update := bson.M{"$set": person}

	result, err := repo.Collection.UpdateOne(repo.ctx, filter, update)

	if err != nil {
		logger.Error("Failed Update %s", err.Error())
		return
	}

	if result.ModifiedCount == 0 {
		err = errors.New("Not found document for update")
		return
	}

	logger.Info("update a single document: %s", result.UpsertedID)

	return
}

// Delete implements Person
func (repo *PersonRepositoryImpl) Delete(id string) (err error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.Error("id user invalid: %s", err.Error())
		return err
	}

	_, err = repo.Collection.DeleteOne(repo.ctx, bson.M{"_id": objectId})

	if err != nil {
		logger.Error("cannot delete user: %s", err.Error())
		return err
	}

	logger.Info("deleted a single document: %s", id)

	return
}

// FindById implements Person
func (repo *PersonRepositoryImpl) FindById(id string) (person model.Person, err error) {

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.Error("Failed FindById %s", err.Error())
		return model.Person{}, err
	}

	logger.Info("Find document by id: %s", id)

	err = repo.Collection.FindOne(repo.ctx, bson.M{"_id": objectId}).Decode(&person)

	if err != nil {
		return
	}

	return
}

// FindAll implements Person
func (repo *PersonRepositoryImpl) FindAll() []model.Person {
	return repo.Find(&ListPersonParams{})
}

func (repo *PersonRepositoryImpl) Find(params *ListPersonParams) (persons []model.Person) {
	opts := options.Find()
	opts = withListPersonParams(opts, params)

	cur, err := repo.Collection.Find(repo.ctx, bson.M{}, opts)

	if err != nil {
		return persons
	}

	defer cur.Close(repo.ctx)

	persons = make([]model.Person, 0)

	for cur.Next(repo.ctx) {
		person := model.Person{}

		err = cur.Decode(&person)

		function.IfErrorPanic("Failed Find", err)

		persons = append(persons, person)
	}

	return
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
