// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/user/service/internal/data/model/predicate"
	"mall-go/app/user/service/internal/data/model/userpoint"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserPointQuery is the builder for querying UserPoint entities.
type UserPointQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserPoint
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserPointQuery builder.
func (upq *UserPointQuery) Where(ps ...predicate.UserPoint) *UserPointQuery {
	upq.predicates = append(upq.predicates, ps...)
	return upq
}

// Limit adds a limit step to the query.
func (upq *UserPointQuery) Limit(limit int) *UserPointQuery {
	upq.limit = &limit
	return upq
}

// Offset adds an offset step to the query.
func (upq *UserPointQuery) Offset(offset int) *UserPointQuery {
	upq.offset = &offset
	return upq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (upq *UserPointQuery) Unique(unique bool) *UserPointQuery {
	upq.unique = &unique
	return upq
}

// Order adds an order step to the query.
func (upq *UserPointQuery) Order(o ...OrderFunc) *UserPointQuery {
	upq.order = append(upq.order, o...)
	return upq
}

// First returns the first UserPoint entity from the query.
// Returns a *NotFoundError when no UserPoint was found.
func (upq *UserPointQuery) First(ctx context.Context) (*UserPoint, error) {
	nodes, err := upq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userpoint.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (upq *UserPointQuery) FirstX(ctx context.Context) *UserPoint {
	node, err := upq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserPoint ID from the query.
// Returns a *NotFoundError when no UserPoint ID was found.
func (upq *UserPointQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = upq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userpoint.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (upq *UserPointQuery) FirstIDX(ctx context.Context) int64 {
	id, err := upq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserPoint entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserPoint entity is found.
// Returns a *NotFoundError when no UserPoint entities are found.
func (upq *UserPointQuery) Only(ctx context.Context) (*UserPoint, error) {
	nodes, err := upq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userpoint.Label}
	default:
		return nil, &NotSingularError{userpoint.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (upq *UserPointQuery) OnlyX(ctx context.Context) *UserPoint {
	node, err := upq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserPoint ID in the query.
// Returns a *NotSingularError when more than one UserPoint ID is found.
// Returns a *NotFoundError when no entities are found.
func (upq *UserPointQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = upq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userpoint.Label}
	default:
		err = &NotSingularError{userpoint.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (upq *UserPointQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := upq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserPoints.
func (upq *UserPointQuery) All(ctx context.Context) ([]*UserPoint, error) {
	if err := upq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return upq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (upq *UserPointQuery) AllX(ctx context.Context) []*UserPoint {
	nodes, err := upq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserPoint IDs.
func (upq *UserPointQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := upq.Select(userpoint.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (upq *UserPointQuery) IDsX(ctx context.Context) []int64 {
	ids, err := upq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (upq *UserPointQuery) Count(ctx context.Context) (int, error) {
	if err := upq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return upq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (upq *UserPointQuery) CountX(ctx context.Context) int {
	count, err := upq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (upq *UserPointQuery) Exist(ctx context.Context) (bool, error) {
	if err := upq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return upq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (upq *UserPointQuery) ExistX(ctx context.Context) bool {
	exist, err := upq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserPointQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (upq *UserPointQuery) Clone() *UserPointQuery {
	if upq == nil {
		return nil
	}
	return &UserPointQuery{
		config:     upq.config,
		limit:      upq.limit,
		offset:     upq.offset,
		order:      append([]OrderFunc{}, upq.order...),
		predicates: append([]predicate.UserPoint{}, upq.predicates...),
		// clone intermediate query.
		sql:    upq.sql.Clone(),
		path:   upq.path,
		unique: upq.unique,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.UserPoint.Query().
//		GroupBy(userpoint.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (upq *UserPointQuery) GroupBy(field string, fields ...string) *UserPointGroupBy {
	group := &UserPointGroupBy{config: upq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := upq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return upq.sqlQuery(ctx), nil
	}
	return group
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		CreateTime time.Time `json:"create_time,omitempty"`
//	}
//
//	client.UserPoint.Query().
//		Select(userpoint.FieldCreateTime).
//		Scan(ctx, &v)
func (upq *UserPointQuery) Select(fields ...string) *UserPointSelect {
	upq.fields = append(upq.fields, fields...)
	return &UserPointSelect{UserPointQuery: upq}
}

func (upq *UserPointQuery) prepareQuery(ctx context.Context) error {
	for _, f := range upq.fields {
		if !userpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if upq.path != nil {
		prev, err := upq.path(ctx)
		if err != nil {
			return err
		}
		upq.sql = prev
	}
	return nil
}

func (upq *UserPointQuery) sqlAll(ctx context.Context) ([]*UserPoint, error) {
	var (
		nodes = []*UserPoint{}
		_spec = upq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserPoint{config: upq.config}
		nodes = append(nodes, node)
		return node.scanValues(columns)
	}
	_spec.Assign = func(columns []string, values []interface{}) error {
		if len(nodes) == 0 {
			return fmt.Errorf("model: Assign called without calling ScanValues")
		}
		node := nodes[len(nodes)-1]
		return node.assignValues(columns, values)
	}
	if err := sqlgraph.QueryNodes(ctx, upq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (upq *UserPointQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := upq.querySpec()
	_spec.Node.Columns = upq.fields
	if len(upq.fields) > 0 {
		_spec.Unique = upq.unique != nil && *upq.unique
	}
	return sqlgraph.CountNodes(ctx, upq.driver, _spec)
}

func (upq *UserPointQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := upq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (upq *UserPointQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userpoint.Table,
			Columns: userpoint.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userpoint.FieldID,
			},
		},
		From:   upq.sql,
		Unique: true,
	}
	if unique := upq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := upq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userpoint.FieldID)
		for i := range fields {
			if fields[i] != userpoint.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := upq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := upq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := upq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := upq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (upq *UserPointQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(upq.driver.Dialect())
	t1 := builder.Table(userpoint.Table)
	columns := upq.fields
	if len(columns) == 0 {
		columns = userpoint.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if upq.sql != nil {
		selector = upq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if upq.unique != nil && *upq.unique {
		selector.Distinct()
	}
	for _, p := range upq.predicates {
		p(selector)
	}
	for _, p := range upq.order {
		p(selector)
	}
	if offset := upq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := upq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserPointGroupBy is the group-by builder for UserPoint entities.
type UserPointGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (upgb *UserPointGroupBy) Aggregate(fns ...AggregateFunc) *UserPointGroupBy {
	upgb.fns = append(upgb.fns, fns...)
	return upgb
}

// Scan applies the group-by query and scans the result into the given value.
func (upgb *UserPointGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := upgb.path(ctx)
	if err != nil {
		return err
	}
	upgb.sql = query
	return upgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (upgb *UserPointGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := upgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (upgb *UserPointGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(upgb.fields) > 1 {
		return nil, errors.New("model: UserPointGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := upgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (upgb *UserPointGroupBy) StringsX(ctx context.Context) []string {
	v, err := upgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upgb *UserPointGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = upgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpoint.Label}
	default:
		err = fmt.Errorf("model: UserPointGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (upgb *UserPointGroupBy) StringX(ctx context.Context) string {
	v, err := upgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (upgb *UserPointGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(upgb.fields) > 1 {
		return nil, errors.New("model: UserPointGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := upgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (upgb *UserPointGroupBy) IntsX(ctx context.Context) []int {
	v, err := upgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upgb *UserPointGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = upgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpoint.Label}
	default:
		err = fmt.Errorf("model: UserPointGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (upgb *UserPointGroupBy) IntX(ctx context.Context) int {
	v, err := upgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (upgb *UserPointGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(upgb.fields) > 1 {
		return nil, errors.New("model: UserPointGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := upgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (upgb *UserPointGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := upgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upgb *UserPointGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = upgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpoint.Label}
	default:
		err = fmt.Errorf("model: UserPointGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (upgb *UserPointGroupBy) Float64X(ctx context.Context) float64 {
	v, err := upgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (upgb *UserPointGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(upgb.fields) > 1 {
		return nil, errors.New("model: UserPointGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := upgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (upgb *UserPointGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := upgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (upgb *UserPointGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = upgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpoint.Label}
	default:
		err = fmt.Errorf("model: UserPointGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (upgb *UserPointGroupBy) BoolX(ctx context.Context) bool {
	v, err := upgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (upgb *UserPointGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range upgb.fields {
		if !userpoint.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := upgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := upgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (upgb *UserPointGroupBy) sqlQuery() *sql.Selector {
	selector := upgb.sql.Select()
	aggregation := make([]string, 0, len(upgb.fns))
	for _, fn := range upgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(upgb.fields)+len(upgb.fns))
		for _, f := range upgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(upgb.fields...)...)
}

// UserPointSelect is the builder for selecting fields of UserPoint entities.
type UserPointSelect struct {
	*UserPointQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (ups *UserPointSelect) Scan(ctx context.Context, v interface{}) error {
	if err := ups.prepareQuery(ctx); err != nil {
		return err
	}
	ups.sql = ups.UserPointQuery.sqlQuery(ctx)
	return ups.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (ups *UserPointSelect) ScanX(ctx context.Context, v interface{}) {
	if err := ups.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (ups *UserPointSelect) Strings(ctx context.Context) ([]string, error) {
	if len(ups.fields) > 1 {
		return nil, errors.New("model: UserPointSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := ups.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (ups *UserPointSelect) StringsX(ctx context.Context) []string {
	v, err := ups.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (ups *UserPointSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = ups.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpoint.Label}
	default:
		err = fmt.Errorf("model: UserPointSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (ups *UserPointSelect) StringX(ctx context.Context) string {
	v, err := ups.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (ups *UserPointSelect) Ints(ctx context.Context) ([]int, error) {
	if len(ups.fields) > 1 {
		return nil, errors.New("model: UserPointSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := ups.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (ups *UserPointSelect) IntsX(ctx context.Context) []int {
	v, err := ups.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (ups *UserPointSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = ups.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpoint.Label}
	default:
		err = fmt.Errorf("model: UserPointSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (ups *UserPointSelect) IntX(ctx context.Context) int {
	v, err := ups.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (ups *UserPointSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(ups.fields) > 1 {
		return nil, errors.New("model: UserPointSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := ups.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (ups *UserPointSelect) Float64sX(ctx context.Context) []float64 {
	v, err := ups.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (ups *UserPointSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = ups.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpoint.Label}
	default:
		err = fmt.Errorf("model: UserPointSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (ups *UserPointSelect) Float64X(ctx context.Context) float64 {
	v, err := ups.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (ups *UserPointSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(ups.fields) > 1 {
		return nil, errors.New("model: UserPointSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := ups.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (ups *UserPointSelect) BoolsX(ctx context.Context) []bool {
	v, err := ups.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (ups *UserPointSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = ups.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userpoint.Label}
	default:
		err = fmt.Errorf("model: UserPointSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (ups *UserPointSelect) BoolX(ctx context.Context) bool {
	v, err := ups.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (ups *UserPointSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := ups.sql.Query()
	if err := ups.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
