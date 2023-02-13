// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// GENERATED, DO NOT EDIT.

package predicate

import (
	"entgo.io/ent/dialect/sql"
)

// Application is the predicate function for application builders.
type Application func(*sql.Selector)

// ApplicationResource is the predicate function for applicationresource builders.
type ApplicationResource func(*sql.Selector)

// ApplicationRevision is the predicate function for applicationrevision builders.
type ApplicationRevision func(*sql.Selector)

// Connector is the predicate function for connector builders.
type Connector func(*sql.Selector)

// Environment is the predicate function for environment builders.
type Environment func(*sql.Selector)

// Module is the predicate function for module builders.
type Module func(*sql.Selector)

// Project is the predicate function for project builders.
type Project func(*sql.Selector)

// Role is the predicate function for role builders.
type Role func(*sql.Selector)

// Setting is the predicate function for setting builders.
type Setting func(*sql.Selector)

// Subject is the predicate function for subject builders.
type Subject func(*sql.Selector)

// Token is the predicate function for token builders.
type Token func(*sql.Selector)
