package database

import (
	"context"
	"fmt"

	"br.com.charlesrodrigo/ms-person/helper"
	"br.com.charlesrodrigo/ms-person/infra/repository"
	"br.com.charlesrodrigo/ms-person/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type PersonRepositoryImpl struct {
	Db         *mongo.Database
	Collection *mongo.Collection
}

func NewPersonRepositoryImpl(Db *mongo.Database) repository.PersonRepository {
	collection := Db.Collection("person")
	return &PersonRepositoryImpl{Db: Db, Collection: collection}
}

// Save implements Person
func (repo *PersonRepositoryImpl) Create(person *model.Person) {

	result, err := repo.Collection.InsertOne(context.TODO(), person)

	if err != nil {
		helper.ErrorPanic(err)
	}

	fmt.Println("Inserted a single document: ", result.InsertedID.(primitive.ObjectID).Hex())

	id := result.InsertedID.(primitive.ObjectID).Hex()

	*person = repo.FindById(id)

}

// Update implements Person
func (repo *PersonRepositoryImpl) Update(person *model.Person) {

	filter := bson.M{"_id": person.ID}
	update := bson.M{"$set": person}

	result, err := repo.Collection.UpdateOne(context.TODO(), filter, update)

	if err != nil {
		helper.ErrorPanic(err)
	}

	fmt.Println("update a single document: ", person.Name, result.UpsertedID)

}

// Delete implements Person
func (repo *PersonRepositoryImpl) Delete(id string) error {
	objectId, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return fmt.Errorf("cannot delete user: %w", err)
	}

	_, err = repo.Collection.DeleteOne(context.TODO(), bson.M{"_id": objectId})

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

	err = repo.Collection.FindOne(context.TODO(), bson.M{"_id": objectId}).Decode(&person)

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

	cur, err := repo.Collection.Find(context.TODO(), bson.M{}, opts)

	if err != nil {
		return []model.Person{}
	}

	defer cur.Close(context.TODO())

	res := make([]model.Person, 0)

	for cur.Next(context.TODO()) {
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
