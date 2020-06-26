package entity

import "github.com/gofrs/uuid"

type ReagentID uuid.UUID

func NewReagentID(u uuid.UUID) ReagentID {
	return ReagentID(u)
}

func (r ReagentID) String() string {
	return uuid.UUID(r).String()
}
