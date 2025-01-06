package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) VerifyPayment() fiber.Handler {
	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		var request presenter.VerifyPaymentPresenter

		err := c.BodyParser(&request)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		validate, trans := utils.InitTranslator()

		//validate the request using a validator and returns validation error(if any)
		err = validate.Struct(request)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		_, err = handler.usecase.VerifyPayment(request)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		return c.JSON(presenter.SuccessResponse())
	}

}
