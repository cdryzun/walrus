// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/environment"
	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
)

// EnvironmentDelete is the builder for deleting a Environment entity.
type EnvironmentDelete struct {
	config
	hooks    []Hook
	mutation *EnvironmentMutation
}

// Where appends a list predicates to the EnvironmentDelete builder.
func (ed *EnvironmentDelete) Where(ps ...predicate.Environment) *EnvironmentDelete {
	ed.mutation.Where(ps...)
	return ed
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (ed *EnvironmentDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, EnvironmentMutation](ctx, ed.sqlExec, ed.mutation, ed.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (ed *EnvironmentDelete) ExecX(ctx context.Context) int {
	n, err := ed.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (ed *EnvironmentDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: environment.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeOther,
				Column: environment.FieldID,
			},
		},
	}
	_spec.Node.Schema = ed.schemaConfig.Environment
	ctx = internal.NewSchemaConfigContext(ctx, ed.schemaConfig)
	if ps := ed.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, ed.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	ed.mutation.done = true
	return affected, err
}

// EnvironmentDeleteOne is the builder for deleting a single Environment entity.
type EnvironmentDeleteOne struct {
	ed *EnvironmentDelete
}

// Where appends a list predicates to the EnvironmentDelete builder.
func (edo *EnvironmentDeleteOne) Where(ps ...predicate.Environment) *EnvironmentDeleteOne {
	edo.ed.mutation.Where(ps...)
	return edo
}

// Exec executes the deletion query.
func (edo *EnvironmentDeleteOne) Exec(ctx context.Context) error {
	n, err := edo.ed.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{environment.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (edo *EnvironmentDeleteOne) ExecX(ctx context.Context) {
	if err := edo.Exec(ctx); err != nil {
		panic(err)
	}
}
