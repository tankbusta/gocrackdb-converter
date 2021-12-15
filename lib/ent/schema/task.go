package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"entgo.io/ent/schema/index"
	"github.com/google/uuid"
	"github.com/tankbusta/gocrackdb-converter/lib/oldschema"
)

// Task holds the schema definition for the Task entity.
type Task struct {
	ent.Schema
}

// Fields of the Task.
func (Task) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),

		field.String("name"),

		field.Enum("status").
			GoType(oldschema.TaskStatus("")),

		field.JSON("engine_payload", map[string]interface{}{}).
			Optional(),

		field.Time("created_at").
			Immutable().
			Default(time.Now().UTC),

		field.Time("last_updated_at").
			Immutable().
			Default(time.Now().UTC),

		field.String("assigned_to_host").
			Optional(),

		field.String("comment").
			Optional(),

		field.String("case_code").
			Optional(),

		field.Int("number_cracked").
			Default(0),

		field.Int("number_passwords").
			Default(0),

		field.String("error").
			Optional(),
	}
}

// Edges of the Task.
func (Task) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("created_by", User.Type).
			Unique().
			Required(),

		edge.To("cracked_passwords", CrackedPassword.Type),
	}
}

func (Task) Indexes() []ent.Index {
	return []ent.Index{
		// non-unique index.
		index.Fields("status"),
	}
}
