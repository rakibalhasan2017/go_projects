package handlers

import (
	"backend/models"

	"github.com/gofiber/fiber/v2"
)

func GetBooks(c *fiber.Ctx) error {
    books, err := models.GetAllBooks()
    if err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(books)
}

func AddBook(c *fiber.Ctx) error {
    book := new(models.Book)
    if err := c.BodyParser(book); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }

    if err := models.CreateBook(book); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.Status(201).JSON(book)
}


// Add these to backend/handlers/book_handler.go

func UpdateBook(c *fiber.Ctx) error {
    id, _ := c.ParamsInt("id")
    book := new(models.Book)
    if err := c.BodyParser(book); err != nil {
        return c.Status(400).JSON(fiber.Map{"error": "Cannot parse JSON"})
    }
    if err := models.UpdateBook(id, book); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.JSON(book)
}

func DeleteBook(c *fiber.Ctx) error {
    id, _ := c.ParamsInt("id")
    if err := models.DeleteBook(id); err != nil {
        return c.Status(500).JSON(fiber.Map{"error": err.Error()})
    }
    return c.SendStatus(204)
}