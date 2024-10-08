package main

import (
	"context"
	"gofiber-boilerplate/modules/app"
	"gofiber-boilerplate/modules/config"
	"gofiber-boilerplate/modules/db"
	"gofiber-boilerplate/modules/jwt"
	"gofiber-boilerplate/modules/transactions"
	"gofiber-boilerplate/modules/user"
	"log"

	"github.com/gofiber/fiber/v2"
	"go.uber.org/fx"
)

func main() {
	fxApp := fx.New(
		config.FxModule,
		app.FxModule,
		db.FxModule,
		jwt.FxModule,
		user.FxModule,
		transactions.FxModule,
		fx.Invoke(registerWebServer),
	)

	fxApp.Run()
}

func registerWebServer(
	lifeCycle fx.Lifecycle,
	app *fiber.App,
	config config.ConfigService,
) {
	lifeCycle.Append(fx.Hook{
		OnStart: func(_ context.Context) error {
			go func() {
				app.Get("/", func(c *fiber.Ctx) error {
					return fiber.NewError(400, "Error")
				})

				if err := app.Listen(
					config.Getenv("APP_HOST", "") + ":" + config.Getenv("PORT", "3000"),
				); err != nil {
					log.Fatalf("start server error : %v\n", err)
				}

			}()
			return nil
		},
		OnStop: func(_ context.Context) error {
			log.Println("stopping server ...")
			err := app.Shutdown()
			log.Println("stop server success")
			return err
		},
	})
}
