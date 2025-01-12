package serviceresources

import (
	"context"

	"go.uber.org/multierr"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/types"
	optypes "github.com/seal-io/walrus/pkg/operator/types"
)

// Label applies the labels to the given model.ServiceResource list with the given operator.Operator.
func Label(ctx context.Context, op optypes.Operator, candidates []*model.ServiceResource) error {
	if op == nil {
		return nil
	}

	// Merge the errors to return them all at once,
	// instead of returning the first error.
	var berr error

	for i := range candidates {
		// Get label values.
		var (
			// Name.
			svcName     string
			projectName string
			envName     string
		)

		if ins := candidates[i].Edges.Service; ins == nil {
			continue
		} else {
			// Service name.
			svcName = ins.Name

			// Project name.
			if proj := ins.Edges.Project; proj != nil {
				projectName = proj.Name
			}

			// Environment name.
			if env := ins.Edges.Environment; env != nil {
				envName = env.Name
			}
		}

		ls := map[string]string{
			// Name.
			types.LabelWalrusEnvironmentName: envName,
			types.LabelWalrusProjectName:     projectName,
			types.LabelWalrusServiceName:     svcName,
		}

		err := op.Label(ctx, candidates[i], ls)
		if multierr.AppendInto(&berr, err) {
			continue
		}
	}

	return berr
}
