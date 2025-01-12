// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresourcerelationship"
)

// ServiceResourceRelationshipUpdate is the builder for updating ServiceResourceRelationship entities.
type ServiceResourceRelationshipUpdate struct {
	config
	hooks     []Hook
	mutation  *ServiceResourceRelationshipMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ServiceResourceRelationship
}

// Where appends a list predicates to the ServiceResourceRelationshipUpdate builder.
func (srru *ServiceResourceRelationshipUpdate) Where(ps ...predicate.ServiceResourceRelationship) *ServiceResourceRelationshipUpdate {
	srru.mutation.Where(ps...)
	return srru
}

// Mutation returns the ServiceResourceRelationshipMutation object of the builder.
func (srru *ServiceResourceRelationshipUpdate) Mutation() *ServiceResourceRelationshipMutation {
	return srru.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (srru *ServiceResourceRelationshipUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, srru.sqlSave, srru.mutation, srru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (srru *ServiceResourceRelationshipUpdate) SaveX(ctx context.Context) int {
	affected, err := srru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (srru *ServiceResourceRelationshipUpdate) Exec(ctx context.Context) error {
	_, err := srru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srru *ServiceResourceRelationshipUpdate) ExecX(ctx context.Context) {
	if err := srru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (srru *ServiceResourceRelationshipUpdate) check() error {
	if _, ok := srru.mutation.ServiceResourceID(); srru.mutation.ServiceResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceResourceRelationship.serviceResource"`)
	}
	if _, ok := srru.mutation.DependencyID(); srru.mutation.DependencyCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceResourceRelationship.dependency"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (srru *ServiceResourceRelationshipUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServiceResourceRelationshipUpdate {
	srru.modifiers = append(srru.modifiers, modifiers...)
	return srru
}

func (srru *ServiceResourceRelationshipUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := srru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(serviceresourcerelationship.Table, serviceresourcerelationship.Columns, sqlgraph.NewFieldSpec(serviceresourcerelationship.FieldID, field.TypeString))
	if ps := srru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.Node.Schema = srru.schemaConfig.ServiceResourceRelationship
	ctx = internal.NewSchemaConfigContext(ctx, srru.schemaConfig)
	_spec.AddModifiers(srru.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, srru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceresourcerelationship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	srru.mutation.done = true
	return n, nil
}

// ServiceResourceRelationshipUpdateOne is the builder for updating a single ServiceResourceRelationship entity.
type ServiceResourceRelationshipUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *ServiceResourceRelationshipMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *ServiceResourceRelationship
}

// Mutation returns the ServiceResourceRelationshipMutation object of the builder.
func (srruo *ServiceResourceRelationshipUpdateOne) Mutation() *ServiceResourceRelationshipMutation {
	return srruo.mutation
}

// Where appends a list predicates to the ServiceResourceRelationshipUpdate builder.
func (srruo *ServiceResourceRelationshipUpdateOne) Where(ps ...predicate.ServiceResourceRelationship) *ServiceResourceRelationshipUpdateOne {
	srruo.mutation.Where(ps...)
	return srruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (srruo *ServiceResourceRelationshipUpdateOne) Select(field string, fields ...string) *ServiceResourceRelationshipUpdateOne {
	srruo.fields = append([]string{field}, fields...)
	return srruo
}

// Save executes the query and returns the updated ServiceResourceRelationship entity.
func (srruo *ServiceResourceRelationshipUpdateOne) Save(ctx context.Context) (*ServiceResourceRelationship, error) {
	return withHooks(ctx, srruo.sqlSave, srruo.mutation, srruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (srruo *ServiceResourceRelationshipUpdateOne) SaveX(ctx context.Context) *ServiceResourceRelationship {
	node, err := srruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (srruo *ServiceResourceRelationshipUpdateOne) Exec(ctx context.Context) error {
	_, err := srruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (srruo *ServiceResourceRelationshipUpdateOne) ExecX(ctx context.Context) {
	if err := srruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (srruo *ServiceResourceRelationshipUpdateOne) check() error {
	if _, ok := srruo.mutation.ServiceResourceID(); srruo.mutation.ServiceResourceCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceResourceRelationship.serviceResource"`)
	}
	if _, ok := srruo.mutation.DependencyID(); srruo.mutation.DependencyCleared() && !ok {
		return errors.New(`model: clearing a required unique edge "ServiceResourceRelationship.dependency"`)
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (srruo *ServiceResourceRelationshipUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *ServiceResourceRelationshipUpdateOne {
	srruo.modifiers = append(srruo.modifiers, modifiers...)
	return srruo
}

func (srruo *ServiceResourceRelationshipUpdateOne) sqlSave(ctx context.Context) (_node *ServiceResourceRelationship, err error) {
	if err := srruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(serviceresourcerelationship.Table, serviceresourcerelationship.Columns, sqlgraph.NewFieldSpec(serviceresourcerelationship.FieldID, field.TypeString))
	id, ok := srruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "ServiceResourceRelationship.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := srruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, serviceresourcerelationship.FieldID)
		for _, f := range fields {
			if !serviceresourcerelationship.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != serviceresourcerelationship.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := srruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	_spec.Node.Schema = srruo.schemaConfig.ServiceResourceRelationship
	ctx = internal.NewSchemaConfigContext(ctx, srruo.schemaConfig)
	_spec.AddModifiers(srruo.modifiers...)
	_node = &ServiceResourceRelationship{config: srruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, srruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{serviceresourcerelationship.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	srruo.mutation.done = true
	return _node, nil
}
