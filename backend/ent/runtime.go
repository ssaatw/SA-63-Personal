// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/ssaatw/app/ent/personal"
	"github.com/ssaatw/app/ent/schema"
)

// The init function reads all schema descriptors with runtime
// code (default values, validators or hooks) and stitches it
// to their package variables.
func init() {
	personalFields := schema.Personal{}.Fields()
	_ = personalFields
	// personalDescAdded is the schema descriptor for Added field.
	personalDescAdded := personalFields[3].Descriptor()
	// personal.DefaultAdded holds the default value on creation for the Added field.
	personal.DefaultAdded = personalDescAdded.Default.(func() time.Time)
}
