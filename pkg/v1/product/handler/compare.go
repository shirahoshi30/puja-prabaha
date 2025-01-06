package handler

import (
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

func (handler *handler) CompareProducts() fiber.Handler {
	return func(c *fiber.Ctx) error {

		errMap := make(map[string]string)

		var requestBody presenter.ProductCompareRequest

		err := c.BodyParser(&requestBody)
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		//initialize the translation function
		validate, trans := utils.InitTranslator()

		//validate the request using a validator and returns validation error(if any)
		err = validate.Struct(requestBody)
		if err != nil {
			validationErrors := err.(validator.ValidationErrors)
			errMap = utils.TranslateError(validationErrors, trans)
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}

		product, err := handler.usecase.CompareProducts(uint(requestBody.ProductIDs[0]), uint(requestBody.ProductIDs[1]))
		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))
		}
		return c.JSON(presenter.LisRes(product))

	}
}
