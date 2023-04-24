package service

import (
	"context"

	"br.com.charlesrodrigo/ms-person/internal/model"
	"br.com.charlesrodrigo/ms-person/internal/repository"
)

type PersonServiceImpl struct {
	ctx              context.Context
	PersonRepository repository.PersonRepository
}

func NewPersonServiceImpl(c context.Context, personRepository repository.PersonRepository) PersonService {
	return &PersonServiceImpl{
		ctx:              c,
		PersonRepository: personRepository,
	}
}

// Create implements PersonService
func (p *PersonServiceImpl) Create(person *model.Person) error {
	return p.PersonRepository.Create(person)
}

// Update implements PersonService
func (p *PersonServiceImpl) Update(person *model.Person) error {
	return p.PersonRepository.Update(person)
}

// Delete implements PersonService
func (p *PersonServiceImpl) Delete(id string) error {
	return p.PersonRepository.Delete(id)
}

// FindById implements PersonService
func (p *PersonServiceImpl) FindById(id string) (model.Person, error) {
	return p.PersonRepository.FindById(id)
}

// FindAll implements PersonService
func (p *PersonServiceImpl) FindAll() []model.Person {
	return p.PersonRepository.FindAll()
}
