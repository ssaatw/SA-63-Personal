// Code generated by entc, DO NOT EDIT.

package systemmember

import (
	"github.com/facebookincubator/ent/dialect/sql"
	"github.com/facebookincubator/ent/dialect/sql/sqlgraph"
	"github.com/ssaatw/app/ent/predicate"
)

// ID filters vertices based on their identifier.
func ID(id int) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldID), id))
	})
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldID), id))
	})
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
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
func IDNotIn(ids ...int) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
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
func IDGT(id int) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldID), id))
	})
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldID), id))
	})
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldID), id))
	})
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldID), id))
	})
}

// Mail applies equality check predicate on the "Mail" field. It's identical to MailEQ.
func Mail(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMail), v))
	})
}

// Password applies equality check predicate on the "Password" field. It's identical to PasswordEQ.
func Password(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// MailEQ applies the EQ predicate on the "Mail" field.
func MailEQ(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldMail), v))
	})
}

// MailNEQ applies the NEQ predicate on the "Mail" field.
func MailNEQ(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldMail), v))
	})
}

// MailIn applies the In predicate on the "Mail" field.
func MailIn(vs ...string) predicate.Systemmember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemmember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldMail), v...))
	})
}

// MailNotIn applies the NotIn predicate on the "Mail" field.
func MailNotIn(vs ...string) predicate.Systemmember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemmember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldMail), v...))
	})
}

// MailGT applies the GT predicate on the "Mail" field.
func MailGT(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldMail), v))
	})
}

// MailGTE applies the GTE predicate on the "Mail" field.
func MailGTE(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldMail), v))
	})
}

// MailLT applies the LT predicate on the "Mail" field.
func MailLT(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldMail), v))
	})
}

// MailLTE applies the LTE predicate on the "Mail" field.
func MailLTE(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldMail), v))
	})
}

// MailContains applies the Contains predicate on the "Mail" field.
func MailContains(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldMail), v))
	})
}

// MailHasPrefix applies the HasPrefix predicate on the "Mail" field.
func MailHasPrefix(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldMail), v))
	})
}

// MailHasSuffix applies the HasSuffix predicate on the "Mail" field.
func MailHasSuffix(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldMail), v))
	})
}

// MailEqualFold applies the EqualFold predicate on the "Mail" field.
func MailEqualFold(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldMail), v))
	})
}

// MailContainsFold applies the ContainsFold predicate on the "Mail" field.
func MailContainsFold(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldMail), v))
	})
}

// PasswordEQ applies the EQ predicate on the "Password" field.
func PasswordEQ(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.EQ(s.C(FieldPassword), v))
	})
}

// PasswordNEQ applies the NEQ predicate on the "Password" field.
func PasswordNEQ(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.NEQ(s.C(FieldPassword), v))
	})
}

// PasswordIn applies the In predicate on the "Password" field.
func PasswordIn(vs ...string) predicate.Systemmember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemmember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.In(s.C(FieldPassword), v...))
	})
}

// PasswordNotIn applies the NotIn predicate on the "Password" field.
func PasswordNotIn(vs ...string) predicate.Systemmember {
	v := make([]interface{}, len(vs))
	for i := range v {
		v[i] = vs[i]
	}
	return predicate.Systemmember(func(s *sql.Selector) {
		// if not arguments were provided, append the FALSE constants,
		// since we can't apply "IN ()". This will make this predicate falsy.
		if len(v) == 0 {
			s.Where(sql.False())
			return
		}
		s.Where(sql.NotIn(s.C(FieldPassword), v...))
	})
}

// PasswordGT applies the GT predicate on the "Password" field.
func PasswordGT(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.GT(s.C(FieldPassword), v))
	})
}

// PasswordGTE applies the GTE predicate on the "Password" field.
func PasswordGTE(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.GTE(s.C(FieldPassword), v))
	})
}

// PasswordLT applies the LT predicate on the "Password" field.
func PasswordLT(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.LT(s.C(FieldPassword), v))
	})
}

// PasswordLTE applies the LTE predicate on the "Password" field.
func PasswordLTE(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.LTE(s.C(FieldPassword), v))
	})
}

// PasswordContains applies the Contains predicate on the "Password" field.
func PasswordContains(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.Contains(s.C(FieldPassword), v))
	})
}

// PasswordHasPrefix applies the HasPrefix predicate on the "Password" field.
func PasswordHasPrefix(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.HasPrefix(s.C(FieldPassword), v))
	})
}

// PasswordHasSuffix applies the HasSuffix predicate on the "Password" field.
func PasswordHasSuffix(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.HasSuffix(s.C(FieldPassword), v))
	})
}

// PasswordEqualFold applies the EqualFold predicate on the "Password" field.
func PasswordEqualFold(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.EqualFold(s.C(FieldPassword), v))
	})
}

// PasswordContainsFold applies the ContainsFold predicate on the "Password" field.
func PasswordContainsFold(v string) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s.Where(sql.ContainsFold(s.C(FieldPassword), v))
	})
}

// HasPersonal applies the HasEdge predicate on the "personal" edge.
func HasPersonal() predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighbors(s, step)
	})
}

// HasPersonalWith applies the HasEdge predicate on the "personal" edge with a given conditions (other predicates).
func HasPersonalWith(preds ...predicate.Personal) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		step := sqlgraph.NewStep(
			sqlgraph.From(Table, FieldID),
			sqlgraph.To(PersonalInverseTable, FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, PersonalTable, PersonalColumn),
		)
		sqlgraph.HasNeighborsWith(s, step, func(s *sql.Selector) {
			for _, p := range preds {
				p(s)
			}
		})
	})
}

// And groups list of predicates with the AND operator between them.
func And(predicates ...predicate.Systemmember) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups list of predicates with the OR operator between them.
func Or(predicates ...predicate.Systemmember) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
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
func Not(p predicate.Systemmember) predicate.Systemmember {
	return predicate.Systemmember(func(s *sql.Selector) {
		p(s.Not())
	})
}
