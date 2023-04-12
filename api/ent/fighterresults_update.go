// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/fight"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/ent/fighterresults"
	"github.com/nherson/psc/api/ent/predicate"
)

// FighterResultsUpdate is the builder for updating FighterResults entities.
type FighterResultsUpdate struct {
	config
	hooks    []Hook
	mutation *FighterResultsMutation
}

// Where appends a list predicates to the FighterResultsUpdate builder.
func (fru *FighterResultsUpdate) Where(ps ...predicate.FighterResults) *FighterResultsUpdate {
	fru.mutation.Where(ps...)
	return fru
}

// SetFighterID sets the "fighter_id" field.
func (fru *FighterResultsUpdate) SetFighterID(i int) *FighterResultsUpdate {
	fru.mutation.SetFighterID(i)
	return fru
}

// SetFightID sets the "fight_id" field.
func (fru *FighterResultsUpdate) SetFightID(i int) *FighterResultsUpdate {
	fru.mutation.SetFightID(i)
	return fru
}

// SetSignificantStrikesLanded sets the "significant_strikes_landed" field.
func (fru *FighterResultsUpdate) SetSignificantStrikesLanded(i int) *FighterResultsUpdate {
	fru.mutation.ResetSignificantStrikesLanded()
	fru.mutation.SetSignificantStrikesLanded(i)
	return fru
}

// AddSignificantStrikesLanded adds i to the "significant_strikes_landed" field.
func (fru *FighterResultsUpdate) AddSignificantStrikesLanded(i int) *FighterResultsUpdate {
	fru.mutation.AddSignificantStrikesLanded(i)
	return fru
}

// SetTakedowns sets the "takedowns" field.
func (fru *FighterResultsUpdate) SetTakedowns(i int) *FighterResultsUpdate {
	fru.mutation.ResetTakedowns()
	fru.mutation.SetTakedowns(i)
	return fru
}

// AddTakedowns adds i to the "takedowns" field.
func (fru *FighterResultsUpdate) AddTakedowns(i int) *FighterResultsUpdate {
	fru.mutation.AddTakedowns(i)
	return fru
}

// SetKnockdowns sets the "knockdowns" field.
func (fru *FighterResultsUpdate) SetKnockdowns(i int) *FighterResultsUpdate {
	fru.mutation.ResetKnockdowns()
	fru.mutation.SetKnockdowns(i)
	return fru
}

// AddKnockdowns adds i to the "knockdowns" field.
func (fru *FighterResultsUpdate) AddKnockdowns(i int) *FighterResultsUpdate {
	fru.mutation.AddKnockdowns(i)
	return fru
}

// SetControlTimeSeconds sets the "control_time_seconds" field.
func (fru *FighterResultsUpdate) SetControlTimeSeconds(i int) *FighterResultsUpdate {
	fru.mutation.ResetControlTimeSeconds()
	fru.mutation.SetControlTimeSeconds(i)
	return fru
}

// AddControlTimeSeconds adds i to the "control_time_seconds" field.
func (fru *FighterResultsUpdate) AddControlTimeSeconds(i int) *FighterResultsUpdate {
	fru.mutation.AddControlTimeSeconds(i)
	return fru
}

// SetWinByStoppage sets the "win_by_stoppage" field.
func (fru *FighterResultsUpdate) SetWinByStoppage(b bool) *FighterResultsUpdate {
	fru.mutation.SetWinByStoppage(b)
	return fru
}

// SetLossByStoppage sets the "loss_by_stoppage" field.
func (fru *FighterResultsUpdate) SetLossByStoppage(b bool) *FighterResultsUpdate {
	fru.mutation.SetLossByStoppage(b)
	return fru
}

// SetMissedWeight sets the "missed_weight" field.
func (fru *FighterResultsUpdate) SetMissedWeight(b bool) *FighterResultsUpdate {
	fru.mutation.SetMissedWeight(b)
	return fru
}

// SetNillableMissedWeight sets the "missed_weight" field if the given value is not nil.
func (fru *FighterResultsUpdate) SetNillableMissedWeight(b *bool) *FighterResultsUpdate {
	if b != nil {
		fru.SetMissedWeight(*b)
	}
	return fru
}

// ClearMissedWeight clears the value of the "missed_weight" field.
func (fru *FighterResultsUpdate) ClearMissedWeight() *FighterResultsUpdate {
	fru.mutation.ClearMissedWeight()
	return fru
}

// SetFighter sets the "fighter" edge to the Fighter entity.
func (fru *FighterResultsUpdate) SetFighter(f *Fighter) *FighterResultsUpdate {
	return fru.SetFighterID(f.ID)
}

// SetFight sets the "fight" edge to the Fight entity.
func (fru *FighterResultsUpdate) SetFight(f *Fight) *FighterResultsUpdate {
	return fru.SetFightID(f.ID)
}

// Mutation returns the FighterResultsMutation object of the builder.
func (fru *FighterResultsUpdate) Mutation() *FighterResultsMutation {
	return fru.mutation
}

// ClearFighter clears the "fighter" edge to the Fighter entity.
func (fru *FighterResultsUpdate) ClearFighter() *FighterResultsUpdate {
	fru.mutation.ClearFighter()
	return fru
}

// ClearFight clears the "fight" edge to the Fight entity.
func (fru *FighterResultsUpdate) ClearFight() *FighterResultsUpdate {
	fru.mutation.ClearFight()
	return fru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (fru *FighterResultsUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, FighterResultsMutation](ctx, fru.sqlSave, fru.mutation, fru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fru *FighterResultsUpdate) SaveX(ctx context.Context) int {
	affected, err := fru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (fru *FighterResultsUpdate) Exec(ctx context.Context) error {
	_, err := fru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fru *FighterResultsUpdate) ExecX(ctx context.Context) {
	if err := fru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fru *FighterResultsUpdate) check() error {
	if _, ok := fru.mutation.FighterID(); fru.mutation.FighterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FighterResults.fighter"`)
	}
	if _, ok := fru.mutation.FightID(); fru.mutation.FightCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FighterResults.fight"`)
	}
	return nil
}

func (fru *FighterResultsUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := fru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(fighterresults.Table, fighterresults.Columns, sqlgraph.NewFieldSpec(fighterresults.FieldID, field.TypeInt))
	if ps := fru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fru.mutation.SignificantStrikesLanded(); ok {
		_spec.SetField(fighterresults.FieldSignificantStrikesLanded, field.TypeInt, value)
	}
	if value, ok := fru.mutation.AddedSignificantStrikesLanded(); ok {
		_spec.AddField(fighterresults.FieldSignificantStrikesLanded, field.TypeInt, value)
	}
	if value, ok := fru.mutation.Takedowns(); ok {
		_spec.SetField(fighterresults.FieldTakedowns, field.TypeInt, value)
	}
	if value, ok := fru.mutation.AddedTakedowns(); ok {
		_spec.AddField(fighterresults.FieldTakedowns, field.TypeInt, value)
	}
	if value, ok := fru.mutation.Knockdowns(); ok {
		_spec.SetField(fighterresults.FieldKnockdowns, field.TypeInt, value)
	}
	if value, ok := fru.mutation.AddedKnockdowns(); ok {
		_spec.AddField(fighterresults.FieldKnockdowns, field.TypeInt, value)
	}
	if value, ok := fru.mutation.ControlTimeSeconds(); ok {
		_spec.SetField(fighterresults.FieldControlTimeSeconds, field.TypeInt, value)
	}
	if value, ok := fru.mutation.AddedControlTimeSeconds(); ok {
		_spec.AddField(fighterresults.FieldControlTimeSeconds, field.TypeInt, value)
	}
	if value, ok := fru.mutation.WinByStoppage(); ok {
		_spec.SetField(fighterresults.FieldWinByStoppage, field.TypeBool, value)
	}
	if value, ok := fru.mutation.LossByStoppage(); ok {
		_spec.SetField(fighterresults.FieldLossByStoppage, field.TypeBool, value)
	}
	if value, ok := fru.mutation.MissedWeight(); ok {
		_spec.SetField(fighterresults.FieldMissedWeight, field.TypeBool, value)
	}
	if fru.mutation.MissedWeightCleared() {
		_spec.ClearField(fighterresults.FieldMissedWeight, field.TypeBool)
	}
	if fru.mutation.FighterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fighterresults.FighterTable,
			Columns: []string{fighterresults.FighterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fighter.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fru.mutation.FighterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fighterresults.FighterTable,
			Columns: []string{fighterresults.FighterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fighter.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fru.mutation.FightCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fighterresults.FightTable,
			Columns: []string{fighterresults.FightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fru.mutation.FightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fighterresults.FightTable,
			Columns: []string{fighterresults.FightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fight.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, fru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fighterresults.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	fru.mutation.done = true
	return n, nil
}

// FighterResultsUpdateOne is the builder for updating a single FighterResults entity.
type FighterResultsUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *FighterResultsMutation
}

// SetFighterID sets the "fighter_id" field.
func (fruo *FighterResultsUpdateOne) SetFighterID(i int) *FighterResultsUpdateOne {
	fruo.mutation.SetFighterID(i)
	return fruo
}

// SetFightID sets the "fight_id" field.
func (fruo *FighterResultsUpdateOne) SetFightID(i int) *FighterResultsUpdateOne {
	fruo.mutation.SetFightID(i)
	return fruo
}

// SetSignificantStrikesLanded sets the "significant_strikes_landed" field.
func (fruo *FighterResultsUpdateOne) SetSignificantStrikesLanded(i int) *FighterResultsUpdateOne {
	fruo.mutation.ResetSignificantStrikesLanded()
	fruo.mutation.SetSignificantStrikesLanded(i)
	return fruo
}

// AddSignificantStrikesLanded adds i to the "significant_strikes_landed" field.
func (fruo *FighterResultsUpdateOne) AddSignificantStrikesLanded(i int) *FighterResultsUpdateOne {
	fruo.mutation.AddSignificantStrikesLanded(i)
	return fruo
}

// SetTakedowns sets the "takedowns" field.
func (fruo *FighterResultsUpdateOne) SetTakedowns(i int) *FighterResultsUpdateOne {
	fruo.mutation.ResetTakedowns()
	fruo.mutation.SetTakedowns(i)
	return fruo
}

// AddTakedowns adds i to the "takedowns" field.
func (fruo *FighterResultsUpdateOne) AddTakedowns(i int) *FighterResultsUpdateOne {
	fruo.mutation.AddTakedowns(i)
	return fruo
}

// SetKnockdowns sets the "knockdowns" field.
func (fruo *FighterResultsUpdateOne) SetKnockdowns(i int) *FighterResultsUpdateOne {
	fruo.mutation.ResetKnockdowns()
	fruo.mutation.SetKnockdowns(i)
	return fruo
}

// AddKnockdowns adds i to the "knockdowns" field.
func (fruo *FighterResultsUpdateOne) AddKnockdowns(i int) *FighterResultsUpdateOne {
	fruo.mutation.AddKnockdowns(i)
	return fruo
}

// SetControlTimeSeconds sets the "control_time_seconds" field.
func (fruo *FighterResultsUpdateOne) SetControlTimeSeconds(i int) *FighterResultsUpdateOne {
	fruo.mutation.ResetControlTimeSeconds()
	fruo.mutation.SetControlTimeSeconds(i)
	return fruo
}

// AddControlTimeSeconds adds i to the "control_time_seconds" field.
func (fruo *FighterResultsUpdateOne) AddControlTimeSeconds(i int) *FighterResultsUpdateOne {
	fruo.mutation.AddControlTimeSeconds(i)
	return fruo
}

// SetWinByStoppage sets the "win_by_stoppage" field.
func (fruo *FighterResultsUpdateOne) SetWinByStoppage(b bool) *FighterResultsUpdateOne {
	fruo.mutation.SetWinByStoppage(b)
	return fruo
}

// SetLossByStoppage sets the "loss_by_stoppage" field.
func (fruo *FighterResultsUpdateOne) SetLossByStoppage(b bool) *FighterResultsUpdateOne {
	fruo.mutation.SetLossByStoppage(b)
	return fruo
}

// SetMissedWeight sets the "missed_weight" field.
func (fruo *FighterResultsUpdateOne) SetMissedWeight(b bool) *FighterResultsUpdateOne {
	fruo.mutation.SetMissedWeight(b)
	return fruo
}

// SetNillableMissedWeight sets the "missed_weight" field if the given value is not nil.
func (fruo *FighterResultsUpdateOne) SetNillableMissedWeight(b *bool) *FighterResultsUpdateOne {
	if b != nil {
		fruo.SetMissedWeight(*b)
	}
	return fruo
}

// ClearMissedWeight clears the value of the "missed_weight" field.
func (fruo *FighterResultsUpdateOne) ClearMissedWeight() *FighterResultsUpdateOne {
	fruo.mutation.ClearMissedWeight()
	return fruo
}

// SetFighter sets the "fighter" edge to the Fighter entity.
func (fruo *FighterResultsUpdateOne) SetFighter(f *Fighter) *FighterResultsUpdateOne {
	return fruo.SetFighterID(f.ID)
}

// SetFight sets the "fight" edge to the Fight entity.
func (fruo *FighterResultsUpdateOne) SetFight(f *Fight) *FighterResultsUpdateOne {
	return fruo.SetFightID(f.ID)
}

// Mutation returns the FighterResultsMutation object of the builder.
func (fruo *FighterResultsUpdateOne) Mutation() *FighterResultsMutation {
	return fruo.mutation
}

// ClearFighter clears the "fighter" edge to the Fighter entity.
func (fruo *FighterResultsUpdateOne) ClearFighter() *FighterResultsUpdateOne {
	fruo.mutation.ClearFighter()
	return fruo
}

// ClearFight clears the "fight" edge to the Fight entity.
func (fruo *FighterResultsUpdateOne) ClearFight() *FighterResultsUpdateOne {
	fruo.mutation.ClearFight()
	return fruo
}

// Where appends a list predicates to the FighterResultsUpdate builder.
func (fruo *FighterResultsUpdateOne) Where(ps ...predicate.FighterResults) *FighterResultsUpdateOne {
	fruo.mutation.Where(ps...)
	return fruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (fruo *FighterResultsUpdateOne) Select(field string, fields ...string) *FighterResultsUpdateOne {
	fruo.fields = append([]string{field}, fields...)
	return fruo
}

// Save executes the query and returns the updated FighterResults entity.
func (fruo *FighterResultsUpdateOne) Save(ctx context.Context) (*FighterResults, error) {
	return withHooks[*FighterResults, FighterResultsMutation](ctx, fruo.sqlSave, fruo.mutation, fruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (fruo *FighterResultsUpdateOne) SaveX(ctx context.Context) *FighterResults {
	node, err := fruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (fruo *FighterResultsUpdateOne) Exec(ctx context.Context) error {
	_, err := fruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fruo *FighterResultsUpdateOne) ExecX(ctx context.Context) {
	if err := fruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fruo *FighterResultsUpdateOne) check() error {
	if _, ok := fruo.mutation.FighterID(); fruo.mutation.FighterCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FighterResults.fighter"`)
	}
	if _, ok := fruo.mutation.FightID(); fruo.mutation.FightCleared() && !ok {
		return errors.New(`ent: clearing a required unique edge "FighterResults.fight"`)
	}
	return nil
}

func (fruo *FighterResultsUpdateOne) sqlSave(ctx context.Context) (_node *FighterResults, err error) {
	if err := fruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(fighterresults.Table, fighterresults.Columns, sqlgraph.NewFieldSpec(fighterresults.FieldID, field.TypeInt))
	id, ok := fruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "FighterResults.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := fruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, fighterresults.FieldID)
		for _, f := range fields {
			if !fighterresults.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != fighterresults.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := fruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := fruo.mutation.SignificantStrikesLanded(); ok {
		_spec.SetField(fighterresults.FieldSignificantStrikesLanded, field.TypeInt, value)
	}
	if value, ok := fruo.mutation.AddedSignificantStrikesLanded(); ok {
		_spec.AddField(fighterresults.FieldSignificantStrikesLanded, field.TypeInt, value)
	}
	if value, ok := fruo.mutation.Takedowns(); ok {
		_spec.SetField(fighterresults.FieldTakedowns, field.TypeInt, value)
	}
	if value, ok := fruo.mutation.AddedTakedowns(); ok {
		_spec.AddField(fighterresults.FieldTakedowns, field.TypeInt, value)
	}
	if value, ok := fruo.mutation.Knockdowns(); ok {
		_spec.SetField(fighterresults.FieldKnockdowns, field.TypeInt, value)
	}
	if value, ok := fruo.mutation.AddedKnockdowns(); ok {
		_spec.AddField(fighterresults.FieldKnockdowns, field.TypeInt, value)
	}
	if value, ok := fruo.mutation.ControlTimeSeconds(); ok {
		_spec.SetField(fighterresults.FieldControlTimeSeconds, field.TypeInt, value)
	}
	if value, ok := fruo.mutation.AddedControlTimeSeconds(); ok {
		_spec.AddField(fighterresults.FieldControlTimeSeconds, field.TypeInt, value)
	}
	if value, ok := fruo.mutation.WinByStoppage(); ok {
		_spec.SetField(fighterresults.FieldWinByStoppage, field.TypeBool, value)
	}
	if value, ok := fruo.mutation.LossByStoppage(); ok {
		_spec.SetField(fighterresults.FieldLossByStoppage, field.TypeBool, value)
	}
	if value, ok := fruo.mutation.MissedWeight(); ok {
		_spec.SetField(fighterresults.FieldMissedWeight, field.TypeBool, value)
	}
	if fruo.mutation.MissedWeightCleared() {
		_spec.ClearField(fighterresults.FieldMissedWeight, field.TypeBool)
	}
	if fruo.mutation.FighterCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fighterresults.FighterTable,
			Columns: []string{fighterresults.FighterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fighter.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fruo.mutation.FighterIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fighterresults.FighterTable,
			Columns: []string{fighterresults.FighterColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fighter.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if fruo.mutation.FightCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fighterresults.FightTable,
			Columns: []string{fighterresults.FightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fight.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := fruo.mutation.FightIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: false,
			Table:   fighterresults.FightTable,
			Columns: []string{fighterresults.FightColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(fight.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &FighterResults{config: fruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, fruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{fighterresults.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	fruo.mutation.done = true
	return _node, nil
}
