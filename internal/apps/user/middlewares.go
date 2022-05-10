package user

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/marlonmp/go-basic-library/ent/user"
	"github.com/marlonmp/go-basic-library/pkg/db"
)

func checkJSONUser(c *fiber.Ctx) error {
	_user := new(UserSerializer)

	if err := c.BodyParser(_user); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	// user validations

	c.Locals("user", _user)

	return c.Next()
}

func checkUserID(c *fiber.Ctx) error {

	paramUserID := c.Params("userID")

	id, err := uuid.Parse(paramUserID)

	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	client := db.GetClient()

	exist, _ := client.User.
		Query().
		Where(user.ID(id)).
		Exist(context.Background())

	if !exist {
		return c.SendStatus(fiber.StatusNotFound)
	}

	c.Locals("userID", id)

	return c.Next()
}

func userCredentials(c *fiber.Ctx) error {

	return c.Next()
}