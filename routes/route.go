package routes

import (
	"github.com/Arjit801/TheBloggies/controller"
	"github.com/Arjit801/TheBloggies/middleware"
	"github.com/gofiber/fiber/v2"
)

func Setup(app *fiber.App) {
	app.Post("/api/register", controller.Register)
	app.Post("/api/login", controller.Login)
	app.Use(middleware.IsAuthenticated)
	app.Post("/api/post", controller.CreatPost)
	app.Get("/api/getallpost", controller.GetAllPost)
	app.Get("/api/getallpost/:id", controller.DetailPost)
	app.Put("/api/updatepost/:id", controller.UpdatePost)
	app.Get("/api/uniquepost/:id", controller.UniquePost)
	app.Delete("/api/deletepost/:id", controller.DeletePost)
	app.Post("/api/uploads-images", controller.Upload)
	app.Static("/api/uploads", "./uploads")
}