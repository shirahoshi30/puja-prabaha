package handler

import (
	"pujaprabha/internal/interfaces"
	"pujaprabha/pkg/v1/auth/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.CategoryUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, usecase interfaces.CategoryUsecase) {
	handler := &handler{
		usecase: usecase,
	}

	publicCategoryHandler := publicApp.Group("/category/")
	privateCategoryHandler := privateApp.Group("/category/")

	privateCategoryHandler.Post("create/", middleware.AdminMiddleware(handler.CreateCategory()))
	publicCategoryHandler.Get("list/", handler.ListCategory())
	privateCategoryHandler.Patch("update/:id/", middleware.AdminMiddleware(handler.UpdatedCategory()))
	privateCategoryHandler.Delete("delete/:id/", middleware.AdminMiddleware(handler.DeleteCategory()))

}
