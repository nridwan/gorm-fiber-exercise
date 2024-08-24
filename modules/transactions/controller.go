package transactions

import (
	"gofiber-boilerplate/modules/app"
	"gofiber-boilerplate/modules/transactions/transactionsdto"
	"gofiber-boilerplate/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const (
	validationError = "Validation Error"
)

type TransactionController struct {
	service         TransactionService
	responseService app.ResponseService
	validator       *validator.Validate
}

func newTransactionController(service TransactionService, responseService app.ResponseService, validator *validator.Validate) *TransactionController {
	return &TransactionController{
		service:         service,
		responseService: responseService,
		validator:       validator,
	}
}

// handlers start

func (controller *TransactionController) handleTopup(ctx *fiber.Ctx) error {
	request := transactionsdto.TopupDTO{}
	ctx.BodyParser(&request)
	err := controller.validator.Struct(request)

	if err != nil {
		return controller.responseService.SendValidationErrorResponse(ctx, 400, validationError, err.(validator.ValidationErrors))
	}

	id, err := utils.GetFiberJwtUserId(ctx)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	model, err := controller.service.Topup(id, &request)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return controller.responseService.SendSuccessResponse(ctx, 201, model)
}

func (controller *TransactionController) handlePayment(ctx *fiber.Ctx) error {
	request := transactionsdto.PaymentDTO{}
	ctx.BodyParser(&request)
	err := controller.validator.Struct(request)

	if err != nil {
		return controller.responseService.SendValidationErrorResponse(ctx, 400, validationError, err.(validator.ValidationErrors))
	}

	id, err := utils.GetFiberJwtUserId(ctx)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	model, err := controller.service.Payment(id, &request)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return controller.responseService.SendSuccessResponse(ctx, 201, model)
}

func (controller *TransactionController) handleTransfer(ctx *fiber.Ctx) error {
	request := transactionsdto.TransferDTO{}
	ctx.BodyParser(&request)
	err := controller.validator.Struct(request)

	if err != nil {
		return controller.responseService.SendValidationErrorResponse(ctx, 400, validationError, err.(validator.ValidationErrors))
	}

	id, err := utils.GetFiberJwtUserId(ctx)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	model, err := controller.service.Transfer(id, &request)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return controller.responseService.SendSuccessResponse(ctx, 201, model)
}

func (controller *TransactionController) handleReport(ctx *fiber.Ctx) error {
	id, err := utils.GetFiberJwtUserId(ctx)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	model, err := controller.service.Report(id)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return controller.responseService.SendSuccessResponse(ctx, 201, model)
}

// handlers end
