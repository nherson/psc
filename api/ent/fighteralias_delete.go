// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/fighteralias"
	"github.com/nherson/psc/api/ent/predicate"
)

// FighterAliasDelete is the builder for deleting a FighterAlias entity.
type FighterAliasDelete struct {
	config
	hooks    []Hook
	mutation *FighterAliasMutation
}

// Where appends a list predicates to the FighterAliasDelete builder.
func (fad *FighterAliasDelete) Where(ps ...predicate.FighterAlias) *FighterAliasDelete {
	fad.mutation.Where(ps...)
	return fad
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (fad *FighterAliasDelete) Exec(ctx context.Context) (int, error) {
	return withHooks[int, FighterAliasMutation](ctx, fad.sqlExec, fad.mutation, fad.hooks)
}

// ExecX is like Exec, but panics if an error occurs.
func (fad *FighterAliasDelete) ExecX(ctx context.Context) int {
	n, err := fad.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (fad *FighterAliasDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := sqlgraph.NewDeleteSpec(fighteralias.Table, sqlgraph.NewFieldSpec(fighteralias.FieldID, field.TypeInt))
	if ps := fad.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, fad.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	fad.mutation.done = true
	return affected, err
}

// FighterAliasDeleteOne is the builder for deleting a single FighterAlias entity.
type FighterAliasDeleteOne struct {
	fad *FighterAliasDelete
}

// Where appends a list predicates to the FighterAliasDelete builder.
func (fado *FighterAliasDeleteOne) Where(ps ...predicate.FighterAlias) *FighterAliasDeleteOne {
	fado.fad.mutation.Where(ps...)
	return fado
}

// Exec executes the deletion query.
func (fado *FighterAliasDeleteOne) Exec(ctx context.Context) error {
	n, err := fado.fad.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{fighteralias.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (fado *FighterAliasDeleteOne) ExecX(ctx context.Context) {
	if err := fado.Exec(ctx); err != nil {
		panic(err)
	}
}
