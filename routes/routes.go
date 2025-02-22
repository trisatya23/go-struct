package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-struct/controller"
)

func SetupRoutes(app *fiber.App, controller *controller.CustomerController) {
	customerGroup := app.Group("/customers")
	customerGroup.Post("/", controller.CreateCustomer)
	customerGroup.Get("/", controller.GetAllCustomers)
	customerGroup.Get("/:id", controller.GetCustomerByID)
	customerGroup.Put("/:id", controller.UpdateCustomer)
	customerGroup.Delete("/:id", controller.DeleteCustomer)
}
