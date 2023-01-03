// Code generated by ent, DO NOT EDIT.

package wugelucky

const (
	// Label holds the string label denoting the wugelucky type in the database.
	Label = "wu_ge_lucky"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldLastStroke1 holds the string denoting the last_stroke_1 field in the database.
	FieldLastStroke1 = "last_stroke_1"
	// FieldLastStroke2 holds the string denoting the last_stroke_2 field in the database.
	FieldLastStroke2 = "last_stroke_2"
	// FieldFirstStroke1 holds the string denoting the first_stroke_1 field in the database.
	FieldFirstStroke1 = "first_stroke_1"
	// FieldFirstStroke2 holds the string denoting the first_stroke_2 field in the database.
	FieldFirstStroke2 = "first_stroke_2"
	// FieldTianGe holds the string denoting the tian_ge field in the database.
	FieldTianGe = "tian_ge"
	// FieldTianDaYan holds the string denoting the tian_da_yan field in the database.
	FieldTianDaYan = "tian_da_yan"
	// FieldRenGe holds the string denoting the ren_ge field in the database.
	FieldRenGe = "ren_ge"
	// FieldRenDaYan holds the string denoting the ren_da_yan field in the database.
	FieldRenDaYan = "ren_da_yan"
	// FieldDiGe holds the string denoting the di_ge field in the database.
	FieldDiGe = "di_ge"
	// FieldDiDaYan holds the string denoting the di_da_yan field in the database.
	FieldDiDaYan = "di_da_yan"
	// FieldWaiGe holds the string denoting the wai_ge field in the database.
	FieldWaiGe = "wai_ge"
	// FieldWaiDaYan holds the string denoting the wai_da_yan field in the database.
	FieldWaiDaYan = "wai_da_yan"
	// FieldZongGe holds the string denoting the zong_ge field in the database.
	FieldZongGe = "zong_ge"
	// FieldZongDaYan holds the string denoting the zong_da_yan field in the database.
	FieldZongDaYan = "zong_da_yan"
	// FieldZongLucky holds the string denoting the zong_lucky field in the database.
	FieldZongLucky = "zong_lucky"
	// FieldZongSex holds the string denoting the zong_sex field in the database.
	FieldZongSex = "zong_sex"
	// FieldZongMax holds the string denoting the zong_max field in the database.
	FieldZongMax = "zong_max"
	// Table holds the table name of the wugelucky in the database.
	Table = "wu_ge_lucky"
)

// Columns holds all SQL columns for wugelucky fields.
var Columns = []string{
	FieldID,
	FieldLastStroke1,
	FieldLastStroke2,
	FieldFirstStroke1,
	FieldFirstStroke2,
	FieldTianGe,
	FieldTianDaYan,
	FieldRenGe,
	FieldRenDaYan,
	FieldDiGe,
	FieldDiDaYan,
	FieldWaiGe,
	FieldWaiDaYan,
	FieldZongGe,
	FieldZongDaYan,
	FieldZongLucky,
	FieldZongSex,
	FieldZongMax,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}
