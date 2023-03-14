// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"time"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/babyname/fate/ent/wuxing"
)

// WuXingCreate is the builder for creating a WuXing entity.
type WuXingCreate struct {
	config
	mutation *WuXingMutation
	hooks    []Hook
}

// SetCreated sets the "created" field.
func (wxc *WuXingCreate) SetCreated(t time.Time) *WuXingCreate {
	wxc.mutation.SetCreated(t)
	return wxc
}

// SetNillableCreated sets the "created" field if the given value is not nil.
func (wxc *WuXingCreate) SetNillableCreated(t *time.Time) *WuXingCreate {
	if t != nil {
		wxc.SetCreated(*t)
	}
	return wxc
}

// SetUpdated sets the "updated" field.
func (wxc *WuXingCreate) SetUpdated(t time.Time) *WuXingCreate {
	wxc.mutation.SetUpdated(t)
	return wxc
}

// SetNillableUpdated sets the "updated" field if the given value is not nil.
func (wxc *WuXingCreate) SetNillableUpdated(t *time.Time) *WuXingCreate {
	if t != nil {
		wxc.SetUpdated(*t)
	}
	return wxc
}

// SetDeleted sets the "deleted" field.
func (wxc *WuXingCreate) SetDeleted(t time.Time) *WuXingCreate {
	wxc.mutation.SetDeleted(t)
	return wxc
}

// SetNillableDeleted sets the "deleted" field if the given value is not nil.
func (wxc *WuXingCreate) SetNillableDeleted(t *time.Time) *WuXingCreate {
	if t != nil {
		wxc.SetDeleted(*t)
	}
	return wxc
}

// SetVersion sets the "version" field.
func (wxc *WuXingCreate) SetVersion(i int) *WuXingCreate {
	wxc.mutation.SetVersion(i)
	return wxc
}

// SetNillableVersion sets the "version" field if the given value is not nil.
func (wxc *WuXingCreate) SetNillableVersion(i *int) *WuXingCreate {
	if i != nil {
		wxc.SetVersion(*i)
	}
	return wxc
}

// SetFirst sets the "first" field.
func (wxc *WuXingCreate) SetFirst(s string) *WuXingCreate {
	wxc.mutation.SetFirst(s)
	return wxc
}

// SetNillableFirst sets the "first" field if the given value is not nil.
func (wxc *WuXingCreate) SetNillableFirst(s *string) *WuXingCreate {
	if s != nil {
		wxc.SetFirst(*s)
	}
	return wxc
}

// SetSecond sets the "second" field.
func (wxc *WuXingCreate) SetSecond(s string) *WuXingCreate {
	wxc.mutation.SetSecond(s)
	return wxc
}

// SetNillableSecond sets the "second" field if the given value is not nil.
func (wxc *WuXingCreate) SetNillableSecond(s *string) *WuXingCreate {
	if s != nil {
		wxc.SetSecond(*s)
	}
	return wxc
}

// SetThird sets the "third" field.
func (wxc *WuXingCreate) SetThird(s string) *WuXingCreate {
	wxc.mutation.SetThird(s)
	return wxc
}

// SetNillableThird sets the "third" field if the given value is not nil.
func (wxc *WuXingCreate) SetNillableThird(s *string) *WuXingCreate {
	if s != nil {
		wxc.SetThird(*s)
	}
	return wxc
}

// SetFortune sets the "fortune" field.
func (wxc *WuXingCreate) SetFortune(s string) *WuXingCreate {
	wxc.mutation.SetFortune(s)
	return wxc
}

// SetNillableFortune sets the "fortune" field if the given value is not nil.
func (wxc *WuXingCreate) SetNillableFortune(s *string) *WuXingCreate {
	if s != nil {
		wxc.SetFortune(*s)
	}
	return wxc
}

// SetID sets the "id" field.
func (wxc *WuXingCreate) SetID(s string) *WuXingCreate {
	wxc.mutation.SetID(s)
	return wxc
}

// Mutation returns the WuXingMutation object of the builder.
func (wxc *WuXingCreate) Mutation() *WuXingMutation {
	return wxc.mutation
}

// Save creates the WuXing in the database.
func (wxc *WuXingCreate) Save(ctx context.Context) (*WuXing, error) {
	return withHooks[*WuXing, WuXingMutation](ctx, wxc.sqlSave, wxc.mutation, wxc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (wxc *WuXingCreate) SaveX(ctx context.Context) *WuXing {
	v, err := wxc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wxc *WuXingCreate) Exec(ctx context.Context) error {
	_, err := wxc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wxc *WuXingCreate) ExecX(ctx context.Context) {
	if err := wxc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (wxc *WuXingCreate) check() error {
	return nil
}

func (wxc *WuXingCreate) sqlSave(ctx context.Context) (*WuXing, error) {
	if err := wxc.check(); err != nil {
		return nil, err
	}
	_node, _spec := wxc.createSpec()
	if err := sqlgraph.CreateNode(ctx, wxc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	if _spec.ID.Value != nil {
		if id, ok := _spec.ID.Value.(string); ok {
			_node.ID = id
		} else {
			return nil, fmt.Errorf("unexpected WuXing.ID type: %T", _spec.ID.Value)
		}
	}
	wxc.mutation.id = &_node.ID
	wxc.mutation.done = true
	return _node, nil
}

func (wxc *WuXingCreate) createSpec() (*WuXing, *sqlgraph.CreateSpec) {
	var (
		_node = &WuXing{config: wxc.config}
		_spec = sqlgraph.NewCreateSpec(wuxing.Table, sqlgraph.NewFieldSpec(wuxing.FieldID, field.TypeString))
	)
	if id, ok := wxc.mutation.ID(); ok {
		_node.ID = id
		_spec.ID.Value = id
	}
	if value, ok := wxc.mutation.Created(); ok {
		_spec.SetField(wuxing.FieldCreated, field.TypeTime, value)
		_node.Created = value
	}
	if value, ok := wxc.mutation.Updated(); ok {
		_spec.SetField(wuxing.FieldUpdated, field.TypeTime, value)
		_node.Updated = value
	}
	if value, ok := wxc.mutation.Deleted(); ok {
		_spec.SetField(wuxing.FieldDeleted, field.TypeTime, value)
		_node.Deleted = value
	}
	if value, ok := wxc.mutation.Version(); ok {
		_spec.SetField(wuxing.FieldVersion, field.TypeInt, value)
		_node.Version = value
	}
	if value, ok := wxc.mutation.First(); ok {
		_spec.SetField(wuxing.FieldFirst, field.TypeString, value)
		_node.First = value
	}
	if value, ok := wxc.mutation.Second(); ok {
		_spec.SetField(wuxing.FieldSecond, field.TypeString, value)
		_node.Second = value
	}
	if value, ok := wxc.mutation.Third(); ok {
		_spec.SetField(wuxing.FieldThird, field.TypeString, value)
		_node.Third = value
	}
	if value, ok := wxc.mutation.Fortune(); ok {
		_spec.SetField(wuxing.FieldFortune, field.TypeString, value)
		_node.Fortune = value
	}
	return _node, _spec
}

// WuXingCreateBulk is the builder for creating many WuXing entities in bulk.
type WuXingCreateBulk struct {
	config
	builders []*WuXingCreate
}

// Save creates the WuXing entities in the database.
func (wxcb *WuXingCreateBulk) Save(ctx context.Context) ([]*WuXing, error) {
	specs := make([]*sqlgraph.CreateSpec, len(wxcb.builders))
	nodes := make([]*WuXing, len(wxcb.builders))
	mutators := make([]Mutator, len(wxcb.builders))
	for i := range wxcb.builders {
		func(i int, root context.Context) {
			builder := wxcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*WuXingMutation)
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
					_, err = mutators[i+1].Mutate(root, wxcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, wxcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, wxcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (wxcb *WuXingCreateBulk) SaveX(ctx context.Context) []*WuXing {
	v, err := wxcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (wxcb *WuXingCreateBulk) Exec(ctx context.Context) error {
	_, err := wxcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (wxcb *WuXingCreateBulk) ExecX(ctx context.Context) {
	if err := wxcb.Exec(ctx); err != nil {
		panic(err)
	}
}
