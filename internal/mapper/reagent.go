package mapper

import (
	"github.com/yagehu/reactor/internal/entity"
	"github.com/yagehu/reactor/internal/model"
)

func ToReagentModel(x entity.Reagent) model.Reagent {
	return model.NewReagent(
		x.ID.String(),
		x.GUID.String(),
		x.CreatedAt,
		x.DeletedAt,
	)
}
