// SPDX-FileCopyrightText: 2023 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/environment"
	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// Environment is the model entity for the Environment schema.
type Environment struct {
	config `json:"-"`
	// ID of the ent.
	ID object.ID `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `json:"labels,omitempty"`
	// Annotations holds the value of the "annotations" field.
	Annotations map[string]string `json:"annotations,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `json:"create_time,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `json:"update_time,omitempty"`
	// ID of the project to belong.
	ProjectID object.ID `json:"project_id,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the EnvironmentQuery when eager-loading is set.
	Edges        EnvironmentEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// EnvironmentEdges holds the relations/edges for other nodes in the graph.
type EnvironmentEdges struct {
	// Project to which the environment belongs.
	Project *Project `json:"project,omitempty"`
	// Connectors holds the value of the connectors edge.
	Connectors []*EnvironmentConnectorRelationship `json:"connectors,omitempty"`
	// Services that belong to the environment.
	Services []*Service `json:"services,omitempty,cli-ignore"`
	// ServicesRevisions that belong to the environment.
	ServiceRevisions []*ServiceRevision `json:"service_revisions,omitempty"`
	// ServiceResources that belong to the environment.
	ServiceResources []*ServiceResource `json:"service_resources,omitempty"`
	// Variables that belong to the environment.
	Variables []*Variable `json:"variables,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [6]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e EnvironmentEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// ConnectorsOrErr returns the Connectors value or an error if the edge
// was not loaded in eager-loading.
func (e EnvironmentEdges) ConnectorsOrErr() ([]*EnvironmentConnectorRelationship, error) {
	if e.loadedTypes[1] {
		return e.Connectors, nil
	}
	return nil, &NotLoadedError{edge: "connectors"}
}

// ServicesOrErr returns the Services value or an error if the edge
// was not loaded in eager-loading.
func (e EnvironmentEdges) ServicesOrErr() ([]*Service, error) {
	if e.loadedTypes[2] {
		return e.Services, nil
	}
	return nil, &NotLoadedError{edge: "services"}
}

// ServiceRevisionsOrErr returns the ServiceRevisions value or an error if the edge
// was not loaded in eager-loading.
func (e EnvironmentEdges) ServiceRevisionsOrErr() ([]*ServiceRevision, error) {
	if e.loadedTypes[3] {
		return e.ServiceRevisions, nil
	}
	return nil, &NotLoadedError{edge: "service_revisions"}
}

// ServiceResourcesOrErr returns the ServiceResources value or an error if the edge
// was not loaded in eager-loading.
func (e EnvironmentEdges) ServiceResourcesOrErr() ([]*ServiceResource, error) {
	if e.loadedTypes[4] {
		return e.ServiceResources, nil
	}
	return nil, &NotLoadedError{edge: "service_resources"}
}

// VariablesOrErr returns the Variables value or an error if the edge
// was not loaded in eager-loading.
func (e EnvironmentEdges) VariablesOrErr() ([]*Variable, error) {
	if e.loadedTypes[5] {
		return e.Variables, nil
	}
	return nil, &NotLoadedError{edge: "variables"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Environment) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case environment.FieldLabels, environment.FieldAnnotations:
			values[i] = new([]byte)
		case environment.FieldID, environment.FieldProjectID:
			values[i] = new(object.ID)
		case environment.FieldName, environment.FieldDescription:
			values[i] = new(sql.NullString)
		case environment.FieldCreateTime, environment.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Environment fields.
func (e *Environment) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case environment.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				e.ID = *value
			}
		case environment.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				e.Name = value.String
			}
		case environment.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				e.Description = value.String
			}
		case environment.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case environment.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &e.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case environment.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				e.CreateTime = new(time.Time)
				*e.CreateTime = value.Time
			}
		case environment.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				e.UpdateTime = new(time.Time)
				*e.UpdateTime = value.Time
			}
		case environment.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				e.ProjectID = *value
			}
		default:
			e.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Environment.
// This includes values selected through modifiers, order, etc.
func (e *Environment) Value(name string) (ent.Value, error) {
	return e.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the Environment entity.
func (e *Environment) QueryProject() *ProjectQuery {
	return NewEnvironmentClient(e.config).QueryProject(e)
}

// QueryConnectors queries the "connectors" edge of the Environment entity.
func (e *Environment) QueryConnectors() *EnvironmentConnectorRelationshipQuery {
	return NewEnvironmentClient(e.config).QueryConnectors(e)
}

// QueryServices queries the "services" edge of the Environment entity.
func (e *Environment) QueryServices() *ServiceQuery {
	return NewEnvironmentClient(e.config).QueryServices(e)
}

// QueryServiceRevisions queries the "service_revisions" edge of the Environment entity.
func (e *Environment) QueryServiceRevisions() *ServiceRevisionQuery {
	return NewEnvironmentClient(e.config).QueryServiceRevisions(e)
}

// QueryServiceResources queries the "service_resources" edge of the Environment entity.
func (e *Environment) QueryServiceResources() *ServiceResourceQuery {
	return NewEnvironmentClient(e.config).QueryServiceResources(e)
}

// QueryVariables queries the "variables" edge of the Environment entity.
func (e *Environment) QueryVariables() *VariableQuery {
	return NewEnvironmentClient(e.config).QueryVariables(e)
}

// Update returns a builder for updating this Environment.
// Note that you need to call Environment.Unwrap() before calling this method if this Environment
// was returned from a transaction, and the transaction was committed or rolled back.
func (e *Environment) Update() *EnvironmentUpdateOne {
	return NewEnvironmentClient(e.config).UpdateOne(e)
}

// Unwrap unwraps the Environment entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (e *Environment) Unwrap() *Environment {
	_tx, ok := e.config.driver.(*txDriver)
	if !ok {
		panic("model: Environment is not a transactional entity")
	}
	e.config.driver = _tx.drv
	return e
}

// String implements the fmt.Stringer.
func (e *Environment) String() string {
	var builder strings.Builder
	builder.WriteString("Environment(")
	builder.WriteString(fmt.Sprintf("id=%v, ", e.ID))
	builder.WriteString("name=")
	builder.WriteString(e.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(e.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", e.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", e.Annotations))
	builder.WriteString(", ")
	if v := e.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := e.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", e.ProjectID))
	builder.WriteByte(')')
	return builder.String()
}

// Environments is a parsable slice of Environment.
type Environments []*Environment
