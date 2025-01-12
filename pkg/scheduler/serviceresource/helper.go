package serviceresource

import (
	"context"

	"github.com/seal-io/walrus/pkg/dao/model"
	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/types"
)

// listCandidateConnectors gets the all connector candidates for task.
func listCandidateConnectors(ctx context.Context, modelClient model.ClientSet) ([]*model.Connector, error) {
	return modelClient.Connectors().Query().
		Select(
			connector.FieldID,
			connector.FieldName,
			connector.FieldType,
			connector.FieldCategory,
			connector.FieldConfigVersion,
			connector.FieldConfigData).
		Where(
			connector.CategoryNEQ(types.ConnectorCategoryCustom),
			connector.CategoryNEQ(types.ConnectorCategoryVersionControl)).
		All(ctx)
}
