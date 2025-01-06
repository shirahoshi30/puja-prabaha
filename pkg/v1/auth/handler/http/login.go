package http

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

// to login get user and password
func (handler *handler) Login() fiber.Handler {
	err_map := make(map[string]string)
	return func(c *fiber.Ctx) error {
		var requestBody presenter.LoginRequest

		err := c.BodyParser(&requestBody)
		if err != nil {
			err_map["error"] = err.Error()
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err_map))
		}

		validate, trans := utils.InitTranslator()
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			err_map = utils.TranslateError(validationErrors, trans)
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err_map))
		}
		response, err_map := handler.usecase.Login(&requestBody)
		if len(err_map) != 0 {
			c.Status(http.StatusBadRequest)
			return c.JSON(presenter.ErrorResponse(err_map))
		}

		return c.JSON(presenter.LisRes(response))
	}

}
