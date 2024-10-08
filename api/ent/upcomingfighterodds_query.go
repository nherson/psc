// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/nherson/psc/api/ent/fighter"
	"github.com/nherson/psc/api/ent/predicate"
	"github.com/nherson/psc/api/ent/upcomingfight"
	"github.com/nherson/psc/api/ent/upcomingfighterodds"
)

// UpcomingFighterOddsQuery is the builder for querying UpcomingFighterOdds entities.
type UpcomingFighterOddsQuery struct {
	config
	ctx               *QueryContext
	order             []upcomingfighterodds.Order
	inters            []Interceptor
	predicates        []predicate.UpcomingFighterOdds
	withFighter       *FighterQuery
	withUpcomingFight *UpcomingFightQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UpcomingFighterOddsQuery builder.
func (ufoq *UpcomingFighterOddsQuery) Where(ps ...predicate.UpcomingFighterOdds) *UpcomingFighterOddsQuery {
	ufoq.predicates = append(ufoq.predicates, ps...)
	return ufoq
}

// Limit the number of records to be returned by this query.
func (ufoq *UpcomingFighterOddsQuery) Limit(limit int) *UpcomingFighterOddsQuery {
	ufoq.ctx.Limit = &limit
	return ufoq
}

// Offset to start from.
func (ufoq *UpcomingFighterOddsQuery) Offset(offset int) *UpcomingFighterOddsQuery {
	ufoq.ctx.Offset = &offset
	return ufoq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (ufoq *UpcomingFighterOddsQuery) Unique(unique bool) *UpcomingFighterOddsQuery {
	ufoq.ctx.Unique = &unique
	return ufoq
}

// Order specifies how the records should be ordered.
func (ufoq *UpcomingFighterOddsQuery) Order(o ...upcomingfighterodds.Order) *UpcomingFighterOddsQuery {
	ufoq.order = append(ufoq.order, o...)
	return ufoq
}

// QueryFighter chains the current query on the "fighter" edge.
func (ufoq *UpcomingFighterOddsQuery) QueryFighter() *FighterQuery {
	query := (&FighterClient{config: ufoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upcomingfighterodds.Table, upcomingfighterodds.FieldID, selector),
			sqlgraph.To(fighter.Table, fighter.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, upcomingfighterodds.FighterTable, upcomingfighterodds.FighterColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// QueryUpcomingFight chains the current query on the "upcoming_fight" edge.
func (ufoq *UpcomingFighterOddsQuery) QueryUpcomingFight() *UpcomingFightQuery {
	query := (&UpcomingFightClient{config: ufoq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := ufoq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := ufoq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(upcomingfighterodds.Table, upcomingfighterodds.FieldID, selector),
			sqlgraph.To(upcomingfight.Table, upcomingfight.FieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, upcomingfighterodds.UpcomingFightTable, upcomingfighterodds.UpcomingFightColumn),
		)
		fromU = sqlgraph.SetNeighbors(ufoq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first UpcomingFighterOdds entity from the query.
// Returns a *NotFoundError when no UpcomingFighterOdds was found.
func (ufoq *UpcomingFighterOddsQuery) First(ctx context.Context) (*UpcomingFighterOdds, error) {
	nodes, err := ufoq.Limit(1).All(setContextOp(ctx, ufoq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{upcomingfighterodds.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (ufoq *UpcomingFighterOddsQuery) FirstX(ctx context.Context) *UpcomingFighterOdds {
	node, err := ufoq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UpcomingFighterOdds ID from the query.
// Returns a *NotFoundError when no UpcomingFighterOdds ID was found.
func (ufoq *UpcomingFighterOddsQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ufoq.Limit(1).IDs(setContextOp(ctx, ufoq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{upcomingfighterodds.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (ufoq *UpcomingFighterOddsQuery) FirstIDX(ctx context.Context) int {
	id, err := ufoq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UpcomingFighterOdds entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UpcomingFighterOdds entity is found.
// Returns a *NotFoundError when no UpcomingFighterOdds entities are found.
func (ufoq *UpcomingFighterOddsQuery) Only(ctx context.Context) (*UpcomingFighterOdds, error) {
	nodes, err := ufoq.Limit(2).All(setContextOp(ctx, ufoq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{upcomingfighterodds.Label}
	default:
		return nil, &NotSingularError{upcomingfighterodds.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (ufoq *UpcomingFighterOddsQuery) OnlyX(ctx context.Context) *UpcomingFighterOdds {
	node, err := ufoq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UpcomingFighterOdds ID in the query.
// Returns a *NotSingularError when more than one UpcomingFighterOdds ID is found.
// Returns a *NotFoundError when no entities are found.
func (ufoq *UpcomingFighterOddsQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = ufoq.Limit(2).IDs(setContextOp(ctx, ufoq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{upcomingfighterodds.Label}
	default:
		err = &NotSingularError{upcomingfighterodds.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (ufoq *UpcomingFighterOddsQuery) OnlyIDX(ctx context.Context) int {
	id, err := ufoq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UpcomingFighterOddsSlice.
func (ufoq *UpcomingFighterOddsQuery) All(ctx context.Context) ([]*UpcomingFighterOdds, error) {
	ctx = setContextOp(ctx, ufoq.ctx, "All")
	if err := ufoq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*UpcomingFighterOdds, *UpcomingFighterOddsQuery]()
	return withInterceptors[[]*UpcomingFighterOdds](ctx, ufoq, qr, ufoq.inters)
}

// AllX is like All, but panics if an error occurs.
func (ufoq *UpcomingFighterOddsQuery) AllX(ctx context.Context) []*UpcomingFighterOdds {
	nodes, err := ufoq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UpcomingFighterOdds IDs.
func (ufoq *UpcomingFighterOddsQuery) IDs(ctx context.Context) (ids []int, err error) {
	if ufoq.ctx.Unique == nil && ufoq.path != nil {
		ufoq.Unique(true)
	}
	ctx = setContextOp(ctx, ufoq.ctx, "IDs")
	if err = ufoq.Select(upcomingfighterodds.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (ufoq *UpcomingFighterOddsQuery) IDsX(ctx context.Context) []int {
	ids, err := ufoq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (ufoq *UpcomingFighterOddsQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, ufoq.ctx, "Count")
	if err := ufoq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, ufoq, querierCount[*UpcomingFighterOddsQuery](), ufoq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (ufoq *UpcomingFighterOddsQuery) CountX(ctx context.Context) int {
	count, err := ufoq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (ufoq *UpcomingFighterOddsQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, ufoq.ctx, "Exist")
	switch _, err := ufoq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (ufoq *UpcomingFighterOddsQuery) ExistX(ctx context.Context) bool {
	exist, err := ufoq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UpcomingFighterOddsQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (ufoq *UpcomingFighterOddsQuery) Clone() *UpcomingFighterOddsQuery {
	if ufoq == nil {
		return nil
	}
	return &UpcomingFighterOddsQuery{
		config:            ufoq.config,
		ctx:               ufoq.ctx.Clone(),
		order:             append([]upcomingfighterodds.Order{}, ufoq.order...),
		inters:            append([]Interceptor{}, ufoq.inters...),
		predicates:        append([]predicate.UpcomingFighterOdds{}, ufoq.predicates...),
		withFighter:       ufoq.withFighter.Clone(),
		withUpcomingFight: ufoq.withUpcomingFight.Clone(),
		// clone intermediate query.
		sql:  ufoq.sql.Clone(),
		path: ufoq.path,
	}
}

// WithFighter tells the query-builder to eager-load the nodes that are connected to
// the "fighter" edge. The optional arguments are used to configure the query builder of the edge.
func (ufoq *UpcomingFighterOddsQuery) WithFighter(opts ...func(*FighterQuery)) *UpcomingFighterOddsQuery {
	query := (&FighterClient{config: ufoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufoq.withFighter = query
	return ufoq
}

// WithUpcomingFight tells the query-builder to eager-load the nodes that are connected to
// the "upcoming_fight" edge. The optional arguments are used to configure the query builder of the edge.
func (ufoq *UpcomingFighterOddsQuery) WithUpcomingFight(opts ...func(*UpcomingFightQuery)) *UpcomingFighterOddsQuery {
	query := (&UpcomingFightClient{config: ufoq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	ufoq.withUpcomingFight = query
	return ufoq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		FighterID int `json:"fighter_id,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UpcomingFighterOdds.Query().
//		GroupBy(upcomingfighterodds.FieldFighterID).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (ufoq *UpcomingFighterOddsQuery) GroupBy(field string, fields ...string) *UpcomingFighterOddsGroupBy {
	ufoq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &UpcomingFighterOddsGroupBy{build: ufoq}
	grbuild.flds = &ufoq.ctx.Fields
	grbuild.label = upcomingfighterodds.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		FighterID int `json:"fighter_id,omitempty"`
//	}
//
//	client.UpcomingFighterOdds.Query().
//		Select(upcomingfighterodds.FieldFighterID).
//		Scan(ctx, &v)
func (ufoq *UpcomingFighterOddsQuery) Select(fields ...string) *UpcomingFighterOddsSelect {
	ufoq.ctx.Fields = append(ufoq.ctx.Fields, fields...)
	sbuild := &UpcomingFighterOddsSelect{UpcomingFighterOddsQuery: ufoq}
	sbuild.label = upcomingfighterodds.Label
	sbuild.flds, sbuild.scan = &ufoq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a UpcomingFighterOddsSelect configured with the given aggregations.
func (ufoq *UpcomingFighterOddsQuery) Aggregate(fns ...AggregateFunc) *UpcomingFighterOddsSelect {
	return ufoq.Select().Aggregate(fns...)
}

func (ufoq *UpcomingFighterOddsQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range ufoq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, ufoq); err != nil {
				return err
			}
		}
	}
	for _, f := range ufoq.ctx.Fields {
		if !upcomingfighterodds.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if ufoq.path != nil {
		prev, err := ufoq.path(ctx)
		if err != nil {
			return err
		}
		ufoq.sql = prev
	}
	return nil
}

func (ufoq *UpcomingFighterOddsQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*UpcomingFighterOdds, error) {
	var (
		nodes       = []*UpcomingFighterOdds{}
		_spec       = ufoq.querySpec()
		loadedTypes = [2]bool{
			ufoq.withFighter != nil,
			ufoq.withUpcomingFight != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*UpcomingFighterOdds).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &UpcomingFighterOdds{config: ufoq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, ufoq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := ufoq.withFighter; query != nil {
		if err := ufoq.loadFighter(ctx, query, nodes, nil,
			func(n *UpcomingFighterOdds, e *Fighter) { n.Edges.Fighter = e }); err != nil {
			return nil, err
		}
	}
	if query := ufoq.withUpcomingFight; query != nil {
		if err := ufoq.loadUpcomingFight(ctx, query, nodes, nil,
			func(n *UpcomingFighterOdds, e *UpcomingFight) { n.Edges.UpcomingFight = e }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (ufoq *UpcomingFighterOddsQuery) loadFighter(ctx context.Context, query *FighterQuery, nodes []*UpcomingFighterOdds, init func(*UpcomingFighterOdds), assign func(*UpcomingFighterOdds, *Fighter)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UpcomingFighterOdds)
	for i := range nodes {
		fk := nodes[i].FighterID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(fighter.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "fighter_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}
func (ufoq *UpcomingFighterOddsQuery) loadUpcomingFight(ctx context.Context, query *UpcomingFightQuery, nodes []*UpcomingFighterOdds, init func(*UpcomingFighterOdds), assign func(*UpcomingFighterOdds, *UpcomingFight)) error {
	ids := make([]int, 0, len(nodes))
	nodeids := make(map[int][]*UpcomingFighterOdds)
	for i := range nodes {
		fk := nodes[i].UpcomingFightID
		if _, ok := nodeids[fk]; !ok {
			ids = append(ids, fk)
		}
		nodeids[fk] = append(nodeids[fk], nodes[i])
	}
	if len(ids) == 0 {
		return nil
	}
	query.Where(upcomingfight.IDIn(ids...))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		nodes, ok := nodeids[n.ID]
		if !ok {
			return fmt.Errorf(`unexpected foreign-key "upcoming_fight_id" returned %v`, n.ID)
		}
		for i := range nodes {
			assign(nodes[i], n)
		}
	}
	return nil
}

func (ufoq *UpcomingFighterOddsQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := ufoq.querySpec()
	_spec.Node.Columns = ufoq.ctx.Fields
	if len(ufoq.ctx.Fields) > 0 {
		_spec.Unique = ufoq.ctx.Unique != nil && *ufoq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, ufoq.driver, _spec)
}

func (ufoq *UpcomingFighterOddsQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(upcomingfighterodds.Table, upcomingfighterodds.Columns, sqlgraph.NewFieldSpec(upcomingfighterodds.FieldID, field.TypeInt))
	_spec.From = ufoq.sql
	if unique := ufoq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if ufoq.path != nil {
		_spec.Unique = true
	}
	if fields := ufoq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, upcomingfighterodds.FieldID)
		for i := range fields {
			if fields[i] != upcomingfighterodds.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
		if ufoq.withFighter != nil {
			_spec.Node.AddColumnOnce(upcomingfighterodds.FieldFighterID)
		}
		if ufoq.withUpcomingFight != nil {
			_spec.Node.AddColumnOnce(upcomingfighterodds.FieldUpcomingFightID)
		}
	}
	if ps := ufoq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := ufoq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := ufoq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := ufoq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (ufoq *UpcomingFighterOddsQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(ufoq.driver.Dialect())
	t1 := builder.Table(upcomingfighterodds.Table)
	columns := ufoq.ctx.Fields
	if len(columns) == 0 {
		columns = upcomingfighterodds.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if ufoq.sql != nil {
		selector = ufoq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if ufoq.ctx.Unique != nil && *ufoq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range ufoq.predicates {
		p(selector)
	}
	for _, p := range ufoq.order {
		p(selector)
	}
	if offset := ufoq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := ufoq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UpcomingFighterOddsGroupBy is the group-by builder for UpcomingFighterOdds entities.
type UpcomingFighterOddsGroupBy struct {
	selector
	build *UpcomingFighterOddsQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (ufogb *UpcomingFighterOddsGroupBy) Aggregate(fns ...AggregateFunc) *UpcomingFighterOddsGroupBy {
	ufogb.fns = append(ufogb.fns, fns...)
	return ufogb
}

// Scan applies the selector query and scans the result into the given value.
func (ufogb *UpcomingFighterOddsGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufogb.build.ctx, "GroupBy")
	if err := ufogb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpcomingFighterOddsQuery, *UpcomingFighterOddsGroupBy](ctx, ufogb.build, ufogb, ufogb.build.inters, v)
}

func (ufogb *UpcomingFighterOddsGroupBy) sqlScan(ctx context.Context, root *UpcomingFighterOddsQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(ufogb.fns))
	for _, fn := range ufogb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*ufogb.flds)+len(ufogb.fns))
		for _, f := range *ufogb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*ufogb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufogb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// UpcomingFighterOddsSelect is the builder for selecting fields of UpcomingFighterOdds entities.
type UpcomingFighterOddsSelect struct {
	*UpcomingFighterOddsQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ufos *UpcomingFighterOddsSelect) Aggregate(fns ...AggregateFunc) *UpcomingFighterOddsSelect {
	ufos.fns = append(ufos.fns, fns...)
	return ufos
}

// Scan applies the selector query and scans the result into the given value.
func (ufos *UpcomingFighterOddsSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ufos.ctx, "Select")
	if err := ufos.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*UpcomingFighterOddsQuery, *UpcomingFighterOddsSelect](ctx, ufos.UpcomingFighterOddsQuery, ufos, ufos.inters, v)
}

func (ufos *UpcomingFighterOddsSelect) sqlScan(ctx context.Context, root *UpcomingFighterOddsQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ufos.fns))
	for _, fn := range ufos.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ufos.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ufos.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
