package user

import (
	"gofiber-boilerplate/modules/app"
	"gofiber-boilerplate/modules/user/userdto"
	"gofiber-boilerplate/utils"

	"github.com/go-playground/validator/v10"
	"github.com/gofiber/fiber/v2"
)

const (
	validationError = "Validation Error"
)

type userController struct {
	service         UserService
	responseService app.ResponseService
	validator       *validator.Validate
}

func newUserController(service UserService, responseService app.ResponseService, validator *validator.Validate) *userController {
	return &userController{
		service:         service,
		responseService: responseService,
		validator:       validator,
	}
}

// handlers start

func (controller *userController) handleRegister(ctx *fiber.Ctx) error {
	request := userdto.RegisterDTO{}
	ctx.BodyParser(&request)
	err := controller.validator.Struct(request)

	if err != nil {
		return controller.responseService.SendValidationErrorResponse(ctx, 400, validationError, err.(validator.ValidationErrors))
	}

	model, err := controller.service.Insert(request.ToModel())

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	return controller.responseService.SendSuccessResponse(ctx, 201, model)
}

func (controller *userController) handleLogin(ctx *fiber.Ctx) error {
	request := userdto.LoginDTO{}
	ctx.BodyParser(&request)
	err := controller.validator.Struct(request)

	if err != nil {
		return controller.responseService.SendValidationErrorResponse(ctx, 400, validationError, err.(validator.ValidationErrors))
	}

	response, err := controller.service.Login(&request)

	if err != nil {
		return fiber.NewError(400, err.Error())
	}
	return controller.responseService.SendSuccessResponse(ctx, 200, response)
}

func (controller *userController) handleProfile(ctx *fiber.Ctx) error {
	var user *userdto.UserDTO
	var err error

	id, err := utils.GetFiberJwtUserId(ctx)

	if err == nil {
		user, err = controller.service.AddBalance(id, 10000)
	}

	if err != nil {
		return fiber.NewError(400, err.Error())
	}

	user.CreatedAt = nil

	return controller.responseService.SendSuccessResponse(ctx, 200, user)
}

// handlers end
