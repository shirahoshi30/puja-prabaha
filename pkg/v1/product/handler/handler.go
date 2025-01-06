package handler

import (
	"pujaprabha/internal/interfaces"
	"pujaprabha/pkg/v1/auth/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.ProductUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, usecase interfaces.ProductUsecase) {
	handler := &handler{
		usecase: usecase,
	}
	publicProductHandler := publicApp.Group("/product/")
	privateProductHandler := privateApp.Group("/product/")

	publicProductHandler.Get("list/", handler.ListAllProduct())
	publicProductHandler.Get("details/:id", handler.DetailsOfProduct())
	privateProductHandler.Post("create/", middleware.AdminMiddleware(handler.CreateProduct()))
	privateProductHandler.Patch("update/:id/", middleware.AdminMiddleware(handler.ProductUpdate()))
	privateProductHandler.Delete("/delete/:id/", middleware.AdminMiddleware(handler.DeleteProduct()))
	publicProductHandler.Get("compare/", (handler.CompareProducts()))
}
