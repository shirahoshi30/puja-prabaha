package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListVarient() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		request := presenter.VariantListReq{
			Page:     int(utils.CheckPageInQuery(c)),
			PageSize: int(utils.CheckPageSizeInQuery(c)),
			Max:      (utils.StringToUint(c.Query("max"))),
			Min:      (utils.StringToUint(c.Query("min"))),
			Search:   c.Query("search"),
		}

		products, totalPage, err := handler.usecase.ListVarient(request)
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
