package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListAllProduct() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		catId, _ := strconv.Atoi(c.Query("categoryId"))

		request := presenter.ProductListReq{
			Page:       int(utils.CheckPageInQuery(c)),
			PageSize:   int(utils.CheckPageSizeInQuery(c)),
			CategoryID: catId,
		}

		products, totalPage, err := handler.usecase.ListAllProduct(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if products == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListRes{
			Success:     true,
			CurrentPage: request.Page,
			TotalPage:   totalPage,
			Data:        products,
		}
		if request.Page > totalPage {
			response.CurrentPage = totalPage
		}
		return c.JSON(response)
	}

}
