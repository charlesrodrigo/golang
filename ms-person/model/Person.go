package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Person struct {
	ID      primitive.ObjectID `bson:"_id,omitempty"`
	Audit   Audit
	Name    string `bson:"name,omitempty"`
	Email   string `bson:"email,omitempty"`
	Address Address
}

func NewPerson() *Person {

	address := Address{
		Zipcode:      "29166225",
		Street:       "Av Ribeirao Preto",
		Neighborhood: "Barcelona",
		City:         "Serra",
		State:        "ES",
		Country:      "Brasil"}

	person := Person{
		Name:    "charles rodrigo",
		Email:   "charlesrodrigo@gmail.com",
		Address: address,
	}

	return &person
}

func (person *Person) MarshalBSON() ([]byte, error) {
	if person.Audit.CreatedAt.IsZero() {
		person.Audit.CreatedAt = time.Now()
	}
	person.Audit.UpdatedAt = time.Now()

	type my Person
	return bson.Marshal((*my)(person))
}
