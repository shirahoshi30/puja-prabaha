package handler

import (
	"pujaprabha/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.CartUsecase
}

func New(privateApp fiber.Router, usecase interfaces.CartUsecase) {
	handler := &handler{
		usecase: usecase,
	}
	privateCartApp := privateApp.Group("/cart/")

	privateCartApp.Post("create/", handler.AddToCart())
	privateCartApp.Get("list/", handler.ListCart())
	privateCartApp.Patch("update/:id/", handler.UpdateCart())
	privateCartApp.Delete("delete/:id/", handler.RemoveCart())
}
