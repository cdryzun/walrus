// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package role

import (
	"time"

	"entgo.io/ent"

	"github.com/seal-io/seal/pkg/dao/schema"
)

const (
	// Label holds the string label denoting the role type in the database.
	Label = "role"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateTime holds the string denoting the createtime field in the database.
	FieldCreateTime = "create_time"
	// FieldUpdateTime holds the string denoting the updatetime field in the database.
	FieldUpdateTime = "update_time"
	// FieldDomain holds the string denoting the domain field in the database.
	FieldDomain = "domain"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDescription holds the string denoting the description field in the database.
	FieldDescription = "description"
	// FieldPolicies holds the string denoting the policies field in the database.
	FieldPolicies = "policies"
	// FieldBuiltin holds the string denoting the builtin field in the database.
	FieldBuiltin = "builtin"
	// FieldSession holds the string denoting the session field in the database.
	FieldSession = "session"
	// Table holds the table name of the role in the database.
	Table = "roles"
)

// Columns holds all SQL columns for role fields.
var Columns = []string{
	FieldID,
	FieldCreateTime,
	FieldUpdateTime,
	FieldDomain,
	FieldName,
	FieldDescription,
	FieldPolicies,
	FieldBuiltin,
	FieldSession,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

// Note that the variables below are initialized by the runtime
// package on the initialization of the application. Therefore,
// it should be imported in the main as follows:
//
//	import _ "github.com/seal-io/seal/pkg/dao/model/runtime"
var (
	Hooks [1]ent.Hook
	// DefaultCreateTime holds the default value on creation for the "createTime" field.
	DefaultCreateTime func() time.Time
	// DefaultUpdateTime holds the default value on creation for the "updateTime" field.
	DefaultUpdateTime func() time.Time
	// UpdateDefaultUpdateTime holds the default value on update for the "updateTime" field.
	UpdateDefaultUpdateTime func() time.Time
	// DefaultDomain holds the default value on creation for the "domain" field.
	DefaultDomain string
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultPolicies holds the default value on creation for the "policies" field.
	DefaultPolicies schema.RolePolicies
	// DefaultBuiltin holds the default value on creation for the "builtin" field.
	DefaultBuiltin bool
	// DefaultSession holds the default value on creation for the "session" field.
	DefaultSession bool
)

// WithoutFields returns the fields ignored the given list.
func WithoutFields(ignores ...string) []string {
	if len(ignores) == 0 {
		return Columns
	}

	var s = make(map[string]bool, len(ignores))
	for i := range ignores {
		s[ignores[i]] = true
	}

	var r = make([]string, 0, len(Columns)-len(s))
	for i := range Columns {
		if s[Columns[i]] {
			continue
		}
		r = append(r, Columns[i])
	}
	return r
}
