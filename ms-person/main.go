package main

import (
	"fmt"

	"br.com.charlesrodrigo/ms-person/model"
)

func main() {

	person := model.NewPerson()

	fmt.Println("ola", person.Name, person.Address.Zipcode)
}
