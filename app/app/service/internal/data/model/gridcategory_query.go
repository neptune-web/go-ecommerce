// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/app/service/internal/data/model/gridcategory"
	"mall-go/app/app/service/internal/data/model/predicate"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// GridCategoryQuery is the builder for querying GridCategory entities.
type GridCategoryQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.GridCategory
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the GridCategoryQuery builder.
func (gcq *GridCategoryQuery) Where(ps ...predicate.GridCategory) *GridCategoryQuery {
	gcq.predicates = append(gcq.predicates, ps...)
	return gcq
}

// Limit adds a limit step to the query.
func (gcq *GridCategoryQuery) Limit(limit int) *GridCategoryQuery {
	gcq.limit = &limit
	return gcq
}

// Offset adds an offset step to the query.
func (gcq *GridCategoryQuery) Offset(offset int) *GridCategoryQuery {
	gcq.offset = &offset
	return gcq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (gcq *GridCategoryQuery) Unique(unique bool) *GridCategoryQuery {
	gcq.unique = &unique
	return gcq
}

// Order adds an order step to the query.
func (gcq *GridCategoryQuery) Order(o ...OrderFunc) *GridCategoryQuery {
	gcq.order = append(gcq.order, o...)
	return gcq
}

// First returns the first GridCategory entity from the query.
// Returns a *NotFoundError when no GridCategory was found.
func (gcq *GridCategoryQuery) First(ctx context.Context) (*GridCategory, error) {
	nodes, err := gcq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gridcategory.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (gcq *GridCategoryQuery) FirstX(ctx context.Context) *GridCategory {
	node, err := gcq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first GridCategory ID from the query.
// Returns a *NotFoundError when no GridCategory ID was found.
func (gcq *GridCategoryQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = gcq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gridcategory.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (gcq *GridCategoryQuery) FirstIDX(ctx context.Context) int64 {
	id, err := gcq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Last returns the last GridCategory entity from the query.
// Returns a *NotFoundError when no GridCategory was found.
func (gcq *GridCategoryQuery) Last(ctx context.Context) (*GridCategory, error) {
	nodes, err := gcq.All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{gridcategory.Label}
	}
	return nodes[len(nodes)-1], nil
}

// LastX is like Last, but panics if an error occurs.
func (gcq *GridCategoryQuery) LastX(ctx context.Context) *GridCategory {
	node, err := gcq.Last(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// LastID returns the last GridCategory ID from the query.
// Returns a *NotFoundError when no GridCategory ID was found.
func (gcq *GridCategoryQuery) LastID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = gcq.IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{gridcategory.Label}
		return
	}
	return ids[len(ids)-1], nil
}

// LastIDX is like LastID, but panics if an error occurs.
func (gcq *GridCategoryQuery) LastIDX(ctx context.Context) int64 {
	id, err := gcq.LastID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single GridCategory entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when exactly one GridCategory entity is not found.
// Returns a *NotFoundError when no GridCategory entities are found.
func (gcq *GridCategoryQuery) Only(ctx context.Context) (*GridCategory, error) {
	nodes, err := gcq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{gridcategory.Label}
	default:
		return nil, &NotSingularError{gridcategory.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (gcq *GridCategoryQuery) OnlyX(ctx context.Context) *GridCategory {
	node, err := gcq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only GridCategory ID in the query.
// Returns a *NotSingularError when exactly one GridCategory ID is not found.
// Returns a *NotFoundError when no entities are found.
func (gcq *GridCategoryQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = gcq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{gridcategory.Label}
	default:
		err = &NotSingularError{gridcategory.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (gcq *GridCategoryQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := gcq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of GridCategories.
func (gcq *GridCategoryQuery) All(ctx context.Context) ([]*GridCategory, error) {
	if err := gcq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return gcq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (gcq *GridCategoryQuery) AllX(ctx context.Context) []*GridCategory {
	nodes, err := gcq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of GridCategory IDs.
func (gcq *GridCategoryQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := gcq.Select(gridcategory.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (gcq *GridCategoryQuery) IDsX(ctx context.Context) []int64 {
	ids, err := gcq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (gcq *GridCategoryQuery) Count(ctx context.Context) (int, error) {
	if err := gcq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return gcq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (gcq *GridCategoryQuery) CountX(ctx context.Context) int {
	count, err := gcq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (gcq *GridCategoryQuery) Exist(ctx context.Context) (bool, error) {
	if err := gcq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return gcq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (gcq *GridCategoryQuery) ExistX(ctx context.Context) bool {
	exist, err := gcq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the GridCategoryQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (gcq *GridCategoryQuery) Clone() *GridCategoryQuery {
	if gcq == nil {
		return nil
	}
	return &GridCategoryQuery{
		config:     gcq.config,
		limit:      gcq.limit,
		offset:     gcq.offset,
		order:      append([]OrderFunc{}, gcq.order...),
		predicates: append([]predicate.GridCategory{}, gcq.predicates...),
		// clone intermediate query.
		sql:  gcq.sql.Clone(),
		path: gcq.path,
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
//	client.GridCategory.Query().
//		GroupBy(gridcategory.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
//
func (gcq *GridCategoryQuery) GroupBy(field string, fields ...string) *GridCategoryGroupBy {
	group := &GridCategoryGroupBy{config: gcq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := gcq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return gcq.sqlQuery(ctx), nil
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
//	client.GridCategory.Query().
//		Select(gridcategory.FieldCreateTime).
//		Scan(ctx, &v)
//
func (gcq *GridCategoryQuery) Select(fields ...string) *GridCategorySelect {
	gcq.fields = append(gcq.fields, fields...)
	return &GridCategorySelect{GridCategoryQuery: gcq}
}

func (gcq *GridCategoryQuery) prepareQuery(ctx context.Context) error {
	for _, f := range gcq.fields {
		if !gridcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if gcq.path != nil {
		prev, err := gcq.path(ctx)
		if err != nil {
			return err
		}
		gcq.sql = prev
	}
	return nil
}

func (gcq *GridCategoryQuery) sqlAll(ctx context.Context) ([]*GridCategory, error) {
	var (
		nodes = []*GridCategory{}
		_spec = gcq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &GridCategory{config: gcq.config}
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
	if err := sqlgraph.QueryNodes(ctx, gcq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (gcq *GridCategoryQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := gcq.querySpec()
	return sqlgraph.CountNodes(ctx, gcq.driver, _spec)
}

func (gcq *GridCategoryQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := gcq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (gcq *GridCategoryQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   gridcategory.Table,
			Columns: gridcategory.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: gridcategory.FieldID,
			},
		},
		From:   gcq.sql,
		Unique: true,
	}
	if unique := gcq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := gcq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, gridcategory.FieldID)
		for i := range fields {
			if fields[i] != gridcategory.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := gcq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := gcq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := gcq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := gcq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (gcq *GridCategoryQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(gcq.driver.Dialect())
	t1 := builder.Table(gridcategory.Table)
	columns := gcq.fields
	if len(columns) == 0 {
		columns = gridcategory.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if gcq.sql != nil {
		selector = gcq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	for _, p := range gcq.predicates {
		p(selector)
	}
	for _, p := range gcq.order {
		p(selector)
	}
	if offset := gcq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := gcq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// GridCategoryGroupBy is the group-by builder for GridCategory entities.
type GridCategoryGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (gcgb *GridCategoryGroupBy) Aggregate(fns ...AggregateFunc) *GridCategoryGroupBy {
	gcgb.fns = append(gcgb.fns, fns...)
	return gcgb
}

// Scan applies the group-by query and scans the result into the given value.
func (gcgb *GridCategoryGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := gcgb.path(ctx)
	if err != nil {
		return err
	}
	gcgb.sql = query
	return gcgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gcgb *GridCategoryGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := gcgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GridCategoryGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("model: GridCategoryGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gcgb *GridCategoryGroupBy) StringsX(ctx context.Context) []string {
	v, err := gcgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GridCategoryGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gcgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gridcategory.Label}
	default:
		err = fmt.Errorf("model: GridCategoryGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gcgb *GridCategoryGroupBy) StringX(ctx context.Context) string {
	v, err := gcgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GridCategoryGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("model: GridCategoryGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gcgb *GridCategoryGroupBy) IntsX(ctx context.Context) []int {
	v, err := gcgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GridCategoryGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gcgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gridcategory.Label}
	default:
		err = fmt.Errorf("model: GridCategoryGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gcgb *GridCategoryGroupBy) IntX(ctx context.Context) int {
	v, err := gcgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GridCategoryGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("model: GridCategoryGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gcgb *GridCategoryGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := gcgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GridCategoryGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gcgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gridcategory.Label}
	default:
		err = fmt.Errorf("model: GridCategoryGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gcgb *GridCategoryGroupBy) Float64X(ctx context.Context) float64 {
	v, err := gcgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GridCategoryGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(gcgb.fields) > 1 {
		return nil, errors.New("model: GridCategoryGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := gcgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gcgb *GridCategoryGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := gcgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (gcgb *GridCategoryGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gcgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gridcategory.Label}
	default:
		err = fmt.Errorf("model: GridCategoryGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gcgb *GridCategoryGroupBy) BoolX(ctx context.Context) bool {
	v, err := gcgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gcgb *GridCategoryGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range gcgb.fields {
		if !gridcategory.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := gcgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := gcgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (gcgb *GridCategoryGroupBy) sqlQuery() *sql.Selector {
	selector := gcgb.sql.Select()
	aggregation := make([]string, 0, len(gcgb.fns))
	for _, fn := range gcgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(gcgb.fields)+len(gcgb.fns))
		for _, f := range gcgb.fields {
			columns = append(columns, selector.C(f))
		}
		for _, c := range aggregation {
			columns = append(columns, c)
		}
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(gcgb.fields...)...)
}

// GridCategorySelect is the builder for selecting fields of GridCategory entities.
type GridCategorySelect struct {
	*GridCategoryQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (gcs *GridCategorySelect) Scan(ctx context.Context, v interface{}) error {
	if err := gcs.prepareQuery(ctx); err != nil {
		return err
	}
	gcs.sql = gcs.GridCategoryQuery.sqlQuery(ctx)
	return gcs.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (gcs *GridCategorySelect) ScanX(ctx context.Context, v interface{}) {
	if err := gcs.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (gcs *GridCategorySelect) Strings(ctx context.Context) ([]string, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("model: GridCategorySelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (gcs *GridCategorySelect) StringsX(ctx context.Context) []string {
	v, err := gcs.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (gcs *GridCategorySelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = gcs.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gridcategory.Label}
	default:
		err = fmt.Errorf("model: GridCategorySelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (gcs *GridCategorySelect) StringX(ctx context.Context) string {
	v, err := gcs.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (gcs *GridCategorySelect) Ints(ctx context.Context) ([]int, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("model: GridCategorySelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (gcs *GridCategorySelect) IntsX(ctx context.Context) []int {
	v, err := gcs.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (gcs *GridCategorySelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = gcs.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gridcategory.Label}
	default:
		err = fmt.Errorf("model: GridCategorySelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (gcs *GridCategorySelect) IntX(ctx context.Context) int {
	v, err := gcs.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (gcs *GridCategorySelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("model: GridCategorySelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (gcs *GridCategorySelect) Float64sX(ctx context.Context) []float64 {
	v, err := gcs.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (gcs *GridCategorySelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = gcs.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gridcategory.Label}
	default:
		err = fmt.Errorf("model: GridCategorySelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (gcs *GridCategorySelect) Float64X(ctx context.Context) float64 {
	v, err := gcs.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (gcs *GridCategorySelect) Bools(ctx context.Context) ([]bool, error) {
	if len(gcs.fields) > 1 {
		return nil, errors.New("model: GridCategorySelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := gcs.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (gcs *GridCategorySelect) BoolsX(ctx context.Context) []bool {
	v, err := gcs.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (gcs *GridCategorySelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = gcs.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{gridcategory.Label}
	default:
		err = fmt.Errorf("model: GridCategorySelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (gcs *GridCategorySelect) BoolX(ctx context.Context) bool {
	v, err := gcs.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (gcs *GridCategorySelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := gcs.sql.Query()
	if err := gcs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
