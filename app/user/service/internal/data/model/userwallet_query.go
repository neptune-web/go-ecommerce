// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/user/service/internal/data/model/predicate"
	"mall-go/app/user/service/internal/data/model/userwallet"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserWalletQuery is the builder for querying UserWallet entities.
type UserWalletQuery struct {
	config
	limit      *int
	offset     *int
	unique     *bool
	order      []OrderFunc
	fields     []string
	predicates []predicate.UserWallet
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the UserWalletQuery builder.
func (uwq *UserWalletQuery) Where(ps ...predicate.UserWallet) *UserWalletQuery {
	uwq.predicates = append(uwq.predicates, ps...)
	return uwq
}

// Limit adds a limit step to the query.
func (uwq *UserWalletQuery) Limit(limit int) *UserWalletQuery {
	uwq.limit = &limit
	return uwq
}

// Offset adds an offset step to the query.
func (uwq *UserWalletQuery) Offset(offset int) *UserWalletQuery {
	uwq.offset = &offset
	return uwq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (uwq *UserWalletQuery) Unique(unique bool) *UserWalletQuery {
	uwq.unique = &unique
	return uwq
}

// Order adds an order step to the query.
func (uwq *UserWalletQuery) Order(o ...OrderFunc) *UserWalletQuery {
	uwq.order = append(uwq.order, o...)
	return uwq
}

// First returns the first UserWallet entity from the query.
// Returns a *NotFoundError when no UserWallet was found.
func (uwq *UserWalletQuery) First(ctx context.Context) (*UserWallet, error) {
	nodes, err := uwq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{userwallet.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (uwq *UserWalletQuery) FirstX(ctx context.Context) *UserWallet {
	node, err := uwq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first UserWallet ID from the query.
// Returns a *NotFoundError when no UserWallet ID was found.
func (uwq *UserWalletQuery) FirstID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = uwq.Limit(1).IDs(ctx); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{userwallet.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (uwq *UserWalletQuery) FirstIDX(ctx context.Context) int64 {
	id, err := uwq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single UserWallet entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one UserWallet entity is found.
// Returns a *NotFoundError when no UserWallet entities are found.
func (uwq *UserWalletQuery) Only(ctx context.Context) (*UserWallet, error) {
	nodes, err := uwq.Limit(2).All(ctx)
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{userwallet.Label}
	default:
		return nil, &NotSingularError{userwallet.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (uwq *UserWalletQuery) OnlyX(ctx context.Context) *UserWallet {
	node, err := uwq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only UserWallet ID in the query.
// Returns a *NotSingularError when more than one UserWallet ID is found.
// Returns a *NotFoundError when no entities are found.
func (uwq *UserWalletQuery) OnlyID(ctx context.Context) (id int64, err error) {
	var ids []int64
	if ids, err = uwq.Limit(2).IDs(ctx); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{userwallet.Label}
	default:
		err = &NotSingularError{userwallet.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (uwq *UserWalletQuery) OnlyIDX(ctx context.Context) int64 {
	id, err := uwq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of UserWallets.
func (uwq *UserWalletQuery) All(ctx context.Context) ([]*UserWallet, error) {
	if err := uwq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	return uwq.sqlAll(ctx)
}

// AllX is like All, but panics if an error occurs.
func (uwq *UserWalletQuery) AllX(ctx context.Context) []*UserWallet {
	nodes, err := uwq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of UserWallet IDs.
func (uwq *UserWalletQuery) IDs(ctx context.Context) ([]int64, error) {
	var ids []int64
	if err := uwq.Select(userwallet.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (uwq *UserWalletQuery) IDsX(ctx context.Context) []int64 {
	ids, err := uwq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (uwq *UserWalletQuery) Count(ctx context.Context) (int, error) {
	if err := uwq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return uwq.sqlCount(ctx)
}

// CountX is like Count, but panics if an error occurs.
func (uwq *UserWalletQuery) CountX(ctx context.Context) int {
	count, err := uwq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (uwq *UserWalletQuery) Exist(ctx context.Context) (bool, error) {
	if err := uwq.prepareQuery(ctx); err != nil {
		return false, err
	}
	return uwq.sqlExist(ctx)
}

// ExistX is like Exist, but panics if an error occurs.
func (uwq *UserWalletQuery) ExistX(ctx context.Context) bool {
	exist, err := uwq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the UserWalletQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (uwq *UserWalletQuery) Clone() *UserWalletQuery {
	if uwq == nil {
		return nil
	}
	return &UserWalletQuery{
		config:     uwq.config,
		limit:      uwq.limit,
		offset:     uwq.offset,
		order:      append([]OrderFunc{}, uwq.order...),
		predicates: append([]predicate.UserWallet{}, uwq.predicates...),
		// clone intermediate query.
		sql:    uwq.sql.Clone(),
		path:   uwq.path,
		unique: uwq.unique,
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
//	client.UserWallet.Query().
//		GroupBy(userwallet.FieldCreateTime).
//		Aggregate(model.Count()).
//		Scan(ctx, &v)
func (uwq *UserWalletQuery) GroupBy(field string, fields ...string) *UserWalletGroupBy {
	group := &UserWalletGroupBy{config: uwq.config}
	group.fields = append([]string{field}, fields...)
	group.path = func(ctx context.Context) (prev *sql.Selector, err error) {
		if err := uwq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		return uwq.sqlQuery(ctx), nil
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
//	client.UserWallet.Query().
//		Select(userwallet.FieldCreateTime).
//		Scan(ctx, &v)
func (uwq *UserWalletQuery) Select(fields ...string) *UserWalletSelect {
	uwq.fields = append(uwq.fields, fields...)
	return &UserWalletSelect{UserWalletQuery: uwq}
}

func (uwq *UserWalletQuery) prepareQuery(ctx context.Context) error {
	for _, f := range uwq.fields {
		if !userwallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
		}
	}
	if uwq.path != nil {
		prev, err := uwq.path(ctx)
		if err != nil {
			return err
		}
		uwq.sql = prev
	}
	return nil
}

func (uwq *UserWalletQuery) sqlAll(ctx context.Context) ([]*UserWallet, error) {
	var (
		nodes = []*UserWallet{}
		_spec = uwq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]interface{}, error) {
		node := &UserWallet{config: uwq.config}
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
	if err := sqlgraph.QueryNodes(ctx, uwq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	return nodes, nil
}

func (uwq *UserWalletQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := uwq.querySpec()
	_spec.Node.Columns = uwq.fields
	if len(uwq.fields) > 0 {
		_spec.Unique = uwq.unique != nil && *uwq.unique
	}
	return sqlgraph.CountNodes(ctx, uwq.driver, _spec)
}

func (uwq *UserWalletQuery) sqlExist(ctx context.Context) (bool, error) {
	n, err := uwq.sqlCount(ctx)
	if err != nil {
		return false, fmt.Errorf("model: check existence: %w", err)
	}
	return n > 0, nil
}

func (uwq *UserWalletQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := &sqlgraph.QuerySpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userwallet.Table,
			Columns: userwallet.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userwallet.FieldID,
			},
		},
		From:   uwq.sql,
		Unique: true,
	}
	if unique := uwq.unique; unique != nil {
		_spec.Unique = *unique
	}
	if fields := uwq.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userwallet.FieldID)
		for i := range fields {
			if fields[i] != userwallet.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := uwq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := uwq.limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := uwq.offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := uwq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (uwq *UserWalletQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(uwq.driver.Dialect())
	t1 := builder.Table(userwallet.Table)
	columns := uwq.fields
	if len(columns) == 0 {
		columns = userwallet.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if uwq.sql != nil {
		selector = uwq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if uwq.unique != nil && *uwq.unique {
		selector.Distinct()
	}
	for _, p := range uwq.predicates {
		p(selector)
	}
	for _, p := range uwq.order {
		p(selector)
	}
	if offset := uwq.offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := uwq.limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// UserWalletGroupBy is the group-by builder for UserWallet entities.
type UserWalletGroupBy struct {
	config
	fields []string
	fns    []AggregateFunc
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Aggregate adds the given aggregation functions to the group-by query.
func (uwgb *UserWalletGroupBy) Aggregate(fns ...AggregateFunc) *UserWalletGroupBy {
	uwgb.fns = append(uwgb.fns, fns...)
	return uwgb
}

// Scan applies the group-by query and scans the result into the given value.
func (uwgb *UserWalletGroupBy) Scan(ctx context.Context, v interface{}) error {
	query, err := uwgb.path(ctx)
	if err != nil {
		return err
	}
	uwgb.sql = query
	return uwgb.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uwgb *UserWalletGroupBy) ScanX(ctx context.Context, v interface{}) {
	if err := uwgb.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from group-by.
// It is only allowed when executing a group-by query with one field.
func (uwgb *UserWalletGroupBy) Strings(ctx context.Context) ([]string, error) {
	if len(uwgb.fields) > 1 {
		return nil, errors.New("model: UserWalletGroupBy.Strings is not achievable when grouping more than 1 field")
	}
	var v []string
	if err := uwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uwgb *UserWalletGroupBy) StringsX(ctx context.Context) []string {
	v, err := uwgb.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uwgb *UserWalletGroupBy) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uwgb.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userwallet.Label}
	default:
		err = fmt.Errorf("model: UserWalletGroupBy.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uwgb *UserWalletGroupBy) StringX(ctx context.Context) string {
	v, err := uwgb.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from group-by.
// It is only allowed when executing a group-by query with one field.
func (uwgb *UserWalletGroupBy) Ints(ctx context.Context) ([]int, error) {
	if len(uwgb.fields) > 1 {
		return nil, errors.New("model: UserWalletGroupBy.Ints is not achievable when grouping more than 1 field")
	}
	var v []int
	if err := uwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uwgb *UserWalletGroupBy) IntsX(ctx context.Context) []int {
	v, err := uwgb.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uwgb *UserWalletGroupBy) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uwgb.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userwallet.Label}
	default:
		err = fmt.Errorf("model: UserWalletGroupBy.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uwgb *UserWalletGroupBy) IntX(ctx context.Context) int {
	v, err := uwgb.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from group-by.
// It is only allowed when executing a group-by query with one field.
func (uwgb *UserWalletGroupBy) Float64s(ctx context.Context) ([]float64, error) {
	if len(uwgb.fields) > 1 {
		return nil, errors.New("model: UserWalletGroupBy.Float64s is not achievable when grouping more than 1 field")
	}
	var v []float64
	if err := uwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uwgb *UserWalletGroupBy) Float64sX(ctx context.Context) []float64 {
	v, err := uwgb.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uwgb *UserWalletGroupBy) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uwgb.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userwallet.Label}
	default:
		err = fmt.Errorf("model: UserWalletGroupBy.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uwgb *UserWalletGroupBy) Float64X(ctx context.Context) float64 {
	v, err := uwgb.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from group-by.
// It is only allowed when executing a group-by query with one field.
func (uwgb *UserWalletGroupBy) Bools(ctx context.Context) ([]bool, error) {
	if len(uwgb.fields) > 1 {
		return nil, errors.New("model: UserWalletGroupBy.Bools is not achievable when grouping more than 1 field")
	}
	var v []bool
	if err := uwgb.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uwgb *UserWalletGroupBy) BoolsX(ctx context.Context) []bool {
	v, err := uwgb.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a group-by query.
// It is only allowed when executing a group-by query with one field.
func (uwgb *UserWalletGroupBy) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uwgb.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userwallet.Label}
	default:
		err = fmt.Errorf("model: UserWalletGroupBy.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uwgb *UserWalletGroupBy) BoolX(ctx context.Context) bool {
	v, err := uwgb.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uwgb *UserWalletGroupBy) sqlScan(ctx context.Context, v interface{}) error {
	for _, f := range uwgb.fields {
		if !userwallet.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("invalid field %q for group-by", f)}
		}
	}
	selector := uwgb.sqlQuery()
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := uwgb.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

func (uwgb *UserWalletGroupBy) sqlQuery() *sql.Selector {
	selector := uwgb.sql.Select()
	aggregation := make([]string, 0, len(uwgb.fns))
	for _, fn := range uwgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	// If no columns were selected in a custom aggregation function, the default
	// selection is the fields used for "group-by", and the aggregation functions.
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(uwgb.fields)+len(uwgb.fns))
		for _, f := range uwgb.fields {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	return selector.GroupBy(selector.Columns(uwgb.fields...)...)
}

// UserWalletSelect is the builder for selecting fields of UserWallet entities.
type UserWalletSelect struct {
	*UserWalletQuery
	// intermediate query (i.e. traversal path).
	sql *sql.Selector
}

// Scan applies the selector query and scans the result into the given value.
func (uws *UserWalletSelect) Scan(ctx context.Context, v interface{}) error {
	if err := uws.prepareQuery(ctx); err != nil {
		return err
	}
	uws.sql = uws.UserWalletQuery.sqlQuery(ctx)
	return uws.sqlScan(ctx, v)
}

// ScanX is like Scan, but panics if an error occurs.
func (uws *UserWalletSelect) ScanX(ctx context.Context, v interface{}) {
	if err := uws.Scan(ctx, v); err != nil {
		panic(err)
	}
}

// Strings returns list of strings from a selector. It is only allowed when selecting one field.
func (uws *UserWalletSelect) Strings(ctx context.Context) ([]string, error) {
	if len(uws.fields) > 1 {
		return nil, errors.New("model: UserWalletSelect.Strings is not achievable when selecting more than 1 field")
	}
	var v []string
	if err := uws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// StringsX is like Strings, but panics if an error occurs.
func (uws *UserWalletSelect) StringsX(ctx context.Context) []string {
	v, err := uws.Strings(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// String returns a single string from a selector. It is only allowed when selecting one field.
func (uws *UserWalletSelect) String(ctx context.Context) (_ string, err error) {
	var v []string
	if v, err = uws.Strings(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userwallet.Label}
	default:
		err = fmt.Errorf("model: UserWalletSelect.Strings returned %d results when one was expected", len(v))
	}
	return
}

// StringX is like String, but panics if an error occurs.
func (uws *UserWalletSelect) StringX(ctx context.Context) string {
	v, err := uws.String(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Ints returns list of ints from a selector. It is only allowed when selecting one field.
func (uws *UserWalletSelect) Ints(ctx context.Context) ([]int, error) {
	if len(uws.fields) > 1 {
		return nil, errors.New("model: UserWalletSelect.Ints is not achievable when selecting more than 1 field")
	}
	var v []int
	if err := uws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// IntsX is like Ints, but panics if an error occurs.
func (uws *UserWalletSelect) IntsX(ctx context.Context) []int {
	v, err := uws.Ints(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Int returns a single int from a selector. It is only allowed when selecting one field.
func (uws *UserWalletSelect) Int(ctx context.Context) (_ int, err error) {
	var v []int
	if v, err = uws.Ints(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userwallet.Label}
	default:
		err = fmt.Errorf("model: UserWalletSelect.Ints returned %d results when one was expected", len(v))
	}
	return
}

// IntX is like Int, but panics if an error occurs.
func (uws *UserWalletSelect) IntX(ctx context.Context) int {
	v, err := uws.Int(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64s returns list of float64s from a selector. It is only allowed when selecting one field.
func (uws *UserWalletSelect) Float64s(ctx context.Context) ([]float64, error) {
	if len(uws.fields) > 1 {
		return nil, errors.New("model: UserWalletSelect.Float64s is not achievable when selecting more than 1 field")
	}
	var v []float64
	if err := uws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// Float64sX is like Float64s, but panics if an error occurs.
func (uws *UserWalletSelect) Float64sX(ctx context.Context) []float64 {
	v, err := uws.Float64s(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Float64 returns a single float64 from a selector. It is only allowed when selecting one field.
func (uws *UserWalletSelect) Float64(ctx context.Context) (_ float64, err error) {
	var v []float64
	if v, err = uws.Float64s(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userwallet.Label}
	default:
		err = fmt.Errorf("model: UserWalletSelect.Float64s returned %d results when one was expected", len(v))
	}
	return
}

// Float64X is like Float64, but panics if an error occurs.
func (uws *UserWalletSelect) Float64X(ctx context.Context) float64 {
	v, err := uws.Float64(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bools returns list of bools from a selector. It is only allowed when selecting one field.
func (uws *UserWalletSelect) Bools(ctx context.Context) ([]bool, error) {
	if len(uws.fields) > 1 {
		return nil, errors.New("model: UserWalletSelect.Bools is not achievable when selecting more than 1 field")
	}
	var v []bool
	if err := uws.Scan(ctx, &v); err != nil {
		return nil, err
	}
	return v, nil
}

// BoolsX is like Bools, but panics if an error occurs.
func (uws *UserWalletSelect) BoolsX(ctx context.Context) []bool {
	v, err := uws.Bools(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Bool returns a single bool from a selector. It is only allowed when selecting one field.
func (uws *UserWalletSelect) Bool(ctx context.Context) (_ bool, err error) {
	var v []bool
	if v, err = uws.Bools(ctx); err != nil {
		return
	}
	switch len(v) {
	case 1:
		return v[0], nil
	case 0:
		err = &NotFoundError{userwallet.Label}
	default:
		err = fmt.Errorf("model: UserWalletSelect.Bools returned %d results when one was expected", len(v))
	}
	return
}

// BoolX is like Bool, but panics if an error occurs.
func (uws *UserWalletSelect) BoolX(ctx context.Context) bool {
	v, err := uws.Bool(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

func (uws *UserWalletSelect) sqlScan(ctx context.Context, v interface{}) error {
	rows := &sql.Rows{}
	query, args := uws.sql.Query()
	if err := uws.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
