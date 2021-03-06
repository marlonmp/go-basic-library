// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"
	"time"

	"entgo.io/ent/dialect/sql"
	"github.com/google/uuid"
	"github.com/marlonmp/go-basic-library/ent/book"
	"github.com/marlonmp/go-basic-library/ent/bookloan"
	"github.com/marlonmp/go-basic-library/ent/user"
)

// BookLoan is the model entity for the BookLoan schema.
type BookLoan struct {
	config `json:"-"`
	// ID of the ent.
	ID uuid.UUID `json:"id,omitempty"`
	// CreateAt holds the value of the "create_at" field.
	CreateAt time.Time `json:"create_at,omitempty"`
	// UpdatedAt holds the value of the "updated_at" field.
	UpdatedAt time.Time `json:"updated_at,omitempty"`
	// LoanDate holds the value of the "loan_date" field.
	LoanDate time.Time `json:"loan_date,omitempty"`
	// DeliveryDate holds the value of the "delivery_date" field.
	DeliveryDate time.Time `json:"delivery_date,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the BookLoanQuery when eager-loading is set.
	Edges         BookLoanEdges `json:"edges"`
	book_on_loans *uuid.UUID
	user_loans    *uuid.UUID
}

// BookLoanEdges holds the relations/edges for other nodes in the graph.
type BookLoanEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// Book holds the value of the book edge.
	Book *Book `json:"book,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [2]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookLoanEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// BookOrErr returns the Book value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e BookLoanEdges) BookOrErr() (*Book, error) {
	if e.loadedTypes[1] {
		if e.Book == nil {
			// The edge book was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: book.Label}
		}
		return e.Book, nil
	}
	return nil, &NotLoadedError{edge: "book"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*BookLoan) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case bookloan.FieldCreateAt, bookloan.FieldUpdatedAt, bookloan.FieldLoanDate, bookloan.FieldDeliveryDate:
			values[i] = new(sql.NullTime)
		case bookloan.FieldID:
			values[i] = new(uuid.UUID)
		case bookloan.ForeignKeys[0]: // book_on_loans
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		case bookloan.ForeignKeys[1]: // user_loans
			values[i] = &sql.NullScanner{S: new(uuid.UUID)}
		default:
			return nil, fmt.Errorf("unexpected column %q for type BookLoan", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the BookLoan fields.
func (bl *BookLoan) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case bookloan.FieldID:
			if value, ok := values[i].(*uuid.UUID); !ok {
				return fmt.Errorf("unexpected type %T for field id", values[i])
			} else if value != nil {
				bl.ID = *value
			}
		case bookloan.FieldCreateAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field create_at", values[i])
			} else if value.Valid {
				bl.CreateAt = value.Time
			}
		case bookloan.FieldUpdatedAt:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field updated_at", values[i])
			} else if value.Valid {
				bl.UpdatedAt = value.Time
			}
		case bookloan.FieldLoanDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field loan_date", values[i])
			} else if value.Valid {
				bl.LoanDate = value.Time
			}
		case bookloan.FieldDeliveryDate:
			if value, ok := values[i].(*sql.NullTime); !ok {
				return fmt.Errorf("unexpected type %T for field delivery_date", values[i])
			} else if value.Valid {
				bl.DeliveryDate = value.Time
			}
		case bookloan.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field book_on_loans", values[i])
			} else if value.Valid {
				bl.book_on_loans = new(uuid.UUID)
				*bl.book_on_loans = *value.S.(*uuid.UUID)
			}
		case bookloan.ForeignKeys[1]:
			if value, ok := values[i].(*sql.NullScanner); !ok {
				return fmt.Errorf("unexpected type %T for field user_loans", values[i])
			} else if value.Valid {
				bl.user_loans = new(uuid.UUID)
				*bl.user_loans = *value.S.(*uuid.UUID)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the BookLoan entity.
func (bl *BookLoan) QueryUser() *UserQuery {
	return (&BookLoanClient{config: bl.config}).QueryUser(bl)
}

// QueryBook queries the "book" edge of the BookLoan entity.
func (bl *BookLoan) QueryBook() *BookQuery {
	return (&BookLoanClient{config: bl.config}).QueryBook(bl)
}

// Update returns a builder for updating this BookLoan.
// Note that you need to call BookLoan.Unwrap() before calling this method if this BookLoan
// was returned from a transaction, and the transaction was committed or rolled back.
func (bl *BookLoan) Update() *BookLoanUpdateOne {
	return (&BookLoanClient{config: bl.config}).UpdateOne(bl)
}

// Unwrap unwraps the BookLoan entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (bl *BookLoan) Unwrap() *BookLoan {
	tx, ok := bl.config.driver.(*txDriver)
	if !ok {
		panic("ent: BookLoan is not a transactional entity")
	}
	bl.config.driver = tx.drv
	return bl
}

// String implements the fmt.Stringer.
func (bl *BookLoan) String() string {
	var builder strings.Builder
	builder.WriteString("BookLoan(")
	builder.WriteString(fmt.Sprintf("id=%v", bl.ID))
	builder.WriteString(", create_at=")
	builder.WriteString(bl.CreateAt.Format(time.ANSIC))
	builder.WriteString(", updated_at=")
	builder.WriteString(bl.UpdatedAt.Format(time.ANSIC))
	builder.WriteString(", loan_date=")
	builder.WriteString(bl.LoanDate.Format(time.ANSIC))
	builder.WriteString(", delivery_date=")
	builder.WriteString(bl.DeliveryDate.Format(time.ANSIC))
	builder.WriteByte(')')
	return builder.String()
}

// BookLoans is a parsable slice of BookLoan.
type BookLoans []*BookLoan

func (bl BookLoans) config(cfg config) {
	for _i := range bl {
		bl[_i].config = cfg
	}
}
