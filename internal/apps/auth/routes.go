package auth

import "github.com/gofiber/fiber/v2"

func RoutesUp(api fiber.Router) {

	authGroup := api.Group("/auth")

	tokenGroup := authGroup.Group("/token")

	tokenGroup.Post("", getUserCredentials, JWTGenerate)
	// sessionGroup.Delete("")
	tokenGroup.Post("/test", JWTVerify, getClaims)
}

func getClaims(c *fiber.Ctx) error {

	claims := c.Locals("claims")

	return c.JSON(claims)
}
