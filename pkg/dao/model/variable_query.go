// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"

	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/variable"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// VariableQuery is the builder for querying Variable entities.
type VariableQuery struct {
	config
	ctx             *QueryContext
	order           []variable.OrderOption
	inters          []Interceptor
	predicates      []predicate.Variable
	withProject     *ProjectQuery
	withEnvironment *EnvironmentQuery
	modifiers       []func(*sql.Selector)
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the VariableQuery builder.
func (vq *VariableQuery) Where(ps ...predicate.Variable) *VariableQuery {
	vq.predicates = append(vq.predicates, ps...)
	return vq
}

// Limit the number of records to be returned by this query.
func (vq *VariableQuery) Limit(limit int) *VariableQuery {
	vq.ctx.Limit = &limit
	return vq
}

// Offset to start from.
func (vq *VariableQuery) Offset(offset int) *VariableQuery {
	vq.ctx.Offset = &offset
	return vq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (vq *VariableQuery) Unique(unique bool) *VariableQuery {
	vq.ctx.Unique = &unique
	return vq
}

// Order specifies how the records should be ordered.
func (vq *VariableQuery) Order(o ...variable.OrderOption) *VariableQuery {
	vq.order = append(vq.order, o...)
	return vq
}

// QueryProject chains the current query on the "project" edge.
func (vq *VariableQuery) QueryProject() *ProjectQuery {
	query := (&ProjectClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(variable.Table, variable.FieldID, selector),
			sqlgraph.To(project.Table, project.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, variable.ProjectTable, variable.ProjectColumn),
		)
		schemaConfig := vq.schemaConfig
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.Variable
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryEnvironment chains the current query on the "environment" edge.
func (vq *VariableQuery) QueryEnvironment() *EnvironmentQuery {
	query := (&EnvironmentClient{config: vq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := vq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := vq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(variable.Table, variable.FieldID, selector),
			sqlgraph.To(environment.Table, environment.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, variable.EnvironmentTable, variable.EnvironmentColumn),
		)
		schemaConfig := vq.schemaConfig
		step.To.Schema = schemaConfig.Environment
		step.Edge.Schema = schemaConfig.Variable
		fromU = sqlgraph.SetNeighbors(vq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Variable entity from the query.
// Returns a *NotFoundError when no Variable was found.
func (vq *VariableQuery) First(ctx context.Context) (*Variable, error) {
	nodes, err := vq.Limit(1).All(setContextOp(ctx, vq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{variable.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (vq *VariableQuery) FirstX(ctx context.Context) *Variable {
	node, err := vq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Variable ID from the query.
// Returns a *NotFoundError when no Variable ID was found.
func (vq *VariableQuery) FirstID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = vq.Limit(1).IDs(setContextOp(ctx, vq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{variable.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (vq *VariableQuery) FirstIDX(ctx context.Context) object.ID {
	id, err := vq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Variable entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Variable entity is found.
// Returns a *NotFoundError when no Variable entities are found.
func (vq *VariableQuery) Only(ctx context.Context) (*Variable, error) {
	nodes, err := vq.Limit(2).All(setContextOp(ctx, vq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{variable.Label}
	default:
		return nil, &NotSingularError{variable.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (vq *VariableQuery) OnlyX(ctx context.Context) *Variable {
	node, err := vq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Variable ID in the query.
// Returns a *NotSingularError when more than one Variable ID is found.
// Returns a *NotFoundError when no entities are found.
func (vq *VariableQuery) OnlyID(ctx context.Context) (id object.ID, err error) {
	var ids []object.ID
	if ids, err = vq.Limit(2).IDs(setContextOp(ctx, vq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{variable.Label}
	default:
		err = &NotSingularError{variable.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (vq *VariableQuery) OnlyIDX(ctx context.Context) object.ID {
	id, err := vq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Variables.
func (vq *VariableQuery) All(ctx context.Context) ([]*Variable, error) {
	ctx = setContextOp(ctx, vq.ctx, "All")
	if err := vq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Variable, *VariableQuery]()
	return withInterceptors[[]*Variable](ctx, vq, qr, vq.inters)
}

// AllX is like All, but panics if an error occurs.
func (vq *VariableQuery) AllX(ctx context.Context) []*Variable {
	nodes, err := vq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Variable IDs.
func (vq *VariableQuery) IDs(ctx context.Context) (ids []object.ID, err error) {
	if vq.ctx.Unique == nil && vq.path != nil {
		vq.Unique(true)
	}
	ctx = setContextOp(ctx, vq.ctx, "IDs")
	if err = vq.Select(variable.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (vq *VariableQuery) IDsX(ctx context.Context) []object.ID {
	ids, err := vq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (vq *VariableQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, vq.ctx, "Count")
	if err := vq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, vq, querierCount[*VariableQuery](), vq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (vq *VariableQuery) CountX(ctx context.Context) int {
	count, err := vq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (vq *VariableQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, vq.ctx, "Exist")
	switch _, err := vq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("model: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (vq *VariableQuery) ExistX(ctx context.Context) bool {
	exist, err := vq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the VariableQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (vq *VariableQuery) Clone() *VariableQuery {
	if vq == nil {
		return nil
	}
	return &VariableQuery{
		config:          vq.config,
		ctx:             vq.ctx.Clone(),
		order:           append([]variable.OrderOption{}, vq.order...),
		inters:          append([]Interceptor{}, vq.inters...),
		predicates:      append([]predicate.Variable{}, vq.predicates...),
		withProject:     vq.withProject.Clone(),
		withEnvironment: vq.withEnvironment.Clone(),
		// clone intermediate query.
		sql:  vq.sql.Clone(),
		path: vq.path,
	}
}

// WithProject tells the query-builder to eager-load the nodes that are connected to
// the "project" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VariableQuery) WithProject(opts ...func(*ProjectQuery)) *VariableQuery {
	query := (&ProjectClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withProject = query
	return vq
}

// WithEnvironment tells the query-builder to eager-load the nodes that are connected to
// the "environment" edge. The optional arguments are used to configure the query builder of the edge.
func (vq *VariableQuery) WithEnvironment(opts ...func(*EnvironmentQuery)) *VariableQuery {
	query := (&EnvironmentClient{config: vq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	vq.withEnvironment = query
	return vq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Variable.Query().
//		GroupBy(variable.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (vq *VariableQuery) GroupBy(field string, fields ...string) *VariableGroupBy {
	vq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &VariableGroupBy{build: vq}
	grbuild.flds = &vq.ctx.Fields
	grbuild.label = variable.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.Variable.Query().
//		Select(variable.FieldCreateTime).
//		Scan(ctx, &v)
func (vq *VariableQuery) Select(fields ...string) *VariableSelect {
	vq.ctx.Fields = append(vq.ctx.Fields, fields...)
	sbuild := &VariableSelect{VariableQuery: vq}
	sbuild.label = variable.Label
	sbuild.flds, sbuild.scan = &vq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a VariableSelect configured with the given aggregations.
func (vq *VariableQuery) Aggregate(fns ...AggregateFunc) *VariableSelect {
	return vq.Select().Aggregate(fns...)
}

func (vq *VariableQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range vq.inters {
		if inter == nil {
			return fmt.Errorf("model: uninitialized interceptor (forgotten import model/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, vq); err != nil {
				return err
			}
		}
	}
	for _, f := range vq.ctx.Fields {
		if !variable.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if vq.path != nil {
		prev, err := vq.path(ctx)
		if err != nil {
			return err
		}
		vq.sql = prev
	}
	return nil
}

func (vq *VariableQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Variable, error) {
	var (
		nodes       = []*Variable{}
		_spec       = vq.querySpec()
		loadedTypes = [2]bool{
			vq.withProject != nil,
			vq.withEnvironment != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Variable).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Variable{config: vq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	_spec.Node.Schema = vq.schemaConfig.Variable
	ctx = internal.NewSchemaConfigContext(ctx, vq.schemaConfig)
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, vq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := vq.withProject; query != nil {
		if err := vq.loadProject(ctx, query, nodes, nil,
			func(n *Variable, e *Project) { n.Edges.Project = e }); err != nil {
			return nil, err
		}
	}
	if query := vq.withEnvironment; query != nil {
		if err := vq.loadEnvironment(ctx, query, nodes, nil,
			func(n *Variable, e *Environment) { n.Edges.Environment = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (vq *VariableQuery) loadProject(ctx context.Context, query *ProjectQuery, nodes []*Variable, init func(*Variable), assign func(*Variable, *Project)) error {
	ids := make([]object.ID, 0, len(nodes))
	nodeids := make(map[object.ID][]*Variable)
	for i := range nodes {
		fk := nodes[i].ProjectID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(project.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "project_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (vq *VariableQuery) loadEnvironment(ctx context.Context, query *EnvironmentQuery, nodes []*Variable, init func(*Variable), assign func(*Variable, *Environment)) error {
	ids := make([]object.ID, 0, len(nodes))
	nodeids := make(map[object.ID][]*Variable)
	for i := range nodes {
		fk := nodes[i].EnvironmentID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(environment.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "environment_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (vq *VariableQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := vq.querySpec()
	_spec.Node.Schema = vq.schemaConfig.Variable
	ctx = internal.NewSchemaConfigContext(ctx, vq.schemaConfig)
	if len(vq.modifiers) > 0 {
		_spec.Modifiers = vq.modifiers
	}
	_spec.Node.Columns = vq.ctx.Fields
	if len(vq.ctx.Fields) > 0 {
		_spec.Unique = vq.ctx.Unique != nil && *vq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, vq.driver, _spec)
}

func (vq *VariableQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(variable.Table, variable.Columns, sqlgraph.NewFieldSpec(variable.FieldID, field.TypeString))
	_spec.From = vq.sql
	if unique := vq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if vq.path != nil {
		_spec.Unique = true
	}
	if fields := vq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, variable.FieldID)
		for i := range fields {
			if fields[i] != variable.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if vq.withProject != nil {
			_spec.Node.AddColumnOnce(variable.FieldProjectID)
		}
		if vq.withEnvironment != nil {
			_spec.Node.AddColumnOnce(variable.FieldEnvironmentID)
		}
	}
	if ps := vq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := vq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := vq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := vq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (vq *VariableQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(vq.driver.Dialect())
	t1 := builder.Table(variable.Table)
	columns := vq.ctx.Fields
	if len(columns) == 0 {
		columns = variable.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if vq.sql != nil {
		selector = vq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if vq.ctx.Unique != nil && *vq.ctx.Unique {
		selector.Distinct()
	}
	t1.Schema(vq.schemaConfig.Variable)
	ctx = internal.NewSchemaConfigContext(ctx, vq.schemaConfig)
	selector.WithContext(ctx)
	for _, m := range vq.modifiers {
		m(selector)
	}
	for _, p := range vq.predicates {
		p(selector)
	}
	for _, p := range vq.order {
		p(selector)
	}
	if offset := vq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := vq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ForUpdate locks the selected rows against concurrent updates, and prevent them from being
// updated, deleted or "selected ... for update" by other sessions, until the transaction is
// either committed or rolled-back.
func (vq *VariableQuery) ForUpdate(opts ...sql.LockOption) *VariableQuery {
	if vq.driver.Dialect() == dialect.Postgres {
		vq.Unique(false)
	}
	vq.modifiers = append(vq.modifiers, func(s *sql.Selector) {
		s.ForUpdate(opts...)
	})
	return vq
}

// ForShare behaves similarly to ForUpdate, except that it acquires a shared mode lock
// on any rows that are read. Other sessions can read the rows, but cannot modify them
// until your transaction commits.
func (vq *VariableQuery) ForShare(opts ...sql.LockOption) *VariableQuery {
	if vq.driver.Dialect() == dialect.Postgres {
		vq.Unique(false)
	}
	vq.modifiers = append(vq.modifiers, func(s *sql.Selector) {
		s.ForShare(opts...)
	})
	return vq
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vq *VariableQuery) Modify(modifiers ...func(s *sql.Selector)) *VariableSelect {
	vq.modifiers = append(vq.modifiers, modifiers...)
	return vq.Select()
}

// WhereP appends storage-level predicates to the VariableQuery builder. Using this method,
// users can use type-assertion to append predicates that do not depend on any generated package.
func (vq *VariableQuery) WhereP(ps ...func(*sql.Selector)) {
	var wps = make([]predicate.Variable, 0, len(ps))
	for i := 0; i < len(ps); i++ {
		wps = append(wps, predicate.Variable(ps[i]))
	}
	vq.predicates = append(vq.predicates, wps...)
}

// VariableGroupBy is the group-by builder for Variable entities.
type VariableGroupBy struct {
	selector
	build *VariableQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (vgb *VariableGroupBy) Aggregate(fns ...AggregateFunc) *VariableGroupBy {
	vgb.fns = append(vgb.fns, fns...)
	return vgb
}

// Scan applies the selector query and scans the result into the given value.
func (vgb *VariableGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vgb.build.ctx, "GroupBy")
	if err := vgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VariableQuery, *VariableGroupBy](ctx, vgb.build, vgb, vgb.build.inters, v)
}

func (vgb *VariableGroupBy) sqlScan(ctx context.Context, root *VariableQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(vgb.fns))
	for _, fn := range vgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*vgb.flds)+len(vgb.fns))
		for _, f := range *vgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*vgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// VariableSelect is the builder for selecting fields of Variable entities.
type VariableSelect struct {
	*VariableQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (vs *VariableSelect) Aggregate(fns ...AggregateFunc) *VariableSelect {
	vs.fns = append(vs.fns, fns...)
	return vs
}

// Scan applies the selector query and scans the result into the given value.
func (vs *VariableSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, vs.ctx, "Select")
	if err := vs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*VariableQuery, *VariableSelect](ctx, vs.VariableQuery, vs, vs.inters, v)
}

func (vs *VariableSelect) sqlScan(ctx context.Context, root *VariableQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(vs.fns))
	for _, fn := range vs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*vs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := vs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// Modify adds a query modifier for attaching custom logic to queries.
func (vs *VariableSelect) Modify(modifiers ...func(s *sql.Selector)) *VariableSelect {
	vs.modifiers = append(vs.modifiers, modifiers...)
	return vs
}
