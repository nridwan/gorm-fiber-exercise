package user

import (
	"gofiber-boilerplate/base"
	"gofiber-boilerplate/modules/db"
	"gofiber-boilerplate/modules/jwt"
	"gofiber-boilerplate/modules/user/usermodel"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type UserModule struct {
	Service    UserService
	controller *userController
	jwtService jwt.JwtService
	db         db.DbService
	app        *fiber.App
}

func NewModule(service UserService, controller *userController, jwtService jwt.JwtService, db db.DbService, app *fiber.App) *UserModule {
	return &UserModule{Service: service, jwtService: jwtService, controller: controller, db: db, app: app}
}

func fxRegister(lifeCycle fx.Lifecycle, module *UserModule) {
	base.FxRegister(module, lifeCycle)
}

var FxModule = fx.Module("User", fx.Provide(NewUserService), fx.Provide(newUserController), fx.Provide(NewModule), fx.Invoke(fxRegister))

// implements `BaseModule` of `base/module.go` start

func (module *UserModule) OnStart() error {
	module.db.Default().AutoMigrate(&usermodel.UserModel{})
	module.Service.Init(module.db)
	module.registerRoutes()
	return nil
}

func (module *UserModule) OnStop() error {
	return nil
}

// implements `BaseModule` of `base/module.go` end
