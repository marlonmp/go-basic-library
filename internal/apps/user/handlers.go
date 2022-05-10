package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/marlonmp/go-basic-library/ent/user"
	"github.com/marlonmp/go-basic-library/pkg/db"
	"github.com/marlonmp/go-basic-library/pkg/handler"
	"golang.org/x/crypto/bcrypt"
)

func hashUserPassword(password string) string {
	passwordBytes := []byte(password)

	hashBytes, err := bcrypt.GenerateFromPassword(passwordBytes, bcrypt.MinCost)

	if err != nil {
		panic(err)
	}

	return string(hashBytes)
}

func userCreate(c *fiber.Ctx) error {
	userSerializer := c.Locals("user").(*UserSerializer)

	client := db.GetClient()

	password := hashUserPassword(userSerializer.Password)

	_user, err := client.User.Create().
		SetFirstName(userSerializer.FirstName).
		SetLastName(userSerializer.LastName).
		SetEmail(userSerializer.Email).
		SetPassword(password).
		Save(context.Background())

	if err == nil {
		return c.Status(fiber.StatusCreated).JSON(_user)
	}

	return handler.HandleEntError(c, err)
}

func userList(c *fiber.Ctx) error {
	client := db.GetClient()

	users, err := client.User.Query().All(context.Background())

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(users)
	}

	return handler.HandleEntError(c, err)
}

func userGetByID(c *fiber.Ctx) error {
	client := db.GetClient()

	id := c.Locals("userID").(uuid.UUID)

	_user, err := client.User.
		Query().
		Where(user.ID(id)).
		First(context.Background())

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(_user)
	}

	return handler.HandleEntError(c, err)
}

func userPutByID(c *fiber.Ctx) error {

	userSerializer := c.Locals("user").(*UserSerializer)

	id := c.Locals("userID").(uuid.UUID)

	client := db.GetClient()

	password := hashUserPassword(userSerializer.Password)

	_user, err := client.User.UpdateOneID(id).
		SetFirstName(userSerializer.FirstName).
		SetLastName(userSerializer.LastName).
		SetEmail(userSerializer.Email).
		SetPassword(password).
		Save(context.Background())

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(_user)
	}

	return handler.HandleEntError(c, err)
}

func userPatchByID(c *fiber.Ctx) error {

	userSerializer := c.Locals("user").(*UserSerializer)

	id := c.Locals("userID").(uuid.UUID)

	client := db.GetClient()

	updateQuery := client.User.UpdateOneID(id)

	if userSerializer.FirstName != "" {
		updateQuery.SetFirstName(userSerializer.FirstName)
	}
	if userSerializer.LastName != "" {
		updateQuery.SetLastName(userSerializer.LastName)
	}
	if userSerializer.Email != "" {
		updateQuery.SetEmail(userSerializer.Email)
	}
	if userSerializer.Password != "" {
		password := hashUserPassword(userSerializer.Password)
		updateQuery.SetPassword(password)
	}

	_user, err := updateQuery.Save(context.Background())

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(_user)
	}

	return handler.HandleEntError(c, err)
}

func userDeleteByID(c *fiber.Ctx) error {
	id := c.Locals("userID").(uuid.UUID)

	client := db.GetClient()

	err := client.User.DeleteOneID(id).
		Exec(context.Background())

	if err == nil {
		return c.SendStatus(fiber.StatusOK)
	}

	return handler.HandleEntError(c, err)
}

func userBooksList(c *fiber.Ctx) error {

	client := db.GetClient()

	id := c.Locals("userID").(uuid.UUID)

	ctx := context.Background()

	userBooks, err := client.User.
		Query().
		Where(user.ID(id)).
		QueryBooks().
		All(ctx)

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(userBooks)
	}

	return handler.HandleEntError(c, err)
}

func userLoansList(c *fiber.Ctx) error {

	client := db.GetClient()

	id := c.Locals("userID").(uuid.UUID)

	userLoans, err := client.User.
		Query().
		Where(user.ID(id)).
		QueryLoans().
		All(context.Background())

	if err == nil {
		return c.Status(fiber.StatusOK).JSON(userLoans)
	}

	return handler.HandleEntError(c, err)
}
