// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/babyname/fate/ent/predicate"
	"github.com/babyname/fate/ent/version"
)

// VersionUpdate is the builder for updating Version entities.
type VersionUpdate struct {
	config
	hooks    []Hook
	mutation *VersionMutation
}

// Where appends a list predicates to the VersionUpdate builder.
func (vu *VersionUpdate) Where(ps ...predicate.Version) *VersionUpdate {
	vu.mutation.Where(ps...)
	return vu
}

// SetCurrentVersion sets the "current_version" field.
func (vu *VersionUpdate) SetCurrentVersion(i int) *VersionUpdate {
	vu.mutation.ResetCurrentVersion()
	vu.mutation.SetCurrentVersion(i)
	return vu
}

// AddCurrentVersion adds i to the "current_version" field.
func (vu *VersionUpdate) AddCurrentVersion(i int) *VersionUpdate {
	vu.mutation.AddCurrentVersion(i)
	return vu
}

// SetUpdatedUnix sets the "updated_unix" field.
func (vu *VersionUpdate) SetUpdatedUnix(i int) *VersionUpdate {
	vu.mutation.ResetUpdatedUnix()
	vu.mutation.SetUpdatedUnix(i)
	return vu
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (vu *VersionUpdate) AddUpdatedUnix(i int) *VersionUpdate {
	vu.mutation.AddUpdatedUnix(i)
	return vu
}

// Mutation returns the VersionMutation object of the builder.
func (vu *VersionUpdate) Mutation() *VersionMutation {
	return vu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (vu *VersionUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, VersionMutation](ctx, vu.sqlSave, vu.mutation, vu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vu *VersionUpdate) SaveX(ctx context.Context) int {
	affected, err := vu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (vu *VersionUpdate) Exec(ctx context.Context) error {
	_, err := vu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vu *VersionUpdate) ExecX(ctx context.Context) {
	if err := vu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vu *VersionUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(version.Table, version.Columns, sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt))
	if ps := vu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vu.mutation.CurrentVersion(); ok {
		_spec.SetField(version.FieldCurrentVersion, field.TypeInt, value)
	}
	if value, ok := vu.mutation.AddedCurrentVersion(); ok {
		_spec.AddField(version.FieldCurrentVersion, field.TypeInt, value)
	}
	if value, ok := vu.mutation.UpdatedUnix(); ok {
		_spec.SetField(version.FieldUpdatedUnix, field.TypeInt, value)
	}
	if value, ok := vu.mutation.AddedUpdatedUnix(); ok {
		_spec.AddField(version.FieldUpdatedUnix, field.TypeInt, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, vu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{version.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	vu.mutation.done = true
	return n, nil
}

// VersionUpdateOne is the builder for updating a single Version entity.
type VersionUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *VersionMutation
}

// SetCurrentVersion sets the "current_version" field.
func (vuo *VersionUpdateOne) SetCurrentVersion(i int) *VersionUpdateOne {
	vuo.mutation.ResetCurrentVersion()
	vuo.mutation.SetCurrentVersion(i)
	return vuo
}

// AddCurrentVersion adds i to the "current_version" field.
func (vuo *VersionUpdateOne) AddCurrentVersion(i int) *VersionUpdateOne {
	vuo.mutation.AddCurrentVersion(i)
	return vuo
}

// SetUpdatedUnix sets the "updated_unix" field.
func (vuo *VersionUpdateOne) SetUpdatedUnix(i int) *VersionUpdateOne {
	vuo.mutation.ResetUpdatedUnix()
	vuo.mutation.SetUpdatedUnix(i)
	return vuo
}

// AddUpdatedUnix adds i to the "updated_unix" field.
func (vuo *VersionUpdateOne) AddUpdatedUnix(i int) *VersionUpdateOne {
	vuo.mutation.AddUpdatedUnix(i)
	return vuo
}

// Mutation returns the VersionMutation object of the builder.
func (vuo *VersionUpdateOne) Mutation() *VersionMutation {
	return vuo.mutation
}

// Where appends a list predicates to the VersionUpdate builder.
func (vuo *VersionUpdateOne) Where(ps ...predicate.Version) *VersionUpdateOne {
	vuo.mutation.Where(ps...)
	return vuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (vuo *VersionUpdateOne) Select(field string, fields ...string) *VersionUpdateOne {
	vuo.fields = append([]string{field}, fields...)
	return vuo
}

// Save executes the query and returns the updated Version entity.
func (vuo *VersionUpdateOne) Save(ctx context.Context) (*Version, error) {
	return withHooks[*Version, VersionMutation](ctx, vuo.sqlSave, vuo.mutation, vuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (vuo *VersionUpdateOne) SaveX(ctx context.Context) *Version {
	node, err := vuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (vuo *VersionUpdateOne) Exec(ctx context.Context) error {
	_, err := vuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (vuo *VersionUpdateOne) ExecX(ctx context.Context) {
	if err := vuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (vuo *VersionUpdateOne) sqlSave(ctx context.Context) (_node *Version, err error) {
	_spec := sqlgraph.NewUpdateSpec(version.Table, version.Columns, sqlgraph.NewFieldSpec(version.FieldID, field.TypeInt))
	id, ok := vuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Version.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := vuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, version.FieldID)
		for _, f := range fields {
			if !version.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != version.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := vuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := vuo.mutation.CurrentVersion(); ok {
		_spec.SetField(version.FieldCurrentVersion, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.AddedCurrentVersion(); ok {
		_spec.AddField(version.FieldCurrentVersion, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.UpdatedUnix(); ok {
		_spec.SetField(version.FieldUpdatedUnix, field.TypeInt, value)
	}
	if value, ok := vuo.mutation.AddedUpdatedUnix(); ok {
		_spec.AddField(version.FieldUpdatedUnix, field.TypeInt, value)
	}
	_node = &Version{config: vuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, vuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{version.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	vuo.mutation.done = true
	return _node, nil
}
