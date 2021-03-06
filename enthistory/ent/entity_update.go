// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ariga/edit-twitter-example-app/enthistory/ent/entity"
	"github.com/ariga/edit-twitter-example-app/enthistory/ent/predicate"
)

// EntityUpdate is the builder for updating Entity entities.
type EntityUpdate struct {
	config
	hooks    []Hook
	mutation *EntityMutation
}

// Where appends a list predicates to the EntityUpdate builder.
func (eu *EntityUpdate) Where(ps ...predicate.Entity) *EntityUpdate {
	eu.mutation.Where(ps...)
	return eu
}

// SetData sets the "data" field.
func (eu *EntityUpdate) SetData(s string) *EntityUpdate {
	eu.mutation.SetData(s)
	return eu
}

// SetIsFun sets the "isFun" field.
func (eu *EntityUpdate) SetIsFun(b bool) *EntityUpdate {
	eu.mutation.SetIsFun(b)
	return eu
}

// SetNillableIsFun sets the "isFun" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableIsFun(b *bool) *EntityUpdate {
	if b != nil {
		eu.SetIsFun(*b)
	}
	return eu
}

// ClearIsFun clears the value of the "isFun" field.
func (eu *EntityUpdate) ClearIsFun() *EntityUpdate {
	eu.mutation.ClearIsFun()
	return eu
}

// SetCounter sets the "counter" field.
func (eu *EntityUpdate) SetCounter(i int) *EntityUpdate {
	eu.mutation.ResetCounter()
	eu.mutation.SetCounter(i)
	return eu
}

// SetNillableCounter sets the "counter" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableCounter(i *int) *EntityUpdate {
	if i != nil {
		eu.SetCounter(*i)
	}
	return eu
}

// AddCounter adds i to the "counter" field.
func (eu *EntityUpdate) AddCounter(i int) *EntityUpdate {
	eu.mutation.AddCounter(i)
	return eu
}

// ClearCounter clears the value of the "counter" field.
func (eu *EntityUpdate) ClearCounter() *EntityUpdate {
	eu.mutation.ClearCounter()
	return eu
}

// SetTimestamp sets the "timestamp" field.
func (eu *EntityUpdate) SetTimestamp(t time.Time) *EntityUpdate {
	eu.mutation.SetTimestamp(t)
	return eu
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (eu *EntityUpdate) SetNillableTimestamp(t *time.Time) *EntityUpdate {
	if t != nil {
		eu.SetTimestamp(*t)
	}
	return eu
}

// ClearTimestamp clears the value of the "timestamp" field.
func (eu *EntityUpdate) ClearTimestamp() *EntityUpdate {
	eu.mutation.ClearTimestamp()
	return eu
}

// SetStrings sets the "strings" field.
func (eu *EntityUpdate) SetStrings(s []string) *EntityUpdate {
	eu.mutation.SetStrings(s)
	return eu
}

// ClearStrings clears the value of the "strings" field.
func (eu *EntityUpdate) ClearStrings() *EntityUpdate {
	eu.mutation.ClearStrings()
	return eu
}

// Mutation returns the EntityMutation object of the builder.
func (eu *EntityUpdate) Mutation() *EntityMutation {
	return eu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (eu *EntityUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(eu.hooks) == 0 {
		affected, err = eu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			eu.mutation = mutation
			affected, err = eu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(eu.hooks) - 1; i >= 0; i-- {
			if eu.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = eu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, eu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (eu *EntityUpdate) SaveX(ctx context.Context) int {
	affected, err := eu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (eu *EntityUpdate) Exec(ctx context.Context) error {
	_, err := eu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (eu *EntityUpdate) ExecX(ctx context.Context) {
	if err := eu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (eu *EntityUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entity.Table,
			Columns: entity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entity.FieldID,
			},
		},
	}
	if ps := eu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := eu.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entity.FieldData,
		})
	}
	if value, ok := eu.mutation.IsFun(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entity.FieldIsFun,
		})
	}
	if eu.mutation.IsFunCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entity.FieldIsFun,
		})
	}
	if value, ok := eu.mutation.Counter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entity.FieldCounter,
		})
	}
	if value, ok := eu.mutation.AddedCounter(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entity.FieldCounter,
		})
	}
	if eu.mutation.CounterCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: entity.FieldCounter,
		})
	}
	if value, ok := eu.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entity.FieldTimestamp,
		})
	}
	if eu.mutation.TimestampCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entity.FieldTimestamp,
		})
	}
	if value, ok := eu.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: entity.FieldStrings,
		})
	}
	if eu.mutation.StringsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: entity.FieldStrings,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, eu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// EntityUpdateOne is the builder for updating a single Entity entity.
type EntityUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *EntityMutation
}

// SetData sets the "data" field.
func (euo *EntityUpdateOne) SetData(s string) *EntityUpdateOne {
	euo.mutation.SetData(s)
	return euo
}

// SetIsFun sets the "isFun" field.
func (euo *EntityUpdateOne) SetIsFun(b bool) *EntityUpdateOne {
	euo.mutation.SetIsFun(b)
	return euo
}

// SetNillableIsFun sets the "isFun" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableIsFun(b *bool) *EntityUpdateOne {
	if b != nil {
		euo.SetIsFun(*b)
	}
	return euo
}

// ClearIsFun clears the value of the "isFun" field.
func (euo *EntityUpdateOne) ClearIsFun() *EntityUpdateOne {
	euo.mutation.ClearIsFun()
	return euo
}

// SetCounter sets the "counter" field.
func (euo *EntityUpdateOne) SetCounter(i int) *EntityUpdateOne {
	euo.mutation.ResetCounter()
	euo.mutation.SetCounter(i)
	return euo
}

// SetNillableCounter sets the "counter" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableCounter(i *int) *EntityUpdateOne {
	if i != nil {
		euo.SetCounter(*i)
	}
	return euo
}

// AddCounter adds i to the "counter" field.
func (euo *EntityUpdateOne) AddCounter(i int) *EntityUpdateOne {
	euo.mutation.AddCounter(i)
	return euo
}

// ClearCounter clears the value of the "counter" field.
func (euo *EntityUpdateOne) ClearCounter() *EntityUpdateOne {
	euo.mutation.ClearCounter()
	return euo
}

// SetTimestamp sets the "timestamp" field.
func (euo *EntityUpdateOne) SetTimestamp(t time.Time) *EntityUpdateOne {
	euo.mutation.SetTimestamp(t)
	return euo
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (euo *EntityUpdateOne) SetNillableTimestamp(t *time.Time) *EntityUpdateOne {
	if t != nil {
		euo.SetTimestamp(*t)
	}
	return euo
}

// ClearTimestamp clears the value of the "timestamp" field.
func (euo *EntityUpdateOne) ClearTimestamp() *EntityUpdateOne {
	euo.mutation.ClearTimestamp()
	return euo
}

// SetStrings sets the "strings" field.
func (euo *EntityUpdateOne) SetStrings(s []string) *EntityUpdateOne {
	euo.mutation.SetStrings(s)
	return euo
}

// ClearStrings clears the value of the "strings" field.
func (euo *EntityUpdateOne) ClearStrings() *EntityUpdateOne {
	euo.mutation.ClearStrings()
	return euo
}

// Mutation returns the EntityMutation object of the builder.
func (euo *EntityUpdateOne) Mutation() *EntityMutation {
	return euo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (euo *EntityUpdateOne) Select(field string, fields ...string) *EntityUpdateOne {
	euo.fields = append([]string{field}, fields...)
	return euo
}

// Save executes the query and returns the updated Entity entity.
func (euo *EntityUpdateOne) Save(ctx context.Context) (*Entity, error) {
	var (
		err  error
		node *Entity
	)
	if len(euo.hooks) == 0 {
		node, err = euo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*EntityMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			euo.mutation = mutation
			node, err = euo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(euo.hooks) - 1; i >= 0; i-- {
			if euo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = euo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, euo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (euo *EntityUpdateOne) SaveX(ctx context.Context) *Entity {
	node, err := euo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (euo *EntityUpdateOne) Exec(ctx context.Context) error {
	_, err := euo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (euo *EntityUpdateOne) ExecX(ctx context.Context) {
	if err := euo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (euo *EntityUpdateOne) sqlSave(ctx context.Context) (_node *Entity, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   entity.Table,
			Columns: entity.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt,
				Column: entity.FieldID,
			},
		},
	}
	id, ok := euo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Entity.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := euo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, entity.FieldID)
		for _, f := range fields {
			if !entity.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != entity.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := euo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := euo.mutation.Data(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: entity.FieldData,
		})
	}
	if value, ok := euo.mutation.IsFun(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Value:  value,
			Column: entity.FieldIsFun,
		})
	}
	if euo.mutation.IsFunCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeBool,
			Column: entity.FieldIsFun,
		})
	}
	if value, ok := euo.mutation.Counter(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entity.FieldCounter,
		})
	}
	if value, ok := euo.mutation.AddedCounter(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: entity.FieldCounter,
		})
	}
	if euo.mutation.CounterCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Column: entity.FieldCounter,
		})
	}
	if value, ok := euo.mutation.Timestamp(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: entity.FieldTimestamp,
		})
	}
	if euo.mutation.TimestampCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: entity.FieldTimestamp,
		})
	}
	if value, ok := euo.mutation.Strings(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Value:  value,
			Column: entity.FieldStrings,
		})
	}
	if euo.mutation.StringsCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeJSON,
			Column: entity.FieldStrings,
		})
	}
	_node = &Entity{config: euo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, euo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{entity.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}
