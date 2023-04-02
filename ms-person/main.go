package main

import (
	"fmt"

	"br.com.charlesrodrigo/ms-person/infra/database"
	"br.com.charlesrodrigo/ms-person/infra/database/config"
	"br.com.charlesrodrigo/ms-person/model"
	"br.com.charlesrodrigo/ms-person/service"
)

func main() {

	person := model.NewPerson()

	fmt.Println("Started Server!")

	db := config.DatabaseConnection()

	personRepository := database.NewPersonRepositoryImpl(db)
	personService := service.NewPersonServiceImpl(personRepository)

	personService.Create(person)
}
