package book

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/marlonmp/go-basic-library/ent"
	"github.com/marlonmp/go-basic-library/ent/book"
	"github.com/marlonmp/go-basic-library/pkg/db"
	"github.com/marlonmp/go-basic-library/pkg/handler"
)

func bookCreate(c *fiber.Ctx) error {

	_book := c.Locals("book").(*ent.Book)

	client := db.GetClient()

	_book, err := client.Book.Create().
		SetTitle(_book.Title).
		SetNillableAbstract(&_book.Abstract).
		SetAuthorID(_book.Edges.Author.ID).
		SetCategory(_book.Category).
		SetPages(_book.Pages).
		SetCoverURL(_book.CoverURL).
		Save(context.Background())

	if err == nil {
		return c.Status(fiber.StatusCreated).JSON(_book)
	}

	return handler.HandleEntError(c, err)
}

func bookList(c *fiber.Ctx) error {

	client := db.GetClient()

	books, err := client.Book.
		Query().
		WithAuthor().
		All(context.Background())

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(books)
	}

	return handler.HandleEntError(c, err)
}

func bookGetByID(c *fiber.Ctx) error {

	id := c.Locals("bookID").(uuid.UUID)

	client := db.GetClient()

	_book, err := client.Book.
		Query().
		Where(book.ID(id)).
		First(context.Background())

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(_book)
	}

	return handler.HandleEntError(c, err)
}

func bookPutByID(c *fiber.Ctx) error {

	id := c.Locals("bookID").(uuid.UUID)

	_book := c.Locals("book").(*ent.Book)

	client := db.GetClient()

	_book, err := client.Book.UpdateOneID(id).
		SetTitle(_book.Title).
		SetNillableAbstract(&_book.Abstract).
		SetAuthorID(_book.Edges.Author.ID).
		SetCategory(_book.Category).
		SetPages(_book.Pages).
		SetCoverURL(_book.CoverURL).
		Save(context.Background())

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(_book)
	}

	return handler.HandleEntError(c, err)
}

func bookPatchByID(c *fiber.Ctx) error {

	id := c.Locals("bookID").(uuid.UUID)

	_book := c.Locals("book").(*ent.Book)

	client := db.GetClient()

	updateQuery := client.Book.UpdateOneID(id)

	if _book.Title != "" {
		updateQuery.SetTitle(_book.Title)
	}
	if _book.Abstract != "" {
		updateQuery.SetNillableAbstract(&_book.Abstract)
	}
	if _book.Category != "" {
		updateQuery.SetCategory(_book.Category)
	}
	if _book.Pages > 0 {
		updateQuery.SetPages(_book.Pages)
	}
	if _book.CoverURL != "" {
		updateQuery.SetCoverURL(_book.CoverURL)
	}

	_book, err := updateQuery.Save(context.Background())

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(_book)
	}

	return handler.HandleEntError(c, err)
}

func bookDeleteByID(c *fiber.Ctx) error {

	id := c.Locals("bookID").(uuid.UUID)

	client := db.GetClient()

	err := client.Book.
		DeleteOneID(id).
		Exec(context.Background())

	if err == nil {
		return c.SendStatus(fiber.StatusOK)
	}

	return handler.HandleEntError(c, err)
}
