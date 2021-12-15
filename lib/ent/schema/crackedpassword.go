package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
	"github.com/google/uuid"
)

// CrackedPassword holds the schema definition for the CrackedPassword entity.
type CrackedPassword struct {
	ent.Schema
}

// Fields of the CrackedPassword.
func (CrackedPassword) Fields() []ent.Field {
	return []ent.Field{
		field.UUID("id", uuid.UUID{}).
			Default(uuid.New),

		field.String("hash"),

		field.String("value"),

		field.Time("cracked_at").
			Immutable().
			Default(time.Now().UTC),
	}
}

// Edges of the CrackedPassword.
func (CrackedPassword) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("related_task", Task.Type).
			Ref("cracked_passwords").
			Unique().
			Required(),
	}
}
