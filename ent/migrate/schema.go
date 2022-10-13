// Code generated by entc, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// CharacterColumns holds the columns for the "character" table.
	CharacterColumns = []*schema.Column{
		{Name: "hash", Type: field.TypeString},
		{Name: "pin_yin", Type: field.TypeString},
		{Name: "ch", Type: field.TypeString},
		{Name: "radical", Type: field.TypeString},
		{Name: "radical_stroke", Type: field.TypeInt32},
		{Name: "stroke", Type: field.TypeInt32},
		{Name: "is_kang_xi", Type: field.TypeBool},
		{Name: "kang_xi", Type: field.TypeString},
		{Name: "kang_xi_stroke", Type: field.TypeInt32},
		{Name: "simple_radical", Type: field.TypeString},
		{Name: "simple_radical_stroke", Type: field.TypeInt32},
		{Name: "simple_total_stroke", Type: field.TypeInt32},
		{Name: "traditional_radical", Type: field.TypeString},
		{Name: "traditional_radical_stroke", Type: field.TypeInt32},
		{Name: "traditional_total_stroke", Type: field.TypeInt32},
		{Name: "name_science", Type: field.TypeBool},
		{Name: "wu_xing", Type: field.TypeString},
		{Name: "lucky", Type: field.TypeString},
		{Name: "regular", Type: field.TypeBool},
		{Name: "traditional_character", Type: field.TypeString},
		{Name: "variant_character", Type: field.TypeString},
		{Name: "comment", Type: field.TypeString},
		{Name: "science_stroke", Type: field.TypeInt32},
	}
	// CharacterTable holds the schema information for the "character" table.
	CharacterTable = &schema.Table{
		Name:       "character",
		Columns:    CharacterColumns,
		PrimaryKey: []*schema.Column{CharacterColumns[0]},
	}
	// VersionsColumns holds the columns for the "versions" table.
	VersionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "version", Type: field.TypeInt},
		{Name: "updated_unix", Type: field.TypeInt64},
	}
	// VersionsTable holds the schema information for the "versions" table.
	VersionsTable = &schema.Table{
		Name:       "versions",
		Columns:    VersionsColumns,
		PrimaryKey: []*schema.Column{VersionsColumns[0]},
	}
	// WuGeLuckyColumns holds the columns for the "wu_ge_lucky" table.
	WuGeLuckyColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "last_stroke_1", Type: field.TypeInt32, Nullable: true},
		{Name: "last_stroke_2", Type: field.TypeInt32, Nullable: true},
		{Name: "first_stroke_1", Type: field.TypeInt32, Nullable: true},
		{Name: "first_stroke_2", Type: field.TypeInt32, Nullable: true},
		{Name: "tian_ge", Type: field.TypeInt32, Nullable: true},
		{Name: "tian_da_yan", Type: field.TypeString, Nullable: true},
		{Name: "ren_ge", Type: field.TypeInt32, Nullable: true},
		{Name: "ren_da_yan", Type: field.TypeString, Nullable: true},
		{Name: "di_ge", Type: field.TypeInt32, Nullable: true},
		{Name: "di_da_yan", Type: field.TypeString, Nullable: true},
		{Name: "wai_ge", Type: field.TypeInt32, Nullable: true},
		{Name: "wai_da_yan", Type: field.TypeString, Nullable: true},
		{Name: "zong_ge", Type: field.TypeInt32, Nullable: true},
		{Name: "zong_da_yan", Type: field.TypeString, Nullable: true},
		{Name: "zong_lucky", Type: field.TypeBool, Nullable: true},
		{Name: "zong_sex", Type: field.TypeBool, Nullable: true},
		{Name: "zong_max", Type: field.TypeBool, Nullable: true},
	}
	// WuGeLuckyTable holds the schema information for the "wu_ge_lucky" table.
	WuGeLuckyTable = &schema.Table{
		Name:       "wu_ge_lucky",
		Columns:    WuGeLuckyColumns,
		PrimaryKey: []*schema.Column{WuGeLuckyColumns[0]},
	}
	// WuXingColumns holds the columns for the "wu_xing" table.
	WuXingColumns = []*schema.Column{
		{Name: "id", Type: field.TypeString},
		{Name: "created", Type: field.TypeTime, Nullable: true},
		{Name: "updated", Type: field.TypeTime, Nullable: true},
		{Name: "deleted", Type: field.TypeTime, Nullable: true},
		{Name: "version", Type: field.TypeInt32, Nullable: true},
		{Name: "first", Type: field.TypeString, Nullable: true},
		{Name: "second", Type: field.TypeString, Nullable: true},
		{Name: "third", Type: field.TypeString, Nullable: true},
		{Name: "fortune", Type: field.TypeString, Nullable: true},
	}
	// WuXingTable holds the schema information for the "wu_xing" table.
	WuXingTable = &schema.Table{
		Name:       "wu_xing",
		Columns:    WuXingColumns,
		PrimaryKey: []*schema.Column{WuXingColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		CharacterTable,
		VersionsTable,
		WuGeLuckyTable,
		WuXingTable,
	}
)

func init() {
	CharacterTable.Annotation = &entsql.Annotation{
		Table: "character",
	}
	WuGeLuckyTable.Annotation = &entsql.Annotation{
		Table: "wu_ge_lucky",
	}
	WuXingTable.Annotation = &entsql.Annotation{
		Table: "wu_xing",
	}
}
