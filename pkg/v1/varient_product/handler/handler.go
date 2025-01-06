package handler

import (
	"pujaprabha/internal/interfaces"
	"pujaprabha/pkg/v1/auth/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.VarientProductUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, usecase interfaces.VarientProductUsecase) {
	handler := &handler{
		usecase: usecase,
	}
	publicVarientProductHandler := publicApp.Group("/varient/")
	privateVarientProductHandler := privateApp.Group("/varient/")

	privateVarientProductHandler.Post("create/", middleware.AdminMiddleware(handler.CreateVarientProduct()))
	publicVarientProductHandler.Get("list/", handler.ListVarient())
	privateVarientProductHandler.Patch("update/:id/", middleware.AdminMiddleware(handler.UpdatedVarientProduct()))
	privateVarientProductHandler.Delete("delete/:id/", middleware.AdminMiddleware(handler.DeleteVarientProduct()))
}
