package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karan-vk/mongo-crud/controllers"
	"github.com/karan-vk/mongo-crud/middleware"
)

func UsersRoute(route fiber.Router) {

	route.Get("/", controllers.GetAllUsers)
	route.Get("/:id", controllers.GetUser)
	route.Post("/", controllers.AddUser)
	route.Put("/:id", middleware.JWTProtected(), controllers.UpdateUser)
	route.Delete("/:id", middleware.JWTProtected(), controllers.DeleteUser)
}
