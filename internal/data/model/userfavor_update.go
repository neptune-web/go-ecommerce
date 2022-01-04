// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/internal/data/model/predicate"
	"mall-go/internal/data/model/userfavor"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFavorUpdate is the builder for updating UserFavor entities.
type UserFavorUpdate struct {
	config
	hooks    []Hook
	mutation *UserFavorMutation
}

// Where appends a list predicates to the UserFavorUpdate builder.
func (ufu *UserFavorUpdate) Where(ps ...predicate.UserFavor) *UserFavorUpdate {
	ufu.mutation.Where(ps...)
	return ufu
}

// SetUpdateTime sets the "update_time" field.
func (ufu *UserFavorUpdate) SetUpdateTime(t time.Time) *UserFavorUpdate {
	ufu.mutation.SetUpdateTime(t)
	return ufu
}

// SetDeleteTime sets the "delete_time" field.
func (ufu *UserFavorUpdate) SetDeleteTime(t time.Time) *UserFavorUpdate {
	ufu.mutation.SetDeleteTime(t)
	return ufu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ufu *UserFavorUpdate) SetNillableDeleteTime(t *time.Time) *UserFavorUpdate {
	if t != nil {
		ufu.SetDeleteTime(*t)
	}
	return ufu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ufu *UserFavorUpdate) ClearDeleteTime() *UserFavorUpdate {
	ufu.mutation.ClearDeleteTime()
	return ufu
}

// SetUserID sets the "user_id" field.
func (ufu *UserFavorUpdate) SetUserID(i int64) *UserFavorUpdate {
	ufu.mutation.ResetUserID()
	ufu.mutation.SetUserID(i)
	return ufu
}

// AddUserID adds i to the "user_id" field.
func (ufu *UserFavorUpdate) AddUserID(i int64) *UserFavorUpdate {
	ufu.mutation.AddUserID(i)
	return ufu
}

// SetSpuID sets the "spu_id" field.
func (ufu *UserFavorUpdate) SetSpuID(i int64) *UserFavorUpdate {
	ufu.mutation.ResetSpuID()
	ufu.mutation.SetSpuID(i)
	return ufu
}

// AddSpuID adds i to the "spu_id" field.
func (ufu *UserFavorUpdate) AddSpuID(i int64) *UserFavorUpdate {
	ufu.mutation.AddSpuID(i)
	return ufu
}

// SetStatus sets the "status" field.
func (ufu *UserFavorUpdate) SetStatus(i int) *UserFavorUpdate {
	ufu.mutation.ResetStatus()
	ufu.mutation.SetStatus(i)
	return ufu
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ufu *UserFavorUpdate) SetNillableStatus(i *int) *UserFavorUpdate {
	if i != nil {
		ufu.SetStatus(*i)
	}
	return ufu
}

// AddStatus adds i to the "status" field.
func (ufu *UserFavorUpdate) AddStatus(i int) *UserFavorUpdate {
	ufu.mutation.AddStatus(i)
	return ufu
}

// Mutation returns the UserFavorMutation object of the builder.
func (ufu *UserFavorUpdate) Mutation() *UserFavorMutation {
	return ufu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ufu *UserFavorUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	ufu.defaults()
	if len(ufu.hooks) == 0 {
		affected, err = ufu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserFavorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ufu.mutation = mutation
			affected, err = ufu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(ufu.hooks) - 1; i >= 0; i-- {
			if ufu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ufu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ufu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (ufu *UserFavorUpdate) SaveX(ctx context.Context) int {
	affected, err := ufu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ufu *UserFavorUpdate) Exec(ctx context.Context) error {
	_, err := ufu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufu *UserFavorUpdate) ExecX(ctx context.Context) {
	if err := ufu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufu *UserFavorUpdate) defaults() {
	if _, ok := ufu.mutation.UpdateTime(); !ok {
		v := userfavor.UpdateDefaultUpdateTime()
		ufu.mutation.SetUpdateTime(v)
	}
}

func (ufu *UserFavorUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userfavor.Table,
			Columns: userfavor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userfavor.FieldID,
			},
		},
	}
	if ps := ufu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userfavor.FieldUpdateTime,
		})
	}
	if value, ok := ufu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userfavor.FieldDeleteTime,
		})
	}
	if ufu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userfavor.FieldDeleteTime,
		})
	}
	if value, ok := ufu.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldUserID,
		})
	}
	if value, ok := ufu.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldUserID,
		})
	}
	if value, ok := ufu.mutation.SpuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldSpuID,
		})
	}
	if value, ok := ufu.mutation.AddedSpuID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldSpuID,
		})
	}
	if value, ok := ufu.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userfavor.FieldStatus,
		})
	}
	if value, ok := ufu.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userfavor.FieldStatus,
		})
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ufu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userfavor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// UserFavorUpdateOne is the builder for updating a single UserFavor entity.
type UserFavorUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *UserFavorMutation
}

// SetUpdateTime sets the "update_time" field.
func (ufuo *UserFavorUpdateOne) SetUpdateTime(t time.Time) *UserFavorUpdateOne {
	ufuo.mutation.SetUpdateTime(t)
	return ufuo
}

// SetDeleteTime sets the "delete_time" field.
func (ufuo *UserFavorUpdateOne) SetDeleteTime(t time.Time) *UserFavorUpdateOne {
	ufuo.mutation.SetDeleteTime(t)
	return ufuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ufuo *UserFavorUpdateOne) SetNillableDeleteTime(t *time.Time) *UserFavorUpdateOne {
	if t != nil {
		ufuo.SetDeleteTime(*t)
	}
	return ufuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (ufuo *UserFavorUpdateOne) ClearDeleteTime() *UserFavorUpdateOne {
	ufuo.mutation.ClearDeleteTime()
	return ufuo
}

// SetUserID sets the "user_id" field.
func (ufuo *UserFavorUpdateOne) SetUserID(i int64) *UserFavorUpdateOne {
	ufuo.mutation.ResetUserID()
	ufuo.mutation.SetUserID(i)
	return ufuo
}

// AddUserID adds i to the "user_id" field.
func (ufuo *UserFavorUpdateOne) AddUserID(i int64) *UserFavorUpdateOne {
	ufuo.mutation.AddUserID(i)
	return ufuo
}

// SetSpuID sets the "spu_id" field.
func (ufuo *UserFavorUpdateOne) SetSpuID(i int64) *UserFavorUpdateOne {
	ufuo.mutation.ResetSpuID()
	ufuo.mutation.SetSpuID(i)
	return ufuo
}

// AddSpuID adds i to the "spu_id" field.
func (ufuo *UserFavorUpdateOne) AddSpuID(i int64) *UserFavorUpdateOne {
	ufuo.mutation.AddSpuID(i)
	return ufuo
}

// SetStatus sets the "status" field.
func (ufuo *UserFavorUpdateOne) SetStatus(i int) *UserFavorUpdateOne {
	ufuo.mutation.ResetStatus()
	ufuo.mutation.SetStatus(i)
	return ufuo
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ufuo *UserFavorUpdateOne) SetNillableStatus(i *int) *UserFavorUpdateOne {
	if i != nil {
		ufuo.SetStatus(*i)
	}
	return ufuo
}

// AddStatus adds i to the "status" field.
func (ufuo *UserFavorUpdateOne) AddStatus(i int) *UserFavorUpdateOne {
	ufuo.mutation.AddStatus(i)
	return ufuo
}

// Mutation returns the UserFavorMutation object of the builder.
func (ufuo *UserFavorUpdateOne) Mutation() *UserFavorMutation {
	return ufuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ufuo *UserFavorUpdateOne) Select(field string, fields ...string) *UserFavorUpdateOne {
	ufuo.fields = append([]string{field}, fields...)
	return ufuo
}

// Save executes the query and returns the updated UserFavor entity.
func (ufuo *UserFavorUpdateOne) Save(ctx context.Context) (*UserFavor, error) {
	var (
		err  error
		node *UserFavor
	)
	ufuo.defaults()
	if len(ufuo.hooks) == 0 {
		node, err = ufuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserFavorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			ufuo.mutation = mutation
			node, err = ufuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(ufuo.hooks) - 1; i >= 0; i-- {
			if ufuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ufuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ufuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (ufuo *UserFavorUpdateOne) SaveX(ctx context.Context) *UserFavor {
	node, err := ufuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ufuo *UserFavorUpdateOne) Exec(ctx context.Context) error {
	_, err := ufuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufuo *UserFavorUpdateOne) ExecX(ctx context.Context) {
	if err := ufuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufuo *UserFavorUpdateOne) defaults() {
	if _, ok := ufuo.mutation.UpdateTime(); !ok {
		v := userfavor.UpdateDefaultUpdateTime()
		ufuo.mutation.SetUpdateTime(v)
	}
}

func (ufuo *UserFavorUpdateOne) sqlSave(ctx context.Context) (_node *UserFavor, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   userfavor.Table,
			Columns: userfavor.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userfavor.FieldID,
			},
		},
	}
	id, ok := ufuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing UserFavor.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := ufuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, userfavor.FieldID)
		for _, f := range fields {
			if !userfavor.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != userfavor.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ufuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ufuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userfavor.FieldUpdateTime,
		})
	}
	if value, ok := ufuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userfavor.FieldDeleteTime,
		})
	}
	if ufuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: userfavor.FieldDeleteTime,
		})
	}
	if value, ok := ufuo.mutation.UserID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldUserID,
		})
	}
	if value, ok := ufuo.mutation.AddedUserID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldUserID,
		})
	}
	if value, ok := ufuo.mutation.SpuID(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldSpuID,
		})
	}
	if value, ok := ufuo.mutation.AddedSpuID(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldSpuID,
		})
	}
	if value, ok := ufuo.mutation.Status(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userfavor.FieldStatus,
		})
	}
	if value, ok := ufuo.mutation.AddedStatus(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userfavor.FieldStatus,
		})
	}
	_node = &UserFavor{config: ufuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ufuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{userfavor.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}