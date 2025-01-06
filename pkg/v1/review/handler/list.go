package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListReview() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.ReviewList{
			Page:      int(utils.CheckPageInQuery(c)),
			PageSize:  int(utils.CheckPageSizeInQuery(c)),
			ProductID: uint(utils.StringToUint(c.Query("productId"))),
		}

		review, totalPage, err := handler.usecase.ListReview(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if review == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListRes{
			Success:     true,
			CurrentPage: request.Page,
			TotalPage:   totalPage,
			Data:        review,
		}
		if request.Page > totalPage {
			response.CurrentPage = totalPage
		}
		return c.JSON(response)
	}

}
