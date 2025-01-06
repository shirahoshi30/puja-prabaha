package handler

import (
	"pujaprabha/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.VerifyPaymentUsecase
}

func New(privateApp fiber.Router, usecase interfaces.VerifyPaymentUsecase) {
	handler := &handler{
		usecase: usecase,
	}
	verifyPaymentHandler := privateApp.Group("payment/")

	verifyPaymentHandler.Get("verify/", handler.VerifyPayment())
}
