package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"
	"strconv"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListBlog() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		blogCatId, err := strconv.Atoi(c.Query("blogCategoryId"))

		request := presenter.BlogListReq{
			Page:           int(utils.CheckPageInQuery(c)),
			PageSize:       int(utils.CheckPageSizeInQuery(c)),
			BlogCategoryID: blogCatId,
		}

		blog, totalPage, err := handler.usecase.ListBlog(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if blog == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		response := presenter.ListRes{
			Success:     true,
			CurrentPage: request.Page,
			TotalPage:   totalPage,
			Data:        blog,
		}
		if request.Page > totalPage {
			response.CurrentPage = totalPage
		}
		return c.JSON(response)
	}

}
