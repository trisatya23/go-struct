package routes

import (
	"github.com/gofiber/fiber/v2"
	"go-struct/controller"
)

func SetupRoutes(app *fiber.App,
	customercontroller *controller.CustomerController,
	ordercontroller *controller.OrderController,
	orderItemcontroller *controller.OrderItemController,
	discountcontroller *controller.DiscountController,
	employeecontroller *controller.EmployeeController) {

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

	orderItemGroup := app.Group("/orderItems")
	orderItemGroup.Post("/", orderItemcontroller.CreateOrderItem)
	orderItemGroup.Get("/", orderItemcontroller.GetAllOrderItems)
	orderItemGroup.Get("/:id", orderItemcontroller.GetOrderItemByID)
	orderItemGroup.Put("/:id", orderItemcontroller.UpdateGetOrderItem)
	orderItemGroup.Delete("/:id", orderItemcontroller.DeleteOrderItem)

	discountItemGroup := app.Group("/discounts")
	discountItemGroup.Post("/", discountcontroller.CreateDiscount)
	discountItemGroup.Get("/", discountcontroller.GetAllDiscounts)
	discountItemGroup.Get("/:id", discountcontroller.GetDiscountByID)
	discountItemGroup.Put("/:id", discountcontroller.UpdateDiscount)
	discountItemGroup.Delete("/:id", discountcontroller.DeleteDiscount)

	employeeItemGroup := app.Group("/discounts")
	employeeItemGroup.Post("/", employeecontroller.CreateEmployee)
	employeeItemGroup.Get("/", employeecontroller.GetAllEmployees)
	employeeItemGroup.Get("/:id", employeecontroller.GetEmployeesByID)
	employeeItemGroup.Put("/:id", employeecontroller.UpdateEmployee)
	employeeItemGroup.Delete("/:id", employeecontroller.DeleteEmployee)

}
