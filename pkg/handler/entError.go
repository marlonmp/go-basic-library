package handler

import (
	"log"

	"github.com/gofiber/fiber/v2"
	"github.com/lib/pq"
	"github.com/marlonmp/go-basic-library/ent"
)

func HandleEntError(c *fiber.Ctx, err error) error {

	var status int

	if ent.IsValidationError(err) {
		status = fiber.StatusBadRequest

	} else if ent.IsConstraintError(err) {
		status = fiber.StatusConflict

	} else if pqerr, ok := err.(*pq.Error); ok {

		switch pqerr.Code.Class() {
		case "20":
			status = fiber.StatusNotFound
			break
		case "23":
			status = fiber.StatusConflict
			break
		default:
			status = fiber.StatusInternalServerError
			break
		}
	} else {
		status = fiber.StatusInternalServerError
	}

	// switch {
	// case ent.IsConstraintError(err):
	// 	status = fiber.StatusConflict
	// 	break

	// case ent.IsNotFound(err), ent.IsNotLoaded(err):
	// 	status = fiber.StatusNotFound
	// 	break

	// case ent.IsValidationError(err):
	// 	status = fiber.StatusBadRequest
	// 	break

	// // case ent.IsNotSingular(err):
	// // 	break

	// default:
	// 	status = fiber.StatusInternalServerError
	// }

	log.Println("err:", err)

	return c.SendStatus(status)
}
