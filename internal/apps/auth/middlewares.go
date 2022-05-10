package auth

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/marlonmp/go-basic-library/ent/user"
	"github.com/marlonmp/go-basic-library/pkg/db"
	"github.com/marlonmp/go-basic-library/pkg/jwt"
	"golang.org/x/crypto/bcrypt"
)

func compareHash(hash, password string) bool {

	hashBytes, passwordBytes := []byte(hash), []byte(password)

	err := bcrypt.CompareHashAndPassword(hashBytes, passwordBytes)

	return err == nil
}

func getUserCredentials(c *fiber.Ctx) error {

	userCredentials := new(UserCredentials)

	if err := c.BodyParser(userCredentials); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	client := db.GetClient()

	_user, _ := client.User.
		Query().
		Select(
			user.FieldID,
			user.FieldPassword,
		).
		Where(
			user.Email(userCredentials.Subject),
		).
		First(context.Background())

	if _user == nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	if !compareHash(_user.Password, userCredentials.Password) {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	userCredentials.Password = ""

	c.Locals("userCredentials", userCredentials)

	return c.Next()
}

func JWTVerify(c *fiber.Ctx) error {
	tokenBytes := c.Request().Header.Peek(fiber.HeaderAuthorization)

	if len(tokenBytes) < 5 {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	authType := string(tokenBytes[:3])

	if authType != jwt.TokenType {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	tokenStr := string(tokenBytes[4:])

	claims, err := jwt.Verify(tokenStr)

	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	c.Locals("claims", claims)

	return c.Next()
}
