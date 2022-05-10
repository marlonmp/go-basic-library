package schema

import (
	"entgo.io/ent"
	"entgo.io/ent/dialect/entsql"
	"entgo.io/ent/schema/edge"
	"entgo.io/ent/schema/field"
)

var Categories = []string{
	"Action",
	"Adventure",
	"Fiction",
	"Science",
}

// Book holds the schema definition for the Book entity.
type Book struct {
	ent.Schema
}

func (Book) Mixin() []ent.Mixin {
	return []ent.Mixin{
		BaseModel{},
	}
}

// Fields of the Book.
func (Book) Fields() []ent.Field {
	return []ent.Field{
		field.String("title").
			Annotations(entsql.Annotation{
				Size: 60,
			}),

		field.Uint16("pages"),

		field.Text("abstract").
			Optional(),

		field.Enum("category").
			Values(Categories...),

		field.String("cover_url"),
	}
}

// Edges of the Book.
func (Book) Edges() []ent.Edge {
	return []ent.Edge{
		edge.To("on_loans", BookLoan.Type).
			Annotations(entsql.Annotation{
				OnDelete: entsql.Restrict,
			}),

		edge.From("author", User.Type).
			Ref("books").
			Unique().
			Required(),
	}
}
