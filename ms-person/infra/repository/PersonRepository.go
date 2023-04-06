package repository

import "br.com.charlesrodrigo/ms-person/model"

type PersonRepository interface {
	Save(person model.Person)
	Update(person model.Person)
	Delete(id string) error
	FindById(id string) model.Person
	FindAll() []model.Person
}
