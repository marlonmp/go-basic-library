// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/marlonmp/go-basic-library/ent/book"
	"github.com/marlonmp/go-basic-library/ent/user"
)

// Book is the model entity for the Book schema.
type Book struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// Title holds the value of the "title" field.
	Title string `json:"title,omitempty"`
	// Pages holds the value of the "pages" field.
	Pages uint16 `json:"pages,omitempty"`
	// Abstract holds the value of the "abstract" field.
	Abstract string `json:"abstract,omitempty"`
	// Category holds the value of the "category" field.
	Category book.Category `json:"category,omitempty"`
	// CoverURL holds the value of the "cover_url" field.
	CoverURL string `json:"cover_url,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookQuery when eager-loading is set.
	Edges      BookEdges `json:"edges"`
	user_books *uuid.UUID
}

// BookEdges holds the relations/edges for other nodes in the graph.
type BookEdges struct {
	// OnLoans holds the value of the on_loans edge.
	OnLoans []*BookLoan `json:"on_loans,omitempty"`
	// Author holds the value of the author edge.
	Author *User `json:"author,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// OnLoansOrErr returns the OnLoans value or an error if the edge
// was not loaded in eager-loading.
func (e BookEdges) OnLoansOrErr() ([]*BookLoan, error) {
	if e.loadedTypes[0] {
		return e.OnLoans, nil
	}
	return nil, &NotLoadedError{edge: "on_loans"}
}

// AuthorOrErr returns the Author value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookEdges) AuthorOrErr() (*User, error) {
	if e.loadedTypes[1] {
		if e.Author == nil {
			// The edge author was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.Author, nil
	}
	return nil, &NotLoadedError{edge: "author"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Book) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case book.FieldPages:
			values[i] = new(sql.NullInt64)
		case book.FieldTitle, book.FieldAbstract, book.FieldCategory, book.FieldCoverURL:
			values[i] = new(sql.NullString)
		case book.FieldCreateAt, book.FieldUpdatedAt:
			values[i] = new(sql.NullTime)
		case book.FieldID:
			values[i] = new(uuid.UUID)
		case book.ForeignKeys[0]: // user_books
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type Book", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Book fields.
func (b *Book) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case book.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				b.ID = *value
			}
		case book.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				b.CreateAt = value.Time
			}
		case book.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				b.UpdatedAt = value.Time
			}
		case book.FieldTitle:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field title", values[i])
			} else if value.Valid {
				b.Title = value.String
			}
		case book.FieldPages:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for field pages", values[i])
			} else if value.Valid {
				b.Pages = uint16(value.Int64)
			}
		case book.FieldAbstract:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field abstract", values[i])
			} else if value.Valid {
				b.Abstract = value.String
			}
		case book.FieldCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field category", values[i])
			} else if value.Valid {
				b.Category = book.Category(value.String)
			}
		case book.FieldCoverURL:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field cover_url", values[i])
			} else if value.Valid {
				b.CoverURL = value.String
			}
		case book.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_books", values[i])
			} else if value.Valid {
				b.user_books = new(uuid.UUID)
				*b.user_books = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryOnLoans queries the "on_loans" edge of the Book entity.
func (b *Book) QueryOnLoans() *BookLoanQuery {
	return (&BookClient{config: b.config}).QueryOnLoans(b)
}

// QueryAuthor queries the "author" edge of the Book entity.
func (b *Book) QueryAuthor() *UserQuery {
	return (&BookClient{config: b.config}).QueryAuthor(b)
}

// Update returns a builder for updating this Book.
// Note that you need to call Book.Unwrap() before calling this method if this Book
// was returned from a transaction, and the transaction was committed or rolled back.
func (b *Book) Update() *BookUpdateOne {
	return (&BookClient{config: b.config}).UpdateOne(b)
}

// Unwrap unwraps the Book entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (b *Book) Unwrap() *Book {
	tx, ok := b.config.driver.(*txDriver)
	if !ok {
		panic("ent: Book is not a transactional entity")
	}
	b.config.driver = tx.drv
	return b
}

// String implements the fmt.Stringer.
func (b *Book) String() string {
	var builder strings.Builder
	builder.WriteString("Book(")
	builder.WriteString(fmt.Sprintf("id=%v", b.ID))
	builder.WriteString(", create_at=")
	builder.WriteString(b.CreateAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(b.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", title=")
	builder.WriteString(b.Title)
	builder.WriteString(", pages=")
	builder.WriteString(fmt.Sprintf("%v", b.Pages))
	builder.WriteString(", abstract=")
	builder.WriteString(b.Abstract)
	builder.WriteString(", category=")
	builder.WriteString(fmt.Sprintf("%v", b.Category))
	builder.WriteString(", cover_url=")
	builder.WriteString(b.CoverURL)
	builder.WriteByte(')')
	return builder.String()
}

// Books is a parsable slice of Book.
type Books []*Book

func (b Books) config(cfg config) {
	for _i := range b {
		b[_i].config = cfg
	}
}
