package http

import (
	"pujaprabha/internal/interfaces"

	"github.com/gofiber/fiber/v2"
)

type handler struct {
	usecase interfaces.UserUsecase
}

func New(publicApp fiber.Router, privateApp fiber.Router, usecase interfaces.UserUsecase) {
	handler := &handler{
		usecase: usecase,
	}

	// userHandler := app.Group("/auth")
	publicApp.Post("login/", handler.Login())
	publicApp.Post("create/", handler.RegisterUser())
	publicApp.Delete("delete/:id/", handler.DeleteUser())
	publicApp.Patch("update/:id/", handler.UpdateUser())
	privateApp.Get("list/", handler.ListUsers())
	publicApp.Post("reset/", handler.ForgotPassword())
	publicApp.Post("resend/", handler.ResendUserOTP())
	publicApp.Get("verifyOTP/", handler.VerifyOTP())
	privateApp.Post("change_password/", handler.ChangePassword())
}
