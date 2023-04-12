// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/upcomingevent"
)

// UpcomingEventCreate is the builder for creating a UpcomingEvent entity.
type UpcomingEventCreate struct {
	config
	mutation *UpcomingEventMutation
	hooks    []Hook
	conflict []sql.ConflictOption
}

// Mutation returns the UpcomingEventMutation object of the builder.
func (uec *UpcomingEventCreate) Mutation() *UpcomingEventMutation {
	return uec.mutation
}

// Save creates the UpcomingEvent in the database.
func (uec *UpcomingEventCreate) Save(ctx context.Context) (*UpcomingEvent, error) {
	return withHooks[*UpcomingEvent, UpcomingEventMutation](ctx, uec.sqlSave, uec.mutation, uec.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (uec *UpcomingEventCreate) SaveX(ctx context.Context) *UpcomingEvent {
	v, err := uec.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uec *UpcomingEventCreate) Exec(ctx context.Context) error {
	_, err := uec.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uec *UpcomingEventCreate) ExecX(ctx context.Context) {
	if err := uec.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (uec *UpcomingEventCreate) check() error {
	return nil
}

func (uec *UpcomingEventCreate) sqlSave(ctx context.Context) (*UpcomingEvent, error) {
	if err := uec.check(); err != nil {
		return nil, err
	}
	_node, _spec := uec.createSpec()
	if err := sqlgraph.CreateNode(ctx, uec.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	uec.mutation.id = &_node.ID
	uec.mutation.done = true
	return _node, nil
}

func (uec *UpcomingEventCreate) createSpec() (*UpcomingEvent, *sqlgraph.CreateSpec) {
	var (
		_node = &UpcomingEvent{config: uec.config}
		_spec = sqlgraph.NewCreateSpec(upcomingevent.Table, sqlgraph.NewFieldSpec(upcomingevent.FieldID, field.TypeInt))
	)
	_spec.OnConflict = uec.conflict
	return _node, _spec
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UpcomingEvent.Create().
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (uec *UpcomingEventCreate) OnConflict(opts ...sql.ConflictOption) *UpcomingEventUpsertOne {
	uec.conflict = opts
	return &UpcomingEventUpsertOne{
		create: uec,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UpcomingEvent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uec *UpcomingEventCreate) OnConflictColumns(columns ...string) *UpcomingEventUpsertOne {
	uec.conflict = append(uec.conflict, sql.ConflictColumns(columns...))
	return &UpcomingEventUpsertOne{
		create: uec,
	}
}

type (
	// UpcomingEventUpsertOne is the builder for "upsert"-ing
	//  one UpcomingEvent node.
	UpcomingEventUpsertOne struct {
		create *UpcomingEventCreate
	}

	// UpcomingEventUpsert is the "OnConflict" setter.
	UpcomingEventUpsert struct {
		*sql.UpdateSet
	}
)

// UpdateNewValues updates the mutable fields using the new values that were set on create.
// Using this option is equivalent to using:
//
//	client.UpcomingEvent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UpcomingEventUpsertOne) UpdateNewValues() *UpcomingEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UpcomingEvent.Create().
//	    OnConflict(sql.ResolveWithIgnore()).
//	    Exec(ctx)
func (u *UpcomingEventUpsertOne) Ignore() *UpcomingEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UpcomingEventUpsertOne) DoNothing() *UpcomingEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UpcomingEventCreate.OnConflict
// documentation for more info.
func (u *UpcomingEventUpsertOne) Update(set func(*UpcomingEventUpsert)) *UpcomingEventUpsertOne {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UpcomingEventUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *UpcomingEventUpsertOne) Exec(ctx context.Context) error {
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UpcomingEventCreate.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UpcomingEventUpsertOne) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}

// Exec executes the UPSERT query and returns the inserted/updated ID.
func (u *UpcomingEventUpsertOne) ID(ctx context.Context) (id int, err error) {
	node, err := u.create.Save(ctx)
	if err != nil {
		return id, err
	}
	return node.ID, nil
}

// IDX is like ID, but panics if an error occurs.
func (u *UpcomingEventUpsertOne) IDX(ctx context.Context) int {
	id, err := u.ID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// UpcomingEventCreateBulk is the builder for creating many UpcomingEvent entities in bulk.
type UpcomingEventCreateBulk struct {
	config
	builders []*UpcomingEventCreate
	conflict []sql.ConflictOption
}

// Save creates the UpcomingEvent entities in the database.
func (uecb *UpcomingEventCreateBulk) Save(ctx context.Context) ([]*UpcomingEvent, error) {
	specs := make([]*sqlgraph.CreateSpec, len(uecb.builders))
	nodes := make([]*UpcomingEvent, len(uecb.builders))
	mutators := make([]Mutator, len(uecb.builders))
	for i := range uecb.builders {
		func(i int, root context.Context) {
			builder := uecb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UpcomingEventMutation)
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
					_, err = mutators[i+1].Mutate(root, uecb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					spec.OnConflict = uecb.conflict
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, uecb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, uecb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (uecb *UpcomingEventCreateBulk) SaveX(ctx context.Context) []*UpcomingEvent {
	v, err := uecb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (uecb *UpcomingEventCreateBulk) Exec(ctx context.Context) error {
	_, err := uecb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (uecb *UpcomingEventCreateBulk) ExecX(ctx context.Context) {
	if err := uecb.Exec(ctx); err != nil {
		panic(err)
	}
}

// OnConflict allows configuring the `ON CONFLICT` / `ON DUPLICATE KEY` clause
// of the `INSERT` statement. For example:
//
//	client.UpcomingEvent.CreateBulk(builders...).
//		OnConflict(
//			// Update the row with the new values
//			// the was proposed for insertion.
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (uecb *UpcomingEventCreateBulk) OnConflict(opts ...sql.ConflictOption) *UpcomingEventUpsertBulk {
	uecb.conflict = opts
	return &UpcomingEventUpsertBulk{
		create: uecb,
	}
}

// OnConflictColumns calls `OnConflict` and configures the columns
// as conflict target. Using this option is equivalent to using:
//
//	client.UpcomingEvent.Create().
//		OnConflict(sql.ConflictColumns(columns...)).
//		Exec(ctx)
func (uecb *UpcomingEventCreateBulk) OnConflictColumns(columns ...string) *UpcomingEventUpsertBulk {
	uecb.conflict = append(uecb.conflict, sql.ConflictColumns(columns...))
	return &UpcomingEventUpsertBulk{
		create: uecb,
	}
}

// UpcomingEventUpsertBulk is the builder for "upsert"-ing
// a bulk of UpcomingEvent nodes.
type UpcomingEventUpsertBulk struct {
	create *UpcomingEventCreateBulk
}

// UpdateNewValues updates the mutable fields using the new values that
// were set on create. Using this option is equivalent to using:
//
//	client.UpcomingEvent.Create().
//		OnConflict(
//			sql.ResolveWithNewValues(),
//		).
//		Exec(ctx)
func (u *UpcomingEventUpsertBulk) UpdateNewValues() *UpcomingEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithNewValues())
	return u
}

// Ignore sets each column to itself in case of conflict.
// Using this option is equivalent to using:
//
//	client.UpcomingEvent.Create().
//		OnConflict(sql.ResolveWithIgnore()).
//		Exec(ctx)
func (u *UpcomingEventUpsertBulk) Ignore() *UpcomingEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWithIgnore())
	return u
}

// DoNothing configures the conflict_action to `DO NOTHING`.
// Supported only by SQLite and PostgreSQL.
func (u *UpcomingEventUpsertBulk) DoNothing() *UpcomingEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.DoNothing())
	return u
}

// Update allows overriding fields `UPDATE` values. See the UpcomingEventCreateBulk.OnConflict
// documentation for more info.
func (u *UpcomingEventUpsertBulk) Update(set func(*UpcomingEventUpsert)) *UpcomingEventUpsertBulk {
	u.create.conflict = append(u.create.conflict, sql.ResolveWith(func(update *sql.UpdateSet) {
		set(&UpcomingEventUpsert{UpdateSet: update})
	}))
	return u
}

// Exec executes the query.
func (u *UpcomingEventUpsertBulk) Exec(ctx context.Context) error {
	for i, b := range u.create.builders {
		if len(b.conflict) != 0 {
			return fmt.Errorf("ent: OnConflict was set for builder %d. Set it on the UpcomingEventCreateBulk instead", i)
		}
	}
	if len(u.create.conflict) == 0 {
		return errors.New("ent: missing options for UpcomingEventCreateBulk.OnConflict")
	}
	return u.create.Exec(ctx)
}

// ExecX is like Exec, but panics if an error occurs.
func (u *UpcomingEventUpsertBulk) ExecX(ctx context.Context) {
	if err := u.create.Exec(ctx); err != nil {
		panic(err)
	}
}
