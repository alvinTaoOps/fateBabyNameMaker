package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema"
	"entgo.io/ent/schema/field"
)

type Character struct {
	ent.Schema
}

func (Character) Fields() []ent.Field {
	return []ent.Field{
		field.String("id").StorageKey("hash"),
		field.String("pin_yin"),
		field.String("ch"),
		field.String("radical"),
		field.Int32("radical_stroke"),
		field.Int32("stroke"),
		field.Bool("is_kang_xi"),
		field.String("kang_xi"),
		field.Int32("kang_xi_stroke"),
		field.String("simple_radical"),
		field.Int32("simple_radical_stroke"),
		field.Int32("simple_total_stroke"),
		field.String("traditional_radical"),
		field.Int32("traditional_radical_stroke"),
		field.Int32("traditional_total_stroke"),
		field.Bool("name_science"),
		field.String("wu_xing"),
		field.String("lucky"),
		field.Bool("regular"),
		field.String("traditional_character"),
		field.String("variant_character"),
		field.String("comment"),
		field.Int32("science_stroke")}
}
func (Character) Edges() []ent.Edge {
	return nil
}
func (Character) Annotations() []schema.Annotation {
	return []schema.Annotation{
		entsql.Annotation{Table: "character"},
	}
}
