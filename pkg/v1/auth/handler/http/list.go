package http

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListUsers() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		request := presenter.UserListReq{
			Page:     int(utils.CheckPageInQuery(c)),
			PageSize: int(utils.CheckPageSizeInQuery(c)),
		}

		users, totalPage, err := handler.usecase.ListUsers(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if users == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListRes{
			Success:     true,
			CurrentPage: request.Page,
			TotalPage:   totalPage,
			Data:        users,
		}
		if request.Page > totalPage {
			response.CurrentPage = totalPage
		}
		return c.JSON(response)
	}

}
