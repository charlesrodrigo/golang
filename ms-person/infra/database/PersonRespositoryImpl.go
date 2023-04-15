package database

import (
	"context"
	"fmt"
	"log"

	"br.com.charlesrodrigo/ms-person/helper"
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
	context    context.Context
}

func NewPersonRepositoryImpl(Db *mongo.Database) repository.PersonRepository {
	collection := Db.Collection("person")

	return &PersonRepositoryImpl{Db: Db, Collection: collection, context: context.Background()}
}

// Save implements Person
func (repo *PersonRepositoryImpl) Create(person *model.Person) {

	result, err := repo.Collection.InsertOne(repo.context, person)

	if err != nil {
		helper.ErrorPanic(err)
	}

	log.Println("Inserted a single document: ", result.InsertedID.(primitive.ObjectID).Hex())

	id := result.InsertedID.(primitive.ObjectID).Hex()

	*person = repo.FindById(id)

}

// Update implements Person
func (repo *PersonRepositoryImpl) Update(person *model.Person) {

	filter := bson.M{"_id": person.ID}
	update := bson.M{"$set": person}

	result, err := repo.Collection.UpdateOne(repo.context, filter, update)

	if err != nil {
		helper.ErrorPanic(err)
	}

	log.Println("update a single document: ", result.UpsertedID)

}

// Delete implements Person
func (repo *PersonRepositoryImpl) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("cannot delete user: %w", err)
	}

	_, err = repo.Collection.DeleteOne(repo.context, bson.M{"_id": objectId})

	if err != nil {
		return fmt.Errorf("cannot delete user: %w", err)
	}

	fmt.Println("deleted a single document: ", id)

	return nil
}

// FindById implements Person
func (repo *PersonRepositoryImpl) FindById(id string) model.Person {

	person := model.Person{}

	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		helper.ErrorPanic(err)
	}

	err = repo.Collection.FindOne(repo.context, bson.M{"_id": objectId}).Decode(&person)

	if err != nil {
		return model.Person{}
	}

	return person
}

// FindAll implements Person
func (repo *PersonRepositoryImpl) FindAll() []model.Person {
	return repo.Find(&ListPersonParams{})
}

func (repo *PersonRepositoryImpl) Find(params *ListPersonParams) []model.Person {
	opts := options.Find()
	opts = withListPersonParams(opts, params)

	cur, err := repo.Collection.Find(repo.context, bson.M{}, opts)

	if err != nil {
		return []model.Person{}
	}

	defer cur.Close(repo.context)

	res := make([]model.Person, 0)

	for cur.Next(repo.context) {
		person := model.Person{}

		err = cur.Decode(&person)
		if err != nil {
			helper.ErrorPanic(err)
		}

		fmt.Println(person)

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
