package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-struct/controller"
)

func SetupRoutes(app *fiber.App,
	customercontroller *controller.CustomerController,
	ordercontroller *controller.OrderController) {
	//Customer
	customerGroup := app.Group("/customers")
	customerGroup.Post("/", customercontroller.CreateCustomer)
	customerGroup.Get("/", customercontroller.GetAllCustomers)
	customerGroup.Get("/:id", customercontroller.GetCustomerByID)
	customerGroup.Put("/:id", customercontroller.UpdateCustomer)
	customerGroup.Delete("/:id", customercontroller.DeleteCustomer)

	//Order
	orderGroup := app.Group("/orders")
	orderGroup.Post("/", ordercontroller.CreateOrder)
	orderGroup.Get("/", ordercontroller.GetAllOrders)
	orderGroup.Get("/:id", ordercontroller.GetOrderByID)
	orderGroup.Put("/:id", ordercontroller.UpdateOrder)
	orderGroup.Delete("/:id", ordercontroller.DeleteOrder)

}
