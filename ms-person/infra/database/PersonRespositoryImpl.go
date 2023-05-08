package database

import (
	"context"
	"fmt"

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
func (repo *PersonRepositoryImpl) Create(ctx context.Context, person *model.Person) (string, error) {

	result, err := repo.Collection.InsertOne(ctx, person)

	if err != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("Failed insert %s", err.Error()))
		return "", err
	}

	logger.InfoWithContext(ctx, fmt.Sprintf("Inserted a single document: %s", result.InsertedID.(primitive.ObjectID).Hex()))

	objectID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return "", fmt.Errorf("failed to convert InsertedID to ObjectID")
	}

	if objectID.IsZero() {
		return "", mongo.ErrNilDocument
	}

	return objectID.Hex(), nil

}

// Update implements Person
func (repo *PersonRepositoryImpl) Update(ctx context.Context, person *model.Person) (err error) {

	filter := bson.M{"_id": person.ID}
	update := bson.M{"$set": person}

	result, err := repo.Collection.UpdateOne(ctx, filter, update)

	if err != nil {
		return
	}

	if result.ModifiedCount == 0 {
		return fmt.Errorf("Not found Person for update %s", person.ID.Hex())
	}

	logger.InfoWithContext(ctx, fmt.Sprintf("Update a single document: %s", person.ID.Hex()))

	return
}

// Delete implements Person
func (repo *PersonRepositoryImpl) Delete(ctx context.Context, id string) (err error) {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("id user invalid: %s", err.Error()))
		return err
	}

	result, err := repo.Collection.DeleteOne(ctx, bson.M{"_id": objectId})

	if err != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("cannot delete user: %s", err.Error()))
		return err
	}

	if result.DeletedCount == 0 {
		return fmt.Errorf("Not found Person for delete %s", id)
	}

	logger.InfoWithContext(ctx, fmt.Sprintf("deleted a single document: %s", id))

	return
}

// FindById implements Person
func (repo *PersonRepositoryImpl) FindById(ctx context.Context, id string) (person model.Person, err error) {

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		logger.ErrorWithContext(ctx, fmt.Sprintf("Failed FindById %s", err.Error()))
		return model.Person{}, err
	}

	logger.InfoWithContext(ctx, "Find document by id", "id", id)

	err = repo.Collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&person)

	if err != nil {
		if err == mongo.ErrNoDocuments {
			return model.Person{}, fmt.Errorf("Not found Person for id %s", id)
		}
		return
	}

	return
}

// FindAll implements Person
func (repo *PersonRepositoryImpl) FindAll(ctx context.Context) []model.Person {
	return repo.Find(ctx, &ListPersonParams{})
}

func (repo *PersonRepositoryImpl) Find(ctx context.Context, params *ListPersonParams) (persons []model.Person) {
	opts := options.Find()
	opts = withListPersonParams(opts, params)

	cur, err := repo.Collection.Find(ctx, bson.M{}, opts)

	if err != nil {
		return persons
	}

	defer cur.Close(ctx)

	persons = make([]model.Person, 0)

	for cur.Next(ctx) {
		person := model.Person{}

		err = cur.Decode(&person)

		if err != nil {
			logger.PanicWithContext(ctx, fmt.Sprintf("Failed Decode person %s", err.Error()))
		}

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
