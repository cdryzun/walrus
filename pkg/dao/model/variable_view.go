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
	"github.com/seal-io/walrus/pkg/dao/model/variable"
	"github.com/seal-io/walrus/pkg/dao/schema/intercept"
	"github.com/seal-io/walrus/pkg/dao/types/crypto"
	"github.com/seal-io/walrus/pkg/dao/types/object"
	"github.com/seal-io/walrus/utils/json"
)

// VariableCreateInput holds the creation input of the Variable entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type VariableCreateInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Variable entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create Variable entity CAN under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`

	// The value of variable, store in string.
	Value crypto.String `path:"-" query:"-" json:"value"`
	// The name of variable.
	Name string `path:"-" query:"-" json:"name"`
	// The value is sensitive or not.
	Sensitive bool `path:"-" query:"-" json:"sensitive,omitempty"`
	// Description of the variable.
	Description string `path:"-" query:"-" json:"description,omitempty"`
}

// Model returns the Variable entity for creating,
// after validating.
func (vci *VariableCreateInput) Model() *Variable {
	if vci == nil {
		return nil
	}

	_v := &Variable{
		Value:       vci.Value,
		Name:        vci.Name,
		Sensitive:   vci.Sensitive,
		Description: vci.Description,
	}

	if vci.Project != nil {
		_v.ProjectID = vci.Project.ID
	}
	if vci.Environment != nil {
		_v.EnvironmentID = vci.Environment.ID
	}

	return _v
}

// Validate checks the VariableCreateInput entity.
func (vci *VariableCreateInput) Validate() error {
	if vci == nil {
		return errors.New("nil receiver")
	}

	return vci.ValidateWith(vci.inputConfig.Context, vci.inputConfig.Client, nil)
}

// ValidateWith checks the VariableCreateInput entity with the given context and client set.
func (vci *VariableCreateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if vci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if vci.Project != nil {
		if err := vci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vci.Project = nil
			}
		}
	}
	// Validate when creating under the Environment route.
	if vci.Environment != nil {
		if err := vci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vci.Environment = nil
			}
		}
	}

	return nil
}

// VariableCreateInputs holds the creation input item of the Variable entities.
type VariableCreateInputsItem struct {
	// The value of variable, store in string.
	Value crypto.String `path:"-" query:"-" json:"value"`
	// The name of variable.
	Name string `path:"-" query:"-" json:"name"`
	// The value is sensitive or not.
	Sensitive bool `path:"-" query:"-" json:"sensitive,omitempty"`
	// Description of the variable.
	Description string `path:"-" query:"-" json:"description,omitempty"`
}

// ValidateWith checks the VariableCreateInputsItem entity with the given context and client set.
func (vci *VariableCreateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if vci == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// VariableCreateInputs holds the creation input of the Variable entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type VariableCreateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to create Variable entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to create Variable entity CAN under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*VariableCreateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Variable entities for creating,
// after validating.
func (vci *VariableCreateInputs) Model() []*Variable {
	if vci == nil || len(vci.Items) == 0 {
		return nil
	}

	_vs := make([]*Variable, len(vci.Items))

	for i := range vci.Items {
		_v := &Variable{
			Value:       vci.Items[i].Value,
			Name:        vci.Items[i].Name,
			Sensitive:   vci.Items[i].Sensitive,
			Description: vci.Items[i].Description,
		}

		if vci.Project != nil {
			_v.ProjectID = vci.Project.ID
		}
		if vci.Environment != nil {
			_v.EnvironmentID = vci.Environment.ID
		}

		_vs[i] = _v
	}

	return _vs
}

// Validate checks the VariableCreateInputs entity .
func (vci *VariableCreateInputs) Validate() error {
	if vci == nil {
		return errors.New("nil receiver")
	}

	return vci.ValidateWith(vci.inputConfig.Context, vci.inputConfig.Client, nil)
}

// ValidateWith checks the VariableCreateInputs entity with the given context and client set.
func (vci *VariableCreateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if vci == nil {
		return errors.New("nil receiver")
	}

	if len(vci.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when creating under the Project route.
	if vci.Project != nil {
		if err := vci.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vci.Project = nil
			}
		}
	}
	// Validate when creating under the Environment route.
	if vci.Environment != nil {
		if err := vci.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vci.Environment = nil
			}
		}
	}

	for i := range vci.Items {
		if vci.Items[i] == nil {
			continue
		}

		if err := vci.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// VariableDeleteInput holds the deletion input of the Variable entity,
// please tags with `path:",inline"` if embedding.
type VariableDeleteInput struct {
	VariableQueryInput `path:",inline"`
}

// VariableDeleteInputs holds the deletion input item of the Variable entities.
type VariableDeleteInputsItem struct {
	// ID of the Variable entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Variable entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// VariableDeleteInputs holds the deletion input of the Variable entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type VariableDeleteInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to delete Variable entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to delete Variable entity CAN under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*VariableDeleteInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Variable entities for deleting,
// after validating.
func (vdi *VariableDeleteInputs) Model() []*Variable {
	if vdi == nil || len(vdi.Items) == 0 {
		return nil
	}

	_vs := make([]*Variable, len(vdi.Items))
	for i := range vdi.Items {
		_vs[i] = &Variable{
			ID: vdi.Items[i].ID,
		}
	}
	return _vs
}

// IDs returns the ID list of the Variable entities for deleting,
// after validating.
func (vdi *VariableDeleteInputs) IDs() []object.ID {
	if vdi == nil || len(vdi.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(vdi.Items))
	for i := range vdi.Items {
		ids[i] = vdi.Items[i].ID
	}
	return ids
}

// Validate checks the VariableDeleteInputs entity.
func (vdi *VariableDeleteInputs) Validate() error {
	if vdi == nil {
		return errors.New("nil receiver")
	}

	return vdi.ValidateWith(vdi.inputConfig.Context, vdi.inputConfig.Client, nil)
}

// ValidateWith checks the VariableDeleteInputs entity with the given context and client set.
func (vdi *VariableDeleteInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if vdi == nil {
		return errors.New("nil receiver")
	}

	if len(vdi.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Variables().Query()

	// Validate when deleting under the Project route.
	if vdi.Project != nil {
		if err := vdi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vdi.Project = nil
				q.Where(
					variable.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				variable.ProjectID(vdi.Project.ID))
		}
	} else {
		q.Where(
			variable.ProjectIDIsNil())
	}

	// Validate when deleting under the Environment route.
	if vdi.Environment != nil {
		if err := vdi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vdi.Environment = nil
				q.Where(
					variable.EnvironmentIDIsNil())
			}
		} else {
			q.Where(
				variable.EnvironmentID(vdi.Environment.ID))
		}
	} else {
		q.Where(
			variable.EnvironmentIDIsNil())
	}

	ids := make([]object.ID, 0, len(vdi.Items))
	ors := make([]predicate.Variable, 0, len(vdi.Items))
	indexers := make(map[any][]int)

	for i := range vdi.Items {
		if vdi.Items[i] == nil {
			return errors.New("nil item")
		}

		if vdi.Items[i].ID != "" {
			ids = append(ids, vdi.Items[i].ID)
			ors = append(ors, variable.ID(vdi.Items[i].ID))
			indexers[vdi.Items[i].ID] = append(indexers[vdi.Items[i].ID], i)
		} else if vdi.Items[i].Name != "" {
			ors = append(ors, variable.And(
				variable.Name(vdi.Items[i].Name)))
			indexerKey := fmt.Sprint("/", vdi.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := variable.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = variable.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			variable.FieldID,
			variable.FieldName,
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
			vdi.Items[j].ID = es[i].ID
			vdi.Items[j].Name = es[i].Name
		}
	}

	return nil
}

// VariablePatchInput holds the patch input of the Variable entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type VariablePatchInput struct {
	VariableQueryInput `path:",inline" query:"-" json:"-"`

	// CreateTime holds the value of the "create_time" field.
	CreateTime *time.Time `path:"-" query:"-" json:"createTime,omitempty"`
	// UpdateTime holds the value of the "update_time" field.
	UpdateTime *time.Time `path:"-" query:"-" json:"updateTime,omitempty"`
	// The name of variable.
	Name string `path:"-" query:"-" json:"name,omitempty"`
	// The value of variable, store in string.
	Value crypto.String `path:"-" query:"-" json:"value,omitempty"`
	// The value is sensitive or not.
	Sensitive bool `path:"-" query:"-" json:"sensitive,omitempty"`
	// Description of the variable.
	Description string `path:"-" query:"-" json:"description,omitempty"`

	patchedEntity *Variable `path:"-" query:"-" json:"-"`
}

// PatchModel returns the Variable partition entity for patching.
func (vpi *VariablePatchInput) PatchModel() *Variable {
	if vpi == nil {
		return nil
	}

	_v := &Variable{
		CreateTime:  vpi.CreateTime,
		UpdateTime:  vpi.UpdateTime,
		Name:        vpi.Name,
		Value:       vpi.Value,
		Sensitive:   vpi.Sensitive,
		Description: vpi.Description,
	}

	if vpi.Project != nil {
		_v.ProjectID = vpi.Project.ID
	}
	if vpi.Environment != nil {
		_v.EnvironmentID = vpi.Environment.ID
	}

	return _v
}

// Model returns the Variable patched entity,
// after validating.
func (vpi *VariablePatchInput) Model() *Variable {
	if vpi == nil {
		return nil
	}

	return vpi.patchedEntity
}

// Validate checks the VariablePatchInput entity.
func (vpi *VariablePatchInput) Validate() error {
	if vpi == nil {
		return errors.New("nil receiver")
	}

	return vpi.ValidateWith(vpi.inputConfig.Context, vpi.inputConfig.Client, nil)
}

// ValidateWith checks the VariablePatchInput entity with the given context and client set.
func (vpi *VariablePatchInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := vpi.VariableQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	q := cs.Variables().Query()

	// Validate when querying under the Project route.
	if vpi.Project != nil {
		if err := vpi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vpi.Project = nil
				q.Where(
					variable.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				variable.ProjectID(vpi.Project.ID))
		}
	} else {
		q.Where(
			variable.ProjectIDIsNil())
	}

	// Validate when querying under the Environment route.
	if vpi.Environment != nil {
		if err := vpi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vpi.Environment = nil
				q.Where(
					variable.EnvironmentIDIsNil())
			}
		} else {
			q.Where(
				variable.EnvironmentID(vpi.Environment.ID))
		}
	} else {
		q.Where(
			variable.EnvironmentIDIsNil())
	}

	if vpi.Refer != nil {
		if vpi.Refer.IsID() {
			q.Where(
				variable.ID(vpi.Refer.ID()))
		} else if refers := vpi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				variable.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of variable")
		}
	} else if vpi.ID != "" {
		q.Where(
			variable.ID(vpi.ID))
	} else if vpi.Name != "" {
		q.Where(
			variable.Name(vpi.Name))
	} else {
		return errors.New("invalid identify of variable")
	}

	q.Select(
		variable.WithoutFields(
			variable.FieldCreateTime,
			variable.FieldUpdateTime,
		)...,
	)

	var e *Variable
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
			e = cv.(*Variable)
		}
	}

	_pm := vpi.PatchModel()

	_po, err := json.PatchObject(*e, *_pm)
	if err != nil {
		return err
	}

	_obj := _po.(*Variable)

	if !reflect.DeepEqual(e.CreateTime, _obj.CreateTime) {
		return errors.New("field createTime is immutable")
	}
	if e.Name != _obj.Name {
		return errors.New("field name is immutable")
	}

	vpi.patchedEntity = _obj
	return nil
}

// VariableQueryInput holds the query input of the Variable entity,
// please tags with `path:",inline"` if embedding.
type VariableQueryInput struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Variable entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"project,omitempty"`
	// Environment indicates to query Variable entity CAN under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"environment,omitempty"`

	// Refer holds the route path reference of the Variable entity.
	Refer *object.Refer `path:"variable,default=" query:"-" json:"-"`
	// ID of the Variable entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Variable entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`
}

// Model returns the Variable entity for querying,
// after validating.
func (vqi *VariableQueryInput) Model() *Variable {
	if vqi == nil {
		return nil
	}

	return &Variable{
		ID:   vqi.ID,
		Name: vqi.Name,
	}
}

// Validate checks the VariableQueryInput entity.
func (vqi *VariableQueryInput) Validate() error {
	if vqi == nil {
		return errors.New("nil receiver")
	}

	return vqi.ValidateWith(vqi.inputConfig.Context, vqi.inputConfig.Client, nil)
}

// ValidateWith checks the VariableQueryInput entity with the given context and client set.
func (vqi *VariableQueryInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if vqi == nil {
		return errors.New("nil receiver")
	}

	if vqi.Refer != nil && *vqi.Refer == "" {
		return fmt.Errorf("model: %s : %w", variable.Label, ErrBlankResourceRefer)
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Variables().Query()

	// Validate when querying under the Project route.
	if vqi.Project != nil {
		if err := vqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vqi.Project = nil
				q.Where(
					variable.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				variable.ProjectID(vqi.Project.ID))
		}
	} else {
		q.Where(
			variable.ProjectIDIsNil())
	}

	// Validate when querying under the Environment route.
	if vqi.Environment != nil {
		if err := vqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vqi.Environment = nil
				q.Where(
					variable.EnvironmentIDIsNil())
			}
		} else {
			q.Where(
				variable.EnvironmentID(vqi.Environment.ID))
		}
	} else {
		q.Where(
			variable.EnvironmentIDIsNil())
	}

	if vqi.Refer != nil {
		if vqi.Refer.IsID() {
			q.Where(
				variable.ID(vqi.Refer.ID()))
		} else if refers := vqi.Refer.Split(1); len(refers) == 1 {
			q.Where(
				variable.Name(refers[0].String()))
		} else {
			return errors.New("invalid identify refer of variable")
		}
	} else if vqi.ID != "" {
		q.Where(
			variable.ID(vqi.ID))
	} else if vqi.Name != "" {
		q.Where(
			variable.Name(vqi.Name))
	} else {
		return errors.New("invalid identify of variable")
	}

	q.Select(
		variable.FieldID,
		variable.FieldName,
	)

	var e *Variable
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
			e = cv.(*Variable)
		}
	}

	vqi.ID = e.ID
	vqi.Name = e.Name
	return nil
}

// VariableQueryInputs holds the query input of the Variable entities,
// please tags with `path:",inline" query:",inline"` if embedding.
type VariableQueryInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to query Variable entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to query Variable entity CAN under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`
}

// Validate checks the VariableQueryInputs entity.
func (vqi *VariableQueryInputs) Validate() error {
	if vqi == nil {
		return errors.New("nil receiver")
	}

	return vqi.ValidateWith(vqi.inputConfig.Context, vqi.inputConfig.Client, nil)
}

// ValidateWith checks the VariableQueryInputs entity with the given context and client set.
func (vqi *VariableQueryInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if vqi == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	// Validate when querying under the Project route.
	if vqi.Project != nil {
		if err := vqi.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vqi.Project = nil
			}
		}
	}

	// Validate when querying under the Environment route.
	if vqi.Environment != nil {
		if err := vqi.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vqi.Environment = nil
			}
		}
	}

	return nil
}

// VariableUpdateInput holds the modification input of the Variable entity,
// please tags with `path:",inline" json:",inline"` if embedding.
type VariableUpdateInput struct {
	VariableQueryInput `path:",inline" query:"-" json:"-"`

	// The value of variable, store in string.
	Value crypto.String `path:"-" query:"-" json:"value,omitempty"`
	// The value is sensitive or not.
	Sensitive bool `path:"-" query:"-" json:"sensitive,omitempty"`
	// Description of the variable.
	Description string `path:"-" query:"-" json:"description,omitempty"`
}

// Model returns the Variable entity for modifying,
// after validating.
func (vui *VariableUpdateInput) Model() *Variable {
	if vui == nil {
		return nil
	}

	_v := &Variable{
		ID:          vui.ID,
		Name:        vui.Name,
		Value:       vui.Value,
		Sensitive:   vui.Sensitive,
		Description: vui.Description,
	}

	return _v
}

// Validate checks the VariableUpdateInput entity.
func (vui *VariableUpdateInput) Validate() error {
	if vui == nil {
		return errors.New("nil receiver")
	}

	return vui.ValidateWith(vui.inputConfig.Context, vui.inputConfig.Client, nil)
}

// ValidateWith checks the VariableUpdateInput entity with the given context and client set.
func (vui *VariableUpdateInput) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if cache == nil {
		cache = map[string]any{}
	}

	if err := vui.VariableQueryInput.ValidateWith(ctx, cs, cache); err != nil {
		return err
	}

	return nil
}

// VariableUpdateInputs holds the modification input item of the Variable entities.
type VariableUpdateInputsItem struct {
	// ID of the Variable entity, tries to retrieve the entity with the following unique index parts if no ID provided.
	ID object.ID `path:"-" query:"-" json:"id,omitempty"`
	// Name of the Variable entity, a part of the unique index.
	Name string `path:"-" query:"-" json:"name,omitempty"`

	// The value of variable, store in string.
	Value crypto.String `path:"-" query:"-" json:"value"`
	// The value is sensitive or not.
	Sensitive bool `path:"-" query:"-" json:"sensitive"`
	// Description of the variable.
	Description string `path:"-" query:"-" json:"description,omitempty"`
}

// ValidateWith checks the VariableUpdateInputsItem entity with the given context and client set.
func (vui *VariableUpdateInputsItem) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if vui == nil {
		return errors.New("nil receiver")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	return nil
}

// VariableUpdateInputs holds the modification input of the Variable entities,
// please tags with `path:",inline" json:",inline"` if embedding.
type VariableUpdateInputs struct {
	inputConfig `path:"-" query:"-" json:"-"`

	// Project indicates to update Variable entity CAN under the Project route.
	Project *ProjectQueryInput `path:",inline" query:"-" json:"-"`
	// Environment indicates to update Variable entity CAN under the Environment route.
	Environment *EnvironmentQueryInput `path:",inline" query:"-" json:"-"`

	// Items holds the entities to create, which MUST not be empty.
	Items []*VariableUpdateInputsItem `path:"-" query:"-" json:"items"`
}

// Model returns the Variable entities for modifying,
// after validating.
func (vui *VariableUpdateInputs) Model() []*Variable {
	if vui == nil || len(vui.Items) == 0 {
		return nil
	}

	_vs := make([]*Variable, len(vui.Items))

	for i := range vui.Items {
		_v := &Variable{
			ID:          vui.Items[i].ID,
			Name:        vui.Items[i].Name,
			Value:       vui.Items[i].Value,
			Sensitive:   vui.Items[i].Sensitive,
			Description: vui.Items[i].Description,
		}

		_vs[i] = _v
	}

	return _vs
}

// IDs returns the ID list of the Variable entities for modifying,
// after validating.
func (vui *VariableUpdateInputs) IDs() []object.ID {
	if vui == nil || len(vui.Items) == 0 {
		return nil
	}

	ids := make([]object.ID, len(vui.Items))
	for i := range vui.Items {
		ids[i] = vui.Items[i].ID
	}
	return ids
}

// Validate checks the VariableUpdateInputs entity.
func (vui *VariableUpdateInputs) Validate() error {
	if vui == nil {
		return errors.New("nil receiver")
	}

	return vui.ValidateWith(vui.inputConfig.Context, vui.inputConfig.Client, nil)
}

// ValidateWith checks the VariableUpdateInputs entity with the given context and client set.
func (vui *VariableUpdateInputs) ValidateWith(ctx context.Context, cs ClientSet, cache map[string]any) error {
	if vui == nil {
		return errors.New("nil receiver")
	}

	if len(vui.Items) == 0 {
		return errors.New("empty items")
	}

	if cache == nil {
		cache = map[string]any{}
	}

	q := cs.Variables().Query()

	// Validate when updating under the Project route.
	if vui.Project != nil {
		if err := vui.Project.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vui.Project = nil
				q.Where(
					variable.ProjectIDIsNil())
			}
		} else {
			ctx = valueContext(ctx, intercept.WithProjectInterceptor)
			q.Where(
				variable.ProjectID(vui.Project.ID))
		}
	} else {
		q.Where(
			variable.ProjectIDIsNil())
	}

	// Validate when updating under the Environment route.
	if vui.Environment != nil {
		if err := vui.Environment.ValidateWith(ctx, cs, cache); err != nil {
			if !IsBlankResourceReferError(err) {
				return err
			} else {
				vui.Environment = nil
				q.Where(
					variable.EnvironmentIDIsNil())
			}
		} else {
			q.Where(
				variable.EnvironmentID(vui.Environment.ID))
		}
	} else {
		q.Where(
			variable.EnvironmentIDIsNil())
	}

	ids := make([]object.ID, 0, len(vui.Items))
	ors := make([]predicate.Variable, 0, len(vui.Items))
	indexers := make(map[any][]int)

	for i := range vui.Items {
		if vui.Items[i] == nil {
			return errors.New("nil item")
		}

		if vui.Items[i].ID != "" {
			ids = append(ids, vui.Items[i].ID)
			ors = append(ors, variable.ID(vui.Items[i].ID))
			indexers[vui.Items[i].ID] = append(indexers[vui.Items[i].ID], i)
		} else if vui.Items[i].Name != "" {
			ors = append(ors, variable.And(
				variable.Name(vui.Items[i].Name)))
			indexerKey := fmt.Sprint("/", vui.Items[i].Name)
			indexers[indexerKey] = append(indexers[indexerKey], i)
		} else {
			return errors.New("found item hasn't identify")
		}
	}

	p := variable.IDIn(ids...)
	if len(ids) != cap(ids) {
		p = variable.Or(ors...)
	}

	es, err := q.
		Where(p).
		Select(
			variable.FieldID,
			variable.FieldName,
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
			vui.Items[j].ID = es[i].ID
			vui.Items[j].Name = es[i].Name
		}
	}

	for i := range vui.Items {
		if err := vui.Items[i].ValidateWith(ctx, cs, cache); err != nil {
			return err
		}
	}

	return nil
}

// VariableOutput holds the output of the Variable entity.
type VariableOutput struct {
	ID          object.ID     `json:"id,omitempty"`
	CreateTime  *time.Time    `json:"createTime,omitempty"`
	UpdateTime  *time.Time    `json:"updateTime,omitempty"`
	Name        string        `json:"name,omitempty"`
	Value       crypto.String `json:"value,omitempty"`
	Sensitive   bool          `json:"sensitive,omitempty"`
	Description string        `json:"description,omitempty"`

	Project     *ProjectOutput     `json:"project,omitempty"`
	Environment *EnvironmentOutput `json:"environment,omitempty"`
}

// View returns the output of Variable entity.
func (_v *Variable) View() *VariableOutput {
	return ExposeVariable(_v)
}

// View returns the output of Variable entities.
func (_vs Variables) View() []*VariableOutput {
	return ExposeVariables(_vs)
}

// ExposeVariable converts the Variable to VariableOutput.
func ExposeVariable(_v *Variable) *VariableOutput {
	if _v == nil {
		return nil
	}

	vo := &VariableOutput{
		ID:          _v.ID,
		CreateTime:  _v.CreateTime,
		UpdateTime:  _v.UpdateTime,
		Name:        _v.Name,
		Value:       _v.Value,
		Sensitive:   _v.Sensitive,
		Description: _v.Description,
	}

	if _v.Edges.Project != nil {
		vo.Project = ExposeProject(_v.Edges.Project)
	} else if _v.ProjectID != "" {
		vo.Project = &ProjectOutput{
			ID: _v.ProjectID,
		}
	}
	if _v.Edges.Environment != nil {
		vo.Environment = ExposeEnvironment(_v.Edges.Environment)
	} else if _v.EnvironmentID != "" {
		vo.Environment = &EnvironmentOutput{
			ID: _v.EnvironmentID,
		}
	}
	return vo
}

// ExposeVariables converts the Variable slice to VariableOutput pointer slice.
func ExposeVariables(_vs []*Variable) []*VariableOutput {
	if len(_vs) == 0 {
		return nil
	}

	vos := make([]*VariableOutput, len(_vs))
	for i := range _vs {
		vos[i] = ExposeVariable(_vs[i])
	}
	return vos
}
