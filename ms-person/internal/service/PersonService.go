package service

import (
	"br.com.charlesrodrigo/ms-person/internal/model"
)

type PersonService interface {
	Create(person *model.Person) error
	Update(person *model.Person) error
	Delete(id string) error
	FindById(id string) (model.Person, error)
	FindAll() []model.Person
}
