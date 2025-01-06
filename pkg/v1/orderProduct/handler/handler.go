package handler

import (
	"pujaprabha/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.OrderProductUsecase
}

func New(privateApp fiber.Router, usecase interfaces.OrderProductUsecase) {
	handler := &handler{
		usecase: usecase,
	}
	orderProductHandler := privateApp.Group("/orderProduct/")

	orderProductHandler.Post("create/", handler.CreateOrderProduct())
	orderProductHandler.Get("list/", handler.ListOrderProducts())
	orderProductHandler.Delete("delete/:id/", handler.DeleteOrderProduct())
}
