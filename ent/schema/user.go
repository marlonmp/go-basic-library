package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

// User holds the schema definition for the User entity.
type User struct {
	ent.Schema
}

func (User) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModel{},
	}
}

// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("first_name").
			Annotations(entsql.Annotation{
				Size: 45,
			}),

		field.String("last_name").
			Annotations(entsql.Annotation{
				Size: 45,
			}),

		field.String("email").
			Annotations(entsql.Annotation{
				Size: 60,
			}).
			Unique(),

		field.String("password").
			Sensitive(),
	}
}

// Edges of the User.
func (User) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("books", Book.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),

		edge.To("loans", BookLoan.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),
	}
}
