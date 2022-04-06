// Code generated by entc, DO NOT EDIT.

package entity

const (
	// Label holds the string label denoting the entity type in the database.
	Label = "entity"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldData holds the string denoting the data field in the database.
	FieldData = "data"
	// FieldIsFun holds the string denoting the isfun field in the database.
	FieldIsFun = "is_fun"
	// FieldCounter holds the string denoting the counter field in the database.
	FieldCounter = "counter"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldStrings holds the string denoting the strings field in the database.
	FieldStrings = "strings"
	// Table holds the table name of the entity in the database.
	Table = "entities"
)

// Columns holds all SQL columns for entity fields.
var Columns = []string{
	FieldID,
	FieldData,
	FieldIsFun,
	FieldCounter,
	FieldTimestamp,
	FieldStrings,
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
