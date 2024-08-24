package monitor

import (
	"gofiber-boilerplate/base"
	"gofiber-boilerplate/modules/config"

	"github.com/gofiber/fiber/v2"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
	"go.uber.org/fx"
)

type MonitorModule struct {
	app    *fiber.App
	config config.ConfigService
	tp     *sdktrace.TracerProvider
}

func NewModule(app *fiber.App, config config.ConfigService) *MonitorModule {
	return &MonitorModule{app: app, config: config}
}

func fxRegister(lifeCycle fx.Lifecycle, module *MonitorModule) {
	base.FxRegister(module, lifeCycle)
}

var FxModule = fx.Module("Monitor", fx.Provide(NewModule), fx.Invoke(fxRegister))

// implements `BaseModule` of `base/module.go` start

func (module *MonitorModule) OnStart() error {
	module.initOpentelemetry()
	module.registerRoutes()
	return nil
}

func (module *MonitorModule) OnStop() error {
	module.destroyOpentelemetry()
	return nil
}

// implements `BaseModule` of `base/module.go` end
