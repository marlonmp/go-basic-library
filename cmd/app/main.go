package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/marlonmp/go-basic-library/config"
	"github.com/marlonmp/go-basic-library/internal/apps/auth"
	"github.com/marlonmp/go-basic-library/internal/apps/book"
	"github.com/marlonmp/go-basic-library/internal/apps/user"
	"github.com/marlonmp/go-basic-library/pkg/db"
)

func routesUp(app *fiber.App) {
	api := app.Group("/api/v1")

	user.RoutesUp(api)

	book.RoutesUp(api)

	auth.RoutesUp(api)
}

func main() {
	entConfig, fiberConfig := config.Ent(), config.Fiber()

	db.Open(entConfig)

	client := db.GetClient()

	defer client.Close()

	// if err := client.Schema.Create(context.Background()); err != nil {
	// 	panic(err)
	// }

	app := fiber.New(fiberConfig.Config)

	routesUp(app)

	if err := app.Listen(fiberConfig.Host); err != nil {
		panic(err)
	}
}
