// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/resource"
	"github.com/seal-io/walrus/pkg/dao/model/resourcerelationship"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ResourceRelationshipCreate is the builder for creating a ResourceRelationship entity.
type ResourceRelationshipCreate struct {
	config
	mutation   *ResourceRelationshipMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *ResourceRelationship
	fromUpsert bool
}

// SetCreateTime sets the "create_time" field.
func (rrc *ResourceRelationshipCreate) SetCreateTime(t time.Time) *ResourceRelationshipCreate {
	rrc.mutation.SetCreateTime(t)
	return rrc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (rrc *ResourceRelationshipCreate) SetNillableCreateTime(t *time.Time) *ResourceRelationshipCreate {
	if t != nil {
		rrc.SetCreateTime(*t)
	}
	return rrc
}

// SetResourceID sets the "resource_id" field.
func (rrc *ResourceRelationshipCreate) SetResourceID(o object.ID) *ResourceRelationshipCreate {
	rrc.mutation.SetResourceID(o)
	return rrc
}

// SetDependencyID sets the "dependency_id" field.
func (rrc *ResourceRelationshipCreate) SetDependencyID(o object.ID) *ResourceRelationshipCreate {
	rrc.mutation.SetDependencyID(o)
	return rrc
}

// SetPath sets the "path" field.
func (rrc *ResourceRelationshipCreate) SetPath(o []object.ID) *ResourceRelationshipCreate {
	rrc.mutation.SetPath(o)
	return rrc
}

// SetType sets the "type" field.
func (rrc *ResourceRelationshipCreate) SetType(s string) *ResourceRelationshipCreate {
	rrc.mutation.SetType(s)
	return rrc
}

// SetID sets the "id" field.
func (rrc *ResourceRelationshipCreate) SetID(o object.ID) *ResourceRelationshipCreate {
	rrc.mutation.SetID(o)
	return rrc
}

// SetResource sets the "resource" edge to the Resource entity.
func (rrc *ResourceRelationshipCreate) SetResource(r *Resource) *ResourceRelationshipCreate {
	return rrc.SetResourceID(r.ID)
}

// SetDependency sets the "dependency" edge to the Resource entity.
func (rrc *ResourceRelationshipCreate) SetDependency(r *Resource) *ResourceRelationshipCreate {
	return rrc.SetDependencyID(r.ID)
}

// Mutation returns the ResourceRelationshipMutation object of the builder.
func (rrc *ResourceRelationshipCreate) Mutation() *ResourceRelationshipMutation {
	return rrc.mutation
}

// Save creates the ResourceRelationship in the database.
func (rrc *ResourceRelationshipCreate) Save(ctx context.Context) (*ResourceRelationship, error) {
	if err := rrc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, rrc.sqlSave, rrc.mutation, rrc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rrc *ResourceRelationshipCreate) SaveX(ctx context.Context) *ResourceRelationship {
	v, err := rrc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rrc *ResourceRelationshipCreate) Exec(ctx context.Context) error {
	_, err := rrc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rrc *ResourceRelationshipCreate) ExecX(ctx context.Context) {
	if err := rrc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rrc *ResourceRelationshipCreate) defaults() error {
	if _, ok := rrc.mutation.CreateTime(); !ok {
		if resourcerelationship.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized resourcerelationship.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := resourcerelationship.DefaultCreateTime()
		rrc.mutation.SetCreateTime(v)
	}
	if _, ok := rrc.mutation.Path(); !ok {
		v := resourcerelationship.DefaultPath
		rrc.mutation.SetPath(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (rrc *ResourceRelationshipCreate) check() error {
	if _, ok := rrc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "ResourceRelationship.create_time"`)}
	}
	if _, ok := rrc.mutation.ResourceID(); !ok {
		return &ValidationError{Name: "resource_id", err: errors.New(`model: missing required field "ResourceRelationship.resource_id"`)}
	}
	if v, ok := rrc.mutation.ResourceID(); ok {
		if err := resourcerelationship.ResourceIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "resource_id", err: fmt.Errorf(`model: validator failed for field "ResourceRelationship.resource_id": %w`, err)}
		}
	}
	if _, ok := rrc.mutation.DependencyID(); !ok {
		return &ValidationError{Name: "dependency_id", err: errors.New(`model: missing required field "ResourceRelationship.dependency_id"`)}
	}
	if v, ok := rrc.mutation.DependencyID(); ok {
		if err := resourcerelationship.DependencyIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "dependency_id", err: fmt.Errorf(`model: validator failed for field "ResourceRelationship.dependency_id": %w`, err)}
		}
	}
	if _, ok := rrc.mutation.Path(); !ok {
		return &ValidationError{Name: "path", err: errors.New(`model: missing required field "ResourceRelationship.path"`)}
	}
	if _, ok := rrc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`model: missing required field "ResourceRelationship.type"`)}
	}
	if v, ok := rrc.mutation.GetType(); ok {
		if err := resourcerelationship.TypeValidator(v); err != nil {
			return &ValidationError{Name: "type", err: fmt.Errorf(`model: validator failed for field "ResourceRelationship.type": %w`, err)}
		}
	}
	if _, ok := rrc.mutation.ResourceID(); !ok {
		return &ValidationError{Name: "resource", err: errors.New(`model: missing required edge "ResourceRelationship.resource"`)}
	}
	if _, ok := rrc.mutation.DependencyID(); !ok {
		return &ValidationError{Name: "dependency", err: errors.New(`model: missing required edge "ResourceRelationship.dependency"`)}
	}
	return nil
}

func (rrc *ResourceRelationshipCreate) sqlSave(ctx context.Context) (*ResourceRelationship, error) {
	if err := rrc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rrc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rrc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(*object.ID); ok {
			_node.ID = *id
		} else if err := _node.ID.Scan(_spec.ID.Value); err != nil {
			return nil, err
		}
	}
	rrc.mutation.id = &_node.ID
	rrc.mutation.done = true
	return _node, nil
}

func (rrc *ResourceRelationshipCreate) createSpec() (*ResourceRelationship, *sqlgraph.CreateSpec) {
	var (
		_node = &ResourceRelationship{config: rrc.config}
		_spec = sqlgraph.NewCreateSpec(resourcerelationship.Table, sqlgraph.NewFieldSpec(resourcerelationship.FieldID, field.TypeString))
	)
	_spec.Schema = rrc.schemaConfig.ResourceRelationship
	_spec.OnConflict = rrc.conflict
	if id, ok := rrc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := rrc.mutation.CreateTime(); ok {
		_spec.SetField(resourcerelationship.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := rrc.mutation.Path(); ok {
		_spec.SetField(resourcerelationship.FieldPath, field.TypeJSON, value)
		_node.Path = value
	}
	if value, ok := rrc.mutation.GetType(); ok {
		_spec.SetField(resourcerelationship.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if nodes := rrc.mutation.ResourceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   resourcerelationship.ResourceTable,
			Columns: []string{resourcerelationship.ResourceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		edge.Schema = rrc.schemaConfig.ResourceRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.ResourceID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := rrc.mutation.DependencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   resourcerelationship.DependencyTable,
			Columns: []string{resourcerelationship.DependencyColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(resource.FieldID, field.TypeString),
			},
		}
		edge.Schema = rrc.schemaConfig.ResourceRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.DependencyID = nodes[0]
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// Set is different from other Set* methods,
// it sets the value by judging the definition of each field within the entire object.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (rrc *ResourceRelationshipCreate) Set(obj *ResourceRelationship) *ResourceRelationshipCreate {
	// Required.
	rrc.SetResourceID(obj.ResourceID)
	rrc.SetDependencyID(obj.DependencyID)
	rrc.SetPath(obj.Path)
	rrc.SetType(obj.Type)

	// Optional.
	if obj.CreateTime != nil {
		rrc.SetCreateTime(*obj.CreateTime)
	}

	// Record the given object.
	rrc.object = obj

	return rrc
}

// getClientSet returns the ClientSet for the given builder.
func (rrc *ResourceRelationshipCreate) getClientSet() (mc ClientSet) {
	if _, ok := rrc.config.driver.(*txDriver); ok {
		tx := &Tx{config: rrc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: rrc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the ResourceRelationship entity,
// which is always good for cascading create operations.
func (rrc *ResourceRelationshipCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ResourceRelationship) error) (*ResourceRelationship, error) {
	obj, err := rrc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := rrc.getClientSet()

	if x := rrc.object; x != nil {
		if _, set := rrc.mutation.Field(resourcerelationship.FieldResourceID); set {
			obj.ResourceID = x.ResourceID
		}
		if _, set := rrc.mutation.Field(resourcerelationship.FieldDependencyID); set {
			obj.DependencyID = x.DependencyID
		}
		if _, set := rrc.mutation.Field(resourcerelationship.FieldType); set {
			obj.Type = x.Type
		}
		obj.Edges = x.Edges
	}

	for i := range cbs {
		if err = cbs[i](ctx, mc, obj); err != nil {
			return nil, err
		}
	}

	return obj, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (rrc *ResourceRelationshipCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ResourceRelationship) error) *ResourceRelationship {
	obj, err := rrc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (rrc *ResourceRelationshipCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ResourceRelationship) error) error {
	_, err := rrc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (rrc *ResourceRelationshipCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ResourceRelationship) error) {
	if err := rrc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the ResourceRelationshipCreate Set method,
// it sets the value by judging the definition of each field within the entire item of the given list.
//
// For required fields, Set calls directly.
//
// For optional fields, Set calls if the value is not zero.
//
// For example:
//
//	## Required
//
//	db.SetX(obj.X)
//
//	## Optional or Default
//
//	if _is_zero_value_(obj.X) {
//	   db.SetX(obj.X)
//	}
func (rrcb *ResourceRelationshipCreateBulk) Set(objs ...*ResourceRelationship) *ResourceRelationshipCreateBulk {
	if len(objs) != 0 {
		client := NewResourceRelationshipClient(rrcb.config)

		rrcb.builders = make([]*ResourceRelationshipCreate, len(objs))
		for i := range objs {
			rrcb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		rrcb.objects = objs
	}

	return rrcb
}

// getClientSet returns the ClientSet for the given builder.
func (rrcb *ResourceRelationshipCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := rrcb.config.driver.(*txDriver); ok {
		tx := &Tx{config: rrcb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: rrcb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the ResourceRelationship entities,
// which is always good for cascading create operations.
func (rrcb *ResourceRelationshipCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ResourceRelationship) error) ([]*ResourceRelationship, error) {
	objs, err := rrcb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := rrcb.getClientSet()

	if x := rrcb.objects; x != nil {
		for i := range x {
			if _, set := rrcb.builders[i].mutation.Field(resourcerelationship.FieldResourceID); set {
				objs[i].ResourceID = x[i].ResourceID
			}
			if _, set := rrcb.builders[i].mutation.Field(resourcerelationship.FieldDependencyID); set {
				objs[i].DependencyID = x[i].DependencyID
			}
			if _, set := rrcb.builders[i].mutation.Field(resourcerelationship.FieldType); set {
				objs[i].Type = x[i].Type
			}
			objs[i].Edges = x[i].Edges
		}
	}

	for i := range objs {
		for j := range cbs {
			if err = cbs[j](ctx, mc, objs[i]); err != nil {
				return nil, err
			}
		}
	}

	return objs, nil
}

// SaveEX is like SaveE, but panics if an error occurs.
func (rrcb *ResourceRelationshipCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ResourceRelationship) error) []*ResourceRelationship {
	objs, err := rrcb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (rrcb *ResourceRelationshipCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ResourceRelationship) error) error {
	_, err := rrcb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (rrcb *ResourceRelationshipCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *ResourceRelationship) error) {
	if err := rrcb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *ResourceRelationshipUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRelationship) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ResourceRelationshipUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *ResourceRelationshipUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRelationship) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *ResourceRelationshipUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRelationship) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the ResourceRelationshipUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ResourceRelationshipUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *ResourceRelationshipUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *ResourceRelationship) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ResourceRelationship.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ResourceRelationshipUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (rrc *ResourceRelationshipCreate) OnConflict(opts ...sql.ConflictOption) *ResourceRelationshipUpsertOne {
	rrc.conflict = opts
	return &ResourceRelationshipUpsertOne{
		create: rrc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ResourceRelationship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rrc *ResourceRelationshipCreate) OnConflictColumns(columns ...string) *ResourceRelationshipUpsertOne {
	rrc.conflict = append(rrc.conflict, sql.ConflictColumns(columns...))
	return &ResourceRelationshipUpsertOne{
		create: rrc,
	}
}

type (
	// ResourceRelationshipUpsertOne is the builder for "upsert"-ing
	//  one ResourceRelationship node.
	ResourceRelationshipUpsertOne struct {
		create *ResourceRelationshipCreate
	}

	// ResourceRelationshipUpsert is the "OnConflict" setter.
	ResourceRelationshipUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.ResourceRelationship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(resourcerelationship.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ResourceRelationshipUpsertOne) UpdateNewValues() *ResourceRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(resourcerelationship.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(resourcerelationship.FieldCreateTime)
		}
		if _, exists := u.create.mutation.ResourceID(); exists {
			s.SetIgnore(resourcerelationship.FieldResourceID)
		}
		if _, exists := u.create.mutation.DependencyID(); exists {
			s.SetIgnore(resourcerelationship.FieldDependencyID)
		}
		if _, exists := u.create.mutation.Path(); exists {
			s.SetIgnore(resourcerelationship.FieldPath)
		}
		if _, exists := u.create.mutation.GetType(); exists {
			s.SetIgnore(resourcerelationship.FieldType)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ResourceRelationship.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ResourceRelationshipUpsertOne) Ignore() *ResourceRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ResourceRelationshipUpsertOne) DoNothing() *ResourceRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ResourceRelationshipCreate.OnConflict
// documentation for more info.
func (u *ResourceRelationshipUpsertOne) Update(set func(*ResourceRelationshipUpsert)) *ResourceRelationshipUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ResourceRelationshipUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ResourceRelationshipUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ResourceRelationshipCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ResourceRelationshipUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ResourceRelationshipUpsertOne) ID(ctx context.Context) (id object.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: ResourceRelationshipUpsertOne.ID is not supported by MySQL driver. Use ResourceRelationshipUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ResourceRelationshipUpsertOne) IDX(ctx context.Context) object.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ResourceRelationshipCreateBulk is the builder for creating many ResourceRelationship entities in bulk.
type ResourceRelationshipCreateBulk struct {
	config
	err        error
	builders   []*ResourceRelationshipCreate
	conflict   []sql.ConflictOption
	objects    []*ResourceRelationship
	fromUpsert bool
}

// Save creates the ResourceRelationship entities in the database.
func (rrcb *ResourceRelationshipCreateBulk) Save(ctx context.Context) ([]*ResourceRelationship, error) {
	if rrcb.err != nil {
		return nil, rrcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rrcb.builders))
	nodes := make([]*ResourceRelationship, len(rrcb.builders))
	mutators := make([]Mutator, len(rrcb.builders))
	for i := range rrcb.builders {
		func(i int, root context.Context) {
			builder := rrcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ResourceRelationshipMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rrcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = rrcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rrcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rrcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rrcb *ResourceRelationshipCreateBulk) SaveX(ctx context.Context) []*ResourceRelationship {
	v, err := rrcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rrcb *ResourceRelationshipCreateBulk) Exec(ctx context.Context) error {
	_, err := rrcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rrcb *ResourceRelationshipCreateBulk) ExecX(ctx context.Context) {
	if err := rrcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.ResourceRelationship.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ResourceRelationshipUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (rrcb *ResourceRelationshipCreateBulk) OnConflict(opts ...sql.ConflictOption) *ResourceRelationshipUpsertBulk {
	rrcb.conflict = opts
	return &ResourceRelationshipUpsertBulk{
		create: rrcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.ResourceRelationship.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (rrcb *ResourceRelationshipCreateBulk) OnConflictColumns(columns ...string) *ResourceRelationshipUpsertBulk {
	rrcb.conflict = append(rrcb.conflict, sql.ConflictColumns(columns...))
	return &ResourceRelationshipUpsertBulk{
		create: rrcb,
	}
}

// ResourceRelationshipUpsertBulk is the builder for "upsert"-ing
// a bulk of ResourceRelationship nodes.
type ResourceRelationshipUpsertBulk struct {
	create *ResourceRelationshipCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.ResourceRelationship.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(resourcerelationship.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ResourceRelationshipUpsertBulk) UpdateNewValues() *ResourceRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(resourcerelationship.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(resourcerelationship.FieldCreateTime)
			}
			if _, exists := b.mutation.ResourceID(); exists {
				s.SetIgnore(resourcerelationship.FieldResourceID)
			}
			if _, exists := b.mutation.DependencyID(); exists {
				s.SetIgnore(resourcerelationship.FieldDependencyID)
			}
			if _, exists := b.mutation.Path(); exists {
				s.SetIgnore(resourcerelationship.FieldPath)
			}
			if _, exists := b.mutation.GetType(); exists {
				s.SetIgnore(resourcerelationship.FieldType)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.ResourceRelationship.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ResourceRelationshipUpsertBulk) Ignore() *ResourceRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ResourceRelationshipUpsertBulk) DoNothing() *ResourceRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ResourceRelationshipCreateBulk.OnConflict
// documentation for more info.
func (u *ResourceRelationshipUpsertBulk) Update(set func(*ResourceRelationshipUpsert)) *ResourceRelationshipUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ResourceRelationshipUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *ResourceRelationshipUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the ResourceRelationshipCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ResourceRelationshipCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ResourceRelationshipUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
