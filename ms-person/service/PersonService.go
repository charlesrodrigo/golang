package service

import "br.com.charlesrodrigo/ms-person/model"

type PersonService interface {
	Create(person *model.Person)
	Update(person *model.Person)
	Delete(id string) error
	FindById(id string) model.Person
	FindAll() []model.Person
}
