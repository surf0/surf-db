// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/ent/recordksf"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordKsfCreate is the builder for creating a RecordKsf entity.
type RecordKsfCreate struct {
	config
	mutation *RecordKsfMutation
	hooks    []Hook
}

// SetTimestamp sets the "timestamp" field.
func (rkc *RecordKsfCreate) SetTimestamp(t time.Time) *RecordKsfCreate {
	rkc.mutation.SetTimestamp(t)
	return rkc
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (rkc *RecordKsfCreate) SetNillableTimestamp(t *time.Time) *RecordKsfCreate {
	if t != nil {
		rkc.SetTimestamp(*t)
	}
	return rkc
}

// SetPlayerName sets the "player_name" field.
func (rkc *RecordKsfCreate) SetPlayerName(s string) *RecordKsfCreate {
	rkc.mutation.SetPlayerName(s)
	return rkc
}

// SetMapName sets the "map_name" field.
func (rkc *RecordKsfCreate) SetMapName(s string) *RecordKsfCreate {
	rkc.mutation.SetMapName(s)
	return rkc
}

// SetTime sets the "time" field.
func (rkc *RecordKsfCreate) SetTime(s string) *RecordKsfCreate {
	rkc.mutation.SetTime(s)
	return rkc
}

// SetImprovement sets the "improvement" field.
func (rkc *RecordKsfCreate) SetImprovement(s string) *RecordKsfCreate {
	rkc.mutation.SetImprovement(s)
	return rkc
}

// SetServer sets the "server" field.
func (rkc *RecordKsfCreate) SetServer(s string) *RecordKsfCreate {
	rkc.mutation.SetServer(s)
	return rkc
}

// SetID sets the "id" field.
func (rkc *RecordKsfCreate) SetID(s string) *RecordKsfCreate {
	rkc.mutation.SetID(s)
	return rkc
}

// Mutation returns the RecordKsfMutation object of the builder.
func (rkc *RecordKsfCreate) Mutation() *RecordKsfMutation {
	return rkc.mutation
}

// Save creates the RecordKsf in the database.
func (rkc *RecordKsfCreate) Save(ctx context.Context) (*RecordKsf, error) {
	var (
		err  error
		node *RecordKsf
	)
	rkc.defaults()
	if len(rkc.hooks) == 0 {
		if err = rkc.check(); err != nil {
			return nil, err
		}
		node, err = rkc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordKsfMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rkc.check(); err != nil {
				return nil, err
			}
			rkc.mutation = mutation
			if node, err = rkc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rkc.hooks) - 1; i >= 0; i-- {
			if rkc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rkc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rkc.mutation)
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

// SaveX calls Save and panics if Save returns an error.
func (rkc *RecordKsfCreate) SaveX(ctx context.Context) *RecordKsf {
	v, err := rkc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rkc *RecordKsfCreate) Exec(ctx context.Context) error {
	_, err := rkc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rkc *RecordKsfCreate) ExecX(ctx context.Context) {
	if err := rkc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rkc *RecordKsfCreate) defaults() {
	if _, ok := rkc.mutation.Timestamp(); !ok {
		v := recordksf.DefaultTimestamp()
		rkc.mutation.SetTimestamp(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rkc *RecordKsfCreate) check() error {
	if _, ok := rkc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "RecordKsf.timestamp"`)}
	}
	if _, ok := rkc.mutation.PlayerName(); !ok {
		return &ValidationError{Name: "player_name", err: errors.New(`ent: missing required field "RecordKsf.player_name"`)}
	}
	if _, ok := rkc.mutation.MapName(); !ok {
		return &ValidationError{Name: "map_name", err: errors.New(`ent: missing required field "RecordKsf.map_name"`)}
	}
	if _, ok := rkc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "RecordKsf.time"`)}
	}
	if _, ok := rkc.mutation.Improvement(); !ok {
		return &ValidationError{Name: "improvement", err: errors.New(`ent: missing required field "RecordKsf.improvement"`)}
	}
	if _, ok := rkc.mutation.Server(); !ok {
		return &ValidationError{Name: "server", err: errors.New(`ent: missing required field "RecordKsf.server"`)}
	}
	return nil
}

func (rkc *RecordKsfCreate) sqlSave(ctx context.Context) (*RecordKsf, error) {
	_node, _spec := rkc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rkc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected RecordKsf.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (rkc *RecordKsfCreate) createSpec() (*RecordKsf, *sqlgraph.CreateSpec) {
	var (
		_node = &RecordKsf{config: rkc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recordksf.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: recordksf.FieldID,
			},
		}
	)
	if id, ok := rkc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rkc.mutation.Timestamp(); ok {
		_spec.SetField(recordksf.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if value, ok := rkc.mutation.PlayerName(); ok {
		_spec.SetField(recordksf.FieldPlayerName, field.TypeString, value)
		_node.PlayerName = value
	}
	if value, ok := rkc.mutation.MapName(); ok {
		_spec.SetField(recordksf.FieldMapName, field.TypeString, value)
		_node.MapName = value
	}
	if value, ok := rkc.mutation.Time(); ok {
		_spec.SetField(recordksf.FieldTime, field.TypeString, value)
		_node.Time = value
	}
	if value, ok := rkc.mutation.Improvement(); ok {
		_spec.SetField(recordksf.FieldImprovement, field.TypeString, value)
		_node.Improvement = value
	}
	if value, ok := rkc.mutation.Server(); ok {
		_spec.SetField(recordksf.FieldServer, field.TypeString, value)
		_node.Server = value
	}
	return _node, _spec
}

// RecordKsfCreateBulk is the builder for creating many RecordKsf entities in bulk.
type RecordKsfCreateBulk struct {
	config
	builders []*RecordKsfCreate
}

// Save creates the RecordKsf entities in the database.
func (rkcb *RecordKsfCreateBulk) Save(ctx context.Context) ([]*RecordKsf, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rkcb.builders))
	nodes := make([]*RecordKsf, len(rkcb.builders))
	mutators := make([]Mutator, len(rkcb.builders))
	for i := range rkcb.builders {
		func(i int, root context.Context) {
			builder := rkcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecordKsfMutation)
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
					_, err = mutators[i+1].Mutate(root, rkcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rkcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rkcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rkcb *RecordKsfCreateBulk) SaveX(ctx context.Context) []*RecordKsf {
	v, err := rkcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rkcb *RecordKsfCreateBulk) Exec(ctx context.Context) error {
	_, err := rkcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rkcb *RecordKsfCreateBulk) ExecX(ctx context.Context) {
	if err := rkcb.Exec(ctx); err != nil {
		panic(err)
	}
}