package main

import (
	"fmt"

	"br.com.charlesrodrigo/ms-person/infra/database"
	"br.com.charlesrodrigo/ms-person/infra/database/config"
	"br.com.charlesrodrigo/ms-person/model"
	"br.com.charlesrodrigo/ms-person/service"
)

func main() {

	fmt.Println("Started Server!")

	db := config.DatabaseConnection()

	personRepository := database.NewPersonRepositoryImpl(db)
	personService := service.NewPersonServiceImpl(personRepository)

	//personService.Delete("642f480985f5026dcd885a71")

	person := model.NewPerson()
	personService.Create(&person)

	//person := personService.FindById("642f4a4570be0f6708715342")
	//fmt.Println("person", person)

	//person.Name = "alterando nome"

	//personService.Update(&person)

	//fmt.Println("person", person)

}
