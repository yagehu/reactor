package entity

import (
	"time"

	"github.com/gofrs/uuid"
)

type Reagent struct {
	ID        ReagentID
	GUID      uuid.UUID
	CreatedAt time.Time
	DeletedAt *time.Time
}

func NewReagent(
	id ReagentID, guid uuid.UUID, createdAt time.Time, deletedAt *time.Time,
) Reagent {
	return Reagent{
		ID:        id,
		GUID:      guid,
		CreatedAt: createdAt,
		DeletedAt: deletedAt,
	}
}
