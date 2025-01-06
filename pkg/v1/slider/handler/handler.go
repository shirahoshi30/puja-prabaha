package handler

import (
	"pujaprabha/internal/interfaces"
	"pujaprabha/pkg/v1/auth/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.SliderUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, usecase interfaces.SliderUsecase) {
	handler := &handler{
		usecase: usecase,
	}
	publicSliderHandler := publicApp.Group("/slider/")
	privateSliderHandler := privateApp.Group("/slider/")

	privateSliderHandler.Post("create/", middleware.AdminMiddleware(handler.CreateSlider()))
	privateSliderHandler.Patch("update/:id/", middleware.AdminMiddleware(handler.UpdateSlider()))
	privateSliderHandler.Delete("delete/:id/", middleware.AdminMiddleware(handler.DeleteSlider()))
	publicSliderHandler.Get("list/", handler.ListSlider())

}
