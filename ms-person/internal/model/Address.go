package model

type Address struct {
	Zipcode      string `bson:"zipcode,omitempty"`
	Street       string `bson:"street,omitempty"`
	Neighborhood string `bson:"neighborhood,omitempty"`
	City         string `bson:"city,omitempty"`
	State        string `bson:"state,omitempty"`
	Country      string `bson:"country,omitempty"`
}
