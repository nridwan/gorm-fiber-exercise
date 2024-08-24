package transactions

import (
	"gofiber-boilerplate/base"
	"gofiber-boilerplate/modules/db"
	"gofiber-boilerplate/modules/jwt"
	"gofiber-boilerplate/modules/transactions/transactionsmodel"
	"gofiber-boilerplate/modules/user"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

type TransactionModule struct {
	Service     TransactionService
	controller  *TransactionController
	userService user.UserService
	jwtService  jwt.JwtService
	db          db.DbService
	app         *fiber.App
}

func NewModule(service TransactionService, controller *TransactionController, jwtService jwt.JwtService, userService user.UserService, db db.DbService, app *fiber.App) *TransactionModule {
	return &TransactionModule{Service: service, jwtService: jwtService, userService: userService, controller: controller, db: db, app: app}
}

func fxRegister(lifeCycle fx.Lifecycle, module *TransactionModule) {
	base.FxRegister(module, lifeCycle)
}

var FxModule = fx.Module("Transaction", fx.Provide(NewTransactionService), fx.Provide(newTransactionController), fx.Provide(NewModule), fx.Invoke(fxRegister))

// implements `BaseModule` of `base/module.go` start

func (module *TransactionModule) OnStart() error {
	module.db.Default().AutoMigrate(&transactionsmodel.TransactionModel{})
	module.Service.Init(module.db)
	module.registerRoutes()
	return nil
}

func (module *TransactionModule) OnStop() error {
	module.Service.Destroy()
	return nil
}

// implements `BaseModule` of `base/module.go` end
