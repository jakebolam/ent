// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package relationship

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgeschema/ent/predicate"
)

// Weight applies equality check predicate on the "weight" field. It's identical to WeightEQ.
func Weight(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// RelativeID applies equality check predicate on the "relative_id" field. It's identical to RelativeIDEQ.
func RelativeID(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRelativeID), v))
	})
}

// InfoID applies equality check predicate on the "info_id" field. It's identical to InfoIDEQ.
func InfoID(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfoID), v))
	})
}

// WeightEQ applies the EQ predicate on the "weight" field.
func WeightEQ(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldWeight), v))
	})
}

// WeightNEQ applies the NEQ predicate on the "weight" field.
func WeightNEQ(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldWeight), v))
	})
}

// WeightIn applies the In predicate on the "weight" field.
func WeightIn(vs ...int) predicate.Relationship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldWeight), v...))
	})
}

// WeightNotIn applies the NotIn predicate on the "weight" field.
func WeightNotIn(vs ...int) predicate.Relationship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldWeight), v...))
	})
}

// WeightGT applies the GT predicate on the "weight" field.
func WeightGT(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldWeight), v))
	})
}

// WeightGTE applies the GTE predicate on the "weight" field.
func WeightGTE(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldWeight), v))
	})
}

// WeightLT applies the LT predicate on the "weight" field.
func WeightLT(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldWeight), v))
	})
}

// WeightLTE applies the LTE predicate on the "weight" field.
func WeightLTE(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldWeight), v))
	})
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldUserID), v))
	})
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldUserID), v))
	})
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Relationship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldUserID), v...))
	})
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Relationship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldUserID), v...))
	})
}

// RelativeIDEQ applies the EQ predicate on the "relative_id" field.
func RelativeIDEQ(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldRelativeID), v))
	})
}

// RelativeIDNEQ applies the NEQ predicate on the "relative_id" field.
func RelativeIDNEQ(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldRelativeID), v))
	})
}

// RelativeIDIn applies the In predicate on the "relative_id" field.
func RelativeIDIn(vs ...int) predicate.Relationship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldRelativeID), v...))
	})
}

// RelativeIDNotIn applies the NotIn predicate on the "relative_id" field.
func RelativeIDNotIn(vs ...int) predicate.Relationship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldRelativeID), v...))
	})
}

// InfoIDEQ applies the EQ predicate on the "info_id" field.
func InfoIDEQ(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldInfoID), v))
	})
}

// InfoIDNEQ applies the NEQ predicate on the "info_id" field.
func InfoIDNEQ(v int) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldInfoID), v))
	})
}

// InfoIDIn applies the In predicate on the "info_id" field.
func InfoIDIn(vs ...int) predicate.Relationship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.In(s.C(FieldInfoID), v...))
	})
}

// InfoIDNotIn applies the NotIn predicate on the "info_id" field.
func InfoIDNotIn(vs ...int) predicate.Relationship {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.NotIn(s.C(FieldInfoID), v...))
	})
}

// InfoIDIsNil applies the IsNil predicate on the "info_id" field.
func InfoIDIsNil() predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.IsNull(s.C(FieldInfoID)))
	})
}

// InfoIDNotNil applies the NotNil predicate on the "info_id" field.
func InfoIDNotNil() predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s.Where(sql.NotNull(s.C(FieldInfoID)))
	})
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.To(UserInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, UserColumn),
			sqlgraph.To(UserInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, UserTable, UserColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasRelative applies the HasEdge predicate on the "relative" edge.
func HasRelative() predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, RelativeColumn),
			sqlgraph.To(RelativeInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RelativeTable, RelativeColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasRelativeWith applies the HasEdge predicate on the "relative" edge with a given conditions (other predicates).
func HasRelativeWith(preds ...predicate.User) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, RelativeColumn),
			sqlgraph.To(RelativeInverseTable, UserFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, RelativeTable, RelativeColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasInfo applies the HasEdge predicate on the "info" edge.
func HasInfo() predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, InfoColumn),
			sqlgraph.To(InfoInverseTable, RelationshipInfoFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InfoTable, InfoColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasInfoWith applies the HasEdge predicate on the "info" edge with a given conditions (other predicates).
func HasInfoWith(preds ...predicate.RelationshipInfo) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, InfoColumn),
			sqlgraph.To(InfoInverseTable, RelationshipInfoFieldID),
			sqlgraph.Edge(sqlgraph.M2O, false, InfoTable, InfoColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Relationship) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Relationship) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
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
func Not(p predicate.Relationship) predicate.Relationship {
	return predicate.Relationship(func(s *sql.Selector) {
		p(s.Not())
	})
}