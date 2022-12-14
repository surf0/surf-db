// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"
	"server/ent/recordsh"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordShCreate is the builder for creating a RecordSh entity.
type RecordShCreate struct {
	config
	mutation *RecordShMutation
	hooks    []Hook
}

// SetTimestamp sets the "timestamp" field.
func (rsc *RecordShCreate) SetTimestamp(t time.Time) *RecordShCreate {
	rsc.mutation.SetTimestamp(t)
	return rsc
}

// SetNillableTimestamp sets the "timestamp" field if the given value is not nil.
func (rsc *RecordShCreate) SetNillableTimestamp(t *time.Time) *RecordShCreate {
	if t != nil {
		rsc.SetTimestamp(*t)
	}
	return rsc
}

// SetPlayerName sets the "player_name" field.
func (rsc *RecordShCreate) SetPlayerName(s string) *RecordShCreate {
	rsc.mutation.SetPlayerName(s)
	return rsc
}

// SetPlayerID sets the "player_id" field.
func (rsc *RecordShCreate) SetPlayerID(s string) *RecordShCreate {
	rsc.mutation.SetPlayerID(s)
	return rsc
}

// SetType sets the "type" field.
func (rsc *RecordShCreate) SetType(s string) *RecordShCreate {
	rsc.mutation.SetType(s)
	return rsc
}

// SetTrack sets the "track" field.
func (rsc *RecordShCreate) SetTrack(i int) *RecordShCreate {
	rsc.mutation.SetTrack(i)
	return rsc
}

// SetMapName sets the "map_name" field.
func (rsc *RecordShCreate) SetMapName(s string) *RecordShCreate {
	rsc.mutation.SetMapName(s)
	return rsc
}

// SetTime sets the "time" field.
func (rsc *RecordShCreate) SetTime(s string) *RecordShCreate {
	rsc.mutation.SetTime(s)
	return rsc
}

// SetImprovement sets the "improvement" field.
func (rsc *RecordShCreate) SetImprovement(s string) *RecordShCreate {
	rsc.mutation.SetImprovement(s)
	return rsc
}

// SetServer sets the "server" field.
func (rsc *RecordShCreate) SetServer(s string) *RecordShCreate {
	rsc.mutation.SetServer(s)
	return rsc
}

// SetID sets the "id" field.
func (rsc *RecordShCreate) SetID(s string) *RecordShCreate {
	rsc.mutation.SetID(s)
	return rsc
}

// Mutation returns the RecordShMutation object of the builder.
func (rsc *RecordShCreate) Mutation() *RecordShMutation {
	return rsc.mutation
}

// Save creates the RecordSh in the database.
func (rsc *RecordShCreate) Save(ctx context.Context) (*RecordSh, error) {
	var (
		err  error
		node *RecordSh
	)
	rsc.defaults()
	if len(rsc.hooks) == 0 {
		if err = rsc.check(); err != nil {
			return nil, err
		}
		node, err = rsc.sqlSave(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordShMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			if err = rsc.check(); err != nil {
				return nil, err
			}
			rsc.mutation = mutation
			if node, err = rsc.sqlSave(ctx); err != nil {
				return nil, err
			}
			mutation.id = &node.ID
			mutation.done = true
			return node, err
		})
		for i := len(rsc.hooks) - 1; i >= 0; i-- {
			if rsc.hooks[i] == nil {
				return nil, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rsc.hooks[i](mut)
		}
		v, err := mut.Mutate(ctx, rsc.mutation)
		if err != nil {
			return nil, err
		}
		nv, ok := v.(*RecordSh)
		if !ok {
			return nil, fmt.Errorf("unexpected node type %T returned from RecordShMutation", v)
		}
		node = nv
	}
	return node, err
}

// SaveX calls Save and panics if Save returns an error.
func (rsc *RecordShCreate) SaveX(ctx context.Context) *RecordSh {
	v, err := rsc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rsc *RecordShCreate) Exec(ctx context.Context) error {
	_, err := rsc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rsc *RecordShCreate) ExecX(ctx context.Context) {
	if err := rsc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rsc *RecordShCreate) defaults() {
	if _, ok := rsc.mutation.Timestamp(); !ok {
		v := recordsh.DefaultTimestamp()
		rsc.mutation.SetTimestamp(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rsc *RecordShCreate) check() error {
	if _, ok := rsc.mutation.Timestamp(); !ok {
		return &ValidationError{Name: "timestamp", err: errors.New(`ent: missing required field "RecordSh.timestamp"`)}
	}
	if _, ok := rsc.mutation.PlayerName(); !ok {
		return &ValidationError{Name: "player_name", err: errors.New(`ent: missing required field "RecordSh.player_name"`)}
	}
	if _, ok := rsc.mutation.PlayerID(); !ok {
		return &ValidationError{Name: "player_id", err: errors.New(`ent: missing required field "RecordSh.player_id"`)}
	}
	if _, ok := rsc.mutation.GetType(); !ok {
		return &ValidationError{Name: "type", err: errors.New(`ent: missing required field "RecordSh.type"`)}
	}
	if _, ok := rsc.mutation.Track(); !ok {
		return &ValidationError{Name: "track", err: errors.New(`ent: missing required field "RecordSh.track"`)}
	}
	if _, ok := rsc.mutation.MapName(); !ok {
		return &ValidationError{Name: "map_name", err: errors.New(`ent: missing required field "RecordSh.map_name"`)}
	}
	if _, ok := rsc.mutation.Time(); !ok {
		return &ValidationError{Name: "time", err: errors.New(`ent: missing required field "RecordSh.time"`)}
	}
	if _, ok := rsc.mutation.Improvement(); !ok {
		return &ValidationError{Name: "improvement", err: errors.New(`ent: missing required field "RecordSh.improvement"`)}
	}
	if _, ok := rsc.mutation.Server(); !ok {
		return &ValidationError{Name: "server", err: errors.New(`ent: missing required field "RecordSh.server"`)}
	}
	return nil
}

func (rsc *RecordShCreate) sqlSave(ctx context.Context) (*RecordSh, error) {
	_node, _spec := rsc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rsc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected RecordSh.ID type: %T", _spec.ID.Value)
		}
	}
	return _node, nil
}

func (rsc *RecordShCreate) createSpec() (*RecordSh, *sqlgraph.CreateSpec) {
	var (
		_node = &RecordSh{config: rsc.config}
		_spec = &sqlgraph.CreateSpec{
			Table: recordsh.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: recordsh.FieldID,
			},
		}
	)
	if id, ok := rsc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := rsc.mutation.Timestamp(); ok {
		_spec.SetField(recordsh.FieldTimestamp, field.TypeTime, value)
		_node.Timestamp = value
	}
	if value, ok := rsc.mutation.PlayerName(); ok {
		_spec.SetField(recordsh.FieldPlayerName, field.TypeString, value)
		_node.PlayerName = value
	}
	if value, ok := rsc.mutation.PlayerID(); ok {
		_spec.SetField(recordsh.FieldPlayerID, field.TypeString, value)
		_node.PlayerID = value
	}
	if value, ok := rsc.mutation.GetType(); ok {
		_spec.SetField(recordsh.FieldType, field.TypeString, value)
		_node.Type = value
	}
	if value, ok := rsc.mutation.Track(); ok {
		_spec.SetField(recordsh.FieldTrack, field.TypeInt, value)
		_node.Track = value
	}
	if value, ok := rsc.mutation.MapName(); ok {
		_spec.SetField(recordsh.FieldMapName, field.TypeString, value)
		_node.MapName = value
	}
	if value, ok := rsc.mutation.Time(); ok {
		_spec.SetField(recordsh.FieldTime, field.TypeString, value)
		_node.Time = value
	}
	if value, ok := rsc.mutation.Improvement(); ok {
		_spec.SetField(recordsh.FieldImprovement, field.TypeString, value)
		_node.Improvement = value
	}
	if value, ok := rsc.mutation.Server(); ok {
		_spec.SetField(recordsh.FieldServer, field.TypeString, value)
		_node.Server = value
	}
	return _node, _spec
}

// RecordShCreateBulk is the builder for creating many RecordSh entities in bulk.
type RecordShCreateBulk struct {
	config
	builders []*RecordShCreate
}

// Save creates the RecordSh entities in the database.
func (rscb *RecordShCreateBulk) Save(ctx context.Context) ([]*RecordSh, error) {
	specs := make([]*sqlgraph.CreateSpec, len(rscb.builders))
	nodes := make([]*RecordSh, len(rscb.builders))
	mutators := make([]Mutator, len(rscb.builders))
	for i := range rscb.builders {
		func(i int, root context.Context) {
			builder := rscb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RecordShMutation)
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
					_, err = mutators[i+1].Mutate(root, rscb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rscb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, rscb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rscb *RecordShCreateBulk) SaveX(ctx context.Context) []*RecordSh {
	v, err := rscb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rscb *RecordShCreateBulk) Exec(ctx context.Context) error {
	_, err := rscb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rscb *RecordShCreateBulk) ExecX(ctx context.Context) {
	if err := rscb.Exec(ctx); err != nil {
		panic(err)
	}
}
