package model

import (
	"time"
)

type Audit struct {
	CreatedAt time.Time `bson:"createdAt,omitempty"`
	UpdatedAt time.Time `bson:"updatedAt,omitempty"`
}
