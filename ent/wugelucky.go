// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/babyname/fate/ent/wugelucky"
)

// WuGeLucky is the model entity for the WuGeLucky schema.
type WuGeLucky struct {
	config `json:"-"`
	// ID of the ent.
	ID string `json:"id,omitempty"`
	// LastStroke1 holds the value of the "last_stroke_1" field.
	LastStroke1 int32 `json:"last_stroke_1,omitempty"`
	// LastStroke2 holds the value of the "last_stroke_2" field.
	LastStroke2 int32 `json:"last_stroke_2,omitempty"`
	// FirstStroke1 holds the value of the "first_stroke_1" field.
	FirstStroke1 int32 `json:"first_stroke_1,omitempty"`
	// FirstStroke2 holds the value of the "first_stroke_2" field.
	FirstStroke2 int32 `json:"first_stroke_2,omitempty"`
	// TianGe holds the value of the "tian_ge" field.
	TianGe int32 `json:"tian_ge,omitempty"`
	// TianDaYan holds the value of the "tian_da_yan" field.
	TianDaYan string `json:"tian_da_yan,omitempty"`
	// RenGe holds the value of the "ren_ge" field.
	RenGe int32 `json:"ren_ge,omitempty"`
	// RenDaYan holds the value of the "ren_da_yan" field.
	RenDaYan string `json:"ren_da_yan,omitempty"`
	// DiGe holds the value of the "di_ge" field.
	DiGe int32 `json:"di_ge,omitempty"`
	// DiDaYan holds the value of the "di_da_yan" field.
	DiDaYan string `json:"di_da_yan,omitempty"`
	// WaiGe holds the value of the "wai_ge" field.
	WaiGe int32 `json:"wai_ge,omitempty"`
	// WaiDaYan holds the value of the "wai_da_yan" field.
	WaiDaYan string `json:"wai_da_yan,omitempty"`
	// ZongGe holds the value of the "zong_ge" field.
	ZongGe int32 `json:"zong_ge,omitempty"`
	// ZongDaYan holds the value of the "zong_da_yan" field.
	ZongDaYan string `json:"zong_da_yan,omitempty"`
	// ZongLucky holds the value of the "zong_lucky" field.
	ZongLucky bool `json:"zong_lucky,omitempty"`
	// ZongSex holds the value of the "zong_sex" field.
	ZongSex bool `json:"zong_sex,omitempty"`
	// ZongMax holds the value of the "zong_max" field.
	ZongMax bool `json:"zong_max,omitempty"`
}

// scanValues returns the types for scanning values from sql.Rows.
func (*WuGeLucky) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case wugelucky.FieldZongLucky, wugelucky.FieldZongSex, wugelucky.FieldZongMax:
			values[i] = new(sql.NullBool)
		case wugelucky.FieldLastStroke1, wugelucky.FieldLastStroke2, wugelucky.FieldFirstStroke1, wugelucky.FieldFirstStroke2, wugelucky.FieldTianGe, wugelucky.FieldRenGe, wugelucky.FieldDiGe, wugelucky.FieldWaiGe, wugelucky.FieldZongGe:
			values[i] = new(sql.NullInt64)
		case wugelucky.FieldID, wugelucky.FieldTianDaYan, wugelucky.FieldRenDaYan, wugelucky.FieldDiDaYan, wugelucky.FieldWaiDaYan, wugelucky.FieldZongDaYan:
			values[i] = new(sql.NullString)
		default:
			return nil, fmt.Errorf("unexpected column %q for type WuGeLucky", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the WuGeLucky fields.
func (wgl *WuGeLucky) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case wugelucky.FieldID:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value.Valid {
				wgl.ID = value.String
			}
		case wugelucky.FieldLastStroke1:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_stroke_1", values[i])
			} else if value.Valid {
				wgl.LastStroke1 = int32(value.Int64)
			}
		case wugelucky.FieldLastStroke2:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field last_stroke_2", values[i])
			} else if value.Valid {
				wgl.LastStroke2 = int32(value.Int64)
			}
		case wugelucky.FieldFirstStroke1:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field first_stroke_1", values[i])
			} else if value.Valid {
				wgl.FirstStroke1 = int32(value.Int64)
			}
		case wugelucky.FieldFirstStroke2:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field first_stroke_2", values[i])
			} else if value.Valid {
				wgl.FirstStroke2 = int32(value.Int64)
			}
		case wugelucky.FieldTianGe:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field tian_ge", values[i])
			} else if value.Valid {
				wgl.TianGe = int32(value.Int64)
			}
		case wugelucky.FieldTianDaYan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field tian_da_yan", values[i])
			} else if value.Valid {
				wgl.TianDaYan = value.String
			}
		case wugelucky.FieldRenGe:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field ren_ge", values[i])
			} else if value.Valid {
				wgl.RenGe = int32(value.Int64)
			}
		case wugelucky.FieldRenDaYan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field ren_da_yan", values[i])
			} else if value.Valid {
				wgl.RenDaYan = value.String
			}
		case wugelucky.FieldDiGe:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field di_ge", values[i])
			} else if value.Valid {
				wgl.DiGe = int32(value.Int64)
			}
		case wugelucky.FieldDiDaYan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field di_da_yan", values[i])
			} else if value.Valid {
				wgl.DiDaYan = value.String
			}
		case wugelucky.FieldWaiGe:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field wai_ge", values[i])
			} else if value.Valid {
				wgl.WaiGe = int32(value.Int64)
			}
		case wugelucky.FieldWaiDaYan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field wai_da_yan", values[i])
			} else if value.Valid {
				wgl.WaiDaYan = value.String
			}
		case wugelucky.FieldZongGe:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field zong_ge", values[i])
			} else if value.Valid {
				wgl.ZongGe = int32(value.Int64)
			}
		case wugelucky.FieldZongDaYan:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field zong_da_yan", values[i])
			} else if value.Valid {
				wgl.ZongDaYan = value.String
			}
		case wugelucky.FieldZongLucky:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field zong_lucky", values[i])
			} else if value.Valid {
				wgl.ZongLucky = value.Bool
			}
		case wugelucky.FieldZongSex:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field zong_sex", values[i])
			} else if value.Valid {
				wgl.ZongSex = value.Bool
			}
		case wugelucky.FieldZongMax:
			if value, ok := values[i].(*sql.NullBool); !ok {
				return fmt.Errorf("unexpected type %T for field zong_max", values[i])
			} else if value.Valid {
				wgl.ZongMax = value.Bool
			}
		}
	}
	return nil
}

// Update returns a builder for updating this WuGeLucky.
// Note that you need to call WuGeLucky.Unwrap() before calling this method if this WuGeLucky
// was returned from a transaction, and the transaction was committed or rolled back.
func (wgl *WuGeLucky) Update() *WuGeLuckyUpdateOne {
	return (&WuGeLuckyClient{config: wgl.config}).UpdateOne(wgl)
}

// Unwrap unwraps the WuGeLucky entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (wgl *WuGeLucky) Unwrap() *WuGeLucky {
	tx, ok := wgl.config.driver.(*txDriver)
	if !ok {
		panic("ent: WuGeLucky is not a transactional entity")
	}
	wgl.config.driver = tx.drv
	return wgl
}

// String implements the fmt.Stringer.
func (wgl *WuGeLucky) String() string {
	var builder strings.Builder
	builder.WriteString("WuGeLucky(")
	builder.WriteString(fmt.Sprintf("id=%v", wgl.ID))
	builder.WriteString(", last_stroke_1=")
	builder.WriteString(fmt.Sprintf("%v", wgl.LastStroke1))
	builder.WriteString(", last_stroke_2=")
	builder.WriteString(fmt.Sprintf("%v", wgl.LastStroke2))
	builder.WriteString(", first_stroke_1=")
	builder.WriteString(fmt.Sprintf("%v", wgl.FirstStroke1))
	builder.WriteString(", first_stroke_2=")
	builder.WriteString(fmt.Sprintf("%v", wgl.FirstStroke2))
	builder.WriteString(", tian_ge=")
	builder.WriteString(fmt.Sprintf("%v", wgl.TianGe))
	builder.WriteString(", tian_da_yan=")
	builder.WriteString(wgl.TianDaYan)
	builder.WriteString(", ren_ge=")
	builder.WriteString(fmt.Sprintf("%v", wgl.RenGe))
	builder.WriteString(", ren_da_yan=")
	builder.WriteString(wgl.RenDaYan)
	builder.WriteString(", di_ge=")
	builder.WriteString(fmt.Sprintf("%v", wgl.DiGe))
	builder.WriteString(", di_da_yan=")
	builder.WriteString(wgl.DiDaYan)
	builder.WriteString(", wai_ge=")
	builder.WriteString(fmt.Sprintf("%v", wgl.WaiGe))
	builder.WriteString(", wai_da_yan=")
	builder.WriteString(wgl.WaiDaYan)
	builder.WriteString(", zong_ge=")
	builder.WriteString(fmt.Sprintf("%v", wgl.ZongGe))
	builder.WriteString(", zong_da_yan=")
	builder.WriteString(wgl.ZongDaYan)
	builder.WriteString(", zong_lucky=")
	builder.WriteString(fmt.Sprintf("%v", wgl.ZongLucky))
	builder.WriteString(", zong_sex=")
	builder.WriteString(fmt.Sprintf("%v", wgl.ZongSex))
	builder.WriteString(", zong_max=")
	builder.WriteString(fmt.Sprintf("%v", wgl.ZongMax))
	builder.WriteByte(')')
	return builder.String()
}

// WuGeLuckies is a parsable slice of WuGeLucky.
type WuGeLuckies []*WuGeLucky

func (wgl WuGeLuckies) config(cfg config) {
	for _i := range wgl {
		wgl[_i].config = cfg
	}
}