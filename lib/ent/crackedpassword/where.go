// Code generated by entc, DO NOT EDIT.

package crackedpassword

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/google/uuid"
	"github.com/tankbusta/gocrackdb-converter/lib/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id uuid.UUID) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id uuid.UUID) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id uuid.UUID) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...uuid.UUID) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.In(s.C(FieldID), v...))
	})
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...uuid.UUID) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(ids) == 0 {
			s.Where(sql.False())
			return
		}
		v := make([]interface{}, len(ids))
		for i := range v {
			v[i] = ids[i]
		}
		s.Where(sql.NotIn(s.C(FieldID), v...))
	})
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id uuid.UUID) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id uuid.UUID) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id uuid.UUID) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id uuid.UUID) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Hash applies equality check predicate on the "hash" field. It's identical to HashEQ.
func Hash(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// Value applies equality check predicate on the "value" field. It's identical to ValueEQ.
func Value(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// CrackedAt applies equality check predicate on the "cracked_at" field. It's identical to CrackedAtEQ.
func CrackedAt(v time.Time) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCrackedAt), v))
	})
}

// HashEQ applies the EQ predicate on the "hash" field.
func HashEQ(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldHash), v))
	})
}

// HashNEQ applies the NEQ predicate on the "hash" field.
func HashNEQ(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldHash), v))
	})
}

// HashIn applies the In predicate on the "hash" field.
func HashIn(vs ...string) predicate.CrackedPassword {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CrackedPassword(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldHash), v...))
	})
}

// HashNotIn applies the NotIn predicate on the "hash" field.
func HashNotIn(vs ...string) predicate.CrackedPassword {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CrackedPassword(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldHash), v...))
	})
}

// HashGT applies the GT predicate on the "hash" field.
func HashGT(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldHash), v))
	})
}

// HashGTE applies the GTE predicate on the "hash" field.
func HashGTE(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldHash), v))
	})
}

// HashLT applies the LT predicate on the "hash" field.
func HashLT(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldHash), v))
	})
}

// HashLTE applies the LTE predicate on the "hash" field.
func HashLTE(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldHash), v))
	})
}

// HashContains applies the Contains predicate on the "hash" field.
func HashContains(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldHash), v))
	})
}

// HashHasPrefix applies the HasPrefix predicate on the "hash" field.
func HashHasPrefix(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldHash), v))
	})
}

// HashHasSuffix applies the HasSuffix predicate on the "hash" field.
func HashHasSuffix(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldHash), v))
	})
}

// HashEqualFold applies the EqualFold predicate on the "hash" field.
func HashEqualFold(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldHash), v))
	})
}

// HashContainsFold applies the ContainsFold predicate on the "hash" field.
func HashContainsFold(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldHash), v))
	})
}

// ValueEQ applies the EQ predicate on the "value" field.
func ValueEQ(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldValue), v))
	})
}

// ValueNEQ applies the NEQ predicate on the "value" field.
func ValueNEQ(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldValue), v))
	})
}

// ValueIn applies the In predicate on the "value" field.
func ValueIn(vs ...string) predicate.CrackedPassword {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CrackedPassword(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldValue), v...))
	})
}

// ValueNotIn applies the NotIn predicate on the "value" field.
func ValueNotIn(vs ...string) predicate.CrackedPassword {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CrackedPassword(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldValue), v...))
	})
}

// ValueGT applies the GT predicate on the "value" field.
func ValueGT(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldValue), v))
	})
}

// ValueGTE applies the GTE predicate on the "value" field.
func ValueGTE(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldValue), v))
	})
}

// ValueLT applies the LT predicate on the "value" field.
func ValueLT(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldValue), v))
	})
}

// ValueLTE applies the LTE predicate on the "value" field.
func ValueLTE(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldValue), v))
	})
}

// ValueContains applies the Contains predicate on the "value" field.
func ValueContains(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldValue), v))
	})
}

// ValueHasPrefix applies the HasPrefix predicate on the "value" field.
func ValueHasPrefix(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldValue), v))
	})
}

// ValueHasSuffix applies the HasSuffix predicate on the "value" field.
func ValueHasSuffix(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldValue), v))
	})
}

// ValueEqualFold applies the EqualFold predicate on the "value" field.
func ValueEqualFold(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldValue), v))
	})
}

// ValueContainsFold applies the ContainsFold predicate on the "value" field.
func ValueContainsFold(v string) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldValue), v))
	})
}

// CrackedAtEQ applies the EQ predicate on the "cracked_at" field.
func CrackedAtEQ(v time.Time) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldCrackedAt), v))
	})
}

// CrackedAtNEQ applies the NEQ predicate on the "cracked_at" field.
func CrackedAtNEQ(v time.Time) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldCrackedAt), v))
	})
}

// CrackedAtIn applies the In predicate on the "cracked_at" field.
func CrackedAtIn(vs ...time.Time) predicate.CrackedPassword {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CrackedPassword(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldCrackedAt), v...))
	})
}

// CrackedAtNotIn applies the NotIn predicate on the "cracked_at" field.
func CrackedAtNotIn(vs ...time.Time) predicate.CrackedPassword {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.CrackedPassword(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldCrackedAt), v...))
	})
}

// CrackedAtGT applies the GT predicate on the "cracked_at" field.
func CrackedAtGT(v time.Time) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldCrackedAt), v))
	})
}

// CrackedAtGTE applies the GTE predicate on the "cracked_at" field.
func CrackedAtGTE(v time.Time) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldCrackedAt), v))
	})
}

// CrackedAtLT applies the LT predicate on the "cracked_at" field.
func CrackedAtLT(v time.Time) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldCrackedAt), v))
	})
}

// CrackedAtLTE applies the LTE predicate on the "cracked_at" field.
func CrackedAtLTE(v time.Time) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldCrackedAt), v))
	})
}

// HasRelatedTask applies the HasEdge predicate on the "related_task" edge.
func HasRelatedTask() predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RelatedTaskTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RelatedTaskTable, RelatedTaskColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRelatedTaskWith applies the HasEdge predicate on the "related_task" edge with a given conditions (other predicates).
func HasRelatedTaskWith(preds ...predicate.Task) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(RelatedTaskInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, RelatedTaskTable, RelatedTaskColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.CrackedPassword) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.CrackedPassword) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for i, p := range predicates {
			if i > 0 {
				s1.Or()
			}
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Not applies the not operator on the given predicate.
func Not(p predicate.CrackedPassword) predicate.CrackedPassword {
	return predicate.CrackedPassword(func(s *sql.Selector) {
		p(s.Not())
	})
}