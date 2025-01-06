package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListOrders() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.OrderListReq{
			Page:     int(utils.CheckPageInQuery(c)),
			PageSize: int(utils.CheckPageSizeInQuery(c)),
		}

		order, totalPage, err := handler.usecase.ListOrder(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if order == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListRes{
			Success:     true,
			CurrentPage: request.Page,
			TotalPage:   totalPage,
			Data:        order,
		}
		if request.Page > totalPage {
			response.CurrentPage = totalPage
		}
		return c.JSON(response)
	}
}
