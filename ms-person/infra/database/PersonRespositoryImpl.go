package database

import (
	"context"
	"fmt"

	"br.com.charlesrodrigo/ms-person/helper"
	"br.com.charlesrodrigo/ms-person/infra/repository"
	"br.com.charlesrodrigo/ms-person/model"
	"go.mongodb.org/mongo-driver/mongo"
)

type PersonRepositoryImpl struct {
	Db *mongo.Database
}

func NewPersonRepositoryImpl(Db *mongo.Database) repository.PersonRepository {
	return &PersonRepositoryImpl{Db: Db}
}

// Save implements PersonRepository
func (repo *PersonRepositoryImpl) Save(person *model.Person) {

	personCollection := repo.Db.Collection("person")

	result, err := personCollection.InsertOne(context.TODO(), person)

	if err != nil {
		helper.ErrorPanic(err)
	}

	fmt.Println("Inserted a single document: ", result.InsertedID)

}
