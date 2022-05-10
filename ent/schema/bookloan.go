package schema

import (
	"time"

	"entgo.io/ent"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// BookLoan holds the schema definition for the BookLoan entity.
type BookLoan struct {
	ent.Schema
}

func (BookLoan) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModel{},
	}
}

// Fields of the BookLoan.
func (BookLoan) Fields() []ent.Field {
	return []ent.Field{
		field.Time("loan_date").
			Default(time.Now).
			Immutable(),

		field.Time("delivery_date").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}

// Edges of the BookLoan.
func (BookLoan) Edges() []ent.Edge {
	return []ent.Edge{
		edge.From("user", User.Type).
			Ref("loans").
			Unique().
			Required(),

		edge.From("book", Book.Type).
			Ref("on_loans").
			Unique().
			Required(),
	}
}
