package service

import (
	"context"

	"br.com.charlesrodrigo/ms-person/internal/model"
)

type PersonService interface {
	Create(ctx context.Context, person *model.Person) error
	Update(ctx context.Context, person *model.Person) error
	Delete(ctx context.Context, id string) error
	FindById(ctx context.Context, id string) (model.Person, error)
	FindAll(ctx context.Context) []model.Person
}
