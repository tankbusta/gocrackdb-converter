// Code generated by entc, DO NOT EDIT.

package ent

import (
	"time"

	"github.com/google/uuid"
	"github.com/tankbusta/gocrackdb-converter/lib/ent/crackedpassword"
	"github.com/tankbusta/gocrackdb-converter/lib/ent/schema"
	"github.com/tankbusta/gocrackdb-converter/lib/ent/task"
	"github.com/tankbusta/gocrackdb-converter/lib/ent/user"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	crackedpasswordFields := schema.CrackedPassword{}.Fields()
	_ = crackedpasswordFields
	// crackedpasswordDescCrackedAt is the schema descriptor for cracked_at field.
	crackedpasswordDescCrackedAt := crackedpasswordFields[3].Descriptor()
	// crackedpassword.DefaultCrackedAt holds the default value on creation for the cracked_at field.
	crackedpassword.DefaultCrackedAt = crackedpasswordDescCrackedAt.Default.(func() time.Time)
	// crackedpasswordDescID is the schema descriptor for id field.
	crackedpasswordDescID := crackedpasswordFields[0].Descriptor()
	// crackedpassword.DefaultID holds the default value on creation for the id field.
	crackedpassword.DefaultID = crackedpasswordDescID.Default.(func() uuid.UUID)
	taskFields := schema.Task{}.Fields()
	_ = taskFields
	// taskDescCreatedAt is the schema descriptor for created_at field.
	taskDescCreatedAt := taskFields[4].Descriptor()
	// task.DefaultCreatedAt holds the default value on creation for the created_at field.
	task.DefaultCreatedAt = taskDescCreatedAt.Default.(func() time.Time)
	// taskDescLastUpdatedAt is the schema descriptor for last_updated_at field.
	taskDescLastUpdatedAt := taskFields[5].Descriptor()
	// task.DefaultLastUpdatedAt holds the default value on creation for the last_updated_at field.
	task.DefaultLastUpdatedAt = taskDescLastUpdatedAt.Default.(func() time.Time)
	// taskDescNumberCracked is the schema descriptor for number_cracked field.
	taskDescNumberCracked := taskFields[9].Descriptor()
	// task.DefaultNumberCracked holds the default value on creation for the number_cracked field.
	task.DefaultNumberCracked = taskDescNumberCracked.Default.(int)
	// taskDescNumberPasswords is the schema descriptor for number_passwords field.
	taskDescNumberPasswords := taskFields[10].Descriptor()
	// task.DefaultNumberPasswords holds the default value on creation for the number_passwords field.
	task.DefaultNumberPasswords = taskDescNumberPasswords.Default.(int)
	// taskDescID is the schema descriptor for id field.
	taskDescID := taskFields[0].Descriptor()
	// task.DefaultID holds the default value on creation for the id field.
	task.DefaultID = taskDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescCreatedAt is the schema descriptor for created_at field.
	userDescCreatedAt := userFields[5].Descriptor()
	// user.DefaultCreatedAt holds the default value on creation for the created_at field.
	user.DefaultCreatedAt = userDescCreatedAt.Default.(func() time.Time)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
}