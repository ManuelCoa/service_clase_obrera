package handlers

import (
	"claseobrera/app/responses"

	"github.com/gofiber/fiber/v2"
)

func CreateErrorResponse(c *fiber.Ctx, statusCode int, message string, err error) error {
	errorDetail := &responses.ErrorDetail{
		Code:    statusCode,
		Message: err.Error(),
	}
	meta := &responses.Meta{
		RequestID:  c.Locals("request_id"),
		APIVersion: "1.0.0",
	}
	return responses.CreateResponse(c, "error", statusCode, message, nil, errorDetail, meta)
}

func CreateSuccessResponse(c *fiber.Ctx, statusCode int, message string, data interface{}) error {
	meta := &responses.Meta{
		RequestID:  c.Locals("request_id"),
		APIVersion: "1.0.0",
	}
	return responses.CreateResponse(c, "success", statusCode, message, data, nil, meta)
}
