// Code generated by entc, DO NOT EDIT.

package jobtitle

const (
	// Label holds the string label denoting the jobtitle type in the database.
	Label = "jobtitle"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldJobtitlename holds the string denoting the jobtitlename field in the database.
	FieldJobtitlename = "jobtitlename"

	// EdgePersonal holds the string denoting the personal edge name in mutations.
	EdgePersonal = "personal"

	// Table holds the table name of the jobtitle in the database.
	Table = "jobtitles"
	// PersonalTable is the table the holds the personal relation/edge.
	PersonalTable = "personals"
	// PersonalInverseTable is the table name for the Personal entity.
	// It exists in this package in order to avoid circular dependency with the "personal" package.
	PersonalInverseTable = "personals"
	// PersonalColumn is the table column denoting the personal relation/edge.
	PersonalColumn = "jobtitle_id"
)

// Columns holds all SQL columns for jobtitle fields.
var Columns = []string{
	FieldID,
	FieldJobtitlename,
}
