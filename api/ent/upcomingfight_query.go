// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/ent/predicate"
	"github.com/nherson/psc/api/ent/upcomingevent"
	"github.com/nherson/psc/api/ent/upcomingfight"
	"github.com/nherson/psc/api/ent/upcomingfighterodds"
)

// UpcomingFightQuery is the builder for querying UpcomingFight entities.
type UpcomingFightQuery struct {
	config
	ctx                     *QueryContext
	order                   []upcomingfight.Order
	inters                  []Interceptor
	predicates              []predicate.UpcomingFight
	withUpcomingEvent       *UpcomingEventQuery
	withFighters            *FighterQuery
	withUpcomingFighterOdds *UpcomingFighterOddsQuery
	withFKs                 bool
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UpcomingFightQuery builder.
func (ufq *UpcomingFightQuery) Where(ps ...predicate.UpcomingFight) *UpcomingFightQuery {
	ufq.predicates = append(ufq.predicates, ps...)
	return ufq
}

// Limit the number of records to be returned by this query.
func (ufq *UpcomingFightQuery) Limit(limit int) *UpcomingFightQuery {
	ufq.ctx.Limit = &limit
	return ufq
}

// Offset to start from.
func (ufq *UpcomingFightQuery) Offset(offset int) *UpcomingFightQuery {
	ufq.ctx.Offset = &offset
	return ufq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ufq *UpcomingFightQuery) Unique(unique bool) *UpcomingFightQuery {
	ufq.ctx.Unique = &unique
	return ufq
}

// Order specifies how the records should be ordered.
func (ufq *UpcomingFightQuery) Order(o ...upcomingfight.Order) *UpcomingFightQuery {
	ufq.order = append(ufq.order, o...)
	return ufq
}

// QueryUpcomingEvent chains the current query on the "upcoming_event" edge.
func (ufq *UpcomingFightQuery) QueryUpcomingEvent() *UpcomingEventQuery {
	query := (&UpcomingEventClient{config: ufq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upcomingfight.Table, upcomingfight.FieldID, selector),
			sqlgraph.To(upcomingevent.Table, upcomingevent.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, upcomingfight.UpcomingEventTable, upcomingfight.UpcomingEventColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryFighters chains the current query on the "fighters" edge.
func (ufq *UpcomingFightQuery) QueryFighters() *FighterQuery {
	query := (&FighterClient{config: ufq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upcomingfight.Table, upcomingfight.FieldID, selector),
			sqlgraph.To(fighter.Table, fighter.FieldID),
			sqlgraph.Edge(sqlgraph.M2M, false, upcomingfight.FightersTable, upcomingfight.FightersPrimaryKey...),
		)
		fromU = sqlgraph.SetNeighbors(ufq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUpcomingFighterOdds chains the current query on the "upcoming_fighter_odds" edge.
func (ufq *UpcomingFightQuery) QueryUpcomingFighterOdds() *UpcomingFighterOddsQuery {
	query := (&UpcomingFighterOddsClient{config: ufq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upcomingfight.Table, upcomingfight.FieldID, selector),
			sqlgraph.To(upcomingfighterodds.Table, upcomingfighterodds.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, true, upcomingfight.UpcomingFighterOddsTable, upcomingfight.UpcomingFighterOddsColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UpcomingFight entity from the query.
// Returns a *NotFoundError when no UpcomingFight was found.
func (ufq *UpcomingFightQuery) First(ctx context.Context) (*UpcomingFight, error) {
	nodes, err := ufq.Limit(1).All(setContextOp(ctx, ufq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{upcomingfight.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ufq *UpcomingFightQuery) FirstX(ctx context.Context) *UpcomingFight {
	node, err := ufq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UpcomingFight ID from the query.
// Returns a *NotFoundError when no UpcomingFight ID was found.
func (ufq *UpcomingFightQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ufq.Limit(1).IDs(setContextOp(ctx, ufq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{upcomingfight.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ufq *UpcomingFightQuery) FirstIDX(ctx context.Context) int {
	id, err := ufq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UpcomingFight entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UpcomingFight entity is found.
// Returns a *NotFoundError when no UpcomingFight entities are found.
func (ufq *UpcomingFightQuery) Only(ctx context.Context) (*UpcomingFight, error) {
	nodes, err := ufq.Limit(2).All(setContextOp(ctx, ufq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{upcomingfight.Label}
	default:
		return nil, &NotSingularError{upcomingfight.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ufq *UpcomingFightQuery) OnlyX(ctx context.Context) *UpcomingFight {
	node, err := ufq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UpcomingFight ID in the query.
// Returns a *NotSingularError when more than one UpcomingFight ID is found.
// Returns a *NotFoundError when no entities are found.
func (ufq *UpcomingFightQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ufq.Limit(2).IDs(setContextOp(ctx, ufq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{upcomingfight.Label}
	default:
		err = &NotSingularError{upcomingfight.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ufq *UpcomingFightQuery) OnlyIDX(ctx context.Context) int {
	id, err := ufq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UpcomingFights.
func (ufq *UpcomingFightQuery) All(ctx context.Context) ([]*UpcomingFight, error) {
	ctx = setContextOp(ctx, ufq.ctx, "All")
	if err := ufq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UpcomingFight, *UpcomingFightQuery]()
	return withInterceptors[[]*UpcomingFight](ctx, ufq, qr, ufq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ufq *UpcomingFightQuery) AllX(ctx context.Context) []*UpcomingFight {
	nodes, err := ufq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UpcomingFight IDs.
func (ufq *UpcomingFightQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ufq.ctx.Unique == nil && ufq.path != nil {
		ufq.Unique(true)
	}
	ctx = setContextOp(ctx, ufq.ctx, "IDs")
	if err = ufq.Select(upcomingfight.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ufq *UpcomingFightQuery) IDsX(ctx context.Context) []int {
	ids, err := ufq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ufq *UpcomingFightQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ufq.ctx, "Count")
	if err := ufq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ufq, querierCount[*UpcomingFightQuery](), ufq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ufq *UpcomingFightQuery) CountX(ctx context.Context) int {
	count, err := ufq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ufq *UpcomingFightQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ufq.ctx, "Exist")
	switch _, err := ufq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ufq *UpcomingFightQuery) ExistX(ctx context.Context) bool {
	exist, err := ufq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UpcomingFightQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ufq *UpcomingFightQuery) Clone() *UpcomingFightQuery {
	if ufq == nil {
		return nil
	}
	return &UpcomingFightQuery{
		config:                  ufq.config,
		ctx:                     ufq.ctx.Clone(),
		order:                   append([]upcomingfight.Order{}, ufq.order...),
		inters:                  append([]Interceptor{}, ufq.inters...),
		predicates:              append([]predicate.UpcomingFight{}, ufq.predicates...),
		withUpcomingEvent:       ufq.withUpcomingEvent.Clone(),
		withFighters:            ufq.withFighters.Clone(),
		withUpcomingFighterOdds: ufq.withUpcomingFighterOdds.Clone(),
		// clone intermediate query.
		sql:  ufq.sql.Clone(),
		path: ufq.path,
	}
}

// WithUpcomingEvent tells the query-builder to eager-load the nodes that are connected to
// the "upcoming_event" edge. The optional arguments are used to configure the query builder of the edge.
func (ufq *UpcomingFightQuery) WithUpcomingEvent(opts ...func(*UpcomingEventQuery)) *UpcomingFightQuery {
	query := (&UpcomingEventClient{config: ufq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufq.withUpcomingEvent = query
	return ufq
}

// WithFighters tells the query-builder to eager-load the nodes that are connected to
// the "fighters" edge. The optional arguments are used to configure the query builder of the edge.
func (ufq *UpcomingFightQuery) WithFighters(opts ...func(*FighterQuery)) *UpcomingFightQuery {
	query := (&FighterClient{config: ufq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufq.withFighters = query
	return ufq
}

// WithUpcomingFighterOdds tells the query-builder to eager-load the nodes that are connected to
// the "upcoming_fighter_odds" edge. The optional arguments are used to configure the query builder of the edge.
func (ufq *UpcomingFightQuery) WithUpcomingFighterOdds(opts ...func(*UpcomingFighterOddsQuery)) *UpcomingFightQuery {
	query := (&UpcomingFighterOddsClient{config: ufq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufq.withUpcomingFighterOdds = query
	return ufq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UpcomingFight.Query().
//		GroupBy(upcomingfight.FieldCreatedAt).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ufq *UpcomingFightQuery) GroupBy(field string, fields ...string) *UpcomingFightGroupBy {
	ufq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UpcomingFightGroupBy{build: ufq}
	grbuild.flds = &ufq.ctx.Fields
	grbuild.label = upcomingfight.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreatedAt time.Time `json:"created_at,omitempty"`
//	}
//
//	client.UpcomingFight.Query().
//		Select(upcomingfight.FieldCreatedAt).
//		Scan(ctx, &v)
func (ufq *UpcomingFightQuery) Select(fields ...string) *UpcomingFightSelect {
	ufq.ctx.Fields = append(ufq.ctx.Fields, fields...)
	sbuild := &UpcomingFightSelect{UpcomingFightQuery: ufq}
	sbuild.label = upcomingfight.Label
	sbuild.flds, sbuild.scan = &ufq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UpcomingFightSelect configured with the given aggregations.
func (ufq *UpcomingFightQuery) Aggregate(fns ...AggregateFunc) *UpcomingFightSelect {
	return ufq.Select().Aggregate(fns...)
}

func (ufq *UpcomingFightQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ufq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ufq); err != nil {
				return err
			}
		}
	}
	for _, f := range ufq.ctx.Fields {
		if !upcomingfight.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ufq.path != nil {
		prev, err := ufq.path(ctx)
		if err != nil {
			return err
		}
		ufq.sql = prev
	}
	return nil
}

func (ufq *UpcomingFightQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UpcomingFight, error) {
	var (
		nodes       = []*UpcomingFight{}
		withFKs     = ufq.withFKs
		_spec       = ufq.querySpec()
		loadedTypes = [3]bool{
			ufq.withUpcomingEvent != nil,
			ufq.withFighters != nil,
			ufq.withUpcomingFighterOdds != nil,
		}
	)
	if ufq.withUpcomingEvent != nil {
		withFKs = true
	}
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, upcomingfight.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UpcomingFight).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UpcomingFight{config: ufq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ufq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ufq.withUpcomingEvent; query != nil {
		if err := ufq.loadUpcomingEvent(ctx, query, nodes, nil,
			func(n *UpcomingFight, e *UpcomingEvent) { n.Edges.UpcomingEvent = e }); err != nil {
			return nil, err
		}
	}
	if query := ufq.withFighters; query != nil {
		if err := ufq.loadFighters(ctx, query, nodes,
			func(n *UpcomingFight) { n.Edges.Fighters = []*Fighter{} },
			func(n *UpcomingFight, e *Fighter) { n.Edges.Fighters = append(n.Edges.Fighters, e) }); err != nil {
			return nil, err
		}
	}
	if query := ufq.withUpcomingFighterOdds; query != nil {
		if err := ufq.loadUpcomingFighterOdds(ctx, query, nodes,
			func(n *UpcomingFight) { n.Edges.UpcomingFighterOdds = []*UpcomingFighterOdds{} },
			func(n *UpcomingFight, e *UpcomingFighterOdds) {
				n.Edges.UpcomingFighterOdds = append(n.Edges.UpcomingFighterOdds, e)
			}); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ufq *UpcomingFightQuery) loadUpcomingEvent(ctx context.Context, query *UpcomingEventQuery, nodes []*UpcomingFight, init func(*UpcomingFight), assign func(*UpcomingFight, *UpcomingEvent)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UpcomingFight)
	for i := range nodes {
		if nodes[i].upcoming_event_upcoming_fights == nil {
			continue
		}
		fk := *nodes[i].upcoming_event_upcoming_fights
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(upcomingevent.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "upcoming_event_upcoming_fights" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ufq *UpcomingFightQuery) loadFighters(ctx context.Context, query *FighterQuery, nodes []*UpcomingFight, init func(*UpcomingFight), assign func(*UpcomingFight, *Fighter)) error {
	edgeIDs := make([]driver.Value, len(nodes))
	byID := make(map[int]*UpcomingFight)
	nids := make(map[int]map[*UpcomingFight]struct{})
	for i, node := range nodes {
		edgeIDs[i] = node.ID
		byID[node.ID] = node
		if init != nil {
			init(node)
		}
	}
	query.Where(func(s *sql.Selector) {
		joinT := sql.Table(upcomingfight.FightersTable)
		s.Join(joinT).On(s.C(fighter.FieldID), joinT.C(upcomingfight.FightersPrimaryKey[1]))
		s.Where(sql.InValues(joinT.C(upcomingfight.FightersPrimaryKey[0]), edgeIDs...))
		columns := s.SelectedColumns()
		s.Select(joinT.C(upcomingfight.FightersPrimaryKey[0]))
		s.AppendSelect(columns...)
		s.SetDistinct(false)
	})
	if err := query.prepareQuery(ctx); err != nil {
		return err
	}
	qr := QuerierFunc(func(ctx context.Context, q Query) (Value, error) {
		return query.sqlAll(ctx, func(_ context.Context, spec *sqlgraph.QuerySpec) {
			assign := spec.Assign
			values := spec.ScanValues
			spec.ScanValues = func(columns []string) ([]any, error) {
				values, err := values(columns[1:])
				if err != nil {
					return nil, err
				}
				return append([]any{new(sql.NullInt64)}, values...), nil
			}
			spec.Assign = func(columns []string, values []any) error {
				outValue := int(values[0].(*sql.NullInt64).Int64)
				inValue := int(values[1].(*sql.NullInt64).Int64)
				if nids[inValue] == nil {
					nids[inValue] = map[*UpcomingFight]struct{}{byID[outValue]: {}}
					return assign(columns[1:], values[1:])
				}
				nids[inValue][byID[outValue]] = struct{}{}
				return nil
			}
		})
	})
	neighbors, err := withInterceptors[[]*Fighter](ctx, query, qr, query.inters)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected "fighters" node returned %v`, n.ID)
		}
		for kn := range nodes {
			assign(kn, n)
		}
	}
	return nil
}
func (ufq *UpcomingFightQuery) loadUpcomingFighterOdds(ctx context.Context, query *UpcomingFighterOddsQuery, nodes []*UpcomingFight, init func(*UpcomingFight), assign func(*UpcomingFight, *UpcomingFighterOdds)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*UpcomingFight)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.Where(predicate.UpcomingFighterOdds(func(s *sql.Selector) {
		s.Where(sql.InValues(upcomingfight.UpcomingFighterOddsColumn, fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.UpcomingFightID
		node, ok := nodeids[fk]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "upcoming_fight_id" returned %v for node %v`, fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (ufq *UpcomingFightQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ufq.querySpec()
	_spec.Node.Columns = ufq.ctx.Fields
	if len(ufq.ctx.Fields) > 0 {
		_spec.Unique = ufq.ctx.Unique != nil && *ufq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ufq.driver, _spec)
}

func (ufq *UpcomingFightQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(upcomingfight.Table, upcomingfight.Columns, sqlgraph.NewFieldSpec(upcomingfight.FieldID, field.TypeInt))
	_spec.From = ufq.sql
	if unique := ufq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ufq.path != nil {
		_spec.Unique = true
	}
	if fields := ufq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upcomingfight.FieldID)
		for i := range fields {
			if fields[i] != upcomingfight.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := ufq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ufq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ufq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ufq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ufq *UpcomingFightQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ufq.driver.Dialect())
	t1 := builder.Table(upcomingfight.Table)
	columns := ufq.ctx.Fields
	if len(columns) == 0 {
		columns = upcomingfight.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ufq.sql != nil {
		selector = ufq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ufq.ctx.Unique != nil && *ufq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ufq.predicates {
		p(selector)
	}
	for _, p := range ufq.order {
		p(selector)
	}
	if offset := ufq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ufq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UpcomingFightGroupBy is the group-by builder for UpcomingFight entities.
type UpcomingFightGroupBy struct {
	selector
	build *UpcomingFightQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ufgb *UpcomingFightGroupBy) Aggregate(fns ...AggregateFunc) *UpcomingFightGroupBy {
	ufgb.fns = append(ufgb.fns, fns...)
	return ufgb
}

// Scan applies the selector query and scans the result into the given value.
func (ufgb *UpcomingFightGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufgb.build.ctx, "GroupBy")
	if err := ufgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpcomingFightQuery, *UpcomingFightGroupBy](ctx, ufgb.build, ufgb, ufgb.build.inters, v)
}

func (ufgb *UpcomingFightGroupBy) sqlScan(ctx context.Context, root *UpcomingFightQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ufgb.fns))
	for _, fn := range ufgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ufgb.flds)+len(ufgb.fns))
		for _, f := range *ufgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ufgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UpcomingFightSelect is the builder for selecting fields of UpcomingFight entities.
type UpcomingFightSelect struct {
	*UpcomingFightQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ufs *UpcomingFightSelect) Aggregate(fns ...AggregateFunc) *UpcomingFightSelect {
	ufs.fns = append(ufs.fns, fns...)
	return ufs
}

// Scan applies the selector query and scans the result into the given value.
func (ufs *UpcomingFightSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufs.ctx, "Select")
	if err := ufs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpcomingFightQuery, *UpcomingFightSelect](ctx, ufs.UpcomingFightQuery, ufs, ufs.inters, v)
}

func (ufs *UpcomingFightSelect) sqlScan(ctx context.Context, root *UpcomingFightQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ufs.fns))
	for _, fn := range ufs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ufs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
