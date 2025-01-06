package handler

import (
	"pujaprabha/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.ReviewUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, usecase interfaces.ReviewUsecase) {
	handler := &handler{
		usecase: usecase,
	}
	publicReviewHandler := publicApp.Group("/review/")
	privateReviewHandler := privateApp.Group("/review/")

	privateReviewHandler.Post("create/", handler.CreateReview())
	publicReviewHandler.Get("list/", handler.ListReview())
	privateReviewHandler.Patch("update/:id/", handler.UpdateReview())
	privateReviewHandler.Delete("delete/:id/", handler.DeleteReview())
}
