package handler

import (
	"pujaprabha/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.OrderUsecase
}

func New(privateApp fiber.Router, usecase interfaces.OrderUsecase) {
	handler := &handler{
		usecase: usecase,
	}
	privateOrderHandler := privateApp.Group("/order/")

	privateOrderHandler.Post("create/", handler.CreateOrder())
	privateOrderHandler.Get("list/", handler.ListOrders())
	privateOrderHandler.Patch("update/:id/", handler.UpdateOrder())
	privateOrderHandler.Delete("delete/:id/", handler.DeleteOrder())
}
