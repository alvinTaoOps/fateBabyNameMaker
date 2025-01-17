// Code generated by ent, DO NOT EDIT.

package version

import (
	"entgo.io/ent/dialect/sql"
)

const (
	// Label holds the string label denoting the version type in the database.
	Label = "version"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCurrentVersion holds the string denoting the current_version field in the database.
	FieldCurrentVersion = "current_version"
	// FieldUpdatedUnix holds the string denoting the updated_unix field in the database.
	FieldUpdatedUnix = "updated_unix"
	// Table holds the table name of the version in the database.
	Table = "versions"
)

// Columns holds all SQL columns for version fields.
var Columns = []string{
	FieldID,
	FieldCurrentVersion,
	FieldUpdatedUnix,
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

// OrderOption defines the ordering options for the Version queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByCurrentVersion orders the results by the current_version field.
func ByCurrentVersion(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCurrentVersion, opts...).ToFunc()
}

// ByUpdatedUnix orders the results by the updated_unix field.
func ByUpdatedUnix(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldUpdatedUnix, opts...).ToFunc()
}
