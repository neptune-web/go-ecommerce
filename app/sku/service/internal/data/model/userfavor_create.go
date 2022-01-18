// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"errors"
	"fmt"
	"mall-go/app/sku/service/internal/data/model/userfavor"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// UserFavorCreate is the builder for creating a UserFavor entity.
type UserFavorCreate struct {
	config
	mutation *UserFavorMutation
	hooks    []Hook
}

// SetCreateTime sets the "create_time" field.
func (ufc *UserFavorCreate) SetCreateTime(t time.Time) *UserFavorCreate {
	ufc.mutation.SetCreateTime(t)
	return ufc
}

// SetNillableCreateTime sets the "create_time" field if the given value is not nil.
func (ufc *UserFavorCreate) SetNillableCreateTime(t *time.Time) *UserFavorCreate {
	if t != nil {
		ufc.SetCreateTime(*t)
	}
	return ufc
}

// SetUpdateTime sets the "update_time" field.
func (ufc *UserFavorCreate) SetUpdateTime(t time.Time) *UserFavorCreate {
	ufc.mutation.SetUpdateTime(t)
	return ufc
}

// SetNillableUpdateTime sets the "update_time" field if the given value is not nil.
func (ufc *UserFavorCreate) SetNillableUpdateTime(t *time.Time) *UserFavorCreate {
	if t != nil {
		ufc.SetUpdateTime(*t)
	}
	return ufc
}

// SetDeleteTime sets the "delete_time" field.
func (ufc *UserFavorCreate) SetDeleteTime(t time.Time) *UserFavorCreate {
	ufc.mutation.SetDeleteTime(t)
	return ufc
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (ufc *UserFavorCreate) SetNillableDeleteTime(t *time.Time) *UserFavorCreate {
	if t != nil {
		ufc.SetDeleteTime(*t)
	}
	return ufc
}

// SetUserID sets the "user_id" field.
func (ufc *UserFavorCreate) SetUserID(i int64) *UserFavorCreate {
	ufc.mutation.SetUserID(i)
	return ufc
}

// SetSpuID sets the "spu_id" field.
func (ufc *UserFavorCreate) SetSpuID(i int64) *UserFavorCreate {
	ufc.mutation.SetSpuID(i)
	return ufc
}

// SetStatus sets the "status" field.
func (ufc *UserFavorCreate) SetStatus(i int) *UserFavorCreate {
	ufc.mutation.SetStatus(i)
	return ufc
}

// SetNillableStatus sets the "status" field if the given value is not nil.
func (ufc *UserFavorCreate) SetNillableStatus(i *int) *UserFavorCreate {
	if i != nil {
		ufc.SetStatus(*i)
	}
	return ufc
}

// Mutation returns the UserFavorMutation object of the builder.
func (ufc *UserFavorCreate) Mutation() *UserFavorMutation {
	return ufc.mutation
}

// Save creates the UserFavor in the database.
func (ufc *UserFavorCreate) Save(ctx context.Context) (*UserFavor, error) {
	var (
		err  error
		node *UserFavor
	)
	ufc.defaults()
	if len(ufc.hooks) == 0 {
		if err = ufc.check(); err != nil {
			return nil, err
		}
		node, err = ufc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*UserFavorMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = ufc.check(); err != nil {
				return nil, err
			}
			ufc.mutation = mutation
			if node, err = ufc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(ufc.hooks) - 1; i >= 0; i-- {
			if ufc.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = ufc.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, ufc.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (ufc *UserFavorCreate) SaveX(ctx context.Context) *UserFavor {
	v, err := ufc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufc *UserFavorCreate) Exec(ctx context.Context) error {
	_, err := ufc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufc *UserFavorCreate) ExecX(ctx context.Context) {
	if err := ufc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (ufc *UserFavorCreate) defaults() {
	if _, ok := ufc.mutation.CreateTime(); !ok {
		v := userfavor.DefaultCreateTime()
		ufc.mutation.SetCreateTime(v)
	}
	if _, ok := ufc.mutation.UpdateTime(); !ok {
		v := userfavor.DefaultUpdateTime()
		ufc.mutation.SetUpdateTime(v)
	}
	if _, ok := ufc.mutation.Status(); !ok {
		v := userfavor.DefaultStatus
		ufc.mutation.SetStatus(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ufc *UserFavorCreate) check() error {
	if _, ok := ufc.mutation.CreateTime(); !ok {
		return &ValidationError{Name: "create_time", err: errors.New(`model: missing required field "create_time"`)}
	}
	if _, ok := ufc.mutation.UpdateTime(); !ok {
		return &ValidationError{Name: "update_time", err: errors.New(`model: missing required field "update_time"`)}
	}
	if _, ok := ufc.mutation.UserID(); !ok {
		return &ValidationError{Name: "user_id", err: errors.New(`model: missing required field "user_id"`)}
	}
	if _, ok := ufc.mutation.SpuID(); !ok {
		return &ValidationError{Name: "spu_id", err: errors.New(`model: missing required field "spu_id"`)}
	}
	if _, ok := ufc.mutation.Status(); !ok {
		return &ValidationError{Name: "status", err: errors.New(`model: missing required field "status"`)}
	}
	return nil
}

func (ufc *UserFavorCreate) sqlSave(ctx context.Context) (*UserFavor, error) {
	_node, _spec := ufc.createSpec()
	if err := sqlgraph.CreateNode(ctx, ufc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int64(id)
	return _node, nil
}

func (ufc *UserFavorCreate) createSpec() (*UserFavor, *sqlgraph.CreateSpec) {
	var (
		_node = &UserFavor{config: ufc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: userfavor.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: userfavor.FieldID,
			},
		}
	)
	if value, ok := ufc.mutation.CreateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userfavor.FieldCreateTime,
		})
		_node.CreateTime = value
	}
	if value, ok := ufc.mutation.UpdateTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userfavor.FieldUpdateTime,
		})
		_node.UpdateTime = value
	}
	if value, ok := ufc.mutation.DeleteTime(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: userfavor.FieldDeleteTime,
		})
		_node.DeleteTime = value
	}
	if value, ok := ufc.mutation.UserID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldUserID,
		})
		_node.UserID = value
	}
	if value, ok := ufc.mutation.SpuID(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt64,
			Value:  value,
			Column: userfavor.FieldSpuID,
		})
		_node.SpuID = value
	}
	if value, ok := ufc.mutation.Status(); ok {
		_spec.Fields = append(_spec.Fields, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: userfavor.FieldStatus,
		})
		_node.Status = value
	}
	return _node, _spec
}

// UserFavorCreateBulk is the builder for creating many UserFavor entities in bulk.
type UserFavorCreateBulk struct {
	config
	builders []*UserFavorCreate
}

// Save creates the UserFavor entities in the database.
func (ufcb *UserFavorCreateBulk) Save(ctx context.Context) ([]*UserFavor, error) {
	specs := make([]*sqlgraph.CreateSpec, len(ufcb.builders))
	nodes := make([]*UserFavor, len(ufcb.builders))
	mutators := make([]Mutator, len(ufcb.builders))
	for i := range ufcb.builders {
		func(i int, root context.Context) {
			builder := ufcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*UserFavorMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				nodes[i], specs[i] = builder.createSpec()
				var err error
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, ufcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, ufcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{err.Error(), err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int64(id)
				}
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, ufcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (ufcb *UserFavorCreateBulk) SaveX(ctx context.Context) []*UserFavor {
	v, err := ufcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (ufcb *UserFavorCreateBulk) Exec(ctx context.Context) error {
	_, err := ufcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ufcb *UserFavorCreateBulk) ExecX(ctx context.Context) {
	if err := ufcb.Exec(ctx); err != nil {
		panic(err)
	}
}