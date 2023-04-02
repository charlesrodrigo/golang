package repository

import "br.com.charlesrodrigo/ms-person/model"

type PersonRepository interface {
	Save(person *model.Person)
}
