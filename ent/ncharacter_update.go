// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/babyname/fate/ent/ncharacter"
	"github.com/babyname/fate/ent/predicate"
)

// NCharacterUpdate is the builder for updating NCharacter entities.
type NCharacterUpdate struct {
	config
	hooks    []Hook
	mutation *NCharacterMutation
}

// Where appends a list predicates to the NCharacterUpdate builder.
func (nu *NCharacterUpdate) Where(ps ...predicate.NCharacter) *NCharacterUpdate {
	nu.mutation.Where(ps...)
	return nu
}

// SetPinYin sets the "pin_yin" field.
func (nu *NCharacterUpdate) SetPinYin(s string) *NCharacterUpdate {
	nu.mutation.SetPinYin(s)
	return nu
}

// SetCh sets the "ch" field.
func (nu *NCharacterUpdate) SetCh(s string) *NCharacterUpdate {
	nu.mutation.SetCh(s)
	return nu
}

// SetChStroke sets the "ch_stroke" field.
func (nu *NCharacterUpdate) SetChStroke(i int) *NCharacterUpdate {
	nu.mutation.ResetChStroke()
	nu.mutation.SetChStroke(i)
	return nu
}

// AddChStroke adds i to the "ch_stroke" field.
func (nu *NCharacterUpdate) AddChStroke(i int) *NCharacterUpdate {
	nu.mutation.AddChStroke(i)
	return nu
}

// SetChType sets the "ch_type" field.
func (nu *NCharacterUpdate) SetChType(i int) *NCharacterUpdate {
	nu.mutation.ResetChType()
	nu.mutation.SetChType(i)
	return nu
}

// SetNillableChType sets the "ch_type" field if the given value is not nil.
func (nu *NCharacterUpdate) SetNillableChType(i *int) *NCharacterUpdate {
	if i != nil {
		nu.SetChType(*i)
	}
	return nu
}

// AddChType adds i to the "ch_type" field.
func (nu *NCharacterUpdate) AddChType(i int) *NCharacterUpdate {
	nu.mutation.AddChType(i)
	return nu
}

// SetRadical sets the "radical" field.
func (nu *NCharacterUpdate) SetRadical(s string) *NCharacterUpdate {
	nu.mutation.SetRadical(s)
	return nu
}

// SetRadicalStroke sets the "radical_stroke" field.
func (nu *NCharacterUpdate) SetRadicalStroke(i int) *NCharacterUpdate {
	nu.mutation.ResetRadicalStroke()
	nu.mutation.SetRadicalStroke(i)
	return nu
}

// AddRadicalStroke adds i to the "radical_stroke" field.
func (nu *NCharacterUpdate) AddRadicalStroke(i int) *NCharacterUpdate {
	nu.mutation.AddRadicalStroke(i)
	return nu
}

// SetRelate sets the "relate" field.
func (nu *NCharacterUpdate) SetRelate(i int32) *NCharacterUpdate {
	nu.mutation.ResetRelate()
	nu.mutation.SetRelate(i)
	return nu
}

// AddRelate adds i to the "relate" field.
func (nu *NCharacterUpdate) AddRelate(i int32) *NCharacterUpdate {
	nu.mutation.AddRelate(i)
	return nu
}

// SetRelateKangXi sets the "relate_kang_xi" field.
func (nu *NCharacterUpdate) SetRelateKangXi(i int32) *NCharacterUpdate {
	nu.mutation.ResetRelateKangXi()
	nu.mutation.SetRelateKangXi(i)
	return nu
}

// AddRelateKangXi adds i to the "relate_kang_xi" field.
func (nu *NCharacterUpdate) AddRelateKangXi(i int32) *NCharacterUpdate {
	nu.mutation.AddRelateKangXi(i)
	return nu
}

// SetRelateTraditional sets the "relate_traditional" field.
func (nu *NCharacterUpdate) SetRelateTraditional(i int32) *NCharacterUpdate {
	nu.mutation.ResetRelateTraditional()
	nu.mutation.SetRelateTraditional(i)
	return nu
}

// AddRelateTraditional adds i to the "relate_traditional" field.
func (nu *NCharacterUpdate) AddRelateTraditional(i int32) *NCharacterUpdate {
	nu.mutation.AddRelateTraditional(i)
	return nu
}

// SetRelateVariant sets the "relate_variant" field.
func (nu *NCharacterUpdate) SetRelateVariant(s []string) *NCharacterUpdate {
	nu.mutation.SetRelateVariant(s)
	return nu
}

// AppendRelateVariant appends s to the "relate_variant" field.
func (nu *NCharacterUpdate) AppendRelateVariant(s []string) *NCharacterUpdate {
	nu.mutation.AppendRelateVariant(s)
	return nu
}

// SetIsNameScience sets the "is_name_science" field.
func (nu *NCharacterUpdate) SetIsNameScience(b bool) *NCharacterUpdate {
	nu.mutation.SetIsNameScience(b)
	return nu
}

// SetNameScienceChStroke sets the "name_science_ch_stroke" field.
func (nu *NCharacterUpdate) SetNameScienceChStroke(i int) *NCharacterUpdate {
	nu.mutation.ResetNameScienceChStroke()
	nu.mutation.SetNameScienceChStroke(i)
	return nu
}

// AddNameScienceChStroke adds i to the "name_science_ch_stroke" field.
func (nu *NCharacterUpdate) AddNameScienceChStroke(i int) *NCharacterUpdate {
	nu.mutation.AddNameScienceChStroke(i)
	return nu
}

// SetIsRegular sets the "is_regular" field.
func (nu *NCharacterUpdate) SetIsRegular(b bool) *NCharacterUpdate {
	nu.mutation.SetIsRegular(b)
	return nu
}

// SetWuXing sets the "wu_xing" field.
func (nu *NCharacterUpdate) SetWuXing(s string) *NCharacterUpdate {
	nu.mutation.SetWuXing(s)
	return nu
}

// SetLucky sets the "lucky" field.
func (nu *NCharacterUpdate) SetLucky(s string) *NCharacterUpdate {
	nu.mutation.SetLucky(s)
	return nu
}

// SetComment sets the "comment" field.
func (nu *NCharacterUpdate) SetComment(s string) *NCharacterUpdate {
	nu.mutation.SetComment(s)
	return nu
}

// Mutation returns the NCharacterMutation object of the builder.
func (nu *NCharacterUpdate) Mutation() *NCharacterMutation {
	return nu.mutation
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (nu *NCharacterUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, nu.sqlSave, nu.mutation, nu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nu *NCharacterUpdate) SaveX(ctx context.Context) int {
	affected, err := nu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (nu *NCharacterUpdate) Exec(ctx context.Context) error {
	_, err := nu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nu *NCharacterUpdate) ExecX(ctx context.Context) {
	if err := nu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nu *NCharacterUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(ncharacter.Table, ncharacter.Columns, sqlgraph.NewFieldSpec(ncharacter.FieldID, field.TypeInt32))
	if ps := nu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nu.mutation.PinYin(); ok {
		_spec.SetField(ncharacter.FieldPinYin, field.TypeString, value)
	}
	if value, ok := nu.mutation.Ch(); ok {
		_spec.SetField(ncharacter.FieldCh, field.TypeString, value)
	}
	if value, ok := nu.mutation.ChStroke(); ok {
		_spec.SetField(ncharacter.FieldChStroke, field.TypeInt, value)
	}
	if value, ok := nu.mutation.AddedChStroke(); ok {
		_spec.AddField(ncharacter.FieldChStroke, field.TypeInt, value)
	}
	if value, ok := nu.mutation.ChType(); ok {
		_spec.SetField(ncharacter.FieldChType, field.TypeInt, value)
	}
	if value, ok := nu.mutation.AddedChType(); ok {
		_spec.AddField(ncharacter.FieldChType, field.TypeInt, value)
	}
	if value, ok := nu.mutation.Radical(); ok {
		_spec.SetField(ncharacter.FieldRadical, field.TypeString, value)
	}
	if value, ok := nu.mutation.RadicalStroke(); ok {
		_spec.SetField(ncharacter.FieldRadicalStroke, field.TypeInt, value)
	}
	if value, ok := nu.mutation.AddedRadicalStroke(); ok {
		_spec.AddField(ncharacter.FieldRadicalStroke, field.TypeInt, value)
	}
	if value, ok := nu.mutation.Relate(); ok {
		_spec.SetField(ncharacter.FieldRelate, field.TypeInt32, value)
	}
	if value, ok := nu.mutation.AddedRelate(); ok {
		_spec.AddField(ncharacter.FieldRelate, field.TypeInt32, value)
	}
	if value, ok := nu.mutation.RelateKangXi(); ok {
		_spec.SetField(ncharacter.FieldRelateKangXi, field.TypeInt32, value)
	}
	if value, ok := nu.mutation.AddedRelateKangXi(); ok {
		_spec.AddField(ncharacter.FieldRelateKangXi, field.TypeInt32, value)
	}
	if value, ok := nu.mutation.RelateTraditional(); ok {
		_spec.SetField(ncharacter.FieldRelateTraditional, field.TypeInt32, value)
	}
	if value, ok := nu.mutation.AddedRelateTraditional(); ok {
		_spec.AddField(ncharacter.FieldRelateTraditional, field.TypeInt32, value)
	}
	if value, ok := nu.mutation.RelateVariant(); ok {
		_spec.SetField(ncharacter.FieldRelateVariant, field.TypeJSON, value)
	}
	if value, ok := nu.mutation.AppendedRelateVariant(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, ncharacter.FieldRelateVariant, value)
		})
	}
	if value, ok := nu.mutation.IsNameScience(); ok {
		_spec.SetField(ncharacter.FieldIsNameScience, field.TypeBool, value)
	}
	if value, ok := nu.mutation.NameScienceChStroke(); ok {
		_spec.SetField(ncharacter.FieldNameScienceChStroke, field.TypeInt, value)
	}
	if value, ok := nu.mutation.AddedNameScienceChStroke(); ok {
		_spec.AddField(ncharacter.FieldNameScienceChStroke, field.TypeInt, value)
	}
	if value, ok := nu.mutation.IsRegular(); ok {
		_spec.SetField(ncharacter.FieldIsRegular, field.TypeBool, value)
	}
	if value, ok := nu.mutation.WuXing(); ok {
		_spec.SetField(ncharacter.FieldWuXing, field.TypeString, value)
	}
	if value, ok := nu.mutation.Lucky(); ok {
		_spec.SetField(ncharacter.FieldLucky, field.TypeString, value)
	}
	if value, ok := nu.mutation.Comment(); ok {
		_spec.SetField(ncharacter.FieldComment, field.TypeString, value)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, nu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ncharacter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	nu.mutation.done = true
	return n, nil
}

// NCharacterUpdateOne is the builder for updating a single NCharacter entity.
type NCharacterUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *NCharacterMutation
}

// SetPinYin sets the "pin_yin" field.
func (nuo *NCharacterUpdateOne) SetPinYin(s string) *NCharacterUpdateOne {
	nuo.mutation.SetPinYin(s)
	return nuo
}

// SetCh sets the "ch" field.
func (nuo *NCharacterUpdateOne) SetCh(s string) *NCharacterUpdateOne {
	nuo.mutation.SetCh(s)
	return nuo
}

// SetChStroke sets the "ch_stroke" field.
func (nuo *NCharacterUpdateOne) SetChStroke(i int) *NCharacterUpdateOne {
	nuo.mutation.ResetChStroke()
	nuo.mutation.SetChStroke(i)
	return nuo
}

// AddChStroke adds i to the "ch_stroke" field.
func (nuo *NCharacterUpdateOne) AddChStroke(i int) *NCharacterUpdateOne {
	nuo.mutation.AddChStroke(i)
	return nuo
}

// SetChType sets the "ch_type" field.
func (nuo *NCharacterUpdateOne) SetChType(i int) *NCharacterUpdateOne {
	nuo.mutation.ResetChType()
	nuo.mutation.SetChType(i)
	return nuo
}

// SetNillableChType sets the "ch_type" field if the given value is not nil.
func (nuo *NCharacterUpdateOne) SetNillableChType(i *int) *NCharacterUpdateOne {
	if i != nil {
		nuo.SetChType(*i)
	}
	return nuo
}

// AddChType adds i to the "ch_type" field.
func (nuo *NCharacterUpdateOne) AddChType(i int) *NCharacterUpdateOne {
	nuo.mutation.AddChType(i)
	return nuo
}

// SetRadical sets the "radical" field.
func (nuo *NCharacterUpdateOne) SetRadical(s string) *NCharacterUpdateOne {
	nuo.mutation.SetRadical(s)
	return nuo
}

// SetRadicalStroke sets the "radical_stroke" field.
func (nuo *NCharacterUpdateOne) SetRadicalStroke(i int) *NCharacterUpdateOne {
	nuo.mutation.ResetRadicalStroke()
	nuo.mutation.SetRadicalStroke(i)
	return nuo
}

// AddRadicalStroke adds i to the "radical_stroke" field.
func (nuo *NCharacterUpdateOne) AddRadicalStroke(i int) *NCharacterUpdateOne {
	nuo.mutation.AddRadicalStroke(i)
	return nuo
}

// SetRelate sets the "relate" field.
func (nuo *NCharacterUpdateOne) SetRelate(i int32) *NCharacterUpdateOne {
	nuo.mutation.ResetRelate()
	nuo.mutation.SetRelate(i)
	return nuo
}

// AddRelate adds i to the "relate" field.
func (nuo *NCharacterUpdateOne) AddRelate(i int32) *NCharacterUpdateOne {
	nuo.mutation.AddRelate(i)
	return nuo
}

// SetRelateKangXi sets the "relate_kang_xi" field.
func (nuo *NCharacterUpdateOne) SetRelateKangXi(i int32) *NCharacterUpdateOne {
	nuo.mutation.ResetRelateKangXi()
	nuo.mutation.SetRelateKangXi(i)
	return nuo
}

// AddRelateKangXi adds i to the "relate_kang_xi" field.
func (nuo *NCharacterUpdateOne) AddRelateKangXi(i int32) *NCharacterUpdateOne {
	nuo.mutation.AddRelateKangXi(i)
	return nuo
}

// SetRelateTraditional sets the "relate_traditional" field.
func (nuo *NCharacterUpdateOne) SetRelateTraditional(i int32) *NCharacterUpdateOne {
	nuo.mutation.ResetRelateTraditional()
	nuo.mutation.SetRelateTraditional(i)
	return nuo
}

// AddRelateTraditional adds i to the "relate_traditional" field.
func (nuo *NCharacterUpdateOne) AddRelateTraditional(i int32) *NCharacterUpdateOne {
	nuo.mutation.AddRelateTraditional(i)
	return nuo
}

// SetRelateVariant sets the "relate_variant" field.
func (nuo *NCharacterUpdateOne) SetRelateVariant(s []string) *NCharacterUpdateOne {
	nuo.mutation.SetRelateVariant(s)
	return nuo
}

// AppendRelateVariant appends s to the "relate_variant" field.
func (nuo *NCharacterUpdateOne) AppendRelateVariant(s []string) *NCharacterUpdateOne {
	nuo.mutation.AppendRelateVariant(s)
	return nuo
}

// SetIsNameScience sets the "is_name_science" field.
func (nuo *NCharacterUpdateOne) SetIsNameScience(b bool) *NCharacterUpdateOne {
	nuo.mutation.SetIsNameScience(b)
	return nuo
}

// SetNameScienceChStroke sets the "name_science_ch_stroke" field.
func (nuo *NCharacterUpdateOne) SetNameScienceChStroke(i int) *NCharacterUpdateOne {
	nuo.mutation.ResetNameScienceChStroke()
	nuo.mutation.SetNameScienceChStroke(i)
	return nuo
}

// AddNameScienceChStroke adds i to the "name_science_ch_stroke" field.
func (nuo *NCharacterUpdateOne) AddNameScienceChStroke(i int) *NCharacterUpdateOne {
	nuo.mutation.AddNameScienceChStroke(i)
	return nuo
}

// SetIsRegular sets the "is_regular" field.
func (nuo *NCharacterUpdateOne) SetIsRegular(b bool) *NCharacterUpdateOne {
	nuo.mutation.SetIsRegular(b)
	return nuo
}

// SetWuXing sets the "wu_xing" field.
func (nuo *NCharacterUpdateOne) SetWuXing(s string) *NCharacterUpdateOne {
	nuo.mutation.SetWuXing(s)
	return nuo
}

// SetLucky sets the "lucky" field.
func (nuo *NCharacterUpdateOne) SetLucky(s string) *NCharacterUpdateOne {
	nuo.mutation.SetLucky(s)
	return nuo
}

// SetComment sets the "comment" field.
func (nuo *NCharacterUpdateOne) SetComment(s string) *NCharacterUpdateOne {
	nuo.mutation.SetComment(s)
	return nuo
}

// Mutation returns the NCharacterMutation object of the builder.
func (nuo *NCharacterUpdateOne) Mutation() *NCharacterMutation {
	return nuo.mutation
}

// Where appends a list predicates to the NCharacterUpdate builder.
func (nuo *NCharacterUpdateOne) Where(ps ...predicate.NCharacter) *NCharacterUpdateOne {
	nuo.mutation.Where(ps...)
	return nuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (nuo *NCharacterUpdateOne) Select(field string, fields ...string) *NCharacterUpdateOne {
	nuo.fields = append([]string{field}, fields...)
	return nuo
}

// Save executes the query and returns the updated NCharacter entity.
func (nuo *NCharacterUpdateOne) Save(ctx context.Context) (*NCharacter, error) {
	return withHooks(ctx, nuo.sqlSave, nuo.mutation, nuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (nuo *NCharacterUpdateOne) SaveX(ctx context.Context) *NCharacter {
	node, err := nuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (nuo *NCharacterUpdateOne) Exec(ctx context.Context) error {
	_, err := nuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (nuo *NCharacterUpdateOne) ExecX(ctx context.Context) {
	if err := nuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (nuo *NCharacterUpdateOne) sqlSave(ctx context.Context) (_node *NCharacter, err error) {
	_spec := sqlgraph.NewUpdateSpec(ncharacter.Table, ncharacter.Columns, sqlgraph.NewFieldSpec(ncharacter.FieldID, field.TypeInt32))
	id, ok := nuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "NCharacter.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := nuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, ncharacter.FieldID)
		for _, f := range fields {
			if !ncharacter.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != ncharacter.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := nuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := nuo.mutation.PinYin(); ok {
		_spec.SetField(ncharacter.FieldPinYin, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Ch(); ok {
		_spec.SetField(ncharacter.FieldCh, field.TypeString, value)
	}
	if value, ok := nuo.mutation.ChStroke(); ok {
		_spec.SetField(ncharacter.FieldChStroke, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.AddedChStroke(); ok {
		_spec.AddField(ncharacter.FieldChStroke, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.ChType(); ok {
		_spec.SetField(ncharacter.FieldChType, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.AddedChType(); ok {
		_spec.AddField(ncharacter.FieldChType, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.Radical(); ok {
		_spec.SetField(ncharacter.FieldRadical, field.TypeString, value)
	}
	if value, ok := nuo.mutation.RadicalStroke(); ok {
		_spec.SetField(ncharacter.FieldRadicalStroke, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.AddedRadicalStroke(); ok {
		_spec.AddField(ncharacter.FieldRadicalStroke, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.Relate(); ok {
		_spec.SetField(ncharacter.FieldRelate, field.TypeInt32, value)
	}
	if value, ok := nuo.mutation.AddedRelate(); ok {
		_spec.AddField(ncharacter.FieldRelate, field.TypeInt32, value)
	}
	if value, ok := nuo.mutation.RelateKangXi(); ok {
		_spec.SetField(ncharacter.FieldRelateKangXi, field.TypeInt32, value)
	}
	if value, ok := nuo.mutation.AddedRelateKangXi(); ok {
		_spec.AddField(ncharacter.FieldRelateKangXi, field.TypeInt32, value)
	}
	if value, ok := nuo.mutation.RelateTraditional(); ok {
		_spec.SetField(ncharacter.FieldRelateTraditional, field.TypeInt32, value)
	}
	if value, ok := nuo.mutation.AddedRelateTraditional(); ok {
		_spec.AddField(ncharacter.FieldRelateTraditional, field.TypeInt32, value)
	}
	if value, ok := nuo.mutation.RelateVariant(); ok {
		_spec.SetField(ncharacter.FieldRelateVariant, field.TypeJSON, value)
	}
	if value, ok := nuo.mutation.AppendedRelateVariant(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, ncharacter.FieldRelateVariant, value)
		})
	}
	if value, ok := nuo.mutation.IsNameScience(); ok {
		_spec.SetField(ncharacter.FieldIsNameScience, field.TypeBool, value)
	}
	if value, ok := nuo.mutation.NameScienceChStroke(); ok {
		_spec.SetField(ncharacter.FieldNameScienceChStroke, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.AddedNameScienceChStroke(); ok {
		_spec.AddField(ncharacter.FieldNameScienceChStroke, field.TypeInt, value)
	}
	if value, ok := nuo.mutation.IsRegular(); ok {
		_spec.SetField(ncharacter.FieldIsRegular, field.TypeBool, value)
	}
	if value, ok := nuo.mutation.WuXing(); ok {
		_spec.SetField(ncharacter.FieldWuXing, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Lucky(); ok {
		_spec.SetField(ncharacter.FieldLucky, field.TypeString, value)
	}
	if value, ok := nuo.mutation.Comment(); ok {
		_spec.SetField(ncharacter.FieldComment, field.TypeString, value)
	}
	_node = &NCharacter{config: nuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, nuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{ncharacter.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	nuo.mutation.done = true
	return _node, nil
}
