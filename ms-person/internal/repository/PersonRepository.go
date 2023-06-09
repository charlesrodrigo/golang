package repository

import (
	"context"

	"br.com.charlesrodrigo/ms-person/internal/model"
)

type PersonRepository interface {
	Create(ctx context.Context, person *model.Person) (string, error)
	Update(ctx context.Context, person *model.Person) error
	Delete(ctx context.Context, id string) error
	FindById(context context.Context, id string) (model.Person, error)
	FindAll(ctx context.Context) []model.Person
}
