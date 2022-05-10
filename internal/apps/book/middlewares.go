package book

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/marlonmp/go-basic-library/ent"
	"github.com/marlonmp/go-basic-library/pkg/db"
)

func checkJSONBook(c *fiber.Ctx) error {

	_book := new(ent.Book)

	if err := c.BodyParser(_book); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// books validations

	c.Locals("book", _book)

	return c.Next()
}

func checkBookID(c *fiber.Ctx) error {
	paramBookID := c.Params("bookID")

	id, err := uuid.Parse(paramBookID)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	client := db.GetClient()

	exist, _ := client.Book.
		Query().
		Exist(context.Background())

	if !exist {
		return c.SendStatus(fiber.StatusNotFound)
	}

	c.Locals("bookID", id)

	return c.Next()
}
