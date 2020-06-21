package model

import (
	"time"
)

type Reagent struct {
	ID        string
	GUID      string
	CreatedAt time.Time
	DeletedAt *time.Time
}

func NewReagent(id string, guid string, createdAt time.Time) Reagent {
	return Reagent{
		ID:        id,
		GUID:      guid,
		CreatedAt: createdAt,
		DeletedAt: nil,
	}
}
