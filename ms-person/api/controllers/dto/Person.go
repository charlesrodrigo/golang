package dto

import "br.com.charlesrodrigo/ms-person/model"

type Address struct {
	Zipcode      string `json:"zipcode" binding:"required"`
	Street       string `json:"street" binding:"required"`
	Neighborhood string `json:"neighborhood" binding:"required"`
	City         string `json:"city" binding:"required"`
	State        string `json:"state" binding:"required"`
	Country      string `json:"country" binding:"required"`
}

type CreatePersonRequest struct {
	Name    string  `json:"name" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Address Address `json:"address" binding:"required"`
}

func (createPersonRequest *CreatePersonRequest) ParseDTOToModel() model.Person {

	personAddress := model.Address{
		Zipcode:      createPersonRequest.Address.Zipcode,
		Street:       createPersonRequest.Address.Street,
		Neighborhood: createPersonRequest.Address.Neighborhood,
		City:         createPersonRequest.Address.City,
		State:        createPersonRequest.Address.State,
		Country:      createPersonRequest.Address.Country,
	}
	person := model.Person{
		Name:    createPersonRequest.Name,
		Email:   createPersonRequest.Email,
		Address: personAddress,
	}

	return person

}

type GetPersonRequest struct {
	Id      string  `json:"id"`
	Name    string  `json:"name"`
	Email   string  `json:"email"`
	Address Address `json:"address"`
}

func (getPersonRequest *GetPersonRequest) ParseModelToDTO(person model.Person) {

	personAddressDTO := Address{
		Zipcode:      person.Address.Zipcode,
		Street:       person.Address.Street,
		Neighborhood: person.Address.Neighborhood,
		City:         person.Address.City,
		State:        person.Address.State,
		Country:      person.Address.Country,
	}

	getPersonRequest.Id = person.ID.Hex()
	getPersonRequest.Name = person.Name
	getPersonRequest.Email = person.Email
	getPersonRequest.Address = personAddressDTO

}
