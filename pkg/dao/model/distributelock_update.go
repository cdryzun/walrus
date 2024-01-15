// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	stdsql "database/sql"
	"errors"
	"fmt"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/distributelock"
	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
)

// DistributeLockUpdate is the builder for updating DistributeLock entities.
type DistributeLockUpdate struct {
	config
	hooks     []Hook
	mutation  *DistributeLockMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *DistributeLock
}

// Where appends a list predicates to the DistributeLockUpdate builder.
func (dlu *DistributeLockUpdate) Where(ps ...predicate.DistributeLock) *DistributeLockUpdate {
	dlu.mutation.Where(ps...)
	return dlu
}

// SetExpireAt sets the "expireAt" field.
func (dlu *DistributeLockUpdate) SetExpireAt(i int64) *DistributeLockUpdate {
	dlu.mutation.ResetExpireAt()
	dlu.mutation.SetExpireAt(i)
	return dlu
}

// SetNillableExpireAt sets the "expireAt" field if the given value is not nil.
func (dlu *DistributeLockUpdate) SetNillableExpireAt(i *int64) *DistributeLockUpdate {
	if i != nil {
		dlu.SetExpireAt(*i)
	}
	return dlu
}

// AddExpireAt adds i to the "expireAt" field.
func (dlu *DistributeLockUpdate) AddExpireAt(i int64) *DistributeLockUpdate {
	dlu.mutation.AddExpireAt(i)
	return dlu
}

// Mutation returns the DistributeLockMutation object of the builder.
func (dlu *DistributeLockUpdate) Mutation() *DistributeLockMutation {
	return dlu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (dlu *DistributeLockUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, dlu.sqlSave, dlu.mutation, dlu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dlu *DistributeLockUpdate) SaveX(ctx context.Context) int {
	affected, err := dlu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (dlu *DistributeLockUpdate) Exec(ctx context.Context) error {
	_, err := dlu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dlu *DistributeLockUpdate) ExecX(ctx context.Context) {
	if err := dlu.Exec(ctx); err != nil {
		panic(err)
	}
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value is not zero.
//
// For no default but required fields, Set calls directly.
//
// For no default but optional fields, Set calls if the value is not zero,
// or clears if the value is zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (dlu *DistributeLockUpdate) Set(obj *DistributeLock) *DistributeLockUpdate {
	// Without Default.
	dlu.SetExpireAt(obj.ExpireAt)

	// With Default.

	// Record the given object.
	dlu.object = obj

	return dlu
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dlu *DistributeLockUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DistributeLockUpdate {
	dlu.modifiers = append(dlu.modifiers, modifiers...)
	return dlu
}

func (dlu *DistributeLockUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(distributelock.Table, distributelock.Columns, sqlgraph.NewFieldSpec(distributelock.FieldID, field.TypeString))
	if ps := dlu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dlu.mutation.ExpireAt(); ok {
		_spec.SetField(distributelock.FieldExpireAt, field.TypeInt64, value)
	}
	if value, ok := dlu.mutation.AddedExpireAt(); ok {
		_spec.AddField(distributelock.FieldExpireAt, field.TypeInt64, value)
	}
	_spec.Node.Schema = dlu.schemaConfig.DistributeLock
	ctx = internal.NewSchemaConfigContext(ctx, dlu.schemaConfig)
	_spec.AddModifiers(dlu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, dlu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{distributelock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	dlu.mutation.done = true
	return n, nil
}

// DistributeLockUpdateOne is the builder for updating a single DistributeLock entity.
type DistributeLockUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *DistributeLockMutation
	modifiers []func(*sql.UpdateBuilder)
	object    *DistributeLock
}

// SetExpireAt sets the "expireAt" field.
func (dluo *DistributeLockUpdateOne) SetExpireAt(i int64) *DistributeLockUpdateOne {
	dluo.mutation.ResetExpireAt()
	dluo.mutation.SetExpireAt(i)
	return dluo
}

// SetNillableExpireAt sets the "expireAt" field if the given value is not nil.
func (dluo *DistributeLockUpdateOne) SetNillableExpireAt(i *int64) *DistributeLockUpdateOne {
	if i != nil {
		dluo.SetExpireAt(*i)
	}
	return dluo
}

// AddExpireAt adds i to the "expireAt" field.
func (dluo *DistributeLockUpdateOne) AddExpireAt(i int64) *DistributeLockUpdateOne {
	dluo.mutation.AddExpireAt(i)
	return dluo
}

// Mutation returns the DistributeLockMutation object of the builder.
func (dluo *DistributeLockUpdateOne) Mutation() *DistributeLockMutation {
	return dluo.mutation
}

// Where appends a list predicates to the DistributeLockUpdate builder.
func (dluo *DistributeLockUpdateOne) Where(ps ...predicate.DistributeLock) *DistributeLockUpdateOne {
	dluo.mutation.Where(ps...)
	return dluo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (dluo *DistributeLockUpdateOne) Select(field string, fields ...string) *DistributeLockUpdateOne {
	dluo.fields = append([]string{field}, fields...)
	return dluo
}

// Save executes the query and returns the updated DistributeLock entity.
func (dluo *DistributeLockUpdateOne) Save(ctx context.Context) (*DistributeLock, error) {
	return withHooks(ctx, dluo.sqlSave, dluo.mutation, dluo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (dluo *DistributeLockUpdateOne) SaveX(ctx context.Context) *DistributeLock {
	node, err := dluo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (dluo *DistributeLockUpdateOne) Exec(ctx context.Context) error {
	_, err := dluo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dluo *DistributeLockUpdateOne) ExecX(ctx context.Context) {
	if err := dluo.Exec(ctx); err != nil {
		panic(err)
	}
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For default fields, Set calls if the value changes from the original.
//
// For no default but required fields, Set calls if the value changes from the original.
//
// For no default but optional fields, Set calls if the value changes from the original,
// or clears if changes to zero.
//
// For example:
//
//	## Without Default
//
//	### Required
//
//	db.SetX(obj.X)
//
//	### Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   if _is_not_equal_(db.X, obj.X) {
//	      db.SetX(obj.X)
//	   }
//	} else {
//	   db.ClearX()
//	}
//
//	## With Default
//
//	if _is_zero_value_(obj.X) && _is_not_equal_(db.X, obj.X) {
//	   db.SetX(obj.X)
//	}
func (dluo *DistributeLockUpdateOne) Set(obj *DistributeLock) *DistributeLockUpdateOne {
	h := func(n ent.Mutator) ent.Mutator {
		return ent.MutateFunc(func(ctx context.Context, m ent.Mutation) (ent.Value, error) {
			mt := m.(*DistributeLockMutation)
			db, err := mt.Client().DistributeLock.Get(ctx, *mt.id)
			if err != nil {
				return nil, fmt.Errorf("failed getting DistributeLock with id: %v", *mt.id)
			}

			// Without Default.
			if db.ExpireAt != obj.ExpireAt {
				dluo.SetExpireAt(obj.ExpireAt)
			}

			// With Default.

			// Record the given object.
			dluo.object = obj

			return n.Mutate(ctx, m)
		})
	}

	dluo.hooks = append(dluo.hooks, h)

	return dluo
}

// getClientSet returns the ClientSet for the given builder.
func (dluo *DistributeLockUpdateOne) getClientSet() (mc ClientSet) {
	if _, ok := dluo.config.driver.(*txDriver); ok {
		tx := &Tx{config: dluo.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: dluo.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after updated the DistributeLock entity,
// which is always good for cascading update operations.
func (dluo *DistributeLockUpdateOne) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *DistributeLock) error) (*DistributeLock, error) {
	obj, err := dluo.Save(ctx)
	if err != nil &&
		(dluo.object == nil || !errors.Is(err, stdsql.ErrNoRows)) {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := dluo.getClientSet()

	if obj == nil {
		obj = dluo.object
	} else if x := dluo.object; x != nil {
		if _, set := dluo.mutation.Field(distributelock.FieldExpireAt); set {
			obj.ExpireAt = x.ExpireAt
		}
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (dluo *DistributeLockUpdateOne) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *DistributeLock) error) *DistributeLock {
	obj, err := dluo.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading update operations.
func (dluo *DistributeLockUpdateOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *DistributeLock) error) error {
	_, err := dluo.SaveE(ctx, cbs...)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (dluo *DistributeLockUpdateOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *DistributeLock) error) {
	if err := dluo.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (dluo *DistributeLockUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *DistributeLockUpdateOne {
	dluo.modifiers = append(dluo.modifiers, modifiers...)
	return dluo
}

func (dluo *DistributeLockUpdateOne) sqlSave(ctx context.Context) (_node *DistributeLock, err error) {
	_spec := sqlgraph.NewUpdateSpec(distributelock.Table, distributelock.Columns, sqlgraph.NewFieldSpec(distributelock.FieldID, field.TypeString))
	id, ok := dluo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "DistributeLock.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := dluo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, distributelock.FieldID)
		for _, f := range fields {
			if !distributelock.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != distributelock.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := dluo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := dluo.mutation.ExpireAt(); ok {
		_spec.SetField(distributelock.FieldExpireAt, field.TypeInt64, value)
	}
	if value, ok := dluo.mutation.AddedExpireAt(); ok {
		_spec.AddField(distributelock.FieldExpireAt, field.TypeInt64, value)
	}
	_spec.Node.Schema = dluo.schemaConfig.DistributeLock
	ctx = internal.NewSchemaConfigContext(ctx, dluo.schemaConfig)
	_spec.AddModifiers(dluo.modifiers...)
	_node = &DistributeLock{config: dluo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, dluo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{distributelock.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	dluo.mutation.done = true
	return _node, nil
}
