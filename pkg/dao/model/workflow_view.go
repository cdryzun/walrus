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

	"github.com/seal-io/walrus/pkg/dao/model/predicate"
	"github.com/seal-io/walrus/pkg/dao/model/workflow"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/types"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// WorkflowCreateInput holds the creation input of the Workflow entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Workflow entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Type of the workflow.
	Type string `path:"-" query:"-" json:"type"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// ID of the environment that this workflow belongs to.
	EnvironmentID object.ID `path:"-" query:"-" json:"environmentID,omitempty"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `path:"-" query:"-" json:"parallelism,omitempty"`
	// Timeout seconds of the workflow.
	Timeout int `path:"-" query:"-" json:"timeout,omitempty"`
	// Configs of workflow variables.
	Variables types.WorkflowVariables `path:"-" query:"-" json:"variables,omitempty"`

	// Stages specifies full inserting the new WorkflowStage entities of the Workflow entity.
	Stages []*WorkflowStageCreateInput `uri:"-" query:"-" json:"stages,omitempty"`
}

// Model returns the Workflow entity for creating,
// after validating.
func (wci *WorkflowCreateInput) Model() *Workflow {
	if wci == nil {
		return nil
	}

	_w := &Workflow{
		Type:          wci.Type,
		Name:          wci.Name,
		Description:   wci.Description,
		Labels:        wci.Labels,
		EnvironmentID: wci.EnvironmentID,
		Parallelism:   wci.Parallelism,
		Timeout:       wci.Timeout,
		Variables:     wci.Variables,
	}

	if wci.Project != nil {
		_w.ProjectID = wci.Project.ID
	}

	if wci.Stages != nil {
		// Empty slice is used for clearing the edge.
		_w.Edges.Stages = make([]*WorkflowStage, 0, len(wci.Stages))
	}
	for j := range wci.Stages {
		if wci.Stages[j] == nil {
			continue
		}
		_w.Edges.Stages = append(_w.Edges.Stages,
			wci.Stages[j].Model())
	}
	return _w
}

// Validate checks the WorkflowCreateInput entity.
func (wci *WorkflowCreateInput) Validate() error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	return wci.ValidateWith(wci.inputConfig.Context, wci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowCreateInput entity with the given context and client set.
func (wci *WorkflowCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if wci.Project != nil {
		if err := wci.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	for i := range wci.Stages {
		if wci.Stages[i] == nil {
			continue
		}

		if err := wci.Stages[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wci.Stages[i] = nil
			}
		}
	}

	return nil
}

// WorkflowCreateInputs holds the creation input item of the Workflow entities.
type WorkflowCreateInputsItem struct {
	// Type of the workflow.
	Type string `path:"-" query:"-" json:"type"`
	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// ID of the environment that this workflow belongs to.
	EnvironmentID object.ID `path:"-" query:"-" json:"environmentID,omitempty"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `path:"-" query:"-" json:"parallelism,omitempty"`
	// Timeout seconds of the workflow.
	Timeout int `path:"-" query:"-" json:"timeout,omitempty"`
	// Configs of workflow variables.
	Variables types.WorkflowVariables `path:"-" query:"-" json:"variables,omitempty"`

	// Stages specifies full inserting the new WorkflowStage entities.
	Stages []*WorkflowStageCreateInput `uri:"-" query:"-" json:"stages,omitempty"`
}

// ValidateWith checks the WorkflowCreateInputsItem entity with the given context and client set.
func (wci *WorkflowCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range wci.Stages {
		if wci.Stages[i] == nil {
			continue
		}

		if err := wci.Stages[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wci.Stages[i] = nil
			}
		}
	}

	return nil
}

// WorkflowCreateInputs holds the creation input of the Workflow entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Workflow entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Workflow entities for creating,
// after validating.
func (wci *WorkflowCreateInputs) Model() []*Workflow {
	if wci == nil || len(wci.Items) == 0 {
		return nil
	}

	_ws := make([]*Workflow, len(wci.Items))

	for i := range wci.Items {
		_w := &Workflow{
			Type:          wci.Items[i].Type,
			Name:          wci.Items[i].Name,
			Description:   wci.Items[i].Description,
			Labels:        wci.Items[i].Labels,
			EnvironmentID: wci.Items[i].EnvironmentID,
			Parallelism:   wci.Items[i].Parallelism,
			Timeout:       wci.Items[i].Timeout,
			Variables:     wci.Items[i].Variables,
		}

		if wci.Project != nil {
			_w.ProjectID = wci.Project.ID
		}

		if wci.Items[i].Stages != nil {
			// Empty slice is used for clearing the edge.
			_w.Edges.Stages = make([]*WorkflowStage, 0, len(wci.Items[i].Stages))
		}
		for j := range wci.Items[i].Stages {
			if wci.Items[i].Stages[j] == nil {
				continue
			}
			_w.Edges.Stages = append(_w.Edges.Stages,
				wci.Items[i].Stages[j].Model())
		}

		_ws[i] = _w
	}

	return _ws
}

// Validate checks the WorkflowCreateInputs entity .
func (wci *WorkflowCreateInputs) Validate() error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	return wci.ValidateWith(wci.inputConfig.Context, wci.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowCreateInputs entity with the given context and client set.
func (wci *WorkflowCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wci == nil {
		return errors.New("nil receiver")
	}

	if len(wci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if wci.Project != nil {
		if err := wci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wci.Project = nil
			}
		}
	}

	for i := range wci.Items {
		if wci.Items[i] == nil {
			continue
		}

		if err := wci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowDeleteInput holds the deletion input of the Workflow entity,
// please tags with `path:",inline"` if embedding.
type WorkflowDeleteInput struct {
	WorkflowQueryInput `path:",inline"`
}

// WorkflowDeleteInputs holds the deletion input item of the Workflow entities.
type WorkflowDeleteInputsItem struct {
	// ID of the Workflow entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Workflow entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// WorkflowDeleteInputs holds the deletion input of the Workflow entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete Workflow entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Workflow entities for deleting,
// after validating.
func (wdi *WorkflowDeleteInputs) Model() []*Workflow {
	if wdi == nil || len(wdi.Items) == 0 {
		return nil
	}

	_ws := make([]*Workflow, len(wdi.Items))
	for i := range wdi.Items {
		_ws[i] = &Workflow{
			ID: wdi.Items[i].ID,
		}
	}
	return _ws
}

// IDs returns the ID list of the Workflow entities for deleting,
// after validating.
func (wdi *WorkflowDeleteInputs) IDs() []object.ID {
	if wdi == nil || len(wdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wdi.Items))
	for i := range wdi.Items {
		ids[i] = wdi.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowDeleteInputs entity.
func (wdi *WorkflowDeleteInputs) Validate() error {
	if wdi == nil {
		return errors.New("nil receiver")
	}

	return wdi.ValidateWith(wdi.inputConfig.Context, wdi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowDeleteInputs entity with the given context and client set.
func (wdi *WorkflowDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wdi == nil {
		return errors.New("nil receiver")
	}

	if len(wdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Workflows().Query()

	// Validate when deleting under the Project route.
	if wdi.Project != nil {
		if err := wdi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				workflow.ProjectID(wdi.Project.ID))
		}
	}

	ids := make([]object.ID, 0, len(wdi.Items))
	ors := make([]predicate.Workflow, 0, len(wdi.Items))
	indexers := make(map[any][]int)

	for i := range wdi.Items {
		if wdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if wdi.Items[i].ID != "" {
			ids = append(ids, wdi.Items[i].ID)
			ors = append(ors, workflow.ID(wdi.Items[i].ID))
			indexers[wdi.Items[i].ID] = append(indexers[wdi.Items[i].ID], i)
		} else if wdi.Items[i].Name != "" {
			ors = append(ors, workflow.And(
				workflow.Name(wdi.Items[i].Name)))
			indexerKey := fmt.Sprint("/", wdi.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := workflow.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = workflow.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			workflow.FieldID,
			workflow.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			wdi.Items[j].ID = es[i].ID
			wdi.Items[j].Name = es[i].Name
		}
	}

	return nil
}

// WorkflowPatchInput holds the patch input of the Workflow entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowPatchInput struct {
	WorkflowQueryInput `path:",inline" query:"-" json:"-"`

	// Name holds the value of the "name" field.
	Name string `path:"-" query:"-" json:"name,omitempty"`
	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Annotations holds the value of the "annotations" field.
	Annotations map[string]string `path:"-" query:"-" json:"annotations,omitempty"`
	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `path:"-" query:"-" json:"createTime,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `path:"-" query:"-" json:"updateTime,omitempty"`
	// ID of the environment that this workflow belongs to.
	EnvironmentID object.ID `path:"-" query:"-" json:"environmentID,omitempty"`
	// Type of the workflow.
	Type string `path:"-" query:"-" json:"type,omitempty"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `path:"-" query:"-" json:"parallelism,omitempty"`
	// Timeout seconds of the workflow.
	Timeout int `path:"-" query:"-" json:"timeout,omitempty"`
	// Execution version of the workflow.
	Version int `path:"-" query:"-" json:"version,omitempty"`
	// Configs of workflow variables.
	Variables types.WorkflowVariables `path:"-" query:"-" json:"variables,omitempty"`

	// Stages indicates replacing the stale WorkflowStage entities.
	Stages []*WorkflowStageCreateInput `uri:"-" query:"-" json:"stages,omitempty"`

	patchedEntity *Workflow `path:"-" query:"-" json:"-"`
}

// PatchModel returns the Workflow partition entity for patching.
func (wpi *WorkflowPatchInput) PatchModel() *Workflow {
	if wpi == nil {
		return nil
	}

	_w := &Workflow{
		Name:          wpi.Name,
		Description:   wpi.Description,
		Labels:        wpi.Labels,
		Annotations:   wpi.Annotations,
		CreateTime:    wpi.CreateTime,
		UpdateTime:    wpi.UpdateTime,
		EnvironmentID: wpi.EnvironmentID,
		Type:          wpi.Type,
		Parallelism:   wpi.Parallelism,
		Timeout:       wpi.Timeout,
		Version:       wpi.Version,
		Variables:     wpi.Variables,
	}

	if wpi.Project != nil {
		_w.ProjectID = wpi.Project.ID
	}

	if wpi.Stages != nil {
		// Empty slice is used for clearing the edge.
		_w.Edges.Stages = make([]*WorkflowStage, 0, len(wpi.Stages))
	}
	for j := range wpi.Stages {
		if wpi.Stages[j] == nil {
			continue
		}
		_w.Edges.Stages = append(_w.Edges.Stages,
			wpi.Stages[j].Model())
	}
	return _w
}

// Model returns the Workflow patched entity,
// after validating.
func (wpi *WorkflowPatchInput) Model() *Workflow {
	if wpi == nil {
		return nil
	}

	return wpi.patchedEntity
}

// Validate checks the WorkflowPatchInput entity.
func (wpi *WorkflowPatchInput) Validate() error {
	if wpi == nil {
		return errors.New("nil receiver")
	}

	return wpi.ValidateWith(wpi.inputConfig.Context, wpi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowPatchInput entity with the given context and client set.
func (wpi *WorkflowPatchInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := wpi.WorkflowQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	q := cs.Workflows().Query()

	// Validate when querying under the Project route.
	if wpi.Project != nil {
		if err := wpi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				workflow.ProjectID(wpi.Project.ID))
		}
	}

	for i := range wpi.Stages {
		if wpi.Stages[i] == nil {
			continue
		}

		if err := wpi.Stages[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wpi.Stages[i] = nil
			}
		}
	}

	if wpi.Refer != nil {
		if wpi.Refer.IsID() {
			q.Where(
				workflow.ID(wpi.Refer.ID()))
		} else if refers := wpi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				workflow.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of workflow")
		}
	} else if wpi.ID != "" {
		q.Where(
			workflow.ID(wpi.ID))
	} else if wpi.Name != "" {
		q.Where(
			workflow.Name(wpi.Name))
	} else {
		return errors.New("invalid identify of workflow")
	}

	q.Select(
		workflow.WithoutFields(
			workflow.FieldAnnotations,
			workflow.FieldCreateTime,
			workflow.FieldUpdateTime,
			workflow.FieldVersion,
		)...,
	)

	var e *Workflow
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*Workflow)
		}
	}

	_pm := wpi.PatchModel()

	_po, err := json.PatchObject(*e, *_pm)
	if err != nil {
		return err
	}

	_obj := _po.(*Workflow)

	if e.Name != _obj.Name {
		return errors.New("field name is immutable")
	}
	if !reflect.DeepEqual(e.CreateTime, _obj.CreateTime) {
		return errors.New("field createTime is immutable")
	}
	if e.EnvironmentID != _obj.EnvironmentID {
		return errors.New("field environmentID is immutable")
	}
	if e.Type != _obj.Type {
		return errors.New("field type is immutable")
	}

	wpi.patchedEntity = _obj
	return nil
}

// WorkflowQueryInput holds the query input of the Workflow entity,
// please tags with `path:",inline"` if embedding.
type WorkflowQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Workflow entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project"`

	// Refer holds the route path reference of the Workflow entity.
	Refer *object.Refer `path:"workflow,default=" query:"-" json:"-"`
	// ID of the Workflow entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Workflow entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// Model returns the Workflow entity for querying,
// after validating.
func (wqi *WorkflowQueryInput) Model() *Workflow {
	if wqi == nil {
		return nil
	}

	return &Workflow{
		ID:   wqi.ID,
		Name: wqi.Name,
	}
}

// Validate checks the WorkflowQueryInput entity.
func (wqi *WorkflowQueryInput) Validate() error {
	if wqi == nil {
		return errors.New("nil receiver")
	}

	return wqi.ValidateWith(wqi.inputConfig.Context, wqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowQueryInput entity with the given context and client set.
func (wqi *WorkflowQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wqi == nil {
		return errors.New("nil receiver")
	}

	if wqi.Refer != nil && *wqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", workflow.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Workflows().Query()

	// Validate when querying under the Project route.
	if wqi.Project != nil {
		if err := wqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				workflow.ProjectID(wqi.Project.ID))
		}
	}

	if wqi.Refer != nil {
		if wqi.Refer.IsID() {
			q.Where(
				workflow.ID(wqi.Refer.ID()))
		} else if refers := wqi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				workflow.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of workflow")
		}
	} else if wqi.ID != "" {
		q.Where(
			workflow.ID(wqi.ID))
	} else if wqi.Name != "" {
		q.Where(
			workflow.Name(wqi.Name))
	} else {
		return errors.New("invalid identify of workflow")
	}

	q.Select(
		workflow.FieldID,
		workflow.FieldName,
	)

	var e *Workflow
	{
		// Get cache from previous validation.
		queryStmt, queryArgs := q.sqlQuery(setContextOp(ctx, q.ctx, "cache")).Query()
		ck := fmt.Sprintf("stmt=%v, args=%v", queryStmt, queryArgs)
		if cv, existed := cache[ck]; !existed {
			var err error
			e, err = q.Only(ctx)
			if err != nil {
				return err
			}

			// Set cache for other validation.
			cache[ck] = e
		} else {
			e = cv.(*Workflow)
		}
	}

	wqi.ID = e.ID
	wqi.Name = e.Name
	return nil
}

// WorkflowQueryInputs holds the query input of the Workflow entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type WorkflowQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Workflow entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the WorkflowQueryInputs entity.
func (wqi *WorkflowQueryInputs) Validate() error {
	if wqi == nil {
		return errors.New("nil receiver")
	}

	return wqi.ValidateWith(wqi.inputConfig.Context, wqi.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowQueryInputs entity with the given context and client set.
func (wqi *WorkflowQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if wqi.Project != nil {
		if err := wqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowUpdateInput holds the modification input of the Workflow entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowUpdateInput struct {
	WorkflowQueryInput `path:",inline" query:"-" json:"-"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `path:"-" query:"-" json:"parallelism,omitempty"`
	// Timeout seconds of the workflow.
	Timeout int `path:"-" query:"-" json:"timeout,omitempty"`
	// Configs of workflow variables.
	Variables types.WorkflowVariables `path:"-" query:"-" json:"variables,omitempty"`

	// Stages indicates replacing the stale WorkflowStage entities.
	Stages []*WorkflowStageCreateInput `uri:"-" query:"-" json:"stages,omitempty"`
}

// Model returns the Workflow entity for modifying,
// after validating.
func (wui *WorkflowUpdateInput) Model() *Workflow {
	if wui == nil {
		return nil
	}

	_w := &Workflow{
		ID:          wui.ID,
		Name:        wui.Name,
		Description: wui.Description,
		Labels:      wui.Labels,
		Parallelism: wui.Parallelism,
		Timeout:     wui.Timeout,
		Variables:   wui.Variables,
	}

	if wui.Stages != nil {
		// Empty slice is used for clearing the edge.
		_w.Edges.Stages = make([]*WorkflowStage, 0, len(wui.Stages))
	}
	for j := range wui.Stages {
		if wui.Stages[j] == nil {
			continue
		}
		_w.Edges.Stages = append(_w.Edges.Stages,
			wui.Stages[j].Model())
	}
	return _w
}

// Validate checks the WorkflowUpdateInput entity.
func (wui *WorkflowUpdateInput) Validate() error {
	if wui == nil {
		return errors.New("nil receiver")
	}

	return wui.ValidateWith(wui.inputConfig.Context, wui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowUpdateInput entity with the given context and client set.
func (wui *WorkflowUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := wui.WorkflowQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	for i := range wui.Stages {
		if wui.Stages[i] == nil {
			continue
		}

		if err := wui.Stages[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wui.Stages[i] = nil
			}
		}
	}

	return nil
}

// WorkflowUpdateInputs holds the modification input item of the Workflow entities.
type WorkflowUpdateInputsItem struct {
	// ID of the Workflow entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Workflow entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`

	// Description holds the value of the "description" field.
	Description string `path:"-" query:"-" json:"description,omitempty"`
	// Labels holds the value of the "labels" field.
	Labels map[string]string `path:"-" query:"-" json:"labels,omitempty"`
	// Number of task pods that can be executed in parallel of workflow.
	Parallelism int `path:"-" query:"-" json:"parallelism"`
	// Timeout seconds of the workflow.
	Timeout int `path:"-" query:"-" json:"timeout"`
	// Configs of workflow variables.
	Variables types.WorkflowVariables `path:"-" query:"-" json:"variables,omitempty"`

	// Stages indicates replacing the stale WorkflowStage entities.
	Stages []*WorkflowStageCreateInput `uri:"-" query:"-" json:"stages,omitempty"`
}

// ValidateWith checks the WorkflowUpdateInputsItem entity with the given context and client set.
func (wui *WorkflowUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	for i := range wui.Stages {
		if wui.Stages[i] == nil {
			continue
		}

		if err := wui.Stages[i].ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				wui.Stages[i] = nil
			}
		}
	}

	return nil
}

// WorkflowUpdateInputs holds the modification input of the Workflow entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type WorkflowUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update Workflow entity MUST under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*WorkflowUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Workflow entities for modifying,
// after validating.
func (wui *WorkflowUpdateInputs) Model() []*Workflow {
	if wui == nil || len(wui.Items) == 0 {
		return nil
	}

	_ws := make([]*Workflow, len(wui.Items))

	for i := range wui.Items {
		_w := &Workflow{
			ID:          wui.Items[i].ID,
			Name:        wui.Items[i].Name,
			Description: wui.Items[i].Description,
			Labels:      wui.Items[i].Labels,
			Parallelism: wui.Items[i].Parallelism,
			Timeout:     wui.Items[i].Timeout,
			Variables:   wui.Items[i].Variables,
		}

		if wui.Items[i].Stages != nil {
			// Empty slice is used for clearing the edge.
			_w.Edges.Stages = make([]*WorkflowStage, 0, len(wui.Items[i].Stages))
		}
		for j := range wui.Items[i].Stages {
			if wui.Items[i].Stages[j] == nil {
				continue
			}
			_w.Edges.Stages = append(_w.Edges.Stages,
				wui.Items[i].Stages[j].Model())
		}

		_ws[i] = _w
	}

	return _ws
}

// IDs returns the ID list of the Workflow entities for modifying,
// after validating.
func (wui *WorkflowUpdateInputs) IDs() []object.ID {
	if wui == nil || len(wui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(wui.Items))
	for i := range wui.Items {
		ids[i] = wui.Items[i].ID
	}
	return ids
}

// Validate checks the WorkflowUpdateInputs entity.
func (wui *WorkflowUpdateInputs) Validate() error {
	if wui == nil {
		return errors.New("nil receiver")
	}

	return wui.ValidateWith(wui.inputConfig.Context, wui.inputConfig.Client, nil)
}

// ValidateWith checks the WorkflowUpdateInputs entity with the given context and client set.
func (wui *WorkflowUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if wui == nil {
		return errors.New("nil receiver")
	}

	if len(wui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Workflows().Query()

	// Validate when updating under the Project route.
	if wui.Project != nil {
		if err := wui.Project.ValidateWith(ctx, cs, cache); err != nil {
			return err
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				workflow.ProjectID(wui.Project.ID))
		}
	}

	ids := make([]object.ID, 0, len(wui.Items))
	ors := make([]predicate.Workflow, 0, len(wui.Items))
	indexers := make(map[any][]int)

	for i := range wui.Items {
		if wui.Items[i] == nil {
			return errors.New("nil item")
		}

		if wui.Items[i].ID != "" {
			ids = append(ids, wui.Items[i].ID)
			ors = append(ors, workflow.ID(wui.Items[i].ID))
			indexers[wui.Items[i].ID] = append(indexers[wui.Items[i].ID], i)
		} else if wui.Items[i].Name != "" {
			ors = append(ors, workflow.And(
				workflow.Name(wui.Items[i].Name)))
			indexerKey := fmt.Sprint("/", wui.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := workflow.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = workflow.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			workflow.FieldID,
			workflow.FieldName,
		).
		All(ctx)
	if err != nil {
		return err
	}

	if len(es) != cap(ids) {
		return errors.New("found unrecognized item")
	}

	for i := range es {
		indexer := indexers[es[i].ID]
		if indexer == nil {
			indexerKey := fmt.Sprint("/", es[i].Name)
			indexer = indexers[indexerKey]
		}
		for _, j := range indexer {
			wui.Items[j].ID = es[i].ID
			wui.Items[j].Name = es[i].Name
		}
	}

	for i := range wui.Items {
		if err := wui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// WorkflowOutput holds the output of the Workflow entity.
type WorkflowOutput struct {
	ID            object.ID               `json:"id,omitempty"`
	Name          string                  `json:"name,omitempty"`
	Description   string                  `json:"description,omitempty"`
	Labels        map[string]string       `json:"labels,omitempty"`
	CreateTime    *time.Time              `json:"createTime,omitempty"`
	UpdateTime    *time.Time              `json:"updateTime,omitempty"`
	EnvironmentID object.ID               `json:"environmentID,omitempty"`
	Type          string                  `json:"type,omitempty"`
	Parallelism   int                     `json:"parallelism,omitempty"`
	Timeout       int                     `json:"timeout,omitempty"`
	Version       int                     `json:"version,omitempty"`
	Variables     types.WorkflowVariables `json:"variables,omitempty"`

	Project    *ProjectOutput             `json:"project,omitempty"`
	Stages     []*WorkflowStageOutput     `json:"stages,omitempty"`
	Executions []*WorkflowExecutionOutput `json:"executions,omitempty"`
}

// View returns the output of Workflow entity.
func (_w *Workflow) View() *WorkflowOutput {
	return ExposeWorkflow(_w)
}

// View returns the output of Workflow entities.
func (_ws Workflows) View() []*WorkflowOutput {
	return ExposeWorkflows(_ws)
}

// ExposeWorkflow converts the Workflow to WorkflowOutput.
func ExposeWorkflow(_w *Workflow) *WorkflowOutput {
	if _w == nil {
		return nil
	}

	wo := &WorkflowOutput{
		ID:            _w.ID,
		Name:          _w.Name,
		Description:   _w.Description,
		Labels:        _w.Labels,
		CreateTime:    _w.CreateTime,
		UpdateTime:    _w.UpdateTime,
		EnvironmentID: _w.EnvironmentID,
		Type:          _w.Type,
		Parallelism:   _w.Parallelism,
		Timeout:       _w.Timeout,
		Version:       _w.Version,
		Variables:     _w.Variables,
	}

	if _w.Edges.Project != nil {
		wo.Project = ExposeProject(_w.Edges.Project)
	} else if _w.ProjectID != "" {
		wo.Project = &ProjectOutput{
			ID: _w.ProjectID,
		}
	}
	if _w.Edges.Stages != nil {
		wo.Stages = ExposeWorkflowStages(_w.Edges.Stages)
	}
	if _w.Edges.Executions != nil {
		wo.Executions = ExposeWorkflowExecutions(_w.Edges.Executions)
	}
	return wo
}

// ExposeWorkflows converts the Workflow slice to WorkflowOutput pointer slice.
func ExposeWorkflows(_ws []*Workflow) []*WorkflowOutput {
	if len(_ws) == 0 {
		return nil
	}

	wos := make([]*WorkflowOutput, len(_ws))
	for i := range _ws {
		wos[i] = ExposeWorkflow(_ws[i])
	}
	return wos
}
