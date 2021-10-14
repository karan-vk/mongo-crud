package routes

import (
	"github.com/gofiber/fiber/v2"
	"github.com/karan-vk/mongo-crud/controllers"
)

func AuthRoute(route fiber.Router) {
	route.Get("/", controllers.GetToken)
	route.Post("/", controllers.AddUser)
}
