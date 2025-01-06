package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"

	"github.com/gofiber/fiber/v2"
)

func (handler *handler) ListSlider() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)

		slider, err := handler.usecase.ListSlider()
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		} else if slider == nil {
			return c.JSON(presenter.EmptyResponse{Data: nil, Success: true})
		}
		return c.JSON(presenter.LisRes(slider))
	}

}
