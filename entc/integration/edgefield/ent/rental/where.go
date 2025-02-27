// Copyright 2019-present Facebook Inc. All rights reserved.
// This source code is licensed under the Apache 2.0 license found
// in the LICENSE file in the root directory of this source tree.

// Code generated by ent, DO NOT EDIT.

package rental

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/entc/integration/edgefield/ent/predicate"
	"github.com/google/uuid"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Rental {
	return predicate.Rental(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Rental {
	return predicate.Rental(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Rental {
	return predicate.Rental(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Rental {
	return predicate.Rental(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Rental {
	return predicate.Rental(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Rental {
	return predicate.Rental(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Rental {
	return predicate.Rental(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Rental {
	return predicate.Rental(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Rental {
	return predicate.Rental(sql.FieldLTE(FieldID, id))
}

// Date applies equality check predicate on the "date" field. It's identical to DateEQ.
func Date(v time.Time) predicate.Rental {
	return predicate.Rental(sql.FieldEQ(FieldDate, v))
}

// UserID applies equality check predicate on the "user_id" field. It's identical to UserIDEQ.
func UserID(v int) predicate.Rental {
	return predicate.Rental(sql.FieldEQ(FieldUserID, v))
}

// CarID applies equality check predicate on the "car_id" field. It's identical to CarIDEQ.
func CarID(v uuid.UUID) predicate.Rental {
	return predicate.Rental(sql.FieldEQ(FieldCarID, v))
}

// DateEQ applies the EQ predicate on the "date" field.
func DateEQ(v time.Time) predicate.Rental {
	return predicate.Rental(sql.FieldEQ(FieldDate, v))
}

// DateNEQ applies the NEQ predicate on the "date" field.
func DateNEQ(v time.Time) predicate.Rental {
	return predicate.Rental(sql.FieldNEQ(FieldDate, v))
}

// DateIn applies the In predicate on the "date" field.
func DateIn(vs ...time.Time) predicate.Rental {
	return predicate.Rental(sql.FieldIn(FieldDate, vs...))
}

// DateNotIn applies the NotIn predicate on the "date" field.
func DateNotIn(vs ...time.Time) predicate.Rental {
	return predicate.Rental(sql.FieldNotIn(FieldDate, vs...))
}

// DateGT applies the GT predicate on the "date" field.
func DateGT(v time.Time) predicate.Rental {
	return predicate.Rental(sql.FieldGT(FieldDate, v))
}

// DateGTE applies the GTE predicate on the "date" field.
func DateGTE(v time.Time) predicate.Rental {
	return predicate.Rental(sql.FieldGTE(FieldDate, v))
}

// DateLT applies the LT predicate on the "date" field.
func DateLT(v time.Time) predicate.Rental {
	return predicate.Rental(sql.FieldLT(FieldDate, v))
}

// DateLTE applies the LTE predicate on the "date" field.
func DateLTE(v time.Time) predicate.Rental {
	return predicate.Rental(sql.FieldLTE(FieldDate, v))
}

// UserIDEQ applies the EQ predicate on the "user_id" field.
func UserIDEQ(v int) predicate.Rental {
	return predicate.Rental(sql.FieldEQ(FieldUserID, v))
}

// UserIDNEQ applies the NEQ predicate on the "user_id" field.
func UserIDNEQ(v int) predicate.Rental {
	return predicate.Rental(sql.FieldNEQ(FieldUserID, v))
}

// UserIDIn applies the In predicate on the "user_id" field.
func UserIDIn(vs ...int) predicate.Rental {
	return predicate.Rental(sql.FieldIn(FieldUserID, vs...))
}

// UserIDNotIn applies the NotIn predicate on the "user_id" field.
func UserIDNotIn(vs ...int) predicate.Rental {
	return predicate.Rental(sql.FieldNotIn(FieldUserID, vs...))
}

// CarIDEQ applies the EQ predicate on the "car_id" field.
func CarIDEQ(v uuid.UUID) predicate.Rental {
	return predicate.Rental(sql.FieldEQ(FieldCarID, v))
}

// CarIDNEQ applies the NEQ predicate on the "car_id" field.
func CarIDNEQ(v uuid.UUID) predicate.Rental {
	return predicate.Rental(sql.FieldNEQ(FieldCarID, v))
}

// CarIDIn applies the In predicate on the "car_id" field.
func CarIDIn(vs ...uuid.UUID) predicate.Rental {
	return predicate.Rental(sql.FieldIn(FieldCarID, vs...))
}

// CarIDNotIn applies the NotIn predicate on the "car_id" field.
func CarIDNotIn(vs ...uuid.UUID) predicate.Rental {
	return predicate.Rental(sql.FieldNotIn(FieldCarID, vs...))
}

// HasUser applies the HasEdge predicate on the "user" edge.
func HasUser() predicate.Rental {
	return predicate.Rental(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, UserTable, UserColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasUserWith applies the HasEdge predicate on the "user" edge with a given conditions (other predicates).
func HasUserWith(preds ...predicate.User) predicate.Rental {
	return predicate.Rental(func(s *sql.Selector) {
		step := newUserStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// HasCar applies the HasEdge predicate on the "car" edge.
func HasCar() predicate.Rental {
	return predicate.Rental(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.Edge(sqlgraph.M2O, true, CarTable, CarColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasCarWith applies the HasEdge predicate on the "car" edge with a given conditions (other predicates).
func HasCarWith(preds ...predicate.Car) predicate.Rental {
	return predicate.Rental(func(s *sql.Selector) {
		step := newCarStep()
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Rental) predicate.Rental {
	return predicate.Rental(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Rental) predicate.Rental {
	return predicate.Rental(func(s *sql.Selector) {
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
func Not(p predicate.Rental) predicate.Rental {
	return predicate.Rental(func(s *sql.Selector) {
		p(s.Not())
	})
}
