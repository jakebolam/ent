// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"math/big"
	"net/url"

	"entgo.io/ent/dialect/gremlin"
	"entgo.io/ent/dialect/gremlin/graph/dsl"
	"entgo.io/ent/dialect/gremlin/graph/dsl/g"
	"entgo.io/ent/entc/integration/gremlin/ent/exvaluescan"
)

// ExValueScanCreate is the builder for creating a ExValueScan entity.
type ExValueScanCreate struct {
	config
	mutation *ExValueScanMutation
	hooks    []Hook
}

// SetBinary sets the "binary" field.
func (evsc *ExValueScanCreate) SetBinary(u *url.URL) *ExValueScanCreate {
	evsc.mutation.SetBinary(u)
	return evsc
}

// SetBinaryOptional sets the "binary_optional" field.
func (evsc *ExValueScanCreate) SetBinaryOptional(u *url.URL) *ExValueScanCreate {
	evsc.mutation.SetBinaryOptional(u)
	return evsc
}

// SetText sets the "text" field.
func (evsc *ExValueScanCreate) SetText(b *big.Int) *ExValueScanCreate {
	evsc.mutation.SetText(b)
	return evsc
}

// SetTextOptional sets the "text_optional" field.
func (evsc *ExValueScanCreate) SetTextOptional(b *big.Int) *ExValueScanCreate {
	evsc.mutation.SetTextOptional(b)
	return evsc
}

// SetBase64 sets the "base64" field.
func (evsc *ExValueScanCreate) SetBase64(s string) *ExValueScanCreate {
	evsc.mutation.SetBase64(s)
	return evsc
}

// SetCustom sets the "custom" field.
func (evsc *ExValueScanCreate) SetCustom(s string) *ExValueScanCreate {
	evsc.mutation.SetCustom(s)
	return evsc
}

// SetCustomOptional sets the "custom_optional" field.
func (evsc *ExValueScanCreate) SetCustomOptional(s string) *ExValueScanCreate {
	evsc.mutation.SetCustomOptional(s)
	return evsc
}

// SetNillableCustomOptional sets the "custom_optional" field if the given value is not nil.
func (evsc *ExValueScanCreate) SetNillableCustomOptional(s *string) *ExValueScanCreate {
	if s != nil {
		evsc.SetCustomOptional(*s)
	}
	return evsc
}

// Mutation returns the ExValueScanMutation object of the builder.
func (evsc *ExValueScanCreate) Mutation() *ExValueScanMutation {
	return evsc.mutation
}

// Save creates the ExValueScan in the database.
func (evsc *ExValueScanCreate) Save(ctx context.Context) (*ExValueScan, error) {
	return withHooks(ctx, evsc.gremlinSave, evsc.mutation, evsc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (evsc *ExValueScanCreate) SaveX(ctx context.Context) *ExValueScan {
	v, err := evsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (evsc *ExValueScanCreate) Exec(ctx context.Context) error {
	_, err := evsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (evsc *ExValueScanCreate) ExecX(ctx context.Context) {
	if err := evsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (evsc *ExValueScanCreate) check() error {
	if _, ok := evsc.mutation.Binary(); !ok {
		return &ValidationError{Name: "binary", err: errors.New(`ent: missing required field "ExValueScan.binary"`)}
	}
	if _, ok := evsc.mutation.Text(); !ok {
		return &ValidationError{Name: "text", err: errors.New(`ent: missing required field "ExValueScan.text"`)}
	}
	if _, ok := evsc.mutation.Base64(); !ok {
		return &ValidationError{Name: "base64", err: errors.New(`ent: missing required field "ExValueScan.base64"`)}
	}
	if _, ok := evsc.mutation.Custom(); !ok {
		return &ValidationError{Name: "custom", err: errors.New(`ent: missing required field "ExValueScan.custom"`)}
	}
	return nil
}

func (evsc *ExValueScanCreate) gremlinSave(ctx context.Context) (*ExValueScan, error) {
	if err := evsc.check(); err != nil {
		return nil, err
	}
	res := &gremlin.Response{}
	query, bindings := evsc.gremlin().Query()
	if err := evsc.driver.Exec(ctx, query, bindings, res); err != nil {
		return nil, err
	}
	if err, ok := isConstantError(res); ok {
		return nil, err
	}
	evs := &ExValueScan{config: evsc.config}
	if err := evs.FromResponse(res); err != nil {
		return nil, err
	}
	evsc.mutation.id = &evs.ID
	evsc.mutation.done = true
	return evs, nil
}

func (evsc *ExValueScanCreate) gremlin() *dsl.Traversal {
	v := g.AddV(exvaluescan.Label)
	if value, ok := evsc.mutation.Binary(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinary, value)
	}
	if value, ok := evsc.mutation.BinaryOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBinaryOptional, value)
	}
	if value, ok := evsc.mutation.Text(); ok {
		v.Property(dsl.Single, exvaluescan.FieldText, value)
	}
	if value, ok := evsc.mutation.TextOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldTextOptional, value)
	}
	if value, ok := evsc.mutation.Base64(); ok {
		v.Property(dsl.Single, exvaluescan.FieldBase64, value)
	}
	if value, ok := evsc.mutation.Custom(); ok {
		v.Property(dsl.Single, exvaluescan.FieldCustom, value)
	}
	if value, ok := evsc.mutation.CustomOptional(); ok {
		v.Property(dsl.Single, exvaluescan.FieldCustomOptional, value)
	}
	return v.ValueMap(true)
}

// ExValueScanCreateBulk is the builder for creating many ExValueScan entities in bulk.
type ExValueScanCreateBulk struct {
	config
	builders []*ExValueScanCreate
}
