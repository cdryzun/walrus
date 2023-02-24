// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/seal/pkg/dao/model/internal"
	"github.com/seal-io/seal/pkg/dao/model/predicate"
	"github.com/seal-io/seal/pkg/dao/model/token"
)

// TokenUpdate is the builder for updating Token entities.
type TokenUpdate struct {
	config
	hooks     []Hook
	mutation  *TokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// Where appends a list predicates to the TokenUpdate builder.
func (tu *TokenUpdate) Where(ps ...predicate.Token) *TokenUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdateTime sets the "updateTime" field.
func (tu *TokenUpdate) SetUpdateTime(t time.Time) *TokenUpdate {
	tu.mutation.SetUpdateTime(t)
	return tu
}

// SetCasdoorTokenName sets the "casdoorTokenName" field.
func (tu *TokenUpdate) SetCasdoorTokenName(s string) *TokenUpdate {
	tu.mutation.SetCasdoorTokenName(s)
	return tu
}

// SetCasdoorTokenOwner sets the "casdoorTokenOwner" field.
func (tu *TokenUpdate) SetCasdoorTokenOwner(s string) *TokenUpdate {
	tu.mutation.SetCasdoorTokenOwner(s)
	return tu
}

// SetName sets the "name" field.
func (tu *TokenUpdate) SetName(s string) *TokenUpdate {
	tu.mutation.SetName(s)
	return tu
}

// SetExpiration sets the "expiration" field.
func (tu *TokenUpdate) SetExpiration(i int) *TokenUpdate {
	tu.mutation.ResetExpiration()
	tu.mutation.SetExpiration(i)
	return tu
}

// SetNillableExpiration sets the "expiration" field if the given value is not nil.
func (tu *TokenUpdate) SetNillableExpiration(i *int) *TokenUpdate {
	if i != nil {
		tu.SetExpiration(*i)
	}
	return tu
}

// AddExpiration adds i to the "expiration" field.
func (tu *TokenUpdate) AddExpiration(i int) *TokenUpdate {
	tu.mutation.AddExpiration(i)
	return tu
}

// ClearExpiration clears the value of the "expiration" field.
func (tu *TokenUpdate) ClearExpiration() *TokenUpdate {
	tu.mutation.ClearExpiration()
	return tu
}

// Mutation returns the TokenMutation object of the builder.
func (tu *TokenUpdate) Mutation() *TokenMutation {
	return tu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TokenUpdate) Save(ctx context.Context) (int, error) {
	if err := tu.defaults(); err != nil {
		return 0, err
	}
	return withHooks[int, TokenMutation](ctx, tu.sqlSave, tu.mutation, tu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TokenUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TokenUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TokenUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TokenUpdate) defaults() error {
	if _, ok := tu.mutation.UpdateTime(); !ok {
		if token.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized token.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := token.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tu *TokenUpdate) check() error {
	if v, ok := tu.mutation.CasdoorTokenName(); ok {
		if err := token.CasdoorTokenNameValidator(v); err != nil {
			return &ValidationError{Name: "casdoorTokenName", err: fmt.Errorf(`model: validator failed for field "Token.casdoorTokenName": %w`, err)}
		}
	}
	if v, ok := tu.mutation.CasdoorTokenOwner(); ok {
		if err := token.CasdoorTokenOwnerValidator(v); err != nil {
			return &ValidationError{Name: "casdoorTokenOwner", err: fmt.Errorf(`model: validator failed for field "Token.casdoorTokenOwner": %w`, err)}
		}
	}
	if v, ok := tu.mutation.Name(); ok {
		if err := token.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Token.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tu *TokenUpdate) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TokenUpdate {
	tu.modifiers = append(tu.modifiers, modifiers...)
	return tu
}

func (tu *TokenUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := tu.check(); err != nil {
		return n, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: token.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.SetField(token.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tu.mutation.CasdoorTokenName(); ok {
		_spec.SetField(token.FieldCasdoorTokenName, field.TypeString, value)
	}
	if value, ok := tu.mutation.CasdoorTokenOwner(); ok {
		_spec.SetField(token.FieldCasdoorTokenOwner, field.TypeString, value)
	}
	if value, ok := tu.mutation.Name(); ok {
		_spec.SetField(token.FieldName, field.TypeString, value)
	}
	if value, ok := tu.mutation.Expiration(); ok {
		_spec.SetField(token.FieldExpiration, field.TypeInt, value)
	}
	if value, ok := tu.mutation.AddedExpiration(); ok {
		_spec.AddField(token.FieldExpiration, field.TypeInt, value)
	}
	if tu.mutation.ExpirationCleared() {
		_spec.ClearField(token.FieldExpiration, field.TypeInt)
	}
	_spec.Node.Schema = tu.schemaConfig.Token
	ctx = internal.NewSchemaConfigContext(ctx, tu.schemaConfig)
	_spec.AddModifiers(tu.modifiers...)
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	tu.mutation.done = true
	return n, nil
}

// TokenUpdateOne is the builder for updating a single Token entity.
type TokenUpdateOne struct {
	config
	fields    []string
	hooks     []Hook
	mutation  *TokenMutation
	modifiers []func(*sql.UpdateBuilder)
}

// SetUpdateTime sets the "updateTime" field.
func (tuo *TokenUpdateOne) SetUpdateTime(t time.Time) *TokenUpdateOne {
	tuo.mutation.SetUpdateTime(t)
	return tuo
}

// SetCasdoorTokenName sets the "casdoorTokenName" field.
func (tuo *TokenUpdateOne) SetCasdoorTokenName(s string) *TokenUpdateOne {
	tuo.mutation.SetCasdoorTokenName(s)
	return tuo
}

// SetCasdoorTokenOwner sets the "casdoorTokenOwner" field.
func (tuo *TokenUpdateOne) SetCasdoorTokenOwner(s string) *TokenUpdateOne {
	tuo.mutation.SetCasdoorTokenOwner(s)
	return tuo
}

// SetName sets the "name" field.
func (tuo *TokenUpdateOne) SetName(s string) *TokenUpdateOne {
	tuo.mutation.SetName(s)
	return tuo
}

// SetExpiration sets the "expiration" field.
func (tuo *TokenUpdateOne) SetExpiration(i int) *TokenUpdateOne {
	tuo.mutation.ResetExpiration()
	tuo.mutation.SetExpiration(i)
	return tuo
}

// SetNillableExpiration sets the "expiration" field if the given value is not nil.
func (tuo *TokenUpdateOne) SetNillableExpiration(i *int) *TokenUpdateOne {
	if i != nil {
		tuo.SetExpiration(*i)
	}
	return tuo
}

// AddExpiration adds i to the "expiration" field.
func (tuo *TokenUpdateOne) AddExpiration(i int) *TokenUpdateOne {
	tuo.mutation.AddExpiration(i)
	return tuo
}

// ClearExpiration clears the value of the "expiration" field.
func (tuo *TokenUpdateOne) ClearExpiration() *TokenUpdateOne {
	tuo.mutation.ClearExpiration()
	return tuo
}

// Mutation returns the TokenMutation object of the builder.
func (tuo *TokenUpdateOne) Mutation() *TokenMutation {
	return tuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TokenUpdateOne) Select(field string, fields ...string) *TokenUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Token entity.
func (tuo *TokenUpdateOne) Save(ctx context.Context) (*Token, error) {
	if err := tuo.defaults(); err != nil {
		return nil, err
	}
	return withHooks[*Token, TokenMutation](ctx, tuo.sqlSave, tuo.mutation, tuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TokenUpdateOne) SaveX(ctx context.Context) *Token {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TokenUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TokenUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TokenUpdateOne) defaults() error {
	if _, ok := tuo.mutation.UpdateTime(); !ok {
		if token.UpdateDefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized token.UpdateDefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := token.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (tuo *TokenUpdateOne) check() error {
	if v, ok := tuo.mutation.CasdoorTokenName(); ok {
		if err := token.CasdoorTokenNameValidator(v); err != nil {
			return &ValidationError{Name: "casdoorTokenName", err: fmt.Errorf(`model: validator failed for field "Token.casdoorTokenName": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.CasdoorTokenOwner(); ok {
		if err := token.CasdoorTokenOwnerValidator(v); err != nil {
			return &ValidationError{Name: "casdoorTokenOwner", err: fmt.Errorf(`model: validator failed for field "Token.casdoorTokenOwner": %w`, err)}
		}
	}
	if v, ok := tuo.mutation.Name(); ok {
		if err := token.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Token.name": %w`, err)}
		}
	}
	return nil
}

// Modify adds a statement modifier for attaching custom logic to the UPDATE statement.
func (tuo *TokenUpdateOne) Modify(modifiers ...func(u *sql.UpdateBuilder)) *TokenUpdateOne {
	tuo.modifiers = append(tuo.modifiers, modifiers...)
	return tuo
}

func (tuo *TokenUpdateOne) sqlSave(ctx context.Context) (_node *Token, err error) {
	if err := tuo.check(); err != nil {
		return _node, err
	}
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   token.Table,
			Columns: token.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: token.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`model: missing "Token.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, token.FieldID)
		for _, f := range fields {
			if !token.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != token.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.SetField(token.FieldUpdateTime, field.TypeTime, value)
	}
	if value, ok := tuo.mutation.CasdoorTokenName(); ok {
		_spec.SetField(token.FieldCasdoorTokenName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.CasdoorTokenOwner(); ok {
		_spec.SetField(token.FieldCasdoorTokenOwner, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Name(); ok {
		_spec.SetField(token.FieldName, field.TypeString, value)
	}
	if value, ok := tuo.mutation.Expiration(); ok {
		_spec.SetField(token.FieldExpiration, field.TypeInt, value)
	}
	if value, ok := tuo.mutation.AddedExpiration(); ok {
		_spec.AddField(token.FieldExpiration, field.TypeInt, value)
	}
	if tuo.mutation.ExpirationCleared() {
		_spec.ClearField(token.FieldExpiration, field.TypeInt)
	}
	_spec.Node.Schema = tuo.schemaConfig.Token
	ctx = internal.NewSchemaConfigContext(ctx, tuo.schemaConfig)
	_spec.AddModifiers(tuo.modifiers...)
	_node = &Token{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{token.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	tuo.mutation.done = true
	return _node, nil
}
