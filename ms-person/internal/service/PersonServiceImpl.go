package service

import (
	"context"

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
func (p *PersonServiceImpl) Create(ctx context.Context, person *model.Person) error {
	return p.PersonRepository.Create(ctx, person)
}

// Update implements PersonService
func (p *PersonServiceImpl) Update(ctx context.Context, person *model.Person) error {
	return p.PersonRepository.Update(ctx, person)
}

// Delete implements PersonService
func (p *PersonServiceImpl) Delete(ctx context.Context, id string) error {
	return p.PersonRepository.Delete(ctx, id)
}

// FindById implements PersonService
func (p *PersonServiceImpl) FindById(ctx context.Context, id string) (model.Person, error) {
	return p.PersonRepository.FindById(ctx, id)
}

// FindAll implements PersonService
func (p *PersonServiceImpl) FindAll(ctx context.Context) []model.Person {
	return p.PersonRepository.FindAll(ctx)
}
