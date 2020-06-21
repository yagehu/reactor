package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Reagent struct {
	ID        uuid.UUID
	GUID      uuid.UUID
	CreatedAt time.Time
	DeletedAt *time.Time
}

func NewReagent(id, guid uuid.UUID, createdAt time.Time) Reagent {
	return Reagent{
		ID:        id,
		GUID:      guid,
		CreatedAt: createdAt,
		DeletedAt: nil,
	}
}
