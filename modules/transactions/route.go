package transactions

func (module *TransactionModule) registerRoutes() {
	module.app.Post("/topup", module.jwtService.GetHandler(), module.userService.CanAccess, module.controller.handleTopup)
	module.app.Post("/payment", module.jwtService.GetHandler(), module.userService.CanAccess, module.controller.handlePayment)
	module.app.Post("/transfer", module.jwtService.GetHandler(), module.userService.CanAccess, module.controller.handleTransfer)
	module.app.Get("/transactions", module.jwtService.GetHandler(), module.userService.CanAccess, module.controller.handleReport)
}
