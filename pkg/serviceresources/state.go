package serviceresources

import (
	"context"

	"go.uber.org/multierr"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	optypes "github.com/seal-io/walrus/pkg/operator/types"
)

type StateResult struct {
	Error         bool
	Transitioning bool
}

func (r *StateResult) Merge(s StateResult) {
	r.merge(s.Error, s.Transitioning)
}

func (r *StateResult) merge(isError, isTransitioning bool) {
	switch {
	case isError:
		r.Error = true
		r.Transitioning = false
	case !r.Error && isTransitioning:
		r.Transitioning = true
	}
}

// State gets status of the given model.ServiceResource list with the given operator.Operator,
// and represents is ready if both `Error` and `Transitioning` of StateResult are false.
func State(
	ctx context.Context, op optypes.Operator,
	modelClient model.ClientSet,
	candidates []*model.ServiceResource,
) (StateResult, error) {
	var sr StateResult

	if op == nil {
		return sr, nil
	}

	// Merge the errors to return them all at once,
	// instead of returning the first error.
	var berr error

	for i := range candidates {
		// Get status of the application resource.
		st, err := op.GetStatus(ctx, candidates[i])
		if err != nil {
			berr = multierr.Append(berr, err)
		} else {
			sr.merge(st.Error, st.Transitioning)
		}

		// Get endpoints of the application resource.
		eps, err := op.GetEndpoints(ctx, candidates[i])
		if err != nil {
			berr = multierr.Append(berr, err)
		}

		// New application resource status.
		newStatus := types.ServiceResourceStatus{
			Status:            *st,
			ResourceEndpoints: eps,
		}
		if candidates[i].Status.Equal(newStatus) {
			// Do not update if the status is same as previous.
			continue
		}

		err = modelClient.ServiceResources().UpdateOne(candidates[i]).
			SetStatus(newStatus).
			Exec(ctx)
		if err != nil {
			if model.IsNotFound(err) {
				// Application resource has been deleted by other thread processing.
				continue
			}
			berr = multierr.Append(berr, err)
		}
	}

	return sr, berr
}
