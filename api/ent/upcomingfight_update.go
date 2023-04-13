// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/predicate"
	"github.com/nherson/psc/api/ent/upcomingfight"
)

// UpcomingFightUpdate is the builder for updating UpcomingFight entities.
type UpcomingFightUpdate struct {
	config
	hooks    []Hook
	mutation *UpcomingFightMutation
}

// Where appends a list predicates to the UpcomingFightUpdate builder.
func (ufu *UpcomingFightUpdate) Where(ps ...predicate.UpcomingFight) *UpcomingFightUpdate {
	ufu.mutation.Where(ps...)
	return ufu
}

// SetUpdatedAt sets the "updated_at" field.
func (ufu *UpcomingFightUpdate) SetUpdatedAt(t time.Time) *UpcomingFightUpdate {
	ufu.mutation.SetUpdatedAt(t)
	return ufu
}

// Mutation returns the UpcomingFightMutation object of the builder.
func (ufu *UpcomingFightUpdate) Mutation() *UpcomingFightMutation {
	return ufu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ufu *UpcomingFightUpdate) Save(ctx context.Context) (int, error) {
	ufu.defaults()
	return withHooks[int, UpcomingFightMutation](ctx, ufu.sqlSave, ufu.mutation, ufu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ufu *UpcomingFightUpdate) SaveX(ctx context.Context) int {
	affected, err := ufu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ufu *UpcomingFightUpdate) Exec(ctx context.Context) error {
	_, err := ufu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufu *UpcomingFightUpdate) ExecX(ctx context.Context) {
	if err := ufu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufu *UpcomingFightUpdate) defaults() {
	if _, ok := ufu.mutation.UpdatedAt(); !ok {
		v := upcomingfight.UpdateDefaultUpdatedAt()
		ufu.mutation.SetUpdatedAt(v)
	}
}

func (ufu *UpcomingFightUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(upcomingfight.Table, upcomingfight.Columns, sqlgraph.NewFieldSpec(upcomingfight.FieldID, field.TypeInt))
	if ps := ufu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufu.mutation.UpdatedAt(); ok {
		_spec.SetField(upcomingfight.FieldUpdatedAt, field.TypeTime, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ufu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upcomingfight.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ufu.mutation.done = true
	return n, nil
}

// UpcomingFightUpdateOne is the builder for updating a single UpcomingFight entity.
type UpcomingFightUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UpcomingFightMutation
}

// SetUpdatedAt sets the "updated_at" field.
func (ufuo *UpcomingFightUpdateOne) SetUpdatedAt(t time.Time) *UpcomingFightUpdateOne {
	ufuo.mutation.SetUpdatedAt(t)
	return ufuo
}

// Mutation returns the UpcomingFightMutation object of the builder.
func (ufuo *UpcomingFightUpdateOne) Mutation() *UpcomingFightMutation {
	return ufuo.mutation
}

// Where appends a list predicates to the UpcomingFightUpdate builder.
func (ufuo *UpcomingFightUpdateOne) Where(ps ...predicate.UpcomingFight) *UpcomingFightUpdateOne {
	ufuo.mutation.Where(ps...)
	return ufuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ufuo *UpcomingFightUpdateOne) Select(field string, fields ...string) *UpcomingFightUpdateOne {
	ufuo.fields = append([]string{field}, fields...)
	return ufuo
}

// Save executes the query and returns the updated UpcomingFight entity.
func (ufuo *UpcomingFightUpdateOne) Save(ctx context.Context) (*UpcomingFight, error) {
	ufuo.defaults()
	return withHooks[*UpcomingFight, UpcomingFightMutation](ctx, ufuo.sqlSave, ufuo.mutation, ufuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ufuo *UpcomingFightUpdateOne) SaveX(ctx context.Context) *UpcomingFight {
	node, err := ufuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ufuo *UpcomingFightUpdateOne) Exec(ctx context.Context) error {
	_, err := ufuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufuo *UpcomingFightUpdateOne) ExecX(ctx context.Context) {
	if err := ufuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufuo *UpcomingFightUpdateOne) defaults() {
	if _, ok := ufuo.mutation.UpdatedAt(); !ok {
		v := upcomingfight.UpdateDefaultUpdatedAt()
		ufuo.mutation.SetUpdatedAt(v)
	}
}

func (ufuo *UpcomingFightUpdateOne) sqlSave(ctx context.Context) (_node *UpcomingFight, err error) {
	_spec := sqlgraph.NewUpdateSpec(upcomingfight.Table, upcomingfight.Columns, sqlgraph.NewFieldSpec(upcomingfight.FieldID, field.TypeInt))
	id, ok := ufuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "UpcomingFight.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ufuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upcomingfight.FieldID)
		for _, f := range fields {
			if !upcomingfight.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != upcomingfight.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ufuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufuo.mutation.UpdatedAt(); ok {
		_spec.SetField(upcomingfight.FieldUpdatedAt, field.TypeTime, value)
	}
	_node = &UpcomingFight{config: ufuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ufuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{upcomingfight.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ufuo.mutation.done = true
	return _node, nil
}
