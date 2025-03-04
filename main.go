package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"go-struct/controller"
	"go-struct/database"
	"go-struct/repository"
	"go-struct/routes"
	"go-struct/service"
)

//TIP <p>To run your code, right-click the code and select <b>Run</b>.</p> <p>Alternatively, click
// the <icon src="AllIcons.Actions.Execute"/> icon in the gutter and select the <b>Run</b> menu item from here.</p>

func main() {
	app := fiber.New()

	database.ConnectDatabase()
	customerRepo := repository.NewCustomerRepository(database.DB)
	customerService := service.NewCustomerService(customerRepo)
	customerController := controller.NewCustomerController(customerService)

	orderRepo := repository.NewOrderRepository(database.DB)
	orderService := service.NewOrderService(orderRepo)
	orderController := controller.NewOrderController(orderService)

	orderItemRepo := repository.NewOrderItemRepository(database.DB)
	orderItemService := service.NewOrderItemService(orderItemRepo)
	orderItemController := controller.NewOrderItemController(orderItemService)

	discountRepo := repository.NewDiscountRepository(database.DB)
	discountService := service.NewDiscountService(discountRepo)
	discountController := controller.NewDiscountController(discountService)

	employeeRepo := repository.NewEmployeeRepository(database.DB)
	employeeService := service.NewEmployeeService(employeeRepo)
	employeeController := controller.NewEmployeeController(employeeService)

	routes.SetupRoutes(
		app,
		customerController,
		orderController,
		orderItemController,
		discountController,
		employeeController)

	err := app.Listen(":8000")
	if err != nil {
		fmt.Println(err)
	}
}
