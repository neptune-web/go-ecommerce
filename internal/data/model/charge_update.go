// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/internal/data/model/charge"
	"mall-go/internal/data/model/predicate"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// ChargeUpdate is the builder for updating Charge entities.
type ChargeUpdate struct {
	config
	hooks    []Hook
	mutation *ChargeMutation
}

// Where appends a list predicates to the ChargeUpdate builder.
func (cu *ChargeUpdate) Where(ps ...predicate.Charge) *ChargeUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetUpdateTime sets the "update_time" field.
func (cu *ChargeUpdate) SetUpdateTime(t time.Time) *ChargeUpdate {
	cu.mutation.SetUpdateTime(t)
	return cu
}

// SetDeleteTime sets the "delete_time" field.
func (cu *ChargeUpdate) SetDeleteTime(t time.Time) *ChargeUpdate {
	cu.mutation.SetDeleteTime(t)
	return cu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cu *ChargeUpdate) SetNillableDeleteTime(t *time.Time) *ChargeUpdate {
	if t != nil {
		cu.SetDeleteTime(*t)
	}
	return cu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (cu *ChargeUpdate) ClearDeleteTime() *ChargeUpdate {
	cu.mutation.ClearDeleteTime()
	return cu
}

// SetUserID sets the "user_id" field.
func (cu *ChargeUpdate) SetUserID(i int64) *ChargeUpdate {
	cu.mutation.ResetUserID()
	cu.mutation.SetUserID(i)
	return cu
}

// AddUserID adds i to the "user_id" field.
func (cu *ChargeUpdate) AddUserID(i int64) *ChargeUpdate {
	cu.mutation.AddUserID(i)
	return cu
}

// SetAmount sets the "amount" field.
func (cu *ChargeUpdate) SetAmount(s string) *ChargeUpdate {
	cu.mutation.SetAmount(s)
	return cu
}

// SetChargeNo sets the "charge_no" field.
func (cu *ChargeUpdate) SetChargeNo(s string) *ChargeUpdate {
	cu.mutation.SetChargeNo(s)
	return cu
}

// SetTransactionID sets the "transaction_id" field.
func (cu *ChargeUpdate) SetTransactionID(s string) *ChargeUpdate {
	cu.mutation.SetTransactionID(s)
	return cu
}

// SetPayWay sets the "pay_way" field.
func (cu *ChargeUpdate) SetPayWay(i int) *ChargeUpdate {
	cu.mutation.ResetPayWay()
	cu.mutation.SetPayWay(i)
	return cu
}

// AddPayWay adds i to the "pay_way" field.
func (cu *ChargeUpdate) AddPayWay(i int) *ChargeUpdate {
	cu.mutation.AddPayWay(i)
	return cu
}

// SetClientType sets the "client_type" field.
func (cu *ChargeUpdate) SetClientType(i int) *ChargeUpdate {
	cu.mutation.ResetClientType()
	cu.mutation.SetClientType(i)
	return cu
}

// AddClientType adds i to the "client_type" field.
func (cu *ChargeUpdate) AddClientType(i int) *ChargeUpdate {
	cu.mutation.AddClientType(i)
	return cu
}

// Mutation returns the ChargeMutation object of the builder.
func (cu *ChargeUpdate) Mutation() *ChargeMutation {
	return cu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ChargeUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	cu.defaults()
	if len(cu.hooks) == 0 {
		affected, err = cu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChargeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cu.mutation = mutation
			affected, err = cu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(cu.hooks) - 1; i >= 0; i-- {
			if cu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = cu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ChargeUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ChargeUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ChargeUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cu *ChargeUpdate) defaults() {
	if _, ok := cu.mutation.UpdateTime(); !ok {
		v := charge.UpdateDefaultUpdateTime()
		cu.mutation.SetUpdateTime(v)
	}
}

func (cu *ChargeUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   charge.Table,
			Columns: charge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: charge.FieldID,
			},
		},
	}
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: charge.FieldUpdateTime,
		})
	}
	if value, ok := cu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: charge.FieldDeleteTime,
		})
	}
	if cu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: charge.FieldDeleteTime,
		})
	}
	if value, ok := cu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: charge.FieldUserID,
		})
	}
	if value, ok := cu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: charge.FieldUserID,
		})
	}
	if value, ok := cu.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: charge.FieldAmount,
		})
	}
	if value, ok := cu.mutation.ChargeNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: charge.FieldChargeNo,
		})
	}
	if value, ok := cu.mutation.TransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: charge.FieldTransactionID,
		})
	}
	if value, ok := cu.mutation.PayWay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: charge.FieldPayWay,
		})
	}
	if value, ok := cu.mutation.AddedPayWay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: charge.FieldPayWay,
		})
	}
	if value, ok := cu.mutation.ClientType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: charge.FieldClientType,
		})
	}
	if value, ok := cu.mutation.AddedClientType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: charge.FieldClientType,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{charge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// ChargeUpdateOne is the builder for updating a single Charge entity.
type ChargeUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ChargeMutation
}

// SetUpdateTime sets the "update_time" field.
func (cuo *ChargeUpdateOne) SetUpdateTime(t time.Time) *ChargeUpdateOne {
	cuo.mutation.SetUpdateTime(t)
	return cuo
}

// SetDeleteTime sets the "delete_time" field.
func (cuo *ChargeUpdateOne) SetDeleteTime(t time.Time) *ChargeUpdateOne {
	cuo.mutation.SetDeleteTime(t)
	return cuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (cuo *ChargeUpdateOne) SetNillableDeleteTime(t *time.Time) *ChargeUpdateOne {
	if t != nil {
		cuo.SetDeleteTime(*t)
	}
	return cuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (cuo *ChargeUpdateOne) ClearDeleteTime() *ChargeUpdateOne {
	cuo.mutation.ClearDeleteTime()
	return cuo
}

// SetUserID sets the "user_id" field.
func (cuo *ChargeUpdateOne) SetUserID(i int64) *ChargeUpdateOne {
	cuo.mutation.ResetUserID()
	cuo.mutation.SetUserID(i)
	return cuo
}

// AddUserID adds i to the "user_id" field.
func (cuo *ChargeUpdateOne) AddUserID(i int64) *ChargeUpdateOne {
	cuo.mutation.AddUserID(i)
	return cuo
}

// SetAmount sets the "amount" field.
func (cuo *ChargeUpdateOne) SetAmount(s string) *ChargeUpdateOne {
	cuo.mutation.SetAmount(s)
	return cuo
}

// SetChargeNo sets the "charge_no" field.
func (cuo *ChargeUpdateOne) SetChargeNo(s string) *ChargeUpdateOne {
	cuo.mutation.SetChargeNo(s)
	return cuo
}

// SetTransactionID sets the "transaction_id" field.
func (cuo *ChargeUpdateOne) SetTransactionID(s string) *ChargeUpdateOne {
	cuo.mutation.SetTransactionID(s)
	return cuo
}

// SetPayWay sets the "pay_way" field.
func (cuo *ChargeUpdateOne) SetPayWay(i int) *ChargeUpdateOne {
	cuo.mutation.ResetPayWay()
	cuo.mutation.SetPayWay(i)
	return cuo
}

// AddPayWay adds i to the "pay_way" field.
func (cuo *ChargeUpdateOne) AddPayWay(i int) *ChargeUpdateOne {
	cuo.mutation.AddPayWay(i)
	return cuo
}

// SetClientType sets the "client_type" field.
func (cuo *ChargeUpdateOne) SetClientType(i int) *ChargeUpdateOne {
	cuo.mutation.ResetClientType()
	cuo.mutation.SetClientType(i)
	return cuo
}

// AddClientType adds i to the "client_type" field.
func (cuo *ChargeUpdateOne) AddClientType(i int) *ChargeUpdateOne {
	cuo.mutation.AddClientType(i)
	return cuo
}

// Mutation returns the ChargeMutation object of the builder.
func (cuo *ChargeUpdateOne) Mutation() *ChargeMutation {
	return cuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ChargeUpdateOne) Select(field string, fields ...string) *ChargeUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Charge entity.
func (cuo *ChargeUpdateOne) Save(ctx context.Context) (*Charge, error) {
	var (
		err  error
		node *Charge
	)
	cuo.defaults()
	if len(cuo.hooks) == 0 {
		node, err = cuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*ChargeMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			cuo.mutation = mutation
			node, err = cuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(cuo.hooks) - 1; i >= 0; i-- {
			if cuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = cuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, cuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ChargeUpdateOne) SaveX(ctx context.Context) *Charge {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ChargeUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ChargeUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (cuo *ChargeUpdateOne) defaults() {
	if _, ok := cuo.mutation.UpdateTime(); !ok {
		v := charge.UpdateDefaultUpdateTime()
		cuo.mutation.SetUpdateTime(v)
	}
}

func (cuo *ChargeUpdateOne) sqlSave(ctx context.Context) (_node *Charge, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   charge.Table,
			Columns: charge.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: charge.FieldID,
			},
		},
	}
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Charge.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, charge.FieldID)
		for _, f := range fields {
			if !charge.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != charge.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: charge.FieldUpdateTime,
		})
	}
	if value, ok := cuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: charge.FieldDeleteTime,
		})
	}
	if cuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: charge.FieldDeleteTime,
		})
	}
	if value, ok := cuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: charge.FieldUserID,
		})
	}
	if value, ok := cuo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: charge.FieldUserID,
		})
	}
	if value, ok := cuo.mutation.Amount(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: charge.FieldAmount,
		})
	}
	if value, ok := cuo.mutation.ChargeNo(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: charge.FieldChargeNo,
		})
	}
	if value, ok := cuo.mutation.TransactionID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: charge.FieldTransactionID,
		})
	}
	if value, ok := cuo.mutation.PayWay(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: charge.FieldPayWay,
		})
	}
	if value, ok := cuo.mutation.AddedPayWay(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: charge.FieldPayWay,
		})
	}
	if value, ok := cuo.mutation.ClientType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: charge.FieldClientType,
		})
	}
	if value, ok := cuo.mutation.AddedClientType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: charge.FieldClientType,
		})
	}
	_node = &Charge{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{charge.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}