package user

import (
	"github.com/gofiber/fiber/v2"
)

func RoutesUp(api fiber.Router) {

	userGroup := api.Group("/users")

	userGroup.Post("", checkJSONUser, userCreate)
	userGroup.Get("", userList)

	userGroupID := userGroup.Group("/:userID", checkUserID)

	userGroupID.Get("", userGetByID)

	userGroupID.Put("", checkJSONUser, userPutByID)
	userGroupID.Patch("", checkJSONUser, userPatchByID)

	userGroupID.Delete("", userDeleteByID)

	userGroupID.Get("/books", userBooksList)

	userGroupID.Get("/loans", userLoansList)
}
