// Code generated by entc, DO NOT EDIT.

package model

import (
	"context"
	"fmt"
	"mall-go/internal/data/model/predicate"
	"mall-go/internal/data/model/spu"
	"mall-go/internal/data/model/tag"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// TagUpdate is the builder for updating Tag entities.
type TagUpdate struct {
	config
	hooks    []Hook
	mutation *TagMutation
}

// Where appends a list predicates to the TagUpdate builder.
func (tu *TagUpdate) Where(ps ...predicate.Tag) *TagUpdate {
	tu.mutation.Where(ps...)
	return tu
}

// SetUpdateTime sets the "update_time" field.
func (tu *TagUpdate) SetUpdateTime(t time.Time) *TagUpdate {
	tu.mutation.SetUpdateTime(t)
	return tu
}

// SetDeleteTime sets the "delete_time" field.
func (tu *TagUpdate) SetDeleteTime(t time.Time) *TagUpdate {
	tu.mutation.SetDeleteTime(t)
	return tu
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (tu *TagUpdate) SetNillableDeleteTime(t *time.Time) *TagUpdate {
	if t != nil {
		tu.SetDeleteTime(*t)
	}
	return tu
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (tu *TagUpdate) ClearDeleteTime() *TagUpdate {
	tu.mutation.ClearDeleteTime()
	return tu
}

// SetTitle sets the "title" field.
func (tu *TagUpdate) SetTitle(s string) *TagUpdate {
	tu.mutation.SetTitle(s)
	return tu
}

// SetDescription sets the "description" field.
func (tu *TagUpdate) SetDescription(s string) *TagUpdate {
	tu.mutation.SetDescription(s)
	return tu
}

// SetHighlight sets the "highlight" field.
func (tu *TagUpdate) SetHighlight(i int) *TagUpdate {
	tu.mutation.ResetHighlight()
	tu.mutation.SetHighlight(i)
	return tu
}

// AddHighlight adds i to the "highlight" field.
func (tu *TagUpdate) AddHighlight(i int) *TagUpdate {
	tu.mutation.AddHighlight(i)
	return tu
}

// SetType sets the "type" field.
func (tu *TagUpdate) SetType(i int) *TagUpdate {
	tu.mutation.ResetType()
	tu.mutation.SetType(i)
	return tu
}

// AddType adds i to the "type" field.
func (tu *TagUpdate) AddType(i int) *TagUpdate {
	tu.mutation.AddType(i)
	return tu
}

// AddSpuIDs adds the "spu" edge to the Spu entity by IDs.
func (tu *TagUpdate) AddSpuIDs(ids ...int64) *TagUpdate {
	tu.mutation.AddSpuIDs(ids...)
	return tu
}

// AddSpu adds the "spu" edges to the Spu entity.
func (tu *TagUpdate) AddSpu(s ...*Spu) *TagUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.AddSpuIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tu *TagUpdate) Mutation() *TagMutation {
	return tu.mutation
}

// ClearSpu clears all "spu" edges to the Spu entity.
func (tu *TagUpdate) ClearSpu() *TagUpdate {
	tu.mutation.ClearSpu()
	return tu
}

// RemoveSpuIDs removes the "spu" edge to Spu entities by IDs.
func (tu *TagUpdate) RemoveSpuIDs(ids ...int64) *TagUpdate {
	tu.mutation.RemoveSpuIDs(ids...)
	return tu
}

// RemoveSpu removes "spu" edges to Spu entities.
func (tu *TagUpdate) RemoveSpu(s ...*Spu) *TagUpdate {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tu.RemoveSpuIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (tu *TagUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	tu.defaults()
	if len(tu.hooks) == 0 {
		affected, err = tu.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tu.mutation = mutation
			affected, err = tu.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(tu.hooks) - 1; i >= 0; i-- {
			if tu.hooks[i] == nil {
				return 0, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = tu.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tu.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (tu *TagUpdate) SaveX(ctx context.Context) int {
	affected, err := tu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (tu *TagUpdate) Exec(ctx context.Context) error {
	_, err := tu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tu *TagUpdate) ExecX(ctx context.Context) {
	if err := tu.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tu *TagUpdate) defaults() {
	if _, ok := tu.mutation.UpdateTime(); !ok {
		v := tag.UpdateDefaultUpdateTime()
		tu.mutation.SetUpdateTime(v)
	}
}

func (tu *TagUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: tag.FieldID,
			},
		},
	}
	if ps := tu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tu.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldUpdateTime,
		})
	}
	if value, ok := tu.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldDeleteTime,
		})
	}
	if tu.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tag.FieldDeleteTime,
		})
	}
	if value, ok := tu.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldTitle,
		})
	}
	if value, ok := tu.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldDescription,
		})
	}
	if value, ok := tu.mutation.Highlight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldHighlight,
		})
	}
	if value, ok := tu.mutation.AddedHighlight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldHighlight,
		})
	}
	if value, ok := tu.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldType,
		})
	}
	if value, ok := tu.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldType,
		})
	}
	if tu.mutation.SpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.SpuTable,
			Columns: tag.SpuPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.RemovedSpuIDs(); len(nodes) > 0 && !tu.mutation.SpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.SpuTable,
			Columns: tag.SpuPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tu.mutation.SpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.SpuTable,
			Columns: tag.SpuPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, tu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return 0, err
	}
	return n, nil
}

// TagUpdateOne is the builder for updating a single Tag entity.
type TagUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *TagMutation
}

// SetUpdateTime sets the "update_time" field.
func (tuo *TagUpdateOne) SetUpdateTime(t time.Time) *TagUpdateOne {
	tuo.mutation.SetUpdateTime(t)
	return tuo
}

// SetDeleteTime sets the "delete_time" field.
func (tuo *TagUpdateOne) SetDeleteTime(t time.Time) *TagUpdateOne {
	tuo.mutation.SetDeleteTime(t)
	return tuo
}

// SetNillableDeleteTime sets the "delete_time" field if the given value is not nil.
func (tuo *TagUpdateOne) SetNillableDeleteTime(t *time.Time) *TagUpdateOne {
	if t != nil {
		tuo.SetDeleteTime(*t)
	}
	return tuo
}

// ClearDeleteTime clears the value of the "delete_time" field.
func (tuo *TagUpdateOne) ClearDeleteTime() *TagUpdateOne {
	tuo.mutation.ClearDeleteTime()
	return tuo
}

// SetTitle sets the "title" field.
func (tuo *TagUpdateOne) SetTitle(s string) *TagUpdateOne {
	tuo.mutation.SetTitle(s)
	return tuo
}

// SetDescription sets the "description" field.
func (tuo *TagUpdateOne) SetDescription(s string) *TagUpdateOne {
	tuo.mutation.SetDescription(s)
	return tuo
}

// SetHighlight sets the "highlight" field.
func (tuo *TagUpdateOne) SetHighlight(i int) *TagUpdateOne {
	tuo.mutation.ResetHighlight()
	tuo.mutation.SetHighlight(i)
	return tuo
}

// AddHighlight adds i to the "highlight" field.
func (tuo *TagUpdateOne) AddHighlight(i int) *TagUpdateOne {
	tuo.mutation.AddHighlight(i)
	return tuo
}

// SetType sets the "type" field.
func (tuo *TagUpdateOne) SetType(i int) *TagUpdateOne {
	tuo.mutation.ResetType()
	tuo.mutation.SetType(i)
	return tuo
}

// AddType adds i to the "type" field.
func (tuo *TagUpdateOne) AddType(i int) *TagUpdateOne {
	tuo.mutation.AddType(i)
	return tuo
}

// AddSpuIDs adds the "spu" edge to the Spu entity by IDs.
func (tuo *TagUpdateOne) AddSpuIDs(ids ...int64) *TagUpdateOne {
	tuo.mutation.AddSpuIDs(ids...)
	return tuo
}

// AddSpu adds the "spu" edges to the Spu entity.
func (tuo *TagUpdateOne) AddSpu(s ...*Spu) *TagUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.AddSpuIDs(ids...)
}

// Mutation returns the TagMutation object of the builder.
func (tuo *TagUpdateOne) Mutation() *TagMutation {
	return tuo.mutation
}

// ClearSpu clears all "spu" edges to the Spu entity.
func (tuo *TagUpdateOne) ClearSpu() *TagUpdateOne {
	tuo.mutation.ClearSpu()
	return tuo
}

// RemoveSpuIDs removes the "spu" edge to Spu entities by IDs.
func (tuo *TagUpdateOne) RemoveSpuIDs(ids ...int64) *TagUpdateOne {
	tuo.mutation.RemoveSpuIDs(ids...)
	return tuo
}

// RemoveSpu removes "spu" edges to Spu entities.
func (tuo *TagUpdateOne) RemoveSpu(s ...*Spu) *TagUpdateOne {
	ids := make([]int64, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return tuo.RemoveSpuIDs(ids...)
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (tuo *TagUpdateOne) Select(field string, fields ...string) *TagUpdateOne {
	tuo.fields = append([]string{field}, fields...)
	return tuo
}

// Save executes the query and returns the updated Tag entity.
func (tuo *TagUpdateOne) Save(ctx context.Context) (*Tag, error) {
	var (
		err  error
		node *Tag
	)
	tuo.defaults()
	if len(tuo.hooks) == 0 {
		node, err = tuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*TagMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			tuo.mutation = mutation
			node, err = tuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(tuo.hooks) - 1; i >= 0; i-- {
			if tuo.hooks[i] == nil {
				return nil, fmt.Errorf("model: uninitialized hook (forgotten import model/runtime?)")
			}
			mut = tuo.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, tuo.mutation); err != nil {
			return nil, err
		}
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (tuo *TagUpdateOne) SaveX(ctx context.Context) *Tag {
	node, err := tuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (tuo *TagUpdateOne) Exec(ctx context.Context) error {
	_, err := tuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (tuo *TagUpdateOne) ExecX(ctx context.Context) {
	if err := tuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (tuo *TagUpdateOne) defaults() {
	if _, ok := tuo.mutation.UpdateTime(); !ok {
		v := tag.UpdateDefaultUpdateTime()
		tuo.mutation.SetUpdateTime(v)
	}
}

func (tuo *TagUpdateOne) sqlSave(ctx context.Context) (_node *Tag, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   tag.Table,
			Columns: tag.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeInt64,
				Column: tag.FieldID,
			},
		},
	}
	id, ok := tuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "ID", err: fmt.Errorf("missing Tag.ID for update")}
	}
	_spec.Node.ID.Value = id
	if fields := tuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, tag.FieldID)
		for _, f := range fields {
			if !tag.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("model: invalid field %q for query", f)}
			}
			if f != tag.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := tuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := tuo.mutation.UpdateTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldUpdateTime,
		})
	}
	if value, ok := tuo.mutation.DeleteTime(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Value:  value,
			Column: tag.FieldDeleteTime,
		})
	}
	if tuo.mutation.DeleteTimeCleared() {
		_spec.Fields.Clear = append(_spec.Fields.Clear, &sqlgraph.FieldSpec{
			Type:   field.TypeTime,
			Column: tag.FieldDeleteTime,
		})
	}
	if value, ok := tuo.mutation.Title(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldTitle,
		})
	}
	if value, ok := tuo.mutation.Description(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeString,
			Value:  value,
			Column: tag.FieldDescription,
		})
	}
	if value, ok := tuo.mutation.Highlight(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldHighlight,
		})
	}
	if value, ok := tuo.mutation.AddedHighlight(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldHighlight,
		})
	}
	if value, ok := tuo.mutation.GetType(); ok {
		_spec.Fields.Set = append(_spec.Fields.Set, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldType,
		})
	}
	if value, ok := tuo.mutation.AddedType(); ok {
		_spec.Fields.Add = append(_spec.Fields.Add, &sqlgraph.FieldSpec{
			Type:   field.TypeInt,
			Value:  value,
			Column: tag.FieldType,
		})
	}
	if tuo.mutation.SpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.SpuTable,
			Columns: tag.SpuPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.RemovedSpuIDs(); len(nodes) > 0 && !tuo.mutation.SpuCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.SpuTable,
			Columns: tag.SpuPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := tuo.mutation.SpuIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   tag.SpuTable,
			Columns: tag.SpuPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: &sqlgraph.FieldSpec{
					Type:   field.TypeInt64,
					Column: spu.FieldID,
				},
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Tag{config: tuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, tuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{tag.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{err.Error(), err}
		}
		return nil, err
	}
	return _node, nil
}