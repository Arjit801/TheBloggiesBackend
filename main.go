package main

import (
	"log"
	"os"

	"github.com/Arjit801/TheBloggies/database"
	"github.com/Arjit801/TheBloggies/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/joho/godotenv"
)

func main() {
	database.Connect()
	err := godotenv.Load()
	if err != nil {
		log.Fatal("error while loading .env file")
	}
	port := os.Getenv("PORT")
	app := fiber.New()
	// Apply the CORS middleware
	app.Use(cors.New(cors.Config{
		AllowOrigins: "https://main--thebloggiesfrontend.netlify.app", // specify the frontend domain
		AllowMethods: "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders: "Content-Type,Authorization",
	}))
	routes.Setup(app)
	app.Listen(":"+port)
}