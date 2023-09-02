// Code generated by ent, DO NOT EDIT.

package todo

import (
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/yugo-ibuki/musdok/ent/predicate"
)

// ID filters vertices based on their ID field.
func ID(id int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDEQ applies the EQ predicate on the ID field.
func IDEQ(id int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldID, id))
}

// IDNEQ applies the NEQ predicate on the ID field.
func IDNEQ(id int) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldID, id))
}

// IDIn applies the In predicate on the ID field.
func IDIn(ids ...int) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldID, ids...))
}

// IDNotIn applies the NotIn predicate on the ID field.
func IDNotIn(ids ...int) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldID, ids...))
}

// IDGT applies the GT predicate on the ID field.
func IDGT(id int) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldID, id))
}

// IDGTE applies the GTE predicate on the ID field.
func IDGTE(id int) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldID, id))
}

// IDLT applies the LT predicate on the ID field.
func IDLT(id int) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldID, id))
}

// IDLTE applies the LTE predicate on the ID field.
func IDLTE(id int) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldID, id))
}

// Order applies equality check predicate on the "order" field. It's identical to OrderEQ.
func Order(v int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldOrder, v))
}

// Title applies equality check predicate on the "title" field. It's identical to TitleEQ.
func Title(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTitle, v))
}

// Description applies equality check predicate on the "description" field. It's identical to DescriptionEQ.
func Description(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldDescription, v))
}

// Done applies equality check predicate on the "done" field. It's identical to DoneEQ.
func Done(v bool) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldDone, v))
}

// CreatedAt applies equality check predicate on the "created_at" field. It's identical to CreatedAtEQ.
func CreatedAt(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreatedAt, v))
}

// UpdatedAt applies equality check predicate on the "updated_at" field. It's identical to UpdatedAtEQ.
func UpdatedAt(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUpdatedAt, v))
}

// OrderEQ applies the EQ predicate on the "order" field.
func OrderEQ(v int) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldOrder, v))
}

// OrderNEQ applies the NEQ predicate on the "order" field.
func OrderNEQ(v int) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldOrder, v))
}

// OrderIn applies the In predicate on the "order" field.
func OrderIn(vs ...int) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldOrder, vs...))
}

// OrderNotIn applies the NotIn predicate on the "order" field.
func OrderNotIn(vs ...int) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldOrder, vs...))
}

// OrderGT applies the GT predicate on the "order" field.
func OrderGT(v int) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldOrder, v))
}

// OrderGTE applies the GTE predicate on the "order" field.
func OrderGTE(v int) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldOrder, v))
}

// OrderLT applies the LT predicate on the "order" field.
func OrderLT(v int) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldOrder, v))
}

// OrderLTE applies the LTE predicate on the "order" field.
func OrderLTE(v int) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldOrder, v))
}

// TitleEQ applies the EQ predicate on the "title" field.
func TitleEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldTitle, v))
}

// TitleNEQ applies the NEQ predicate on the "title" field.
func TitleNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldTitle, v))
}

// TitleIn applies the In predicate on the "title" field.
func TitleIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldTitle, vs...))
}

// TitleNotIn applies the NotIn predicate on the "title" field.
func TitleNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldTitle, vs...))
}

// TitleGT applies the GT predicate on the "title" field.
func TitleGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldTitle, v))
}

// TitleGTE applies the GTE predicate on the "title" field.
func TitleGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldTitle, v))
}

// TitleLT applies the LT predicate on the "title" field.
func TitleLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldTitle, v))
}

// TitleLTE applies the LTE predicate on the "title" field.
func TitleLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldTitle, v))
}

// TitleContains applies the Contains predicate on the "title" field.
func TitleContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldTitle, v))
}

// TitleHasPrefix applies the HasPrefix predicate on the "title" field.
func TitleHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldTitle, v))
}

// TitleHasSuffix applies the HasSuffix predicate on the "title" field.
func TitleHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldTitle, v))
}

// TitleEqualFold applies the EqualFold predicate on the "title" field.
func TitleEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldTitle, v))
}

// TitleContainsFold applies the ContainsFold predicate on the "title" field.
func TitleContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldTitle, v))
}

// DescriptionEQ applies the EQ predicate on the "description" field.
func DescriptionEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldDescription, v))
}

// DescriptionNEQ applies the NEQ predicate on the "description" field.
func DescriptionNEQ(v string) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldDescription, v))
}

// DescriptionIn applies the In predicate on the "description" field.
func DescriptionIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldDescription, vs...))
}

// DescriptionNotIn applies the NotIn predicate on the "description" field.
func DescriptionNotIn(vs ...string) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldDescription, vs...))
}

// DescriptionGT applies the GT predicate on the "description" field.
func DescriptionGT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldDescription, v))
}

// DescriptionGTE applies the GTE predicate on the "description" field.
func DescriptionGTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldDescription, v))
}

// DescriptionLT applies the LT predicate on the "description" field.
func DescriptionLT(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldDescription, v))
}

// DescriptionLTE applies the LTE predicate on the "description" field.
func DescriptionLTE(v string) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldDescription, v))
}

// DescriptionContains applies the Contains predicate on the "description" field.
func DescriptionContains(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContains(FieldDescription, v))
}

// DescriptionHasPrefix applies the HasPrefix predicate on the "description" field.
func DescriptionHasPrefix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasPrefix(FieldDescription, v))
}

// DescriptionHasSuffix applies the HasSuffix predicate on the "description" field.
func DescriptionHasSuffix(v string) predicate.Todo {
	return predicate.Todo(sql.FieldHasSuffix(FieldDescription, v))
}

// DescriptionEqualFold applies the EqualFold predicate on the "description" field.
func DescriptionEqualFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldEqualFold(FieldDescription, v))
}

// DescriptionContainsFold applies the ContainsFold predicate on the "description" field.
func DescriptionContainsFold(v string) predicate.Todo {
	return predicate.Todo(sql.FieldContainsFold(FieldDescription, v))
}

// DoneEQ applies the EQ predicate on the "done" field.
func DoneEQ(v bool) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldDone, v))
}

// DoneNEQ applies the NEQ predicate on the "done" field.
func DoneNEQ(v bool) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldDone, v))
}

// CreatedAtEQ applies the EQ predicate on the "created_at" field.
func CreatedAtEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldCreatedAt, v))
}

// CreatedAtNEQ applies the NEQ predicate on the "created_at" field.
func CreatedAtNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldCreatedAt, v))
}

// CreatedAtIn applies the In predicate on the "created_at" field.
func CreatedAtIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldCreatedAt, vs...))
}

// CreatedAtNotIn applies the NotIn predicate on the "created_at" field.
func CreatedAtNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldCreatedAt, vs...))
}

// CreatedAtGT applies the GT predicate on the "created_at" field.
func CreatedAtGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldCreatedAt, v))
}

// CreatedAtGTE applies the GTE predicate on the "created_at" field.
func CreatedAtGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldCreatedAt, v))
}

// CreatedAtLT applies the LT predicate on the "created_at" field.
func CreatedAtLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldCreatedAt, v))
}

// CreatedAtLTE applies the LTE predicate on the "created_at" field.
func CreatedAtLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldCreatedAt, v))
}

// UpdatedAtEQ applies the EQ predicate on the "updated_at" field.
func UpdatedAtEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldEQ(FieldUpdatedAt, v))
}

// UpdatedAtNEQ applies the NEQ predicate on the "updated_at" field.
func UpdatedAtNEQ(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNEQ(FieldUpdatedAt, v))
}

// UpdatedAtIn applies the In predicate on the "updated_at" field.
func UpdatedAtIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldIn(FieldUpdatedAt, vs...))
}

// UpdatedAtNotIn applies the NotIn predicate on the "updated_at" field.
func UpdatedAtNotIn(vs ...time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldNotIn(FieldUpdatedAt, vs...))
}

// UpdatedAtGT applies the GT predicate on the "updated_at" field.
func UpdatedAtGT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGT(FieldUpdatedAt, v))
}

// UpdatedAtGTE applies the GTE predicate on the "updated_at" field.
func UpdatedAtGTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldGTE(FieldUpdatedAt, v))
}

// UpdatedAtLT applies the LT predicate on the "updated_at" field.
func UpdatedAtLT(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLT(FieldUpdatedAt, v))
}

// UpdatedAtLTE applies the LTE predicate on the "updated_at" field.
func UpdatedAtLTE(v time.Time) predicate.Todo {
	return predicate.Todo(sql.FieldLTE(FieldUpdatedAt, v))
}

// And groups predicates with the AND operator between them.
func And(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		s1 := s.Clone().SetP(nil)
		for _, p := range predicates {
			p(s1)
		}
		s.Where(s1.P())
	})
}

// Or groups predicates with the OR operator between them.
func Or(predicates ...predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
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
func Not(p predicate.Todo) predicate.Todo {
	return predicate.Todo(func(s *sql.Selector) {
		p(s.Not())
	})
}
