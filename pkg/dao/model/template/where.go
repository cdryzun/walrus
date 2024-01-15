// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package template

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"

	"github.com/seal-io/walrus/pkg/dao/model/internal"
	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/types/object"
)

// ID filters vertices based on their ID field.
func ID(id object.ID) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id object.ID) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id object.ID) predicate.Template {
	return predicate.Template(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...object.ID) predicate.Template {
	return predicate.Template(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...object.ID) predicate.Template {
	return predicate.Template(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id object.ID) predicate.Template {
	return predicate.Template(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id object.ID) predicate.Template {
	return predicate.Template(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id object.ID) predicate.Template {
	return predicate.Template(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id object.ID) predicate.Template {
	return predicate.Template(sql.FieldLTE(FieldID, id))
}

// Name applies equality check predicate on the "name" field. It's identical to NameEQ.
func Name(v string) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldName, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldDescription, v))
}

// CreateTime applies equality check predicate on the "create_time" field. It's identical to CreateTimeEQ.
func CreateTime(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldCreateTime, v))
}

// UpdateTime applies equality check predicate on the "update_time" field. It's identical to UpdateTimeEQ.
func UpdateTime(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldUpdateTime, v))
}

// Icon applies equality check predicate on the "icon" field. It's identical to IconEQ.
func Icon(v string) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldIcon, v))
}

// Source applies equality check predicate on the "source" field. It's identical to SourceEQ.
func Source(v string) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldSource, v))
}

// CatalogID applies equality check predicate on the "catalog_id" field. It's identical to CatalogIDEQ.
func CatalogID(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldCatalogID, v))
}

// ProjectID applies equality check predicate on the "project_id" field. It's identical to ProjectIDEQ.
func ProjectID(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldProjectID, v))
}

// NameEQ applies the EQ predicate on the "name" field.
func NameEQ(v string) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldName, v))
}

// NameNEQ applies the NEQ predicate on the "name" field.
func NameNEQ(v string) predicate.Template {
	return predicate.Template(sql.FieldNEQ(FieldName, v))
}

// NameIn applies the In predicate on the "name" field.
func NameIn(vs ...string) predicate.Template {
	return predicate.Template(sql.FieldIn(FieldName, vs...))
}

// NameNotIn applies the NotIn predicate on the "name" field.
func NameNotIn(vs ...string) predicate.Template {
	return predicate.Template(sql.FieldNotIn(FieldName, vs...))
}

// NameGT applies the GT predicate on the "name" field.
func NameGT(v string) predicate.Template {
	return predicate.Template(sql.FieldGT(FieldName, v))
}

// NameGTE applies the GTE predicate on the "name" field.
func NameGTE(v string) predicate.Template {
	return predicate.Template(sql.FieldGTE(FieldName, v))
}

// NameLT applies the LT predicate on the "name" field.
func NameLT(v string) predicate.Template {
	return predicate.Template(sql.FieldLT(FieldName, v))
}

// NameLTE applies the LTE predicate on the "name" field.
func NameLTE(v string) predicate.Template {
	return predicate.Template(sql.FieldLTE(FieldName, v))
}

// NameContains applies the Contains predicate on the "name" field.
func NameContains(v string) predicate.Template {
	return predicate.Template(sql.FieldContains(FieldName, v))
}

// NameHasPrefix applies the HasPrefix predicate on the "name" field.
func NameHasPrefix(v string) predicate.Template {
	return predicate.Template(sql.FieldHasPrefix(FieldName, v))
}

// NameHasSuffix applies the HasSuffix predicate on the "name" field.
func NameHasSuffix(v string) predicate.Template {
	return predicate.Template(sql.FieldHasSuffix(FieldName, v))
}

// NameEqualFold applies the EqualFold predicate on the "name" field.
func NameEqualFold(v string) predicate.Template {
	return predicate.Template(sql.FieldEqualFold(FieldName, v))
}

// NameContainsFold applies the ContainsFold predicate on the "name" field.
func NameContainsFold(v string) predicate.Template {
	return predicate.Template(sql.FieldContainsFold(FieldName, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Template {
	return predicate.Template(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Template {
	return predicate.Template(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Template {
	return predicate.Template(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Template {
	return predicate.Template(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Template {
	return predicate.Template(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Template {
	return predicate.Template(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Template {
	return predicate.Template(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Template {
	return predicate.Template(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Template {
	return predicate.Template(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Template {
	return predicate.Template(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionIsNil applies the IsNil predicate on the "description" field.
func DescriptionIsNil() predicate.Template {
	return predicate.Template(sql.FieldIsNull(FieldDescription))
}

// DescriptionNotNil applies the NotNil predicate on the "description" field.
func DescriptionNotNil() predicate.Template {
	return predicate.Template(sql.FieldNotNull(FieldDescription))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Template {
	return predicate.Template(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Template {
	return predicate.Template(sql.FieldContainsFold(FieldDescription, v))
}

// LabelsIsNil applies the IsNil predicate on the "labels" field.
func LabelsIsNil() predicate.Template {
	return predicate.Template(sql.FieldIsNull(FieldLabels))
}

// LabelsNotNil applies the NotNil predicate on the "labels" field.
func LabelsNotNil() predicate.Template {
	return predicate.Template(sql.FieldNotNull(FieldLabels))
}

// CreateTimeEQ applies the EQ predicate on the "create_time" field.
func CreateTimeEQ(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldCreateTime, v))
}

// CreateTimeNEQ applies the NEQ predicate on the "create_time" field.
func CreateTimeNEQ(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldNEQ(FieldCreateTime, v))
}

// CreateTimeIn applies the In predicate on the "create_time" field.
func CreateTimeIn(vs ...time.Time) predicate.Template {
	return predicate.Template(sql.FieldIn(FieldCreateTime, vs...))
}

// CreateTimeNotIn applies the NotIn predicate on the "create_time" field.
func CreateTimeNotIn(vs ...time.Time) predicate.Template {
	return predicate.Template(sql.FieldNotIn(FieldCreateTime, vs...))
}

// CreateTimeGT applies the GT predicate on the "create_time" field.
func CreateTimeGT(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldGT(FieldCreateTime, v))
}

// CreateTimeGTE applies the GTE predicate on the "create_time" field.
func CreateTimeGTE(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldGTE(FieldCreateTime, v))
}

// CreateTimeLT applies the LT predicate on the "create_time" field.
func CreateTimeLT(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldLT(FieldCreateTime, v))
}

// CreateTimeLTE applies the LTE predicate on the "create_time" field.
func CreateTimeLTE(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldLTE(FieldCreateTime, v))
}

// UpdateTimeEQ applies the EQ predicate on the "update_time" field.
func UpdateTimeEQ(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldUpdateTime, v))
}

// UpdateTimeNEQ applies the NEQ predicate on the "update_time" field.
func UpdateTimeNEQ(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldNEQ(FieldUpdateTime, v))
}

// UpdateTimeIn applies the In predicate on the "update_time" field.
func UpdateTimeIn(vs ...time.Time) predicate.Template {
	return predicate.Template(sql.FieldIn(FieldUpdateTime, vs...))
}

// UpdateTimeNotIn applies the NotIn predicate on the "update_time" field.
func UpdateTimeNotIn(vs ...time.Time) predicate.Template {
	return predicate.Template(sql.FieldNotIn(FieldUpdateTime, vs...))
}

// UpdateTimeGT applies the GT predicate on the "update_time" field.
func UpdateTimeGT(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldGT(FieldUpdateTime, v))
}

// UpdateTimeGTE applies the GTE predicate on the "update_time" field.
func UpdateTimeGTE(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldGTE(FieldUpdateTime, v))
}

// UpdateTimeLT applies the LT predicate on the "update_time" field.
func UpdateTimeLT(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldLT(FieldUpdateTime, v))
}

// UpdateTimeLTE applies the LTE predicate on the "update_time" field.
func UpdateTimeLTE(v time.Time) predicate.Template {
	return predicate.Template(sql.FieldLTE(FieldUpdateTime, v))
}

// StatusIsNil applies the IsNil predicate on the "status" field.
func StatusIsNil() predicate.Template {
	return predicate.Template(sql.FieldIsNull(FieldStatus))
}

// StatusNotNil applies the NotNil predicate on the "status" field.
func StatusNotNil() predicate.Template {
	return predicate.Template(sql.FieldNotNull(FieldStatus))
}

// IconEQ applies the EQ predicate on the "icon" field.
func IconEQ(v string) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldIcon, v))
}

// IconNEQ applies the NEQ predicate on the "icon" field.
func IconNEQ(v string) predicate.Template {
	return predicate.Template(sql.FieldNEQ(FieldIcon, v))
}

// IconIn applies the In predicate on the "icon" field.
func IconIn(vs ...string) predicate.Template {
	return predicate.Template(sql.FieldIn(FieldIcon, vs...))
}

// IconNotIn applies the NotIn predicate on the "icon" field.
func IconNotIn(vs ...string) predicate.Template {
	return predicate.Template(sql.FieldNotIn(FieldIcon, vs...))
}

// IconGT applies the GT predicate on the "icon" field.
func IconGT(v string) predicate.Template {
	return predicate.Template(sql.FieldGT(FieldIcon, v))
}

// IconGTE applies the GTE predicate on the "icon" field.
func IconGTE(v string) predicate.Template {
	return predicate.Template(sql.FieldGTE(FieldIcon, v))
}

// IconLT applies the LT predicate on the "icon" field.
func IconLT(v string) predicate.Template {
	return predicate.Template(sql.FieldLT(FieldIcon, v))
}

// IconLTE applies the LTE predicate on the "icon" field.
func IconLTE(v string) predicate.Template {
	return predicate.Template(sql.FieldLTE(FieldIcon, v))
}

// IconContains applies the Contains predicate on the "icon" field.
func IconContains(v string) predicate.Template {
	return predicate.Template(sql.FieldContains(FieldIcon, v))
}

// IconHasPrefix applies the HasPrefix predicate on the "icon" field.
func IconHasPrefix(v string) predicate.Template {
	return predicate.Template(sql.FieldHasPrefix(FieldIcon, v))
}

// IconHasSuffix applies the HasSuffix predicate on the "icon" field.
func IconHasSuffix(v string) predicate.Template {
	return predicate.Template(sql.FieldHasSuffix(FieldIcon, v))
}

// IconIsNil applies the IsNil predicate on the "icon" field.
func IconIsNil() predicate.Template {
	return predicate.Template(sql.FieldIsNull(FieldIcon))
}

// IconNotNil applies the NotNil predicate on the "icon" field.
func IconNotNil() predicate.Template {
	return predicate.Template(sql.FieldNotNull(FieldIcon))
}

// IconEqualFold applies the EqualFold predicate on the "icon" field.
func IconEqualFold(v string) predicate.Template {
	return predicate.Template(sql.FieldEqualFold(FieldIcon, v))
}

// IconContainsFold applies the ContainsFold predicate on the "icon" field.
func IconContainsFold(v string) predicate.Template {
	return predicate.Template(sql.FieldContainsFold(FieldIcon, v))
}

// SourceEQ applies the EQ predicate on the "source" field.
func SourceEQ(v string) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldSource, v))
}

// SourceNEQ applies the NEQ predicate on the "source" field.
func SourceNEQ(v string) predicate.Template {
	return predicate.Template(sql.FieldNEQ(FieldSource, v))
}

// SourceIn applies the In predicate on the "source" field.
func SourceIn(vs ...string) predicate.Template {
	return predicate.Template(sql.FieldIn(FieldSource, vs...))
}

// SourceNotIn applies the NotIn predicate on the "source" field.
func SourceNotIn(vs ...string) predicate.Template {
	return predicate.Template(sql.FieldNotIn(FieldSource, vs...))
}

// SourceGT applies the GT predicate on the "source" field.
func SourceGT(v string) predicate.Template {
	return predicate.Template(sql.FieldGT(FieldSource, v))
}

// SourceGTE applies the GTE predicate on the "source" field.
func SourceGTE(v string) predicate.Template {
	return predicate.Template(sql.FieldGTE(FieldSource, v))
}

// SourceLT applies the LT predicate on the "source" field.
func SourceLT(v string) predicate.Template {
	return predicate.Template(sql.FieldLT(FieldSource, v))
}

// SourceLTE applies the LTE predicate on the "source" field.
func SourceLTE(v string) predicate.Template {
	return predicate.Template(sql.FieldLTE(FieldSource, v))
}

// SourceContains applies the Contains predicate on the "source" field.
func SourceContains(v string) predicate.Template {
	return predicate.Template(sql.FieldContains(FieldSource, v))
}

// SourceHasPrefix applies the HasPrefix predicate on the "source" field.
func SourceHasPrefix(v string) predicate.Template {
	return predicate.Template(sql.FieldHasPrefix(FieldSource, v))
}

// SourceHasSuffix applies the HasSuffix predicate on the "source" field.
func SourceHasSuffix(v string) predicate.Template {
	return predicate.Template(sql.FieldHasSuffix(FieldSource, v))
}

// SourceEqualFold applies the EqualFold predicate on the "source" field.
func SourceEqualFold(v string) predicate.Template {
	return predicate.Template(sql.FieldEqualFold(FieldSource, v))
}

// SourceContainsFold applies the ContainsFold predicate on the "source" field.
func SourceContainsFold(v string) predicate.Template {
	return predicate.Template(sql.FieldContainsFold(FieldSource, v))
}

// CatalogIDEQ applies the EQ predicate on the "catalog_id" field.
func CatalogIDEQ(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldCatalogID, v))
}

// CatalogIDNEQ applies the NEQ predicate on the "catalog_id" field.
func CatalogIDNEQ(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldNEQ(FieldCatalogID, v))
}

// CatalogIDIn applies the In predicate on the "catalog_id" field.
func CatalogIDIn(vs ...object.ID) predicate.Template {
	return predicate.Template(sql.FieldIn(FieldCatalogID, vs...))
}

// CatalogIDNotIn applies the NotIn predicate on the "catalog_id" field.
func CatalogIDNotIn(vs ...object.ID) predicate.Template {
	return predicate.Template(sql.FieldNotIn(FieldCatalogID, vs...))
}

// CatalogIDGT applies the GT predicate on the "catalog_id" field.
func CatalogIDGT(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldGT(FieldCatalogID, v))
}

// CatalogIDGTE applies the GTE predicate on the "catalog_id" field.
func CatalogIDGTE(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldGTE(FieldCatalogID, v))
}

// CatalogIDLT applies the LT predicate on the "catalog_id" field.
func CatalogIDLT(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldLT(FieldCatalogID, v))
}

// CatalogIDLTE applies the LTE predicate on the "catalog_id" field.
func CatalogIDLTE(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldLTE(FieldCatalogID, v))
}

// CatalogIDContains applies the Contains predicate on the "catalog_id" field.
func CatalogIDContains(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldContains(FieldCatalogID, vc))
}

// CatalogIDHasPrefix applies the HasPrefix predicate on the "catalog_id" field.
func CatalogIDHasPrefix(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldHasPrefix(FieldCatalogID, vc))
}

// CatalogIDHasSuffix applies the HasSuffix predicate on the "catalog_id" field.
func CatalogIDHasSuffix(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldHasSuffix(FieldCatalogID, vc))
}

// CatalogIDIsNil applies the IsNil predicate on the "catalog_id" field.
func CatalogIDIsNil() predicate.Template {
	return predicate.Template(sql.FieldIsNull(FieldCatalogID))
}

// CatalogIDNotNil applies the NotNil predicate on the "catalog_id" field.
func CatalogIDNotNil() predicate.Template {
	return predicate.Template(sql.FieldNotNull(FieldCatalogID))
}

// CatalogIDEqualFold applies the EqualFold predicate on the "catalog_id" field.
func CatalogIDEqualFold(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldEqualFold(FieldCatalogID, vc))
}

// CatalogIDContainsFold applies the ContainsFold predicate on the "catalog_id" field.
func CatalogIDContainsFold(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldContainsFold(FieldCatalogID, vc))
}

// ProjectIDEQ applies the EQ predicate on the "project_id" field.
func ProjectIDEQ(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldEQ(FieldProjectID, v))
}

// ProjectIDNEQ applies the NEQ predicate on the "project_id" field.
func ProjectIDNEQ(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldNEQ(FieldProjectID, v))
}

// ProjectIDIn applies the In predicate on the "project_id" field.
func ProjectIDIn(vs ...object.ID) predicate.Template {
	return predicate.Template(sql.FieldIn(FieldProjectID, vs...))
}

// ProjectIDNotIn applies the NotIn predicate on the "project_id" field.
func ProjectIDNotIn(vs ...object.ID) predicate.Template {
	return predicate.Template(sql.FieldNotIn(FieldProjectID, vs...))
}

// ProjectIDGT applies the GT predicate on the "project_id" field.
func ProjectIDGT(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldGT(FieldProjectID, v))
}

// ProjectIDGTE applies the GTE predicate on the "project_id" field.
func ProjectIDGTE(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldGTE(FieldProjectID, v))
}

// ProjectIDLT applies the LT predicate on the "project_id" field.
func ProjectIDLT(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldLT(FieldProjectID, v))
}

// ProjectIDLTE applies the LTE predicate on the "project_id" field.
func ProjectIDLTE(v object.ID) predicate.Template {
	return predicate.Template(sql.FieldLTE(FieldProjectID, v))
}

// ProjectIDContains applies the Contains predicate on the "project_id" field.
func ProjectIDContains(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldContains(FieldProjectID, vc))
}

// ProjectIDHasPrefix applies the HasPrefix predicate on the "project_id" field.
func ProjectIDHasPrefix(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldHasPrefix(FieldProjectID, vc))
}

// ProjectIDHasSuffix applies the HasSuffix predicate on the "project_id" field.
func ProjectIDHasSuffix(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldHasSuffix(FieldProjectID, vc))
}

// ProjectIDIsNil applies the IsNil predicate on the "project_id" field.
func ProjectIDIsNil() predicate.Template {
	return predicate.Template(sql.FieldIsNull(FieldProjectID))
}

// ProjectIDNotNil applies the NotNil predicate on the "project_id" field.
func ProjectIDNotNil() predicate.Template {
	return predicate.Template(sql.FieldNotNull(FieldProjectID))
}

// ProjectIDEqualFold applies the EqualFold predicate on the "project_id" field.
func ProjectIDEqualFold(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldEqualFold(FieldProjectID, vc))
}

// ProjectIDContainsFold applies the ContainsFold predicate on the "project_id" field.
func ProjectIDContainsFold(v object.ID) predicate.Template {
	vc := string(v)
	return predicate.Template(sql.FieldContainsFold(FieldProjectID, vc))
}

// HasVersions applies the HasEdge predicate on the "versions" edge.
func HasVersions() predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, VersionsTable, VersionsColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TemplateVersion
		step.Edge.Schema = schemaConfig.TemplateVersion
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasVersionsWith applies the HasEdge predicate on the "versions" edge with a given conditions (other predicates).
func HasVersionsWith(preds ...predicate.TemplateVersion) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		step := newVersionsStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.TemplateVersion
		step.Edge.Schema = schemaConfig.TemplateVersion
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCatalog applies the HasEdge predicate on the "catalog" edge.
func HasCatalog() predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CatalogTable, CatalogColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Catalog
		step.Edge.Schema = schemaConfig.Template
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCatalogWith applies the HasEdge predicate on the "catalog" edge with a given conditions (other predicates).
func HasCatalogWith(preds ...predicate.Catalog) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		step := newCatalogStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Catalog
		step.Edge.Schema = schemaConfig.Template
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasProject applies the HasEdge predicate on the "project" edge.
func HasProject() predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, ProjectTable, ProjectColumn),
		)
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.Template
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasProjectWith applies the HasEdge predicate on the "project" edge with a given conditions (other predicates).
func HasProjectWith(preds ...predicate.Project) predicate.Template {
	return predicate.Template(func(s *sql.Selector) {
		step := newProjectStep()
		schemaConfig := internal.SchemaConfigFromContext(s.Context())
		step.To.Schema = schemaConfig.Project
		step.Edge.Schema = schemaConfig.Template
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Template) predicate.Template {
	return predicate.Template(sql.AndPredicates(predicates...))
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Template) predicate.Template {
	return predicate.Template(sql.OrPredicates(predicates...))
}

// Not applies the not operator on the given predicate.
func Not(p predicate.Template) predicate.Template {
	return predicate.Template(sql.NotPredicates(p))
}
