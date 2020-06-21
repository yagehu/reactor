package mapper

import (
	"github.com/gofrs/uuid"

	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/model"
)

func FromReagentModel(x model.Reagent) (entity.Reagent, error) {
	id, err := uuid.FromString(x.ID)
	if err != nil {
		return entity.Reagent{}, err
	}

	guid, err := uuid.FromString(x.GUID)
	if err != nil {
		return entity.Reagent{}, err
	}

	return entity.NewReagent(id, guid, x.CreatedAt, x.DeletedAt), nil
}

func FromReagentModelList(xs []model.Reagent) ([]entity.Reagent, error) {
	s := make([]entity.Reagent, len(xs))

	for i, x := range xs {
		reagent, err := FromReagentModel(x)
		if err != nil {
			return nil, err
		}

		s[i] = reagent
	}

	return s, nil
}

func ToReagentModel(x entity.Reagent) model.Reagent {
	return model.NewReagent(
		x.ID.String(),
		x.GUID.String(),
		x.CreatedAt,
		x.DeletedAt,
	)
}
