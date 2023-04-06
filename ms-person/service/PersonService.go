package service

import "br.com.charlesrodrigo/ms-person/model"

type PersonService interface {
	Create(person model.Person)
}
