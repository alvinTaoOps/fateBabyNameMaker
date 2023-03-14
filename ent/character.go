// Code generated by ent, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/babyname/fate/ent/character"
)

// Character is the model entity for the Character schema.
type Character struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// PinYin holds the value of the "pin_yin" field.
	PinYin string `json:"pin_yin,omitempty"`
	// Ch holds the value of the "ch" field.
	Ch string `json:"ch,omitempty"`
	// Radical holds the value of the "radical" field.
	Radical string `json:"radical,omitempty"`
	// RadicalStroke holds the value of the "radical_stroke" field.
	RadicalStroke int `json:"radical_stroke,omitempty"`
	// Stroke holds the value of the "stroke" field.
	Stroke int `json:"stroke,omitempty"`
	// IsKangXi holds the value of the "is_kang_xi" field.
	IsKangXi bool `json:"is_kang_xi,omitempty"`
	// KangXi holds the value of the "kang_xi" field.
	KangXi string `json:"kang_xi,omitempty"`
	// KangXiStroke holds the value of the "kang_xi_stroke" field.
	KangXiStroke int `json:"kang_xi_stroke,omitempty"`
	// SimpleRadical holds the value of the "simple_radical" field.
	SimpleRadical string `json:"simple_radical,omitempty"`
	// SimpleRadicalStroke holds the value of the "simple_radical_stroke" field.
	SimpleRadicalStroke int `json:"simple_radical_stroke,omitempty"`
	// SimpleTotalStroke holds the value of the "simple_total_stroke" field.
	SimpleTotalStroke int `json:"simple_total_stroke,omitempty"`
	// TraditionalRadical holds the value of the "traditional_radical" field.
	TraditionalRadical string `json:"traditional_radical,omitempty"`
	// TraditionalRadicalStroke holds the value of the "traditional_radical_stroke" field.
	TraditionalRadicalStroke int `json:"traditional_radical_stroke,omitempty"`
	// TraditionalTotalStroke holds the value of the "traditional_total_stroke" field.
	TraditionalTotalStroke int `json:"traditional_total_stroke,omitempty"`
	// NameScience holds the value of the "name_science" field.
	NameScience bool `json:"name_science,omitempty"`
	// WuXing holds the value of the "wu_xing" field.
	WuXing string `json:"wu_xing,omitempty"`
	// Lucky holds the value of the "lucky" field.
	Lucky string `json:"lucky,omitempty"`
	// Regular holds the value of the "regular" field.
	Regular bool `json:"regular,omitempty"`
	// TraditionalCharacter holds the value of the "traditional_character" field.
	TraditionalCharacter string `json:"traditional_character,omitempty"`
	// VariantCharacter holds the value of the "variant_character" field.
	VariantCharacter string `json:"variant_character,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment string `json:"comment,omitempty"`
	// ScienceStroke holds the value of the "science_stroke" field.
	ScienceStroke int `json:"science_stroke,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Character) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case character.FieldIsKangXi, character.FieldNameScience, character.FieldRegular:
			values[i] = new(sql.NullBool)
		case character.FieldRadicalStroke, character.FieldStroke, character.FieldKangXiStroke, character.FieldSimpleRadicalStroke, character.FieldSimpleTotalStroke, character.FieldTraditionalRadicalStroke, character.FieldTraditionalTotalStroke, character.FieldScienceStroke:
			values[i] = new(sql.NullInt64)
		case character.FieldID, character.FieldPinYin, character.FieldCh, character.FieldRadical, character.FieldKangXi, character.FieldSimpleRadical, character.FieldTraditionalRadical, character.FieldWuXing, character.FieldLucky, character.FieldTraditionalCharacter, character.FieldVariantCharacter, character.FieldComment:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Character", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Character fields.
func (c *Character) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case character.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				c.ID = value.String
			}
		case character.FieldPinYin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pin_yin", values[i])
			} else if value.Valid {
				c.PinYin = value.String
			}
		case character.FieldCh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ch", values[i])
			} else if value.Valid {
				c.Ch = value.String
			}
		case character.FieldRadical:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field radical", values[i])
			} else if value.Valid {
				c.Radical = value.String
			}
		case character.FieldRadicalStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field radical_stroke", values[i])
			} else if value.Valid {
				c.RadicalStroke = int(value.Int64)
			}
		case character.FieldStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field stroke", values[i])
			} else if value.Valid {
				c.Stroke = int(value.Int64)
			}
		case character.FieldIsKangXi:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_kang_xi", values[i])
			} else if value.Valid {
				c.IsKangXi = value.Bool
			}
		case character.FieldKangXi:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field kang_xi", values[i])
			} else if value.Valid {
				c.KangXi = value.String
			}
		case character.FieldKangXiStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field kang_xi_stroke", values[i])
			} else if value.Valid {
				c.KangXiStroke = int(value.Int64)
			}
		case character.FieldSimpleRadical:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field simple_radical", values[i])
			} else if value.Valid {
				c.SimpleRadical = value.String
			}
		case character.FieldSimpleRadicalStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field simple_radical_stroke", values[i])
			} else if value.Valid {
				c.SimpleRadicalStroke = int(value.Int64)
			}
		case character.FieldSimpleTotalStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field simple_total_stroke", values[i])
			} else if value.Valid {
				c.SimpleTotalStroke = int(value.Int64)
			}
		case character.FieldTraditionalRadical:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field traditional_radical", values[i])
			} else if value.Valid {
				c.TraditionalRadical = value.String
			}
		case character.FieldTraditionalRadicalStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field traditional_radical_stroke", values[i])
			} else if value.Valid {
				c.TraditionalRadicalStroke = int(value.Int64)
			}
		case character.FieldTraditionalTotalStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field traditional_total_stroke", values[i])
			} else if value.Valid {
				c.TraditionalTotalStroke = int(value.Int64)
			}
		case character.FieldNameScience:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field name_science", values[i])
			} else if value.Valid {
				c.NameScience = value.Bool
			}
		case character.FieldWuXing:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wu_xing", values[i])
			} else if value.Valid {
				c.WuXing = value.String
			}
		case character.FieldLucky:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lucky", values[i])
			} else if value.Valid {
				c.Lucky = value.String
			}
		case character.FieldRegular:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field regular", values[i])
			} else if value.Valid {
				c.Regular = value.Bool
			}
		case character.FieldTraditionalCharacter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field traditional_character", values[i])
			} else if value.Valid {
				c.TraditionalCharacter = value.String
			}
		case character.FieldVariantCharacter:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field variant_character", values[i])
			} else if value.Valid {
				c.VariantCharacter = value.String
			}
		case character.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				c.Comment = value.String
			}
		case character.FieldScienceStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field science_stroke", values[i])
			} else if value.Valid {
				c.ScienceStroke = int(value.Int64)
			}
		}
	}
	return nil
}

// Update returns a builder for updating this Character.
// Note that you need to call Character.Unwrap() before calling this method if this Character
// was returned from a transaction, and the transaction was committed or rolled back.
func (c *Character) Update() *CharacterUpdateOne {
	return NewCharacterClient(c.config).UpdateOne(c)
}

// Unwrap unwraps the Character entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (c *Character) Unwrap() *Character {
	_tx, ok := c.config.driver.(*txDriver)
	if !ok {
		panic("ent: Character is not a transactional entity")
	}
	c.config.driver = _tx.drv
	return c
}

// String implements the fmt.Stringer.
func (c *Character) String() string {
	var builder strings.Builder
	builder.WriteString("Character(")
	builder.WriteString(fmt.Sprintf("id=%v, ", c.ID))
	builder.WriteString("pin_yin=")
	builder.WriteString(c.PinYin)
	builder.WriteString(", ")
	builder.WriteString("ch=")
	builder.WriteString(c.Ch)
	builder.WriteString(", ")
	builder.WriteString("radical=")
	builder.WriteString(c.Radical)
	builder.WriteString(", ")
	builder.WriteString("radical_stroke=")
	builder.WriteString(fmt.Sprintf("%v", c.RadicalStroke))
	builder.WriteString(", ")
	builder.WriteString("stroke=")
	builder.WriteString(fmt.Sprintf("%v", c.Stroke))
	builder.WriteString(", ")
	builder.WriteString("is_kang_xi=")
	builder.WriteString(fmt.Sprintf("%v", c.IsKangXi))
	builder.WriteString(", ")
	builder.WriteString("kang_xi=")
	builder.WriteString(c.KangXi)
	builder.WriteString(", ")
	builder.WriteString("kang_xi_stroke=")
	builder.WriteString(fmt.Sprintf("%v", c.KangXiStroke))
	builder.WriteString(", ")
	builder.WriteString("simple_radical=")
	builder.WriteString(c.SimpleRadical)
	builder.WriteString(", ")
	builder.WriteString("simple_radical_stroke=")
	builder.WriteString(fmt.Sprintf("%v", c.SimpleRadicalStroke))
	builder.WriteString(", ")
	builder.WriteString("simple_total_stroke=")
	builder.WriteString(fmt.Sprintf("%v", c.SimpleTotalStroke))
	builder.WriteString(", ")
	builder.WriteString("traditional_radical=")
	builder.WriteString(c.TraditionalRadical)
	builder.WriteString(", ")
	builder.WriteString("traditional_radical_stroke=")
	builder.WriteString(fmt.Sprintf("%v", c.TraditionalRadicalStroke))
	builder.WriteString(", ")
	builder.WriteString("traditional_total_stroke=")
	builder.WriteString(fmt.Sprintf("%v", c.TraditionalTotalStroke))
	builder.WriteString(", ")
	builder.WriteString("name_science=")
	builder.WriteString(fmt.Sprintf("%v", c.NameScience))
	builder.WriteString(", ")
	builder.WriteString("wu_xing=")
	builder.WriteString(c.WuXing)
	builder.WriteString(", ")
	builder.WriteString("lucky=")
	builder.WriteString(c.Lucky)
	builder.WriteString(", ")
	builder.WriteString("regular=")
	builder.WriteString(fmt.Sprintf("%v", c.Regular))
	builder.WriteString(", ")
	builder.WriteString("traditional_character=")
	builder.WriteString(c.TraditionalCharacter)
	builder.WriteString(", ")
	builder.WriteString("variant_character=")
	builder.WriteString(c.VariantCharacter)
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(c.Comment)
	builder.WriteString(", ")
	builder.WriteString("science_stroke=")
	builder.WriteString(fmt.Sprintf("%v", c.ScienceStroke))
	builder.WriteByte(')')
	return builder.String()
}

// Characters is a parsable slice of Character.
type Characters []*Character
