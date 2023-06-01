package main

import (
	"github.com/gofiber/fiber/v2"
	"github.com/somnidev/go-fiber/model"
	"github.com/somnidev/go-fiber/services"
)

var (
	bookService *services.BookService
)

func GetBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	b, found := bookService.GetBookById(id)
	if !found {
		c.Status(fiber.StatusNotFound)
		return nil
	}
	return c.Status(fiber.StatusOK).JSON(b)
}

func GetBooks(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).JSON(bookService.ListBooks())
}

func CreateBook(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	nb := bookService.CreateBook(*b)

	location, _ := c.GetRouteURL("books.id", fiber.Map{"id": nb.ID})
	c.Location(location)
	c.Status(fiber.StatusCreated)
	return nil
}

func DeleteBookById(c *fiber.Ctx) error {
	id := c.Params("id")
	bookService.DeleteBookById(id)
	c.Status(fiber.StatusNoContent)
	return nil
}

func UpdateBookById(c *fiber.Ctx) error {
	b := new(model.Book)
	if err := c.BodyParser(b); err != nil {
		return err
	}
	id := c.Params("id")
	bookService.UpdateBookById(id, *b)
	c.Status(fiber.StatusNoContent)
	return nil
}

func main() {
	app := fiber.New()
	bookService, _ = services.NewBookService()

	app.Get("/books/:id", GetBookById).Name("books.id")
	app.Get("/books", GetBooks)
	app.Post("/books", CreateBook)
	app.Delete("/books/:id", DeleteBookById)
	app.Put("/books/:id", UpdateBookById)
	app.Listen(":3000")
}
