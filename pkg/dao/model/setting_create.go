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

	"github.com/seal-io/walrus/pkg/dao/model/setting"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// SettingCreate is the builder for creating a Setting entity.
type SettingCreate struct {
	config
	mutation   *SettingMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *Setting
	fromUpsert bool
}

// SetCreateTime sets the "create_time" field.
func (sc *SettingCreate) SetCreateTime(t time.Time) *SettingCreate {
	sc.mutation.SetCreateTime(t)
	return sc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (sc *SettingCreate) SetNillableCreateTime(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetCreateTime(*t)
	}
	return sc
}

// SetUpdateTime sets the "update_time" field.
func (sc *SettingCreate) SetUpdateTime(t time.Time) *SettingCreate {
	sc.mutation.SetUpdateTime(t)
	return sc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (sc *SettingCreate) SetNillableUpdateTime(t *time.Time) *SettingCreate {
	if t != nil {
		sc.SetUpdateTime(*t)
	}
	return sc
}

// SetName sets the "name" field.
func (sc *SettingCreate) SetName(s string) *SettingCreate {
	sc.mutation.SetName(s)
	return sc
}

// SetValue sets the "value" field.
func (sc *SettingCreate) SetValue(c crypto.String) *SettingCreate {
	sc.mutation.SetValue(c)
	return sc
}

// SetHidden sets the "hidden" field.
func (sc *SettingCreate) SetHidden(b bool) *SettingCreate {
	sc.mutation.SetHidden(b)
	return sc
}

// SetNillableHidden sets the "hidden" field if the given value is not nil.
func (sc *SettingCreate) SetNillableHidden(b *bool) *SettingCreate {
	if b != nil {
		sc.SetHidden(*b)
	}
	return sc
}

// SetEditable sets the "editable" field.
func (sc *SettingCreate) SetEditable(b bool) *SettingCreate {
	sc.mutation.SetEditable(b)
	return sc
}

// SetNillableEditable sets the "editable" field if the given value is not nil.
func (sc *SettingCreate) SetNillableEditable(b *bool) *SettingCreate {
	if b != nil {
		sc.SetEditable(*b)
	}
	return sc
}

// SetSensitive sets the "sensitive" field.
func (sc *SettingCreate) SetSensitive(b bool) *SettingCreate {
	sc.mutation.SetSensitive(b)
	return sc
}

// SetNillableSensitive sets the "sensitive" field if the given value is not nil.
func (sc *SettingCreate) SetNillableSensitive(b *bool) *SettingCreate {
	if b != nil {
		sc.SetSensitive(*b)
	}
	return sc
}

// SetPrivate sets the "private" field.
func (sc *SettingCreate) SetPrivate(b bool) *SettingCreate {
	sc.mutation.SetPrivate(b)
	return sc
}

// SetNillablePrivate sets the "private" field if the given value is not nil.
func (sc *SettingCreate) SetNillablePrivate(b *bool) *SettingCreate {
	if b != nil {
		sc.SetPrivate(*b)
	}
	return sc
}

// SetID sets the "id" field.
func (sc *SettingCreate) SetID(o object.ID) *SettingCreate {
	sc.mutation.SetID(o)
	return sc
}

// Mutation returns the SettingMutation object of the builder.
func (sc *SettingCreate) Mutation() *SettingMutation {
	return sc.mutation
}

// Save creates the Setting in the database.
func (sc *SettingCreate) Save(ctx context.Context) (*Setting, error) {
	if err := sc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, sc.sqlSave, sc.mutation, sc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (sc *SettingCreate) SaveX(ctx context.Context) *Setting {
	v, err := sc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (sc *SettingCreate) Exec(ctx context.Context) error {
	_, err := sc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (sc *SettingCreate) ExecX(ctx context.Context) {
	if err := sc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (sc *SettingCreate) defaults() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		if setting.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized setting.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := setting.DefaultCreateTime()
		sc.mutation.SetCreateTime(v)
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		if setting.DefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized setting.DefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := setting.DefaultUpdateTime()
		sc.mutation.SetUpdateTime(v)
	}
	if _, ok := sc.mutation.Hidden(); !ok {
		v := setting.DefaultHidden
		sc.mutation.SetHidden(v)
	}
	if _, ok := sc.mutation.Editable(); !ok {
		v := setting.DefaultEditable
		sc.mutation.SetEditable(v)
	}
	if _, ok := sc.mutation.Sensitive(); !ok {
		v := setting.DefaultSensitive
		sc.mutation.SetSensitive(v)
	}
	if _, ok := sc.mutation.Private(); !ok {
		v := setting.DefaultPrivate
		sc.mutation.SetPrivate(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (sc *SettingCreate) check() error {
	if _, ok := sc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "Setting.create_time"`)}
	}
	if _, ok := sc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "Setting.update_time"`)}
	}
	if _, ok := sc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "Setting.name"`)}
	}
	if v, ok := sc.mutation.Name(); ok {
		if err := setting.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Setting.name": %w`, err)}
		}
	}
	if _, ok := sc.mutation.Value(); !ok {
		return &ValidationError{Name: "value", err: errors.New(`model: missing required field "Setting.value"`)}
	}
	return nil
}

func (sc *SettingCreate) sqlSave(ctx context.Context) (*Setting, error) {
	if err := sc.check(); err != nil {
		return nil, err
	}
	_node, _spec := sc.createSpec()
	if err := sqlgraph.CreateNode(ctx, sc.driver, _spec); err != nil {
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
	sc.mutation.id = &_node.ID
	sc.mutation.done = true
	return _node, nil
}

func (sc *SettingCreate) createSpec() (*Setting, *sqlgraph.CreateSpec) {
	var (
		_node = &Setting{config: sc.config}
		_spec = sqlgraph.NewCreateSpec(setting.Table, sqlgraph.NewFieldSpec(setting.FieldID, field.TypeString))
	)
	_spec.Schema = sc.schemaConfig.Setting
	_spec.OnConflict = sc.conflict
	if id, ok := sc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := sc.mutation.CreateTime(); ok {
		_spec.SetField(setting.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := sc.mutation.UpdateTime(); ok {
		_spec.SetField(setting.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if value, ok := sc.mutation.Name(); ok {
		_spec.SetField(setting.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := sc.mutation.Value(); ok {
		_spec.SetField(setting.FieldValue, field.TypeString, value)
		_node.Value = value
	}
	if value, ok := sc.mutation.Hidden(); ok {
		_spec.SetField(setting.FieldHidden, field.TypeBool, value)
		_node.Hidden = &value
	}
	if value, ok := sc.mutation.Editable(); ok {
		_spec.SetField(setting.FieldEditable, field.TypeBool, value)
		_node.Editable = &value
	}
	if value, ok := sc.mutation.Sensitive(); ok {
		_spec.SetField(setting.FieldSensitive, field.TypeBool, value)
		_node.Sensitive = &value
	}
	if value, ok := sc.mutation.Private(); ok {
		_spec.SetField(setting.FieldPrivate, field.TypeBool, value)
		_node.Private = &value
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
func (sc *SettingCreate) Set(obj *Setting) *SettingCreate {
	// Required.
	sc.SetName(obj.Name)
	sc.SetValue(obj.Value)

	// Optional.
	if obj.CreateTime != nil {
		sc.SetCreateTime(*obj.CreateTime)
	}
	if obj.UpdateTime != nil {
		sc.SetUpdateTime(*obj.UpdateTime)
	}
	if obj.Hidden != nil {
		sc.SetHidden(*obj.Hidden)
	}
	if obj.Editable != nil {
		sc.SetEditable(*obj.Editable)
	}
	if obj.Sensitive != nil {
		sc.SetSensitive(*obj.Sensitive)
	}
	if obj.Private != nil {
		sc.SetPrivate(*obj.Private)
	}

	// Record the given object.
	sc.object = obj

	return sc
}

// getClientSet returns the ClientSet for the given builder.
func (sc *SettingCreate) getClientSet() (mc ClientSet) {
	if _, ok := sc.config.driver.(*txDriver); ok {
		tx := &Tx{config: sc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: sc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Setting entity,
// which is always good for cascading create operations.
func (sc *SettingCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Setting) error) (*Setting, error) {
	obj, err := sc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := sc.getClientSet()
	if sc.fromUpsert {
		q := mc.Settings().Query().
			Where(
				setting.Name(obj.Name),
			)
		obj.ID, err = q.OnlyID(ctx)
		if err != nil {
			return nil, fmt.Errorf("model: failed to query id of Setting entity: %w", err)
		}
	}

	if x := sc.object; x != nil {
		if _, set := sc.mutation.Field(setting.FieldName); set {
			obj.Name = x.Name
		}
		if _, set := sc.mutation.Field(setting.FieldValue); set {
			obj.Value = x.Value
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
func (sc *SettingCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Setting) error) *Setting {
	obj, err := sc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (sc *SettingCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Setting) error) error {
	_, err := sc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (sc *SettingCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Setting) error) {
	if err := sc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the SettingCreate Set method,
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
func (scb *SettingCreateBulk) Set(objs ...*Setting) *SettingCreateBulk {
	if len(objs) != 0 {
		client := NewSettingClient(scb.config)

		scb.builders = make([]*SettingCreate, len(objs))
		for i := range objs {
			scb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		scb.objects = objs
	}

	return scb
}

// getClientSet returns the ClientSet for the given builder.
func (scb *SettingCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := scb.config.driver.(*txDriver); ok {
		tx := &Tx{config: scb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: scb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Setting entities,
// which is always good for cascading create operations.
func (scb *SettingCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Setting) error) ([]*Setting, error) {
	objs, err := scb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := scb.getClientSet()
	if scb.fromUpsert {
		for i := range objs {
			obj := objs[i]
			q := mc.Settings().Query().
				Where(
					setting.Name(obj.Name),
				)
			objs[i].ID, err = q.OnlyID(ctx)
			if err != nil {
				return nil, fmt.Errorf("model: failed to query id of Setting entity: %w", err)
			}
		}
	}

	if x := scb.objects; x != nil {
		for i := range x {
			if _, set := scb.builders[i].mutation.Field(setting.FieldName); set {
				objs[i].Name = x[i].Name
			}
			if _, set := scb.builders[i].mutation.Field(setting.FieldValue); set {
				objs[i].Value = x[i].Value
			}
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
func (scb *SettingCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Setting) error) []*Setting {
	objs, err := scb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (scb *SettingCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Setting) error) error {
	_, err := scb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (scb *SettingCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Setting) error) {
	if err := scb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *SettingUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Setting) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SettingUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *SettingUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Setting) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *SettingUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Setting) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the SettingUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SettingUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *SettingUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Setting) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.Create().
//		SetCreateTime(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (sc *SettingCreate) OnConflict(opts ...sql.ConflictOption) *SettingUpsertOne {
	sc.conflict = opts
	return &SettingUpsertOne{
		create: sc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (sc *SettingCreate) OnConflictColumns(columns ...string) *SettingUpsertOne {
	sc.conflict = append(sc.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertOne{
		create: sc,
	}
}

type (
	// SettingUpsertOne is the builder for "upsert"-ing
	//  one Setting node.
	SettingUpsertOne struct {
		create *SettingCreate
	}

	// SettingUpsert is the "OnConflict" setter.
	SettingUpsert struct {
		*sql.UpdateSet
	}
)

// SetUpdateTime sets the "update_time" field.
func (u *SettingUpsert) SetUpdateTime(v time.Time) *SettingUpsert {
	u.Set(setting.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SettingUpsert) UpdateUpdateTime() *SettingUpsert {
	u.SetExcluded(setting.FieldUpdateTime)
	return u
}

// SetValue sets the "value" field.
func (u *SettingUpsert) SetValue(v crypto.String) *SettingUpsert {
	u.Set(setting.FieldValue, v)
	return u
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SettingUpsert) UpdateValue() *SettingUpsert {
	u.SetExcluded(setting.FieldValue)
	return u
}

// SetHidden sets the "hidden" field.
func (u *SettingUpsert) SetHidden(v bool) *SettingUpsert {
	u.Set(setting.FieldHidden, v)
	return u
}

// UpdateHidden sets the "hidden" field to the value that was provided on create.
func (u *SettingUpsert) UpdateHidden() *SettingUpsert {
	u.SetExcluded(setting.FieldHidden)
	return u
}

// ClearHidden clears the value of the "hidden" field.
func (u *SettingUpsert) ClearHidden() *SettingUpsert {
	u.SetNull(setting.FieldHidden)
	return u
}

// SetEditable sets the "editable" field.
func (u *SettingUpsert) SetEditable(v bool) *SettingUpsert {
	u.Set(setting.FieldEditable, v)
	return u
}

// UpdateEditable sets the "editable" field to the value that was provided on create.
func (u *SettingUpsert) UpdateEditable() *SettingUpsert {
	u.SetExcluded(setting.FieldEditable)
	return u
}

// ClearEditable clears the value of the "editable" field.
func (u *SettingUpsert) ClearEditable() *SettingUpsert {
	u.SetNull(setting.FieldEditable)
	return u
}

// SetSensitive sets the "sensitive" field.
func (u *SettingUpsert) SetSensitive(v bool) *SettingUpsert {
	u.Set(setting.FieldSensitive, v)
	return u
}

// UpdateSensitive sets the "sensitive" field to the value that was provided on create.
func (u *SettingUpsert) UpdateSensitive() *SettingUpsert {
	u.SetExcluded(setting.FieldSensitive)
	return u
}

// ClearSensitive clears the value of the "sensitive" field.
func (u *SettingUpsert) ClearSensitive() *SettingUpsert {
	u.SetNull(setting.FieldSensitive)
	return u
}

// SetPrivate sets the "private" field.
func (u *SettingUpsert) SetPrivate(v bool) *SettingUpsert {
	u.Set(setting.FieldPrivate, v)
	return u
}

// UpdatePrivate sets the "private" field to the value that was provided on create.
func (u *SettingUpsert) UpdatePrivate() *SettingUpsert {
	u.SetExcluded(setting.FieldPrivate)
	return u
}

// ClearPrivate clears the value of the "private" field.
func (u *SettingUpsert) ClearPrivate() *SettingUpsert {
	u.SetNull(setting.FieldPrivate)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(setting.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SettingUpsertOne) UpdateNewValues() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(setting.FieldID)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(setting.FieldCreateTime)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(setting.FieldName)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *SettingUpsertOne) Ignore() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertOne) DoNothing() *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreate.OnConflict
// documentation for more info.
func (u *SettingUpsertOne) Update(set func(*SettingUpsert)) *SettingUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *SettingUpsertOne) SetUpdateTime(v time.Time) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateUpdateTime() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetValue sets the "value" field.
func (u *SettingUpsertOne) SetValue(v crypto.String) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateValue() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateValue()
	})
}

// SetHidden sets the "hidden" field.
func (u *SettingUpsertOne) SetHidden(v bool) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetHidden(v)
	})
}

// UpdateHidden sets the "hidden" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateHidden() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateHidden()
	})
}

// ClearHidden clears the value of the "hidden" field.
func (u *SettingUpsertOne) ClearHidden() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearHidden()
	})
}

// SetEditable sets the "editable" field.
func (u *SettingUpsertOne) SetEditable(v bool) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetEditable(v)
	})
}

// UpdateEditable sets the "editable" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateEditable() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateEditable()
	})
}

// ClearEditable clears the value of the "editable" field.
func (u *SettingUpsertOne) ClearEditable() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearEditable()
	})
}

// SetSensitive sets the "sensitive" field.
func (u *SettingUpsertOne) SetSensitive(v bool) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetSensitive(v)
	})
}

// UpdateSensitive sets the "sensitive" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdateSensitive() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateSensitive()
	})
}

// ClearSensitive clears the value of the "sensitive" field.
func (u *SettingUpsertOne) ClearSensitive() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearSensitive()
	})
}

// SetPrivate sets the "private" field.
func (u *SettingUpsertOne) SetPrivate(v bool) *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.SetPrivate(v)
	})
}

// UpdatePrivate sets the "private" field to the value that was provided on create.
func (u *SettingUpsertOne) UpdatePrivate() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.UpdatePrivate()
	})
}

// ClearPrivate clears the value of the "private" field.
func (u *SettingUpsertOne) ClearPrivate() *SettingUpsertOne {
	return u.Update(func(s *SettingUpsert) {
		s.ClearPrivate()
	})
}

// Exec executes the query.
func (u *SettingUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SettingCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *SettingUpsertOne) ID(ctx context.Context) (id object.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: SettingUpsertOne.ID is not supported by MySQL driver. Use SettingUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *SettingUpsertOne) IDX(ctx context.Context) object.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// SettingCreateBulk is the builder for creating many Setting entities in bulk.
type SettingCreateBulk struct {
	config
	err        error
	builders   []*SettingCreate
	conflict   []sql.ConflictOption
	objects    []*Setting
	fromUpsert bool
}

// Save creates the Setting entities in the database.
func (scb *SettingCreateBulk) Save(ctx context.Context) ([]*Setting, error) {
	if scb.err != nil {
		return nil, scb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(scb.builders))
	nodes := make([]*Setting, len(scb.builders))
	mutators := make([]Mutator, len(scb.builders))
	for i := range scb.builders {
		func(i int, root context.Context) {
			builder := scb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*SettingMutation)
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
					_, err = mutators[i+1].Mutate(root, scb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = scb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, scb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, scb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (scb *SettingCreateBulk) SaveX(ctx context.Context) []*Setting {
	v, err := scb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (scb *SettingCreateBulk) Exec(ctx context.Context) error {
	_, err := scb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (scb *SettingCreateBulk) ExecX(ctx context.Context) {
	if err := scb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Setting.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.SettingUpsert) {
//			SetCreateTime(v+v).
//		}).
//		Exec(ctx)
func (scb *SettingCreateBulk) OnConflict(opts ...sql.ConflictOption) *SettingUpsertBulk {
	scb.conflict = opts
	return &SettingUpsertBulk{
		create: scb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (scb *SettingCreateBulk) OnConflictColumns(columns ...string) *SettingUpsertBulk {
	scb.conflict = append(scb.conflict, sql.ConflictColumns(columns...))
	return &SettingUpsertBulk{
		create: scb,
	}
}

// SettingUpsertBulk is the builder for "upsert"-ing
// a bulk of Setting nodes.
type SettingUpsertBulk struct {
	create *SettingCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(setting.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *SettingUpsertBulk) UpdateNewValues() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(setting.FieldID)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(setting.FieldCreateTime)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(setting.FieldName)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Setting.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *SettingUpsertBulk) Ignore() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *SettingUpsertBulk) DoNothing() *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the SettingCreateBulk.OnConflict
// documentation for more info.
func (u *SettingUpsertBulk) Update(set func(*SettingUpsert)) *SettingUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&SettingUpsert{UpdateSet: update})
	}))
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *SettingUpsertBulk) SetUpdateTime(v time.Time) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateUpdateTime() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateUpdateTime()
	})
}

// SetValue sets the "value" field.
func (u *SettingUpsertBulk) SetValue(v crypto.String) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetValue(v)
	})
}

// UpdateValue sets the "value" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateValue() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateValue()
	})
}

// SetHidden sets the "hidden" field.
func (u *SettingUpsertBulk) SetHidden(v bool) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetHidden(v)
	})
}

// UpdateHidden sets the "hidden" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateHidden() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateHidden()
	})
}

// ClearHidden clears the value of the "hidden" field.
func (u *SettingUpsertBulk) ClearHidden() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearHidden()
	})
}

// SetEditable sets the "editable" field.
func (u *SettingUpsertBulk) SetEditable(v bool) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetEditable(v)
	})
}

// UpdateEditable sets the "editable" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateEditable() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateEditable()
	})
}

// ClearEditable clears the value of the "editable" field.
func (u *SettingUpsertBulk) ClearEditable() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearEditable()
	})
}

// SetSensitive sets the "sensitive" field.
func (u *SettingUpsertBulk) SetSensitive(v bool) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetSensitive(v)
	})
}

// UpdateSensitive sets the "sensitive" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdateSensitive() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdateSensitive()
	})
}

// ClearSensitive clears the value of the "sensitive" field.
func (u *SettingUpsertBulk) ClearSensitive() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearSensitive()
	})
}

// SetPrivate sets the "private" field.
func (u *SettingUpsertBulk) SetPrivate(v bool) *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.SetPrivate(v)
	})
}

// UpdatePrivate sets the "private" field to the value that was provided on create.
func (u *SettingUpsertBulk) UpdatePrivate() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.UpdatePrivate()
	})
}

// ClearPrivate clears the value of the "private" field.
func (u *SettingUpsertBulk) ClearPrivate() *SettingUpsertBulk {
	return u.Update(func(s *SettingUpsert) {
		s.ClearPrivate()
	})
}

// Exec executes the query.
func (u *SettingUpsertBulk) Exec(ctx context.Context) error {
	if u.create.err != nil {
		return u.create.err
	}
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the SettingCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for SettingCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *SettingUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
