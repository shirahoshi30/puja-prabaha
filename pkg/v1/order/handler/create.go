package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CreateOrder() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var data presenter.OrderCreateUpdatePresenter

		err := c.BodyParser(&data)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		//initialize the translation function
		validate, trans := utils.InitTranslator()

		//validate the request using a validator and returns validation error(if any)
		err = validate.Struct(data)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		_, _ = handler.usecase.CreateOrder(data)
		if len(errMap) > 0 {
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		return c.JSON(presenter.SuccessResponse())
	}

}
