package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type WuGeLucky struct {
	ent.Schema
}

func (WuGeLucky) Fields() []ent.Field {
	return []ent.Field{
		field.String("id"),
		field.Int32("last_stroke_1").Optional(),
		field.Int32("last_stroke_2").Optional(),
		field.Int32("first_stroke_1").Optional(),
		field.Int32("first_stroke_2").Optional(),
		field.Int32("tian_ge").Optional(),
		field.String("tian_da_yan").Optional(),
		field.Int32("ren_ge").Optional(),
		field.String("ren_da_yan").Optional(),
		field.Int32("di_ge").Optional(),
		field.String("di_da_yan").Optional(),
		field.Int32("wai_ge").Optional(),
		field.String("wai_da_yan").Optional(),
		field.Int32("zong_ge").Optional(),
		field.String("zong_da_yan").Optional(),
		field.Bool("zong_lucky").Optional(),
		field.Bool("zong_sex").Optional(),
		field.Bool("zong_max").Optional()}
}
func (WuGeLucky) Edges() []ent.Edge {
	return nil
}
func (WuGeLucky) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "wu_ge_lucky"},
	}
}
