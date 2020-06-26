package reagent

import (
	"context"
	"sync"

	"go.uber.org/zap"

	"github.com/yagehu/reactor/internal/entity"
)

type ReconcileParams struct {
	NewReagents    []entity.Reagent
	ActiveReagents []entity.Reagent
}

type ReconcileResult struct {
}

func (c *controller) Reconcile(
	ctx context.Context, p *ReconcileParams,
) (*ReconcileResult, error) {
	var (
		newReagents = make(
			map[entity.ReagentID]entity.Reagent, len(p.NewReagents),
		)
		wg sync.WaitGroup
	)

	for _, reagent := range p.NewReagents {
		newReagents[reagent.ID] = reagent
	}

	for _, reagent := range p.ActiveReagents {
		if _, ok := newReagents[reagent.ID]; ok {
			// No need to delete reagent.
			continue
		}

		wg.Add(1)

		go func(reagent entity.Reagent) {
			defer wg.Done()

			_, err := c.DeleteReagent(
				ctx, &DeleteReagentParams{Reagent: reagent},
			)
			if err != nil {
				c.logger.Error("Could not delete reagent.", zap.Error(err))
			}
		}(reagent)
	}

	wg.Wait()

	return &ReconcileResult{}, nil
}
