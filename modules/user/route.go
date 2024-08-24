package user

func (module *UserModule) registerRoutes() {
	module.app.Post("/register", module.controller.handleRegister)
	module.app.Post("/login", module.controller.handleLogin)
	module.app.Get("/profile", module.jwtService.GetHandler(), module.Service.CanAccess, module.controller.handleProfile)
}
