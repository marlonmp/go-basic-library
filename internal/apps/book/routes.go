package book

import "github.com/gofiber/fiber/v2"

func RoutesUp(api fiber.Router) {

	booksGroup := api.Group("/books")

	booksGroup.Post("", checkJSONBook, bookCreate)
	booksGroup.Get("", bookList)

	idBooksGroup := booksGroup.Group("/:bookID", checkBookID)

	idBooksGroup.Get("", bookGetByID)
	idBooksGroup.Put("", checkJSONBook, bookPutByID)
	idBooksGroup.Patch("", checkJSONBook, bookPatchByID)
	idBooksGroup.Delete("", bookDeleteByID)
}
