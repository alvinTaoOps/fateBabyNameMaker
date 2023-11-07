// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/babyname/fate/ent/ncharacter"
)

// NCharacter is the model entity for the NCharacter schema.
type NCharacter struct {
	config `json:"-"`
	// ID of the ent.
	ID int32 `json:"id,omitempty"`
	// PinYin holds the value of the "pin_yin" field.
	PinYin string `json:"pin_yin,omitempty"`
	// Ch holds the value of the "ch" field.
	Ch string `json:"ch,omitempty"`
	// ChStroke holds the value of the "ch_stroke" field.
	ChStroke int `json:"ch_stroke,omitempty"`
	// ChType holds the value of the "ch_type" field.
	ChType int `json:"ch_type,omitempty"`
	// Radical holds the value of the "radical" field.
	Radical string `json:"radical,omitempty"`
	// RadicalStroke holds the value of the "radical_stroke" field.
	RadicalStroke int `json:"radical_stroke,omitempty"`
	// Relate holds the value of the "relate" field.
	Relate int32 `json:"relate,omitempty"`
	// RelateKangXi holds the value of the "relate_kang_xi" field.
	RelateKangXi int32 `json:"relate_kang_xi,omitempty"`
	// RelateTraditional holds the value of the "relate_traditional" field.
	RelateTraditional int32 `json:"relate_traditional,omitempty"`
	// RelateVariant holds the value of the "relate_variant" field.
	RelateVariant []string `json:"relate_variant,omitempty"`
	// IsNameScience holds the value of the "is_name_science" field.
	IsNameScience bool `json:"is_name_science,omitempty"`
	// NameScienceChStroke holds the value of the "name_science_ch_stroke" field.
	NameScienceChStroke int `json:"name_science_ch_stroke,omitempty"`
	// IsRegular holds the value of the "is_regular" field.
	IsRegular bool `json:"is_regular,omitempty"`
	// WuXing holds the value of the "wu_xing" field.
	WuXing string `json:"wu_xing,omitempty"`
	// Lucky holds the value of the "lucky" field.
	Lucky string `json:"lucky,omitempty"`
	// Comment holds the value of the "comment" field.
	Comment      string `json:"comment,omitempty"`
	selectValues sql.SelectValues
}

// scanValues returns the types for scanning values from sql.Rows.
func (*NCharacter) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case ncharacter.FieldRelateVariant:
			values[i] = new([]byte)
		case ncharacter.FieldIsNameScience, ncharacter.FieldIsRegular:
			values[i] = new(sql.NullBool)
		case ncharacter.FieldID, ncharacter.FieldChStroke, ncharacter.FieldChType, ncharacter.FieldRadicalStroke, ncharacter.FieldRelate, ncharacter.FieldRelateKangXi, ncharacter.FieldRelateTraditional, ncharacter.FieldNameScienceChStroke:
			values[i] = new(sql.NullInt64)
		case ncharacter.FieldPinYin, ncharacter.FieldCh, ncharacter.FieldRadical, ncharacter.FieldWuXing, ncharacter.FieldLucky, ncharacter.FieldComment:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the NCharacter fields.
func (n *NCharacter) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case ncharacter.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			n.ID = int32(value.Int64)
		case ncharacter.FieldPinYin:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field pin_yin", values[i])
			} else if value.Valid {
				n.PinYin = value.String
			}
		case ncharacter.FieldCh:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ch", values[i])
			} else if value.Valid {
				n.Ch = value.String
			}
		case ncharacter.FieldChStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ch_stroke", values[i])
			} else if value.Valid {
				n.ChStroke = int(value.Int64)
			}
		case ncharacter.FieldChType:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ch_type", values[i])
			} else if value.Valid {
				n.ChType = int(value.Int64)
			}
		case ncharacter.FieldRadical:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field radical", values[i])
			} else if value.Valid {
				n.Radical = value.String
			}
		case ncharacter.FieldRadicalStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field radical_stroke", values[i])
			} else if value.Valid {
				n.RadicalStroke = int(value.Int64)
			}
		case ncharacter.FieldRelate:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relate", values[i])
			} else if value.Valid {
				n.Relate = int32(value.Int64)
			}
		case ncharacter.FieldRelateKangXi:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relate_kang_xi", values[i])
			} else if value.Valid {
				n.RelateKangXi = int32(value.Int64)
			}
		case ncharacter.FieldRelateTraditional:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field relate_traditional", values[i])
			} else if value.Valid {
				n.RelateTraditional = int32(value.Int64)
			}
		case ncharacter.FieldRelateVariant:
			if value, ok := values[i].(*[]byte); !ok {
				return fmt.Errorf("unexpected type %T for field relate_variant", values[i])
			} else if value != nil && len(*value) > 0 {
				if err := json.Unmarshal(*value, &n.RelateVariant); err != nil {
					return fmt.Errorf("unmarshal field relate_variant: %w", err)
				}
			}
		case ncharacter.FieldIsNameScience:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_name_science", values[i])
			} else if value.Valid {
				n.IsNameScience = value.Bool
			}
		case ncharacter.FieldNameScienceChStroke:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field name_science_ch_stroke", values[i])
			} else if value.Valid {
				n.NameScienceChStroke = int(value.Int64)
			}
		case ncharacter.FieldIsRegular:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field is_regular", values[i])
			} else if value.Valid {
				n.IsRegular = value.Bool
			}
		case ncharacter.FieldWuXing:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wu_xing", values[i])
			} else if value.Valid {
				n.WuXing = value.String
			}
		case ncharacter.FieldLucky:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field lucky", values[i])
			} else if value.Valid {
				n.Lucky = value.String
			}
		case ncharacter.FieldComment:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field comment", values[i])
			} else if value.Valid {
				n.Comment = value.String
			}
		default:
			n.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the NCharacter.
// This includes values selected through modifiers, order, etc.
func (n *NCharacter) Value(name string) (ent.Value, error) {
	return n.selectValues.Get(name)
}

// Update returns a builder for updating this NCharacter.
// Note that you need to call NCharacter.Unwrap() before calling this method if this NCharacter
// was returned from a transaction, and the transaction was committed or rolled back.
func (n *NCharacter) Update() *NCharacterUpdateOne {
	return NewNCharacterClient(n.config).UpdateOne(n)
}

// Unwrap unwraps the NCharacter entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (n *NCharacter) Unwrap() *NCharacter {
	_tx, ok := n.config.driver.(*txDriver)
	if !ok {
		panic("ent: NCharacter is not a transactional entity")
	}
	n.config.driver = _tx.drv
	return n
}

// String implements the fmt.Stringer.
func (n *NCharacter) String() string {
	var builder strings.Builder
	builder.WriteString("NCharacter(")
	builder.WriteString(fmt.Sprintf("id=%v, ", n.ID))
	builder.WriteString("pin_yin=")
	builder.WriteString(n.PinYin)
	builder.WriteString(", ")
	builder.WriteString("ch=")
	builder.WriteString(n.Ch)
	builder.WriteString(", ")
	builder.WriteString("ch_stroke=")
	builder.WriteString(fmt.Sprintf("%v", n.ChStroke))
	builder.WriteString(", ")
	builder.WriteString("ch_type=")
	builder.WriteString(fmt.Sprintf("%v", n.ChType))
	builder.WriteString(", ")
	builder.WriteString("radical=")
	builder.WriteString(n.Radical)
	builder.WriteString(", ")
	builder.WriteString("radical_stroke=")
	builder.WriteString(fmt.Sprintf("%v", n.RadicalStroke))
	builder.WriteString(", ")
	builder.WriteString("relate=")
	builder.WriteString(fmt.Sprintf("%v", n.Relate))
	builder.WriteString(", ")
	builder.WriteString("relate_kang_xi=")
	builder.WriteString(fmt.Sprintf("%v", n.RelateKangXi))
	builder.WriteString(", ")
	builder.WriteString("relate_traditional=")
	builder.WriteString(fmt.Sprintf("%v", n.RelateTraditional))
	builder.WriteString(", ")
	builder.WriteString("relate_variant=")
	builder.WriteString(fmt.Sprintf("%v", n.RelateVariant))
	builder.WriteString(", ")
	builder.WriteString("is_name_science=")
	builder.WriteString(fmt.Sprintf("%v", n.IsNameScience))
	builder.WriteString(", ")
	builder.WriteString("name_science_ch_stroke=")
	builder.WriteString(fmt.Sprintf("%v", n.NameScienceChStroke))
	builder.WriteString(", ")
	builder.WriteString("is_regular=")
	builder.WriteString(fmt.Sprintf("%v", n.IsRegular))
	builder.WriteString(", ")
	builder.WriteString("wu_xing=")
	builder.WriteString(n.WuXing)
	builder.WriteString(", ")
	builder.WriteString("lucky=")
	builder.WriteString(n.Lucky)
	builder.WriteString(", ")
	builder.WriteString("comment=")
	builder.WriteString(n.Comment)
	builder.WriteByte(')')
	return builder.String()
}

// NCharacters is a parsable slice of NCharacter.
type NCharacters []*NCharacter