package service

import (
	"br.com.charlesrodrigo/ms-person/internal/model"
	"br.com.charlesrodrigo/ms-person/internal/repository"
)

type PersonServiceImpl struct {
	PersonRepository repository.PersonRepository
}

func NewPersonServiceImpl(personRepository repository.PersonRepository) PersonService {
	return &PersonServiceImpl{
		PersonRepository: personRepository,
	}
}

// Create implements PersonService
func (p *PersonServiceImpl) Create(person *model.Person) {
	p.PersonRepository.Create(person)
}

// Update implements PersonService
func (p *PersonServiceImpl) Update(person *model.Person) {
	p.PersonRepository.Update(person)
}

// Delete implements PersonService
func (p *PersonServiceImpl) Delete(id string) error {
	return p.PersonRepository.Delete(id)
}

// FindById implements PersonService
func (p *PersonServiceImpl) FindById(id string) model.Person {
	return p.PersonRepository.FindById(id)
}

// FindAll implements PersonService
func (p *PersonServiceImpl) FindAll() []model.Person {
	return p.PersonRepository.FindAll()
}
