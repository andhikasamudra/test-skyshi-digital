package internal

import (
	"github.com/gofiber/fiber/v2"
	"net/http"
)

const (
	FailedStatus   = "Failed"
	NotFoundStatus = "Not Found"
)

var mapStatusToKey = map[int]string{
	http.StatusInternalServerError: FailedStatus,
	http.StatusNotFound:            NotFoundStatus,
}

func BuildResponse(data interface{}) *fiber.Map {
	return &fiber.Map{
		"status":  "Success",
		"message": "Success",
		"data":    data,
	}
}
func BuildErrorResponse(err error) *fiber.Map {
	return &fiber.Map{
		"status":  "Failed",
		"message": err.Error(),
	}
}
