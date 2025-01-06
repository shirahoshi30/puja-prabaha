package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListCart() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		uid := c.Locals("requester").(uint)

		cart, err := handler.usecase.ListCart(uint(uid))
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if cart == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		return c.JSON(presenter.ListResponse(cart))
	}

}
