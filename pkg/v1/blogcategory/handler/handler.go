package handler

import (
	"pujaprabha/internal/interfaces"
	"pujaprabha/pkg/v1/auth/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.BlogCategoryUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, usecase interfaces.BlogCategoryUsecase) {
	handler := &handler{
		usecase: usecase,
	}
	publicBlogCategoryHandler := publicApp.Group("/blogcategory/")
	privateBlogCategoryHandler := privateApp.Group("/blogcategory")

	privateBlogCategoryHandler.Post("create/", middleware.AdminMiddleware(handler.CreateBlogCategory()))
	publicBlogCategoryHandler.Get("list/", handler.ListBlogCategory())
	privateBlogCategoryHandler.Patch("update/:id/", middleware.AdminMiddleware(handler.UpdateBlogCategory()))
	privateBlogCategoryHandler.Delete("delete/:id/", middleware.AdminMiddleware(handler.DeleteBlogCategory()))
}
