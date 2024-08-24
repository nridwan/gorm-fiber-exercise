package app

import (
	"gofiber-boilerplate/modules/app/appmodel"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const (
	StatusSuccess = "SUCCESS"
)

type ResponseService interface {
	CreateErrorResponse(message string, errors []appmodel.Error) *appmodel.ErrorResponse
	CreateResponse(status string, result interface{}) *appmodel.SuccessResponse
	SendErrorResponse(ctx *fiber.Ctx, code int, message string, errors []appmodel.Error) error
	SendValidationErrorResponse(ctx *fiber.Ctx, code int, message string, errors validator.ValidationErrors) error
	SendResponse(ctx *fiber.Ctx, code int, status string, result interface{}) error
	SendSuccessResponse(ctx *fiber.Ctx, code int, result interface{}) error
	ErrorHandler(ctx *fiber.Ctx, err error) error
}

type responseServiceImpl struct {
}

func NewResponseService() ResponseService {
	return &responseServiceImpl{}
}

// impl `ResponseService` start

func (service *responseServiceImpl) CreateErrorResponse(message string, errors []appmodel.Error) *appmodel.ErrorResponse {
	return &appmodel.ErrorResponse{
		Message: message,
		Errors:  errors,
	}
}

func (service *responseServiceImpl) CreateResponse(status string, result interface{}) *appmodel.SuccessResponse {
	return &appmodel.SuccessResponse{
		Status: status,
		Result: result,
	}
}

func (service *responseServiceImpl) SendErrorResponse(ctx *fiber.Ctx, code int, message string, errors []appmodel.Error) error {
	return ctx.Status(code).JSON(service.CreateErrorResponse(message, errors))
}

func (service *responseServiceImpl) SendValidationErrorResponse(ctx *fiber.Ctx, code int, message string, errors validator.ValidationErrors) error {
	mappedError := make([]appmodel.Error, len(errors))
	for i, err := range errors {
		mappedError[i] = appmodel.Error{
			Field:   err.Field(),
			Message: err.Error(),
		}
	}

	return service.SendErrorResponse(ctx, code, message, mappedError)
}

func (service *responseServiceImpl) SendResponse(ctx *fiber.Ctx, code int, status string, result interface{}) error {
	return ctx.Status(code).JSON(service.CreateResponse(status, result))
}

func (service *responseServiceImpl) SendSuccessResponse(ctx *fiber.Ctx, code int, result interface{}) error {
	return service.SendResponse(ctx, 200, StatusSuccess, result)
}

// ErrorHandler check if connection should be continued or not
func (service *responseServiceImpl) ErrorHandler(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's an fiber.*Error
	if e, ok := err.(*fiber.Error); ok {
		code = e.Code
	}

	return ctx.Status(code).JSON(service.CreateErrorResponse(err.Error(), nil))
}

// impl `ResponseService` end
