// Code generated by entc, DO NOT EDIT.

package book

import (
	"fmt"
	"time"

	"github.com/google/uuid"
)

const (
	// Label holds the string label denoting the book type in the database.
	Label = "book"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldCreateAt holds the string denoting the create_at field in the database.
	FieldCreateAt = "create_at"
	// FieldUpdatedAt holds the string denoting the updated_at field in the database.
	FieldUpdatedAt = "updated_at"
	// FieldTitle holds the string denoting the title field in the database.
	FieldTitle = "title"
	// FieldPages holds the string denoting the pages field in the database.
	FieldPages = "pages"
	// FieldAbstract holds the string denoting the abstract field in the database.
	FieldAbstract = "abstract"
	// FieldCategory holds the string denoting the category field in the database.
	FieldCategory = "category"
	// FieldCoverURL holds the string denoting the cover_url field in the database.
	FieldCoverURL = "cover_url"
	// EdgeOnLoans holds the string denoting the on_loans edge name in mutations.
	EdgeOnLoans = "on_loans"
	// EdgeAuthor holds the string denoting the author edge name in mutations.
	EdgeAuthor = "author"
	// Table holds the table name of the book in the database.
	Table = "books"
	// OnLoansTable is the table that holds the on_loans relation/edge.
	OnLoansTable = "book_loans"
	// OnLoansInverseTable is the table name for the BookLoan entity.
	// It exists in this package in order to avoid circular dependency with the "bookloan" package.
	OnLoansInverseTable = "book_loans"
	// OnLoansColumn is the table column denoting the on_loans relation/edge.
	OnLoansColumn = "book_on_loans"
	// AuthorTable is the table that holds the author relation/edge.
	AuthorTable = "books"
	// AuthorInverseTable is the table name for the User entity.
	// It exists in this package in order to avoid circular dependency with the "user" package.
	AuthorInverseTable = "users"
	// AuthorColumn is the table column denoting the author relation/edge.
	AuthorColumn = "user_books"
)

// Columns holds all SQL columns for book fields.
var Columns = []string{
	FieldID,
	FieldCreateAt,
	FieldUpdatedAt,
	FieldTitle,
	FieldPages,
	FieldAbstract,
	FieldCategory,
	FieldCoverURL,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "books"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"user_books",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultCreateAt holds the default value on creation for the "create_at" field.
	DefaultCreateAt func() time.Time
	// DefaultUpdatedAt holds the default value on creation for the "updated_at" field.
	DefaultUpdatedAt func() time.Time
	// UpdateDefaultUpdatedAt holds the default value on update for the "updated_at" field.
	UpdateDefaultUpdatedAt func() time.Time
	// DefaultID holds the default value on creation for the "id" field.
	DefaultID func() uuid.UUID
)

// Category defines the type for the "category" enum field.
type Category string

// Category values.
const (
	CategoryAction    Category = "Action"
	CategoryAdventure Category = "Adventure"
	CategoryFiction   Category = "Fiction"
	CategoryScience   Category = "Science"
)

func (c Category) String() string {
	return string(c)
}

// CategoryValidator is a validator for the "category" field enum values. It is called by the builders before save.
func CategoryValidator(c Category) error {
	switch c {
	case CategoryAction, CategoryAdventure, CategoryFiction, CategoryScience:
		return nil
	default:
		return fmt.Errorf("book: invalid enum value for category field: %q", c)
	}
}
