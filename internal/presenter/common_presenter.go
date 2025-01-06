package presenter

import "github.com/gofiber/fiber/v2"

type EmptyResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

type DetailResponse struct {
	Success bool        `json:"success"`
	Data    interface{} `json:"data"`
}

// SuccessResponse creates a success response in the form of fiber map, commonly used for indicating the success of an operation.
func SuccessResponse() *fiber.Map {
	return &fiber.Map{
		"success": true,
	}
}

// ErrorResponse creates an error response in the form of fiber map, includes a map of errors and sets the success flag to false.
// it is also used to communicate error in response.
func ErrorResponse(err map[string]string) *fiber.Map {
	return &fiber.Map{
		"success": false,
		"errors":  err,
	}
}

func ListResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"success":     true,
		"data":        data,
		"currentPage": 0,
		"totalPage":   0,
	}
}

func LisRes(data interface{}) *fiber.Map {
	return &fiber.Map{
		"success": true,
		"data":    data,
	}
}

type ListRes struct {
	Success     bool        `json:"success"`
	CurrentPage int         `json:"currentPage"`
	TotalPage   int         `json:"totalPage"`
	Data        interface{} `json:"data"`
}
