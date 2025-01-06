package handler

import (
	"pujaprabha/internal/interfaces"
	"pujaprabha/pkg/v1/auth/middleware"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.BlogUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, usecase interfaces.BlogUsecase) {
	handler := &handler{
		usecase: usecase,
	}

	publicBlogHandler := publicApp.Group("/blog/")
	privateBlogHandler := privateApp.Group("/blog/")

	privateBlogHandler.Post("create/", middleware.AdminMiddleware(handler.CreateBlog()))
	publicBlogHandler.Get("list/", handler.ListBlog())
	publicBlogHandler.Get("details/:id/", handler.DetailsOfBlog())
	privateBlogHandler.Patch("update/:id/", middleware.AdminMiddleware(handler.UpdateBlog()))
	privateBlogHandler.Delete("delete/:id/", middleware.AdminMiddleware(handler.DeleteBlog()))
}
