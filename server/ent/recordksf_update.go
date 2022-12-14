// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/ent/predicate"
	"server/ent/recordksf"
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordKsfUpdate is the builder for updating RecordKsf entities.
type RecordKsfUpdate struct {
	config
	hooks    []Hook
	mutation *RecordKsfMutation
}

// Where appends a list predicates to the RecordKsfUpdate builder.
func (rku *RecordKsfUpdate) Where(ps ...predicate.RecordKsf) *RecordKsfUpdate {
	rku.mutation.Where(ps...)
	return rku
}

// SetTimestamp sets the "timestamp" field.
func (rku *RecordKsfUpdate) SetTimestamp(t time.Time) *RecordKsfUpdate {
	rku.mutation.SetTimestamp(t)
	return rku
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (rku *RecordKsfUpdate) SetNillableTimestamp(t *time.Time) *RecordKsfUpdate {
	if t != nil {
		rku.SetTimestamp(*t)
	}
	return rku
}

// SetPlayerName sets the "player_name" field.
func (rku *RecordKsfUpdate) SetPlayerName(s string) *RecordKsfUpdate {
	rku.mutation.SetPlayerName(s)
	return rku
}

// SetMapName sets the "map_name" field.
func (rku *RecordKsfUpdate) SetMapName(s string) *RecordKsfUpdate {
	rku.mutation.SetMapName(s)
	return rku
}

// SetTime sets the "time" field.
func (rku *RecordKsfUpdate) SetTime(s string) *RecordKsfUpdate {
	rku.mutation.SetTime(s)
	return rku
}

// SetImprovement sets the "improvement" field.
func (rku *RecordKsfUpdate) SetImprovement(s string) *RecordKsfUpdate {
	rku.mutation.SetImprovement(s)
	return rku
}

// SetServer sets the "server" field.
func (rku *RecordKsfUpdate) SetServer(s string) *RecordKsfUpdate {
	rku.mutation.SetServer(s)
	return rku
}

// Mutation returns the RecordKsfMutation object of the builder.
func (rku *RecordKsfUpdate) Mutation() *RecordKsfMutation {
	return rku.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (rku *RecordKsfUpdate) Save(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rku.hooks) == 0 {
		affected, err = rku.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordKsfMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rku.mutation = mutation
			affected, err = rku.sqlSave(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rku.hooks) - 1; i >= 0; i-- {
			if rku.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rku.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rku.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// SaveX is like Save, but panics if an error occurs.
func (rku *RecordKsfUpdate) SaveX(ctx context.Context) int {
	affected, err := rku.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (rku *RecordKsfUpdate) Exec(ctx context.Context) error {
	_, err := rku.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rku *RecordKsfUpdate) ExecX(ctx context.Context) {
	if err := rku.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rku *RecordKsfUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recordksf.Table,
			Columns: recordksf.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: recordksf.FieldID,
			},
		},
	}
	if ps := rku.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rku.mutation.Timestamp(); ok {
		_spec.SetField(recordksf.FieldTimestamp, field.TypeTime, value)
	}
	if value, ok := rku.mutation.PlayerName(); ok {
		_spec.SetField(recordksf.FieldPlayerName, field.TypeString, value)
	}
	if value, ok := rku.mutation.MapName(); ok {
		_spec.SetField(recordksf.FieldMapName, field.TypeString, value)
	}
	if value, ok := rku.mutation.Time(); ok {
		_spec.SetField(recordksf.FieldTime, field.TypeString, value)
	}
	if value, ok := rku.mutation.Improvement(); ok {
		_spec.SetField(recordksf.FieldImprovement, field.TypeString, value)
	}
	if value, ok := rku.mutation.Server(); ok {
		_spec.SetField(recordksf.FieldServer, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, rku.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recordksf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	return n, nil
}

// RecordKsfUpdateOne is the builder for updating a single RecordKsf entity.
type RecordKsfUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RecordKsfMutation
}

// SetTimestamp sets the "timestamp" field.
func (rkuo *RecordKsfUpdateOne) SetTimestamp(t time.Time) *RecordKsfUpdateOne {
	rkuo.mutation.SetTimestamp(t)
	return rkuo
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (rkuo *RecordKsfUpdateOne) SetNillableTimestamp(t *time.Time) *RecordKsfUpdateOne {
	if t != nil {
		rkuo.SetTimestamp(*t)
	}
	return rkuo
}

// SetPlayerName sets the "player_name" field.
func (rkuo *RecordKsfUpdateOne) SetPlayerName(s string) *RecordKsfUpdateOne {
	rkuo.mutation.SetPlayerName(s)
	return rkuo
}

// SetMapName sets the "map_name" field.
func (rkuo *RecordKsfUpdateOne) SetMapName(s string) *RecordKsfUpdateOne {
	rkuo.mutation.SetMapName(s)
	return rkuo
}

// SetTime sets the "time" field.
func (rkuo *RecordKsfUpdateOne) SetTime(s string) *RecordKsfUpdateOne {
	rkuo.mutation.SetTime(s)
	return rkuo
}

// SetImprovement sets the "improvement" field.
func (rkuo *RecordKsfUpdateOne) SetImprovement(s string) *RecordKsfUpdateOne {
	rkuo.mutation.SetImprovement(s)
	return rkuo
}

// SetServer sets the "server" field.
func (rkuo *RecordKsfUpdateOne) SetServer(s string) *RecordKsfUpdateOne {
	rkuo.mutation.SetServer(s)
	return rkuo
}

// Mutation returns the RecordKsfMutation object of the builder.
func (rkuo *RecordKsfUpdateOne) Mutation() *RecordKsfMutation {
	return rkuo.mutation
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (rkuo *RecordKsfUpdateOne) Select(field string, fields ...string) *RecordKsfUpdateOne {
	rkuo.fields = append([]string{field}, fields...)
	return rkuo
}

// Save executes the query and returns the updated RecordKsf entity.
func (rkuo *RecordKsfUpdateOne) Save(ctx context.Context) (*RecordKsf, error) {
	var (
		err  error
		node *RecordKsf
	)
	if len(rkuo.hooks) == 0 {
		node, err = rkuo.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordKsfMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rkuo.mutation = mutation
			node, err = rkuo.sqlSave(ctx)
			mutation.done = true
			return node, err
		})
		for i := len(rkuo.hooks) - 1; i >= 0; i-- {
			if rkuo.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rkuo.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rkuo.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RecordKsf)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RecordKsfMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX is like Save, but panics if an error occurs.
func (rkuo *RecordKsfUpdateOne) SaveX(ctx context.Context) *RecordKsf {
	node, err := rkuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (rkuo *RecordKsfUpdateOne) Exec(ctx context.Context) error {
	_, err := rkuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rkuo *RecordKsfUpdateOne) ExecX(ctx context.Context) {
	if err := rkuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (rkuo *RecordKsfUpdateOne) sqlSave(ctx context.Context) (_node *RecordKsf, err error) {
	_spec := &sqlgraph.UpdateSpec{
		Node: &sqlgraph.NodeSpec{
			Table:   recordksf.Table,
			Columns: recordksf.Columns,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: recordksf.FieldID,
			},
		},
	}
	id, ok := rkuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "RecordKsf.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := rkuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, recordksf.FieldID)
		for _, f := range fields {
			if !recordksf.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != recordksf.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := rkuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := rkuo.mutation.Timestamp(); ok {
		_spec.SetField(recordksf.FieldTimestamp, field.TypeTime, value)
	}
	if value, ok := rkuo.mutation.PlayerName(); ok {
		_spec.SetField(recordksf.FieldPlayerName, field.TypeString, value)
	}
	if value, ok := rkuo.mutation.MapName(); ok {
		_spec.SetField(recordksf.FieldMapName, field.TypeString, value)
	}
	if value, ok := rkuo.mutation.Time(); ok {
		_spec.SetField(recordksf.FieldTime, field.TypeString, value)
	}
	if value, ok := rkuo.mutation.Improvement(); ok {
		_spec.SetField(recordksf.FieldImprovement, field.TypeString, value)
	}
	if value, ok := rkuo.mutation.Server(); ok {
		_spec.SetField(recordksf.FieldServer, field.TypeString, value)
	}
	_node = &RecordKsf{config: rkuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, rkuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{recordksf.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	return _node, nil
}
