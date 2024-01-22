// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package resourcecomponentrelationship

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ID filters vertices based on their ID field.
func ID(id object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLTE(FieldID, id))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldCreateTime, v))
}

// ResourceComponentID applies equality check predicate on the "resource_component_id" field. It's identical to ResourceComponentIDEQ.
func ResourceComponentID(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldResourceComponentID, v))
}

// DependencyID applies equality check predicate on the "dependency_id" field. It's identical to DependencyIDEQ.
func DependencyID(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldDependencyID, v))
}

// Type applies equality check predicate on the "type" field. It's identical to TypeEQ.
func Type(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldType, v))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLTE(FieldCreateTime, v))
}

// ResourceComponentIDEQ applies the EQ predicate on the "resource_component_id" field.
func ResourceComponentIDEQ(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldResourceComponentID, v))
}

// ResourceComponentIDNEQ applies the NEQ predicate on the "resource_component_id" field.
func ResourceComponentIDNEQ(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNEQ(FieldResourceComponentID, v))
}

// ResourceComponentIDIn applies the In predicate on the "resource_component_id" field.
func ResourceComponentIDIn(vs ...object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldIn(FieldResourceComponentID, vs...))
}

// ResourceComponentIDNotIn applies the NotIn predicate on the "resource_component_id" field.
func ResourceComponentIDNotIn(vs ...object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNotIn(FieldResourceComponentID, vs...))
}

// ResourceComponentIDGT applies the GT predicate on the "resource_component_id" field.
func ResourceComponentIDGT(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGT(FieldResourceComponentID, v))
}

// ResourceComponentIDGTE applies the GTE predicate on the "resource_component_id" field.
func ResourceComponentIDGTE(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGTE(FieldResourceComponentID, v))
}

// ResourceComponentIDLT applies the LT predicate on the "resource_component_id" field.
func ResourceComponentIDLT(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLT(FieldResourceComponentID, v))
}

// ResourceComponentIDLTE applies the LTE predicate on the "resource_component_id" field.
func ResourceComponentIDLTE(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLTE(FieldResourceComponentID, v))
}

// ResourceComponentIDContains applies the Contains predicate on the "resource_component_id" field.
func ResourceComponentIDContains(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldContains(FieldResourceComponentID, vc))
}

// ResourceComponentIDHasPrefix applies the HasPrefix predicate on the "resource_component_id" field.
func ResourceComponentIDHasPrefix(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldHasPrefix(FieldResourceComponentID, vc))
}

// ResourceComponentIDHasSuffix applies the HasSuffix predicate on the "resource_component_id" field.
func ResourceComponentIDHasSuffix(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldHasSuffix(FieldResourceComponentID, vc))
}

// ResourceComponentIDEqualFold applies the EqualFold predicate on the "resource_component_id" field.
func ResourceComponentIDEqualFold(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldEqualFold(FieldResourceComponentID, vc))
}

// ResourceComponentIDContainsFold applies the ContainsFold predicate on the "resource_component_id" field.
func ResourceComponentIDContainsFold(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldContainsFold(FieldResourceComponentID, vc))
}

// DependencyIDEQ applies the EQ predicate on the "dependency_id" field.
func DependencyIDEQ(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldDependencyID, v))
}

// DependencyIDNEQ applies the NEQ predicate on the "dependency_id" field.
func DependencyIDNEQ(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNEQ(FieldDependencyID, v))
}

// DependencyIDIn applies the In predicate on the "dependency_id" field.
func DependencyIDIn(vs ...object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldIn(FieldDependencyID, vs...))
}

// DependencyIDNotIn applies the NotIn predicate on the "dependency_id" field.
func DependencyIDNotIn(vs ...object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNotIn(FieldDependencyID, vs...))
}

// DependencyIDGT applies the GT predicate on the "dependency_id" field.
func DependencyIDGT(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGT(FieldDependencyID, v))
}

// DependencyIDGTE applies the GTE predicate on the "dependency_id" field.
func DependencyIDGTE(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGTE(FieldDependencyID, v))
}

// DependencyIDLT applies the LT predicate on the "dependency_id" field.
func DependencyIDLT(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLT(FieldDependencyID, v))
}

// DependencyIDLTE applies the LTE predicate on the "dependency_id" field.
func DependencyIDLTE(v object.ID) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLTE(FieldDependencyID, v))
}

// DependencyIDContains applies the Contains predicate on the "dependency_id" field.
func DependencyIDContains(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldContains(FieldDependencyID, vc))
}

// DependencyIDHasPrefix applies the HasPrefix predicate on the "dependency_id" field.
func DependencyIDHasPrefix(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldHasPrefix(FieldDependencyID, vc))
}

// DependencyIDHasSuffix applies the HasSuffix predicate on the "dependency_id" field.
func DependencyIDHasSuffix(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldHasSuffix(FieldDependencyID, vc))
}

// DependencyIDEqualFold applies the EqualFold predicate on the "dependency_id" field.
func DependencyIDEqualFold(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldEqualFold(FieldDependencyID, vc))
}

// DependencyIDContainsFold applies the ContainsFold predicate on the "dependency_id" field.
func DependencyIDContainsFold(v object.ID) predicate.ResourceComponentRelationship {
	vc := string(v)
	return predicate.ResourceComponentRelationship(sql.FieldContainsFold(FieldDependencyID, vc))
}

// TypeEQ applies the EQ predicate on the "type" field.
func TypeEQ(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEQ(FieldType, v))
}

// TypeNEQ applies the NEQ predicate on the "type" field.
func TypeNEQ(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNEQ(FieldType, v))
}

// TypeIn applies the In predicate on the "type" field.
func TypeIn(vs ...string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldIn(FieldType, vs...))
}

// TypeNotIn applies the NotIn predicate on the "type" field.
func TypeNotIn(vs ...string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldNotIn(FieldType, vs...))
}

// TypeGT applies the GT predicate on the "type" field.
func TypeGT(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGT(FieldType, v))
}

// TypeGTE applies the GTE predicate on the "type" field.
func TypeGTE(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldGTE(FieldType, v))
}

// TypeLT applies the LT predicate on the "type" field.
func TypeLT(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLT(FieldType, v))
}

// TypeLTE applies the LTE predicate on the "type" field.
func TypeLTE(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldLTE(FieldType, v))
}

// TypeContains applies the Contains predicate on the "type" field.
func TypeContains(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldContains(FieldType, v))
}

// TypeHasPrefix applies the HasPrefix predicate on the "type" field.
func TypeHasPrefix(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldHasPrefix(FieldType, v))
}

// TypeHasSuffix applies the HasSuffix predicate on the "type" field.
func TypeHasSuffix(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldHasSuffix(FieldType, v))
}

// TypeEqualFold applies the EqualFold predicate on the "type" field.
func TypeEqualFold(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldEqualFold(FieldType, v))
}

// TypeContainsFold applies the ContainsFold predicate on the "type" field.
func TypeContainsFold(v string) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(sql.FieldContainsFold(FieldType, v))
}

// HasResourceComponent applies the HasEdge predicate on the "resource_component" edge.
func HasResourceComponent() predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, ResourceComponentTable, ResourceComponentColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ResourceComponent
		step.Edge.Schema = schemaConfig.ResourceComponentRelationship
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasResourceComponentWith applies the HasEdge predicate on the "resource_component" edge with a given conditions (other predicates).
func HasResourceComponentWith(preds ...predicate.ResourceComponent) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(func(s *sql.Selector) {
		step := newResourceComponentStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ResourceComponent
		step.Edge.Schema = schemaConfig.ResourceComponentRelationship
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasDependency applies the HasEdge predicate on the "dependency" edge.
func HasDependency() predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, DependencyTable, DependencyColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ResourceComponent
		step.Edge.Schema = schemaConfig.ResourceComponentRelationship
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasDependencyWith applies the HasEdge predicate on the "dependency" edge with a given conditions (other predicates).
func HasDependencyWith(preds ...predicate.ResourceComponent) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(func(s *sql.Selector) {
		step := newDependencyStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.ResourceComponent
		step.Edge.Schema = schemaConfig.ResourceComponentRelationship
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.ResourceComponentRelationship) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.ResourceComponentRelationship) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.ResourceComponentRelationship) predicate.ResourceComponentRelationship {
	return predicate.ResourceComponentRelationship(func(s *sql.Selector) {
		p(s.Not())
	})
}
