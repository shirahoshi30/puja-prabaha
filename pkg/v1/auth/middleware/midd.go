package middleware

import (
	"context"
	"errors"
	"net/http"
	"pujaprabha/internal/presenter"
	"pujaprabha/pkg/v1/auth/repository"
	"pujaprabha/pkg/v1/auth/usecase"
	"strings"

	"github.com/gofiber/fiber/v2"
	jwtware "github.com/gofiber/jwt/v2"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// Protected protects the route, validate jwtware.
func Protected() fiber.Handler {

	jwtMW := jwtware.New(jwtware.Config{
		SigningKey:   []byte(viper.GetString("jwt.secret")),
		Claims:       &presenter.JwtCustomClaims{},
		ErrorHandler: jwtError,
	})
	return jwtMW
}

// this function is executed in case of invalid token.
func jwtError(c *fiber.Ctx, err error) error {
	errMap := make(map[string]string)

	if err.Error() == "Missing or malformed JWT" {
		errMap["error"] = err.Error()
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Missing or malformed JWT", "data": nil})
	}
	errMap["error"] = err.Error()
	return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"status": "error", "message": "Invalid or expired JWt", "data": nil})
}

func SetUserRole(userRole string) fiber.Handler {
	return func(c *fiber.Ctx) error {

		ctx := context.WithValue(c.Context(), "userRole", userRole)
		c.SetUserContext(ctx)
		return c.Next()
	}
}

func GetUserRole(ctx context.Context) string {
	userRole := ctx.Value("userRole")
	if userRole != nil {
		return userRole.(string)
	}
	return ""
}

func Contains(urlPaths []string, urlPath string) bool {
	for i := range urlPaths {
		if urlPaths[i] == urlPath {
			return true
		}
	}
	return false
}

// validates the role of the user for authorization.
func AdminMiddleware(handler fiber.Handler) fiber.Handler {
	permissionUrls := []string{"/user/login/", "/product/list/", "/product/details/", "/category/list/", "/varient/list/", "/cart/create/", "/cart/list/",
		"/cart/update/", "/cart/delete/", "/slider/list/", "/blog/list/", "/blog/details/", "/blogCategory/list/"}

	return func(c *fiber.Ctx) error {
		errMap := make(map[string]string)
		requesterID := (c.Locals("requester")).(uint)

		rID := uint(requesterID)

		user, err := repository.GetUserByID(rID)
		if err != nil {
			return err
		}
		if user.Role == ("admin") {
			return handler(c)
		} else if user.Role == "user" {
			url := c.Path()

			split := strings.Split(url, "/api/v1")
			for _, url := range permissionUrls {
				if yes := strings.Contains(split[1], url); yes {
					return handler(c)
				}
			}
			errMap["error"] = "access denied"
			return c.Status(http.StatusBadRequest).JSON(presenter.ErrorResponse(errMap))

		}

		return errors.New("unauthorized action")

	}
}

func ValidateJWT(db *gorm.DB, config jwtware.Config) fiber.Handler {
	errMap := make(map[string]string)

	return func(c *fiber.Ctx) error {
		uc := usecase.New(repository.New(db))

		user, err := uc.ValidateToken(c)

		if err != nil {
			errMap["error"] = err.Error()
			return c.Status(http.StatusForbidden).JSON(presenter.ErrorResponse(errMap))

		}
		c.Locals("requester", user.ID)

		return c.Next()
	}

}
