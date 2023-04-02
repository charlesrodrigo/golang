package service

import (
	"br.com.charlesrodrigo/ms-person/infra/repository"
	"br.com.charlesrodrigo/ms-person/model"
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
	p.PersonRepository.Save(person)
}
