// SPDX-FileCopyrightText: 2024 Seal, Inc
// SPDX-License-Identifier: Apache-2.0

// Code generated by "walrus". DO NOT EDIT.

package model

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"

	"github.com/seal-io/walrus/pkg/dao/model/project"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	"github.com/seal-io/walrus/pkg/dao/model/workflowstage"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// WorkflowStage is the model entity for the WorkflowStage schema.
type WorkflowStage struct {
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
	// ID of the workflow that this workflow stage belongs to.
	WorkflowID object.ID `json:"workflow_id,omitempty"`
	// ID list of the workflow stages that this workflow stage depends on.
	Dependencies []object.ID `json:"dependencies,omitempty"`
	// Order of the workflow stage.
	Order int `json:"order,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the WorkflowStageQuery when eager-loading is set.
	Edges        WorkflowStageEdges `json:"edges,omitempty"`
	selectValues sql.SelectValues
}

// WorkflowStageEdges holds the relations/edges for other nodes in the graph.
type WorkflowStageEdges struct {
	// Project to which the workflow stage belongs.
	Project *Project `json:"project,omitempty"`
	// Workflow steps that belong to this workflow stage.
	Steps []*WorkflowStep `json:"steps,omitempty"`
	// Workflow that this workflow stage belongs to.
	Workflow *Workflow `json:"workflow,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// ProjectOrErr returns the Project value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowStageEdges) ProjectOrErr() (*Project, error) {
	if e.loadedTypes[0] {
		if e.Project == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: project.Label}
		}
		return e.Project, nil
	}
	return nil, &NotLoadedError{edge: "project"}
}

// StepsOrErr returns the Steps value or an error if the edge
// was not loaded in eager-loading.
func (e WorkflowStageEdges) StepsOrErr() ([]*WorkflowStep, error) {
	if e.loadedTypes[1] {
		return e.Steps, nil
	}
	return nil, &NotLoadedError{edge: "steps"}
}

// WorkflowOrErr returns the Workflow value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e WorkflowStageEdges) WorkflowOrErr() (*Workflow, error) {
	if e.loadedTypes[2] {
		if e.Workflow == nil {
			// Edge was loaded but was not found.
			return nil, &NotFoundError{label: workflow.Label}
		}
		return e.Workflow, nil
	}
	return nil, &NotLoadedError{edge: "workflow"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WorkflowStage) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case workflowstage.FieldLabels, workflowstage.FieldAnnotations, workflowstage.FieldDependencies:
			values[i] = new([]byte)
		case workflowstage.FieldID, workflowstage.FieldProjectID, workflowstage.FieldWorkflowID:
			values[i] = new(object.ID)
		case workflowstage.FieldOrder:
			values[i] = new(sql.NullInt64)
		case workflowstage.FieldName, workflowstage.FieldDescription:
			values[i] = new(sql.NullString)
		case workflowstage.FieldCreateTime, workflowstage.FieldUpdateTime:
			values[i] = new(sql.NullTime)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WorkflowStage fields.
func (ws *WorkflowStage) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case workflowstage.FieldID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				ws.ID = *value
			}
		case workflowstage.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				ws.Name = value.String
			}
		case workflowstage.FieldDescription:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field description", values[i])
			} else if value.Valid {
				ws.Description = value.String
			}
		case workflowstage.FieldLabels:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field labels", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Labels); err != nil {
					return fmt.Errorf("unmarshal field labels: %w", err)
				}
			}
		case workflowstage.FieldAnnotations:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field annotations", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Annotations); err != nil {
					return fmt.Errorf("unmarshal field annotations: %w", err)
				}
			}
		case workflowstage.FieldCreateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_time", values[i])
			} else if value.Valid {
				ws.CreateTime = new(time.Time)
				*ws.CreateTime = value.Time
			}
		case workflowstage.FieldUpdateTime:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field update_time", values[i])
			} else if value.Valid {
				ws.UpdateTime = new(time.Time)
				*ws.UpdateTime = value.Time
			}
		case workflowstage.FieldProjectID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field project_id", values[i])
			} else if value != nil {
				ws.ProjectID = *value
			}
		case workflowstage.FieldWorkflowID:
			if value, ok := values[i].(*object.ID); !ok {
				return fmt.Errorf("unexpected type %T for field workflow_id", values[i])
			} else if value != nil {
				ws.WorkflowID = *value
			}
		case workflowstage.FieldDependencies:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field dependencies", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &ws.Dependencies); err != nil {
					return fmt.Errorf("unmarshal field dependencies: %w", err)
				}
			}
		case workflowstage.FieldOrder:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field order", values[i])
			} else if value.Valid {
				ws.Order = int(value.Int64)
			}
		default:
			ws.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the WorkflowStage.
// This includes values selected through modifiers, order, etc.
func (ws *WorkflowStage) Value(name string) (ent.Value, error) {
	return ws.selectValues.Get(name)
}

// QueryProject queries the "project" edge of the WorkflowStage entity.
func (ws *WorkflowStage) QueryProject() *ProjectQuery {
	return NewWorkflowStageClient(ws.config).QueryProject(ws)
}

// QuerySteps queries the "steps" edge of the WorkflowStage entity.
func (ws *WorkflowStage) QuerySteps() *WorkflowStepQuery {
	return NewWorkflowStageClient(ws.config).QuerySteps(ws)
}

// QueryWorkflow queries the "workflow" edge of the WorkflowStage entity.
func (ws *WorkflowStage) QueryWorkflow() *WorkflowQuery {
	return NewWorkflowStageClient(ws.config).QueryWorkflow(ws)
}

// Update returns a builder for updating this WorkflowStage.
// Note that you need to call WorkflowStage.Unwrap() before calling this method if this WorkflowStage
// was returned from a transaction, and the transaction was committed or rolled back.
func (ws *WorkflowStage) Update() *WorkflowStageUpdateOne {
	return NewWorkflowStageClient(ws.config).UpdateOne(ws)
}

// Unwrap unwraps the WorkflowStage entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (ws *WorkflowStage) Unwrap() *WorkflowStage {
	_tx, ok := ws.config.driver.(*txDriver)
	if !ok {
		panic("model: WorkflowStage is not a transactional entity")
	}
	ws.config.driver = _tx.drv
	return ws
}

// String implements the fmt.Stringer.
func (ws *WorkflowStage) String() string {
	var builder strings.Builder
	builder.WriteString("WorkflowStage(")
	builder.WriteString(fmt.Sprintf("id=%v, ", ws.ID))
	builder.WriteString("name=")
	builder.WriteString(ws.Name)
	builder.WriteString(", ")
	builder.WriteString("description=")
	builder.WriteString(ws.Description)
	builder.WriteString(", ")
	builder.WriteString("labels=")
	builder.WriteString(fmt.Sprintf("%v", ws.Labels))
	builder.WriteString(", ")
	builder.WriteString("annotations=")
	builder.WriteString(fmt.Sprintf("%v", ws.Annotations))
	builder.WriteString(", ")
	if v := ws.CreateTime; v != nil {
		builder.WriteString("create_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	if v := ws.UpdateTime; v != nil {
		builder.WriteString("update_time=")
		builder.WriteString(v.Format(time.ANSIC))
	}
	builder.WriteString(", ")
	builder.WriteString("project_id=")
	builder.WriteString(fmt.Sprintf("%v", ws.ProjectID))
	builder.WriteString(", ")
	builder.WriteString("workflow_id=")
	builder.WriteString(fmt.Sprintf("%v", ws.WorkflowID))
	builder.WriteString(", ")
	builder.WriteString("dependencies=")
	builder.WriteString(fmt.Sprintf("%v", ws.Dependencies))
	builder.WriteString(", ")
	builder.WriteString("order=")
	builder.WriteString(fmt.Sprintf("%v", ws.Order))
	builder.WriteByte(')')
	return builder.String()
}

// WorkflowStages is a parsable slice of WorkflowStage.
type WorkflowStages []*WorkflowStage
