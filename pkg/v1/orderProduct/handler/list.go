package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListOrderProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		orderId, _ := strconv.Atoi(c.Query("orderId"))

		request := presenter.OrderProductListReq{
			Page:     int(utils.CheckPageInQuery(c)),
			PageSize: int(utils.CheckPageSizeInQuery(c)),
			OrderID:  orderId,
		}

		orderProducts, totalPage, err := handler.usecase.ListOrderProducts(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if orderProducts == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListRes{
			Success:     true,
			CurrentPage: request.Page,
			TotalPage:   totalPage,
			Data:        orderProducts,
		}
		if request.Page > totalPage {
			response.CurrentPage = totalPage
		}
		return c.JSON(response)
	}
}
