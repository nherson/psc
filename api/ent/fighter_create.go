// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/fight"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/ent/fighteralias"
	"github.com/nherson/psc/api/ent/fighterresults"
)

// FighterCreate is the builder for creating a Fighter entity.
type FighterCreate struct {
	config
	mutation *FighterMutation
	hooks    []Hook
}

// AddFightIDs adds the "fights" edge to the Fight entity by IDs.
func (fc *FighterCreate) AddFightIDs(ids ...int) *FighterCreate {
	fc.mutation.AddFightIDs(ids...)
	return fc
}

// AddFights adds the "fights" edges to the Fight entity.
func (fc *FighterCreate) AddFights(f ...*Fight) *FighterCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddFightIDs(ids...)
}

// AddFighterAliasIDs adds the "fighter_aliases" edge to the FighterAlias entity by IDs.
func (fc *FighterCreate) AddFighterAliasIDs(ids ...int) *FighterCreate {
	fc.mutation.AddFighterAliasIDs(ids...)
	return fc
}

// AddFighterAliases adds the "fighter_aliases" edges to the FighterAlias entity.
func (fc *FighterCreate) AddFighterAliases(f ...*FighterAlias) *FighterCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddFighterAliasIDs(ids...)
}

// AddFighterResultIDs adds the "fighter_results" edge to the FighterResults entity by IDs.
func (fc *FighterCreate) AddFighterResultIDs(ids ...int) *FighterCreate {
	fc.mutation.AddFighterResultIDs(ids...)
	return fc
}

// AddFighterResults adds the "fighter_results" edges to the FighterResults entity.
func (fc *FighterCreate) AddFighterResults(f ...*FighterResults) *FighterCreate {
	ids := make([]int, len(f))
	for i := range f {
		ids[i] = f[i].ID
	}
	return fc.AddFighterResultIDs(ids...)
}

// Mutation returns the FighterMutation object of the builder.
func (fc *FighterCreate) Mutation() *FighterMutation {
	return fc.mutation
}

// Save creates the Fighter in the database.
func (fc *FighterCreate) Save(ctx context.Context) (*Fighter, error) {
	return withHooks[*Fighter, FighterMutation](ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FighterCreate) SaveX(ctx context.Context) *Fighter {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FighterCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FighterCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FighterCreate) check() error {
	return nil
}

func (fc *FighterCreate) sqlSave(ctx context.Context) (*Fighter, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FighterCreate) createSpec() (*Fighter, *sqlgraph.CreateSpec) {
	var (
		_node = &Fighter{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(fighter.Table, sqlgraph.NewFieldSpec(fighter.FieldID, field.TypeInt))
	)
	if nodes := fc.mutation.FightsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   fighter.FightsTable,
			Columns: fighter.FightsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fight.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FighterAliasesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   fighter.FighterAliasesTable,
			Columns: []string{fighter.FighterAliasesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fighteralias.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	if nodes := fc.mutation.FighterResultsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   fighter.FighterResultsTable,
			Columns: []string{fighter.FighterResultsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fighterresults.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FighterCreateBulk is the builder for creating many Fighter entities in bulk.
type FighterCreateBulk struct {
	config
	builders []*FighterCreate
}

// Save creates the Fighter entities in the database.
func (fcb *FighterCreateBulk) Save(ctx context.Context) ([]*Fighter, error) {
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Fighter, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FighterMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FighterCreateBulk) SaveX(ctx context.Context) []*Fighter {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FighterCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FighterCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
