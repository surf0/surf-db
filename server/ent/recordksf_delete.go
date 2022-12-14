// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"server/ent/predicate"
	"server/ent/recordksf"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
)

// RecordKsfDelete is the builder for deleting a RecordKsf entity.
type RecordKsfDelete struct {
	config
	hooks    []Hook
	mutation *RecordKsfMutation
}

// Where appends a list predicates to the RecordKsfDelete builder.
func (rkd *RecordKsfDelete) Where(ps ...predicate.RecordKsf) *RecordKsfDelete {
	rkd.mutation.Where(ps...)
	return rkd
}

// Exec executes the deletion query and returns how many vertices were deleted.
func (rkd *RecordKsfDelete) Exec(ctx context.Context) (int, error) {
	var (
		err      error
		affected int
	)
	if len(rkd.hooks) == 0 {
		affected, err = rkd.sqlExec(ctx)
	} else {
		var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
			mutation, ok := m.(*RecordKsfMutation)
			if !ok {
				return nil, fmt.Errorf("unexpected mutation type %T", m)
			}
			rkd.mutation = mutation
			affected, err = rkd.sqlExec(ctx)
			mutation.done = true
			return affected, err
		})
		for i := len(rkd.hooks) - 1; i >= 0; i-- {
			if rkd.hooks[i] == nil {
				return 0, fmt.Errorf("ent: uninitialized hook (forgotten import ent/runtime?)")
			}
			mut = rkd.hooks[i](mut)
		}
		if _, err := mut.Mutate(ctx, rkd.mutation); err != nil {
			return 0, err
		}
	}
	return affected, err
}

// ExecX is like Exec, but panics if an error occurs.
func (rkd *RecordKsfDelete) ExecX(ctx context.Context) int {
	n, err := rkd.Exec(ctx)
	if err != nil {
		panic(err)
	}
	return n
}

func (rkd *RecordKsfDelete) sqlExec(ctx context.Context) (int, error) {
	_spec := &sqlgraph.DeleteSpec{
		Node: &sqlgraph.NodeSpec{
			Table: recordksf.Table,
			ID: &sqlgraph.FieldSpec{
				Type:   field.TypeString,
				Column: recordksf.FieldID,
			},
		},
	}
	if ps := rkd.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	affected, err := sqlgraph.DeleteNodes(ctx, rkd.driver, _spec)
	if err != nil && sqlgraph.IsConstraintError(err) {
		err = &ConstraintError{msg: err.Error(), wrap: err}
	}
	return affected, err
}

// RecordKsfDeleteOne is the builder for deleting a single RecordKsf entity.
type RecordKsfDeleteOne struct {
	rkd *RecordKsfDelete
}

// Exec executes the deletion query.
func (rkdo *RecordKsfDeleteOne) Exec(ctx context.Context) error {
	n, err := rkdo.rkd.Exec(ctx)
	switch {
	case err != nil:
		return err
	case n == 0:
		return &NotFoundError{recordksf.Label}
	default:
		return nil
	}
}

// ExecX is like Exec, but panics if an error occurs.
func (rkdo *RecordKsfDeleteOne) ExecX(ctx context.Context) {
	rkdo.rkd.ExecX(ctx)
}
