package model

import (
	"github.com/google/uuid"
	"go.starlark.net/lib/time"
)

type BaseEntity struct {
	Id               uuid.UUID
	CreatedDate      time.Time
	LastModifiedDate time.Time
}
