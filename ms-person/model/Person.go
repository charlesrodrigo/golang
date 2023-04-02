package model

type Person struct {
	Base    BaseEntity
	Name    string
	Email   string
	Address Address
}

func NewPerson() Person {
	address := Address{
		Zipcode:      "29166225",
		Street:       "Av Ribeirao Preto",
		Neighborhood: "Barcelona",
		City:         "Serra",
		State:        "ES",
		Country:      "Brasil"}

	person := Person{
		Name:    "charles",
		Email:   "charlesrodrigo@gmail.com",
		Address: address,
	}

	return person
}
