// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"reflect"
	"time"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/connector"
	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/service"
	"github.com/seal-io/walrus/pkg/dao/model/serviceresource"
	"github.com/seal-io/walrus/pkg/dao/model/servicerevision"
	"github.com/seal-io/walrus/pkg/dao/model/subjectrolerelationship"
	"github.com/seal-io/walrus/pkg/dao/model/variable"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ProjectCreate is the builder for creating a Project entity.
type ProjectCreate struct {
	config
	mutation   *ProjectMutation
	hooks      []Hook
	conflict   []sql.ConflictOption
	object     *Project
	fromUpsert bool
}

// SetName sets the "name" field.
func (pc *ProjectCreate) SetName(s string) *ProjectCreate {
	pc.mutation.SetName(s)
	return pc
}

// SetDescription sets the "description" field.
func (pc *ProjectCreate) SetDescription(s string) *ProjectCreate {
	pc.mutation.SetDescription(s)
	return pc
}

// SetNillableDescription sets the "description" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableDescription(s *string) *ProjectCreate {
	if s != nil {
		pc.SetDescription(*s)
	}
	return pc
}

// SetLabels sets the "labels" field.
func (pc *ProjectCreate) SetLabels(m map[string]string) *ProjectCreate {
	pc.mutation.SetLabels(m)
	return pc
}

// SetAnnotations sets the "annotations" field.
func (pc *ProjectCreate) SetAnnotations(m map[string]string) *ProjectCreate {
	pc.mutation.SetAnnotations(m)
	return pc
}

// SetCreateTime sets the "create_time" field.
func (pc *ProjectCreate) SetCreateTime(t time.Time) *ProjectCreate {
	pc.mutation.SetCreateTime(t)
	return pc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableCreateTime(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetCreateTime(*t)
	}
	return pc
}

// SetUpdateTime sets the "update_time" field.
func (pc *ProjectCreate) SetUpdateTime(t time.Time) *ProjectCreate {
	pc.mutation.SetUpdateTime(t)
	return pc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (pc *ProjectCreate) SetNillableUpdateTime(t *time.Time) *ProjectCreate {
	if t != nil {
		pc.SetUpdateTime(*t)
	}
	return pc
}

// SetID sets the "id" field.
func (pc *ProjectCreate) SetID(o object.ID) *ProjectCreate {
	pc.mutation.SetID(o)
	return pc
}

// AddEnvironmentIDs adds the "environments" edge to the Environment entity by IDs.
func (pc *ProjectCreate) AddEnvironmentIDs(ids ...object.ID) *ProjectCreate {
	pc.mutation.AddEnvironmentIDs(ids...)
	return pc
}

// AddEnvironments adds the "environments" edges to the Environment entity.
func (pc *ProjectCreate) AddEnvironments(e ...*Environment) *ProjectCreate {
	ids := make([]object.ID, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return pc.AddEnvironmentIDs(ids...)
}

// AddConnectorIDs adds the "connectors" edge to the Connector entity by IDs.
func (pc *ProjectCreate) AddConnectorIDs(ids ...object.ID) *ProjectCreate {
	pc.mutation.AddConnectorIDs(ids...)
	return pc
}

// AddConnectors adds the "connectors" edges to the Connector entity.
func (pc *ProjectCreate) AddConnectors(c ...*Connector) *ProjectCreate {
	ids := make([]object.ID, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pc.AddConnectorIDs(ids...)
}

// AddSubjectRoleIDs adds the "subject_roles" edge to the SubjectRoleRelationship entity by IDs.
func (pc *ProjectCreate) AddSubjectRoleIDs(ids ...object.ID) *ProjectCreate {
	pc.mutation.AddSubjectRoleIDs(ids...)
	return pc
}

// AddSubjectRoles adds the "subject_roles" edges to the SubjectRoleRelationship entity.
func (pc *ProjectCreate) AddSubjectRoles(s ...*SubjectRoleRelationship) *ProjectCreate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddSubjectRoleIDs(ids...)
}

// AddServiceIDs adds the "services" edge to the Service entity by IDs.
func (pc *ProjectCreate) AddServiceIDs(ids ...object.ID) *ProjectCreate {
	pc.mutation.AddServiceIDs(ids...)
	return pc
}

// AddServices adds the "services" edges to the Service entity.
func (pc *ProjectCreate) AddServices(s ...*Service) *ProjectCreate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddServiceIDs(ids...)
}

// AddServiceResourceIDs adds the "service_resources" edge to the ServiceResource entity by IDs.
func (pc *ProjectCreate) AddServiceResourceIDs(ids ...object.ID) *ProjectCreate {
	pc.mutation.AddServiceResourceIDs(ids...)
	return pc
}

// AddServiceResources adds the "service_resources" edges to the ServiceResource entity.
func (pc *ProjectCreate) AddServiceResources(s ...*ServiceResource) *ProjectCreate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddServiceResourceIDs(ids...)
}

// AddServiceRevisionIDs adds the "service_revisions" edge to the ServiceRevision entity by IDs.
func (pc *ProjectCreate) AddServiceRevisionIDs(ids ...object.ID) *ProjectCreate {
	pc.mutation.AddServiceRevisionIDs(ids...)
	return pc
}

// AddServiceRevisions adds the "service_revisions" edges to the ServiceRevision entity.
func (pc *ProjectCreate) AddServiceRevisions(s ...*ServiceRevision) *ProjectCreate {
	ids := make([]object.ID, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return pc.AddServiceRevisionIDs(ids...)
}

// AddVariableIDs adds the "variables" edge to the Variable entity by IDs.
func (pc *ProjectCreate) AddVariableIDs(ids ...object.ID) *ProjectCreate {
	pc.mutation.AddVariableIDs(ids...)
	return pc
}

// AddVariables adds the "variables" edges to the Variable entity.
func (pc *ProjectCreate) AddVariables(v ...*Variable) *ProjectCreate {
	ids := make([]object.ID, len(v))
	for i := range v {
		ids[i] = v[i].ID
	}
	return pc.AddVariableIDs(ids...)
}

// Mutation returns the ProjectMutation object of the builder.
func (pc *ProjectCreate) Mutation() *ProjectMutation {
	return pc.mutation
}

// Save creates the Project in the database.
func (pc *ProjectCreate) Save(ctx context.Context) (*Project, error) {
	if err := pc.defaults(); err != nil {
		return nil, err
	}
	return withHooks(ctx, pc.sqlSave, pc.mutation, pc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (pc *ProjectCreate) SaveX(ctx context.Context) *Project {
	v, err := pc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pc *ProjectCreate) Exec(ctx context.Context) error {
	_, err := pc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pc *ProjectCreate) ExecX(ctx context.Context) {
	if err := pc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (pc *ProjectCreate) defaults() error {
	if _, ok := pc.mutation.Labels(); !ok {
		v := project.DefaultLabels
		pc.mutation.SetLabels(v)
	}
	if _, ok := pc.mutation.Annotations(); !ok {
		v := project.DefaultAnnotations
		pc.mutation.SetAnnotations(v)
	}
	if _, ok := pc.mutation.CreateTime(); !ok {
		if project.DefaultCreateTime == nil {
			return fmt.Errorf("model: uninitialized project.DefaultCreateTime (forgotten import model/runtime?)")
		}
		v := project.DefaultCreateTime()
		pc.mutation.SetCreateTime(v)
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		if project.DefaultUpdateTime == nil {
			return fmt.Errorf("model: uninitialized project.DefaultUpdateTime (forgotten import model/runtime?)")
		}
		v := project.DefaultUpdateTime()
		pc.mutation.SetUpdateTime(v)
	}
	return nil
}

// check runs all checks and user-defined validators on the builder.
func (pc *ProjectCreate) check() error {
	if _, ok := pc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`model: missing required field "Project.name"`)}
	}
	if v, ok := pc.mutation.Name(); ok {
		if err := project.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`model: validator failed for field "Project.name": %w`, err)}
		}
	}
	if _, ok := pc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "Project.create_time"`)}
	}
	if _, ok := pc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "Project.update_time"`)}
	}
	return nil
}

func (pc *ProjectCreate) sqlSave(ctx context.Context) (*Project, error) {
	if err := pc.check(); err != nil {
		return nil, err
	}
	_node, _spec := pc.createSpec()
	if err := sqlgraph.CreateNode(ctx, pc.driver, _spec); err != nil {
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
	pc.mutation.id = &_node.ID
	pc.mutation.done = true
	return _node, nil
}

func (pc *ProjectCreate) createSpec() (*Project, *sqlgraph.CreateSpec) {
	var (
		_node = &Project{config: pc.config}
		_spec = sqlgraph.NewCreateSpec(project.Table, sqlgraph.NewFieldSpec(project.FieldID, field.TypeString))
	)
	_spec.Schema = pc.schemaConfig.Project
	_spec.OnConflict = pc.conflict
	if id, ok := pc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = &id
	}
	if value, ok := pc.mutation.Name(); ok {
		_spec.SetField(project.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := pc.mutation.Description(); ok {
		_spec.SetField(project.FieldDescription, field.TypeString, value)
		_node.Description = value
	}
	if value, ok := pc.mutation.Labels(); ok {
		_spec.SetField(project.FieldLabels, field.TypeJSON, value)
		_node.Labels = value
	}
	if value, ok := pc.mutation.Annotations(); ok {
		_spec.SetField(project.FieldAnnotations, field.TypeJSON, value)
		_node.Annotations = value
	}
	if value, ok := pc.mutation.CreateTime(); ok {
		_spec.SetField(project.FieldCreateTime, field.TypeTime, value)
		_node.CreateTime = &value
	}
	if value, ok := pc.mutation.UpdateTime(); ok {
		_spec.SetField(project.FieldUpdateTime, field.TypeTime, value)
		_node.UpdateTime = &value
	}
	if nodes := pc.mutation.EnvironmentsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.EnvironmentsTable,
			Columns: []string{project.EnvironmentsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(environment.FieldID, field.TypeString),
			},
		}
		edge.Schema = pc.schemaConfig.Environment
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ConnectorsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ConnectorsTable,
			Columns: []string{project.ConnectorsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(connector.FieldID, field.TypeString),
			},
		}
		edge.Schema = pc.schemaConfig.Connector
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.SubjectRolesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.SubjectRolesTable,
			Columns: []string{project.SubjectRolesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subjectrolerelationship.FieldID, field.TypeString),
			},
		}
		edge.Schema = pc.schemaConfig.SubjectRoleRelationship
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ServicesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ServicesTable,
			Columns: []string{project.ServicesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(service.FieldID, field.TypeString),
			},
		}
		edge.Schema = pc.schemaConfig.Service
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ServiceResourcesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ServiceResourcesTable,
			Columns: []string{project.ServiceResourcesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(serviceresource.FieldID, field.TypeString),
			},
		}
		edge.Schema = pc.schemaConfig.ServiceResource
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.ServiceRevisionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.ServiceRevisionsTable,
			Columns: []string{project.ServiceRevisionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(servicerevision.FieldID, field.TypeString),
			},
		}
		edge.Schema = pc.schemaConfig.ServiceRevision
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := pc.mutation.VariablesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   project.VariablesTable,
			Columns: []string{project.VariablesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(variable.FieldID, field.TypeString),
			},
		}
		edge.Schema = pc.schemaConfig.Variable
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
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
func (pc *ProjectCreate) Set(obj *Project) *ProjectCreate {
	// Required.
	pc.SetName(obj.Name)

	// Optional.
	if obj.Description != "" {
		pc.SetDescription(obj.Description)
	}
	if !reflect.ValueOf(obj.Labels).IsZero() {
		pc.SetLabels(obj.Labels)
	}
	if !reflect.ValueOf(obj.Annotations).IsZero() {
		pc.SetAnnotations(obj.Annotations)
	}
	if obj.CreateTime != nil {
		pc.SetCreateTime(*obj.CreateTime)
	}
	if obj.UpdateTime != nil {
		pc.SetUpdateTime(*obj.UpdateTime)
	}

	// Record the given object.
	pc.object = obj

	return pc
}

// getClientSet returns the ClientSet for the given builder.
func (pc *ProjectCreate) getClientSet() (mc ClientSet) {
	if _, ok := pc.config.driver.(*txDriver); ok {
		tx := &Tx{config: pc.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: pc.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Project entity,
// which is always good for cascading create operations.
func (pc *ProjectCreate) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Project) error) (*Project, error) {
	obj, err := pc.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return obj, err
	}

	mc := pc.getClientSet()
	if pc.fromUpsert {
		q := mc.Projects().Query().
			Where(
				project.Name(obj.Name),
			)
		obj.ID, err = q.OnlyID(ctx)
		if err != nil {
			return nil, fmt.Errorf("model: failed to query id of Project entity: %w", err)
		}
	}

	if x := pc.object; x != nil {
		if _, set := pc.mutation.Field(project.FieldName); set {
			obj.Name = x.Name
		}
		if _, set := pc.mutation.Field(project.FieldDescription); set {
			obj.Description = x.Description
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
func (pc *ProjectCreate) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Project) error) *Project {
	obj, err := pc.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return obj
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (pc *ProjectCreate) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Project) error) error {
	_, err := pc.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (pc *ProjectCreate) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Project) error) {
	if err := pc.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// Set leverages the ProjectCreate Set method,
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
func (pcb *ProjectCreateBulk) Set(objs ...*Project) *ProjectCreateBulk {
	if len(objs) != 0 {
		client := NewProjectClient(pcb.config)

		pcb.builders = make([]*ProjectCreate, len(objs))
		for i := range objs {
			pcb.builders[i] = client.Create().Set(objs[i])
		}

		// Record the given objects.
		pcb.objects = objs
	}

	return pcb
}

// getClientSet returns the ClientSet for the given builder.
func (pcb *ProjectCreateBulk) getClientSet() (mc ClientSet) {
	if _, ok := pcb.config.driver.(*txDriver); ok {
		tx := &Tx{config: pcb.config}
		tx.init()
		mc = tx
	} else {
		cli := &Client{config: pcb.config}
		cli.init()
		mc = cli
	}
	return mc
}

// SaveE calls the given function after created the Project entities,
// which is always good for cascading create operations.
func (pcb *ProjectCreateBulk) SaveE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Project) error) ([]*Project, error) {
	objs, err := pcb.Save(ctx)
	if err != nil {
		return nil, err
	}

	if len(cbs) == 0 {
		return objs, err
	}

	mc := pcb.getClientSet()
	if pcb.fromUpsert {
		for i := range objs {
			obj := objs[i]
			q := mc.Projects().Query().
				Where(
					project.Name(obj.Name),
				)
			objs[i].ID, err = q.OnlyID(ctx)
			if err != nil {
				return nil, fmt.Errorf("model: failed to query id of Project entity: %w", err)
			}
		}
	}

	if x := pcb.objects; x != nil {
		for i := range x {
			if _, set := pcb.builders[i].mutation.Field(project.FieldName); set {
				objs[i].Name = x[i].Name
			}
			if _, set := pcb.builders[i].mutation.Field(project.FieldDescription); set {
				objs[i].Description = x[i].Description
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
func (pcb *ProjectCreateBulk) SaveEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Project) error) []*Project {
	objs, err := pcb.SaveE(ctx, cbs...)
	if err != nil {
		panic(err)
	}
	return objs
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (pcb *ProjectCreateBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Project) error) error {
	_, err := pcb.SaveE(ctx, cbs...)
	return err
}

// ExecEX is like ExecE, but panics if an error occurs.
func (pcb *ProjectCreateBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, created *Project) error) {
	if err := pcb.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *ProjectUpsertOne) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Project) error) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ProjectUpsertOne.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *ProjectUpsertOne) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Project) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// ExecE calls the given function after executed the query,
// which is always good for cascading create operations.
func (u *ProjectUpsertBulk) ExecE(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Project) error) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the ProjectUpsertBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ProjectUpsertBulk.OnConflict")
	}
	u.create.fromUpsert = true
	return u.create.ExecE(ctx, cbs...)
}

// ExecEX is like ExecE, but panics if an error occurs.
func (u *ProjectUpsertBulk) ExecEX(ctx context.Context, cbs ...func(ctx context.Context, mc ClientSet, updated *Project) error) {
	if err := u.ExecE(ctx, cbs...); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Project.Create().
//		SetName(v).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pc *ProjectCreate) OnConflict(opts ...sql.ConflictOption) *ProjectUpsertOne {
	pc.conflict = opts
	return &ProjectUpsertOne{
		create: pc,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pc *ProjectCreate) OnConflictColumns(columns ...string) *ProjectUpsertOne {
	pc.conflict = append(pc.conflict, sql.ConflictColumns(columns...))
	return &ProjectUpsertOne{
		create: pc,
	}
}

type (
	// ProjectUpsertOne is the builder for "upsert"-ing
	//  one Project node.
	ProjectUpsertOne struct {
		create *ProjectCreate
	}

	// ProjectUpsert is the "OnConflict" setter.
	ProjectUpsert struct {
		*sql.UpdateSet
	}
)

// SetDescription sets the "description" field.
func (u *ProjectUpsert) SetDescription(v string) *ProjectUpsert {
	u.Set(project.FieldDescription, v)
	return u
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateDescription() *ProjectUpsert {
	u.SetExcluded(project.FieldDescription)
	return u
}

// ClearDescription clears the value of the "description" field.
func (u *ProjectUpsert) ClearDescription() *ProjectUpsert {
	u.SetNull(project.FieldDescription)
	return u
}

// SetLabels sets the "labels" field.
func (u *ProjectUpsert) SetLabels(v map[string]string) *ProjectUpsert {
	u.Set(project.FieldLabels, v)
	return u
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateLabels() *ProjectUpsert {
	u.SetExcluded(project.FieldLabels)
	return u
}

// ClearLabels clears the value of the "labels" field.
func (u *ProjectUpsert) ClearLabels() *ProjectUpsert {
	u.SetNull(project.FieldLabels)
	return u
}

// SetAnnotations sets the "annotations" field.
func (u *ProjectUpsert) SetAnnotations(v map[string]string) *ProjectUpsert {
	u.Set(project.FieldAnnotations, v)
	return u
}

// UpdateAnnotations sets the "annotations" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateAnnotations() *ProjectUpsert {
	u.SetExcluded(project.FieldAnnotations)
	return u
}

// ClearAnnotations clears the value of the "annotations" field.
func (u *ProjectUpsert) ClearAnnotations() *ProjectUpsert {
	u.SetNull(project.FieldAnnotations)
	return u
}

// SetUpdateTime sets the "update_time" field.
func (u *ProjectUpsert) SetUpdateTime(v time.Time) *ProjectUpsert {
	u.Set(project.FieldUpdateTime, v)
	return u
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ProjectUpsert) UpdateUpdateTime() *ProjectUpsert {
	u.SetExcluded(project.FieldUpdateTime)
	return u
}

// UpdateNewValues updates the mutable fields using the new values that were set on create except the ID field.
// Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(project.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProjectUpsertOne) UpdateNewValues() *ProjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		if _, exists := u.create.mutation.ID(); exists {
			s.SetIgnore(project.FieldID)
		}
		if _, exists := u.create.mutation.Name(); exists {
			s.SetIgnore(project.FieldName)
		}
		if _, exists := u.create.mutation.CreateTime(); exists {
			s.SetIgnore(project.FieldCreateTime)
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Project.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *ProjectUpsertOne) Ignore() *ProjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectUpsertOne) DoNothing() *ProjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectCreate.OnConflict
// documentation for more info.
func (u *ProjectUpsertOne) Update(set func(*ProjectUpsert)) *ProjectUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *ProjectUpsertOne) SetDescription(v string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateDescription() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ProjectUpsertOne) ClearDescription() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearDescription()
	})
}

// SetLabels sets the "labels" field.
func (u *ProjectUpsertOne) SetLabels(v map[string]string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateLabels() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *ProjectUpsertOne) ClearLabels() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearLabels()
	})
}

// SetAnnotations sets the "annotations" field.
func (u *ProjectUpsertOne) SetAnnotations(v map[string]string) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetAnnotations(v)
	})
}

// UpdateAnnotations sets the "annotations" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateAnnotations() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateAnnotations()
	})
}

// ClearAnnotations clears the value of the "annotations" field.
func (u *ProjectUpsertOne) ClearAnnotations() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearAnnotations()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ProjectUpsertOne) SetUpdateTime(v time.Time) *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ProjectUpsertOne) UpdateUpdateTime() *ProjectUpsertOne {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *ProjectUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ProjectCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *ProjectUpsertOne) ID(ctx context.Context) (id object.ID, err error) {
	if u.create.driver.Dialect() == dialect.MySQL {
		// In case of "ON CONFLICT", there is no way to get back non-numeric ID
		// fields from the database since MySQL does not support the RETURNING clause.
		return id, errors.New("model: ProjectUpsertOne.ID is not supported by MySQL driver. Use ProjectUpsertOne.Exec instead")
	}
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *ProjectUpsertOne) IDX(ctx context.Context) object.ID {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// ProjectCreateBulk is the builder for creating many Project entities in bulk.
type ProjectCreateBulk struct {
	config
	builders   []*ProjectCreate
	conflict   []sql.ConflictOption
	objects    []*Project
	fromUpsert bool
}

// Save creates the Project entities in the database.
func (pcb *ProjectCreateBulk) Save(ctx context.Context) ([]*Project, error) {
	specs := make([]*sqlgraph.CreateSpec, len(pcb.builders))
	nodes := make([]*Project, len(pcb.builders))
	mutators := make([]Mutator, len(pcb.builders))
	for i := range pcb.builders {
		func(i int, root context.Context) {
			builder := pcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*ProjectMutation)
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
					_, err = mutators[i+1].Mutate(root, pcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = pcb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, pcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, pcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (pcb *ProjectCreateBulk) SaveX(ctx context.Context) []*Project {
	v, err := pcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (pcb *ProjectCreateBulk) Exec(ctx context.Context) error {
	_, err := pcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcb *ProjectCreateBulk) ExecX(ctx context.Context) {
	if err := pcb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.Project.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		// Override some of the fields with custom
//		// update values.
//		Update(func(u *ent.ProjectUpsert) {
//			SetName(v+v).
//		}).
//		Exec(ctx)
func (pcb *ProjectCreateBulk) OnConflict(opts ...sql.ConflictOption) *ProjectUpsertBulk {
	pcb.conflict = opts
	return &ProjectUpsertBulk{
		create: pcb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (pcb *ProjectCreateBulk) OnConflictColumns(columns ...string) *ProjectUpsertBulk {
	pcb.conflict = append(pcb.conflict, sql.ConflictColumns(columns...))
	return &ProjectUpsertBulk{
		create: pcb,
	}
}

// ProjectUpsertBulk is the builder for "upsert"-ing
// a bulk of Project nodes.
type ProjectUpsertBulk struct {
	create *ProjectCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//			sql.ResolveWith(func(u *sql.UpdateSet) {
//				u.SetIgnore(project.FieldID)
//			}),
//		).
//		Exec(ctx)
func (u *ProjectUpsertBulk) UpdateNewValues() *ProjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(s *sql.UpdateSet) {
		for _, b := range u.create.builders {
			if _, exists := b.mutation.ID(); exists {
				s.SetIgnore(project.FieldID)
			}
			if _, exists := b.mutation.Name(); exists {
				s.SetIgnore(project.FieldName)
			}
			if _, exists := b.mutation.CreateTime(); exists {
				s.SetIgnore(project.FieldCreateTime)
			}
		}
	}))
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.Project.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *ProjectUpsertBulk) Ignore() *ProjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *ProjectUpsertBulk) DoNothing() *ProjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the ProjectCreateBulk.OnConflict
// documentation for more info.
func (u *ProjectUpsertBulk) Update(set func(*ProjectUpsert)) *ProjectUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&ProjectUpsert{UpdateSet: update})
	}))
	return u
}

// SetDescription sets the "description" field.
func (u *ProjectUpsertBulk) SetDescription(v string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetDescription(v)
	})
}

// UpdateDescription sets the "description" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateDescription() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateDescription()
	})
}

// ClearDescription clears the value of the "description" field.
func (u *ProjectUpsertBulk) ClearDescription() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearDescription()
	})
}

// SetLabels sets the "labels" field.
func (u *ProjectUpsertBulk) SetLabels(v map[string]string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetLabels(v)
	})
}

// UpdateLabels sets the "labels" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateLabels() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateLabels()
	})
}

// ClearLabels clears the value of the "labels" field.
func (u *ProjectUpsertBulk) ClearLabels() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearLabels()
	})
}

// SetAnnotations sets the "annotations" field.
func (u *ProjectUpsertBulk) SetAnnotations(v map[string]string) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetAnnotations(v)
	})
}

// UpdateAnnotations sets the "annotations" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateAnnotations() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateAnnotations()
	})
}

// ClearAnnotations clears the value of the "annotations" field.
func (u *ProjectUpsertBulk) ClearAnnotations() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.ClearAnnotations()
	})
}

// SetUpdateTime sets the "update_time" field.
func (u *ProjectUpsertBulk) SetUpdateTime(v time.Time) *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.SetUpdateTime(v)
	})
}

// UpdateUpdateTime sets the "update_time" field to the value that was provided on create.
func (u *ProjectUpsertBulk) UpdateUpdateTime() *ProjectUpsertBulk {
	return u.Update(func(s *ProjectUpsert) {
		s.UpdateUpdateTime()
	})
}

// Exec executes the query.
func (u *ProjectUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("model: OnConflict was set for builder %d. Set it on the ProjectCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("model: missing options for ProjectCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *ProjectUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
