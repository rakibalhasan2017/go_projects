package main

import (
	"log"
	"os"

	"backend/database"
	"backend/handlers"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
    // Load .env for local dev, but don't fail in containerized/prod environments.
    if err := godotenv.Load(); err != nil {
        log.Println("No .env file found, reading environment variables from runtime")
    }

    // 2. Connect to Database
    database.Connect()

    // 3. Initialize Fiber App
    app := fiber.New()

    // 4. Middleware (CORS is vital so React can talk to this API)
    app.Use(cors.New())

    // 5. Routes
    app.Get("/api/books", handlers.GetBooks)
    app.Post("/api/books", handlers.AddBook)
    app.Put("/api/books/:id", handlers.UpdateBook)    // New
    app.Delete("/api/books/:id", handlers.DeleteBook) // New

    // 6. Start Server
    port := os.Getenv("PORT")
    if port == "" {
        port = "3000"
    }
    log.Fatal(app.Listen(":" + port))
}