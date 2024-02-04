// SPDX-FileCopyrightText: 2024 Seal, Inc
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

	"github.com/seal-io/walrus/pkg/dao/model/subject"
	"github.com/seal-io/walrus/pkg/dao/model/token"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// TokenCreate is the builder for creating a Token entity.
type TokenCreate struct {
	config
	mutation   *TokenMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *Token
	fromUpsert bool
}

// SetCreateTime sets the "create_time" field.
func (tc *TokenCreate) SetCreateTime(t time.Time) *TokenCreate {
	tc.mutation.SetCreateTime(t)
	return tc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (tc *TokenCreate) SetNillableCreateTime(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetCreateTime(*t)
	}
	return tc
}

// SetSubjectID sets the "subject_id" field.
func (tc *TokenCreate) SetSubjectID(o object.ID) *TokenCreate {
	tc.mutation.SetSubjectID(o)
	return tc
}

// SetKind sets the "kind" field.
func (tc *TokenCreate) SetKind(s string) *TokenCreate {
	tc.mutation.SetKind(s)
	return tc
}

// SetNillableKind sets the "kind" field if the given value is not nil.
func (tc *TokenCreate) SetNillableKind(s *string) *TokenCreate {
	if s != nil {
		tc.SetKind(*s)
	}
	return tc
}

// SetName sets the "name" field.
func (tc *TokenCreate) SetName(s string) *TokenCreate {
	tc.mutation.SetName(s)
	return tc
}

// SetExpiration sets the "expiration" field.
func (tc *TokenCreate) SetExpiration(t time.Time) *TokenCreate {
	tc.mutation.SetExpiration(t)
	return tc
}

// SetNillableExpiration sets the "expiration" field if the given value is not nil.
func (tc *TokenCreate) SetNillableExpiration(t *time.Time) *TokenCreate {
	if t != nil {
		tc.SetExpiration(*t)
	}
	return tc
}

// SetValue sets the "value" field.
func (tc *TokenCreate) SetValue(c crypto.String) *TokenCreate {
	tc.mutation.SetValue(c)
	return tc
}

// SetID sets the "id" field.
func (tc *TokenCreate) SetID(o object.ID) *TokenCreate {
	tc.mutation.SetID(o)
	return tc
}

// SetSubject sets the "subject" edge to the Subject entity.
func (tc *TokenCreate) SetSubject(s *Subject) *TokenCreate {
	return tc.SetSubjectID(s.ID)
}

// Mutation returns the TokenMutation object of the builder.
func (tc *TokenCreate) Mutation() *TokenMutation {
	return tc.mutation
}

// Save creates the Token in the database.
func (tc *TokenCreate) Save(ctx context.Context) (*Token, error) {
	if err := tc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, tc.sqlSave, tc.mutation, tc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (tc *TokenCreate) SaveX(ctx context.Context) *Token {
	v, err := tc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tc *TokenCreate) Exec(ctx context.Context) error {
	_, err := tc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tc *TokenCreate) ExecX(ctx context.Context) {
	if err := tc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tc *TokenCreate) defaults() error {
	if _, ok := tc.mutation.CreateTime(); !ok {
		if token.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized token.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := token.DefaultCreateTime()
		tc.mutation.SetCreateTime(v)
	}
	if _, ok := tc.mutation.Kind(); !ok {
		v := token.DefaultKind
		tc.mutation.SetKind(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tc *TokenCreate) check() error {
	if _, ok := tc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "Token.create_time"`)}
	}
	if _, ok := tc.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject_id", err: errors.New(`model: missing required field "Token.subject_id"`)}
	}
	if v, ok := tc.mutation.SubjectID(); ok {
		if err := token.SubjectIDValidator(string(v)); err != nil {
			return &ValidationError{Name: "subject_id", err: fmt.Errorf(`model: validator failed for field "Token.subject_id": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Kind(); !ok {
		return &ValidationError{Name: "kind", err: errors.New(`model: missing required field "Token.kind"`)}
	}
	if _, ok := tc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "Token.name"`)}
	}
	if v, ok := tc.mutation.Name(); ok {
		if err := token.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Token.name": %w`, err)}
		}
	}
	if _, ok := tc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`model: missing required field "Token.value"`)}
	}
	if v, ok := tc.mutation.Value(); ok {
		if err := token.ValueValidator(string(v)); err != nil {
			return &ValidationError{Name: "value", err: fmt.Errorf(`model: validator failed for field "Token.value": %w`, err)}
		}
	}
	if _, ok := tc.mutation.SubjectID(); !ok {
		return &ValidationError{Name: "subject", err: errors.New(`model: missing required edge "Token.subject"`)}
	}
	return nil
}

func (tc *TokenCreate) sqlSave(ctx context.Context) (*Token, error) {
	if err := tc.check(); err != nil {
		return nil, err
	}
	_node, _spec := tc.createSpec()
	if err := sqlgraph.CreateNode(ctx, tc.driver, _spec); err != nil {
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
	tc.mutation.id = &_node.ID
	tc.mutation.done = true
	return _node, nil
}

func (tc *TokenCreate) createSpec() (*Token, *sqlgraph.CreateSpec) {
	var (
		_node = &Token{config: tc.config}
		_spec = sqlgraph.NewCreateSpec(token.Table, sqlgraph.NewFieldSpec(token.FieldID, field.TypeString))
	)
	_spec.Schema = tc.schemaConfig.Token
	_spec.OnConflict = tc.conflict
	if id, ok := tc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := tc.mutation.CreateTime(); ok {
		_spec.SetField(token.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := tc.mutation.Kind(); ok {
		_spec.SetField(token.FieldKind, field.TypeString, value)
		_node.Kind = value
	}
	if value, ok := tc.mutation.Name(); ok {
		_spec.SetField(token.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := tc.mutation.Expiration(); ok {
		_spec.SetField(token.FieldExpiration, field.TypeTime, value)
		_node.Expiration = &value
	}
	if value, ok := tc.mutation.Value(); ok {
		_spec.SetField(token.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if nodes := tc.mutation.SubjectIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   token.SubjectTable,
			Columns: []string{token.SubjectColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subject.FieldID, field.TypeString),
			},
		}
		edge.Schema = tc.schemaConfig.Token
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_node.SubjectID = nodes[0]
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
func (tc *TokenCreate) Set(obj *Token) *TokenCreate {
	// Required.
	tc.SetSubjectID(obj.SubjectID)
	tc.SetKind(obj.Kind)
	tc.SetName(obj.Name)
	tc.SetValue(obj.Value)

	// Optional.
	if obj.CreateTime != nil {
		tc.SetCreateTime(*obj.CreateTime)
	}
	if obj.Expiration != nil {
		tc.SetExpiration(*obj.Expiration)
	}

	// Record the given object.
	tc.object = obj

	return tc
}

// getClientSet returns the ClientSet for the given builder.
func (tc *TokenCreate) getClientSet() (mc ClientSet) {
	if _, ok := tc.config.driver.(*txDriver); ok {
		tx := &Tx{config: tc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: tc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Token entity,
// which is always good for cascading create operations.
func (tc *TokenCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Token) error) (*Token, error) {
	obj, err := tc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := tc.getClientSet()
	if tc.fromUpsert {
		q := mc.Tokens().Query().
			Where(
				token.SubjectID(obj.SubjectID),
				token.Name(obj.Name),
			)
		obj.ID, err = q.OnlyID(ctx)
		if err != nil {
			return nil, fmt.Errorf("model: failed to query id of Token entity: %w", err)
		}
	}

	if x := tc.object; x != nil {
		if _, set := tc.mutation.Field(token.FieldSubjectID); set {
			obj.SubjectID = x.SubjectID
		}
		if _, set := tc.mutation.Field(token.FieldName); set {
			obj.Name = x.Name
		}
		if _, set := tc.mutation.Field(token.FieldExpiration); set {
			obj.Expiration = x.Expiration
		}
		if _, set := tc.mutation.Field(token.FieldValue); set {
			obj.Value = x.Value
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
func (tc *TokenCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Token) error) *Token {
	obj, err := tc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (tc *TokenCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Token) error) error {
	_, err := tc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (tc *TokenCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Token) error) {
	if err := tc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the TokenCreate Set method,
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
func (tcb *TokenCreateBulk) Set(objs ...*Token) *TokenCreateBulk {
	if len(objs) != 0 {
		client := NewTokenClient(tcb.config)

		tcb.builders = make([]*TokenCreate, len(objs))
		for i := range objs {
			tcb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		tcb.objects = objs
	}

	return tcb
}

// getClientSet returns the ClientSet for the given builder.
func (tcb *TokenCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := tcb.config.driver.(*txDriver); ok {
		tx := &Tx{config: tcb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: tcb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Token entities,
// which is always good for cascading create operations.
func (tcb *TokenCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Token) error) ([]*Token, error) {
	objs, err := tcb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := tcb.getClientSet()
	if tcb.fromUpsert {
		for i := range objs {
			obj := objs[i]
			q := mc.Tokens().Query().
				Where(
					token.SubjectID(obj.SubjectID),
					token.Name(obj.Name),
				)
			objs[i].ID, err = q.OnlyID(ctx)
			if err != nil {
				return nil, fmt.Errorf("model: failed to query id of Token entity: %w", err)
			}
		}
	}

	if x := tcb.objects; x != nil {
		for i := range x {
			if _, set := tcb.builders[i].mutation.Field(token.FieldSubjectID); set {
				objs[i].SubjectID = x[i].SubjectID
			}
			if _, set := tcb.builders[i].mutation.Field(token.FieldName); set {
				objs[i].Name = x[i].Name
			}
			if _, set := tcb.builders[i].mutation.Field(token.FieldExpiration); set {
				objs[i].Expiration = x[i].Expiration
			}
			if _, set := tcb.builders[i].mutation.Field(token.FieldValue); set {
				objs[i].Value = x[i].Value
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
func (tcb *TokenCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Token) error) []*Token {
	objs, err := tcb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (tcb *TokenCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Token) error) error {
	_, err := tcb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (tcb *TokenCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Token) error) {
	if err := tcb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *TokenUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Token) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TokenUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *TokenUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Token) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *TokenUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Token) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the TokenUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TokenUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *TokenUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Token) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (tc *TokenCreate) OnConflict(opts ...sql.ConflictOption) *TokenUpsertOne {
	tc.conflict = opts
	return &TokenUpsertOne{
		create: tc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tc *TokenCreate) OnConflictColumns(columns ...string) *TokenUpsertOne {
	tc.conflict = append(tc.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertOne{
		create: tc,
	}
}

type (
	// TokenUpsertOne is the builder for "upsert"-ing
	//  one Token node.
	TokenUpsertOne struct {
		create *TokenCreate
	}

	// TokenUpsert is the "OnConflict" setter.
	TokenUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(token.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TokenUpsertOne) UpdateNewValues() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(token.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(token.FieldCreateTime)
		}
		if _, exists := u.create.mutation.SubjectID(); exists {
			s.SetIgnore(token.FieldSubjectID)
		}
		if _, exists := u.create.mutation.Kind(); exists {
			s.SetIgnore(token.FieldKind)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(token.FieldName)
		}
		if _, exists := u.create.mutation.Expiration(); exists {
			s.SetIgnore(token.FieldExpiration)
		}
		if _, exists := u.create.mutation.Value(); exists {
			s.SetIgnore(token.FieldValue)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *TokenUpsertOne) Ignore() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertOne) DoNothing() *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreate.OnConflict
// documentation for more info.
func (u *TokenUpsertOne) Update(set func(*TokenUpsert)) *TokenUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *TokenUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TokenCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *TokenUpsertOne) ID(ctx context.Context) (id object.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: TokenUpsertOne.ID is not supported by MySQL driver. Use TokenUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *TokenUpsertOne) IDX(ctx context.Context) object.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// TokenCreateBulk is the builder for creating many Token entities in bulk.
type TokenCreateBulk struct {
	config
	err        error
	builders   []*TokenCreate
	conflict   []sql.ConflictOption
	objects    []*Token
	fromUpsert bool
}

// Save creates the Token entities in the database.
func (tcb *TokenCreateBulk) Save(ctx context.Context) ([]*Token, error) {
	if tcb.err != nil {
		return nil, tcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(tcb.builders))
	nodes := make([]*Token, len(tcb.builders))
	mutators := make([]Mutator, len(tcb.builders))
	for i := range tcb.builders {
		func(i int, root context.Context) {
			builder := tcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*TokenMutation)
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
					_, err = mutators[i+1].Mutate(root, tcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = tcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, tcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, tcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (tcb *TokenCreateBulk) SaveX(ctx context.Context) []*Token {
	v, err := tcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (tcb *TokenCreateBulk) Exec(ctx context.Context) error {
	_, err := tcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tcb *TokenCreateBulk) ExecX(ctx context.Context) {
	if err := tcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Token.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.TokenUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (tcb *TokenCreateBulk) OnConflict(opts ...sql.ConflictOption) *TokenUpsertBulk {
	tcb.conflict = opts
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (tcb *TokenCreateBulk) OnConflictColumns(columns ...string) *TokenUpsertBulk {
	tcb.conflict = append(tcb.conflict, sql.ConflictColumns(columns...))
	return &TokenUpsertBulk{
		create: tcb,
	}
}

// TokenUpsertBulk is the builder for "upsert"-ing
// a bulk of Token nodes.
type TokenUpsertBulk struct {
	create *TokenCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(token.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *TokenUpsertBulk) UpdateNewValues() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(token.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(token.FieldCreateTime)
			}
			if _, exists := b.mutation.SubjectID(); exists {
				s.SetIgnore(token.FieldSubjectID)
			}
			if _, exists := b.mutation.Kind(); exists {
				s.SetIgnore(token.FieldKind)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(token.FieldName)
			}
			if _, exists := b.mutation.Expiration(); exists {
				s.SetIgnore(token.FieldExpiration)
			}
			if _, exists := b.mutation.Value(); exists {
				s.SetIgnore(token.FieldValue)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Token.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *TokenUpsertBulk) Ignore() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *TokenUpsertBulk) DoNothing() *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the TokenCreateBulk.OnConflict
// documentation for more info.
func (u *TokenUpsertBulk) Update(set func(*TokenUpsert)) *TokenUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&TokenUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *TokenUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the TokenCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for TokenCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *TokenUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
