// Code generated by entc, DO NOT EDIT.

package history

import (
	"fmt"
	"io"
	"strconv"
)

const (
	// Label holds the string label denoting the history type in the database.
	Label = "history"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldEntityName holds the string denoting the entity_name field in the database.
	FieldEntityName = "entity_name"
	// FieldRecordID holds the string denoting the record_id field in the database.
	FieldRecordID = "record_id"
	// FieldTimestamp holds the string denoting the timestamp field in the database.
	FieldTimestamp = "timestamp"
	// FieldAction holds the string denoting the action field in the database.
	FieldAction = "action"
	// EdgeChanges holds the string denoting the changes edge name in mutations.
	EdgeChanges = "changes"
	// Table holds the table name of the history in the database.
	Table = "histories"
	// ChangesTable is the table that holds the changes relation/edge.
	ChangesTable = "changes"
	// ChangesInverseTable is the table name for the Changes entity.
	// It exists in this package in order to avoid circular dependency with the "changes" package.
	ChangesInverseTable = "changes"
	// ChangesColumn is the table column denoting the changes relation/edge.
	ChangesColumn = "history_changes"
)

// Columns holds all SQL columns for history fields.
var Columns = []string{
	FieldID,
	FieldEntityName,
	FieldRecordID,
	FieldTimestamp,
	FieldAction,
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

// Action defines the type for the "action" enum field.
type Action string

// Action values.
const (
	ActionCreate Action = "CREATE"
	ActionUpdate Action = "UPDATE"
	ActionDelete Action = "DELETE"
)

func (a Action) String() string {
	return string(a)
}

// ActionValidator is a validator for the "action" field enum values. It is called by the builders before save.
func ActionValidator(a Action) error {
	switch a {
	case ActionCreate, ActionUpdate, ActionDelete:
		return nil
	default:
		return fmt.Errorf("history: invalid enum value for action field: %q", a)
	}
}

// MarshalGQL implements graphql.Marshaler interface.
func (a Action) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(a.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (a *Action) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*a = Action(str)
	if err := ActionValidator(*a); err != nil {
		return fmt.Errorf("%s is not a valid Action", str)
	}
	return nil
}
