package auth

import "github.com/gofiber/fiber/v2"

func JWTGenerate(c *fiber.Ctx) error {

	subject := c.Locals("userCredentials").(*UserCredentials).Subject

	atr := NewAccesTokenResponse(subject)

	return c.
		Status(fiber.StatusOK).
		JSON(atr)
}
