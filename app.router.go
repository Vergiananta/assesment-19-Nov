package main

import (
	"github.com/gorilla/mux"
	"superapp/connect"
	"superapp/controller"
	"superapp/controller/handler"
	"superapp/manager"
	"superapp/middlewares"
	"superapp/utils/httpparse"
	"superapp/utils/httpresponse"
)

type appRouter struct {
	app                  *superApp
	parse                *httpparse.JsonParse
	responder            httpresponse.IResponder
	connect              connect.Connect
	logRequestMiddleware *middlewares.LogRequestMiddleware
}

type appRoutes struct {
	centerRoutes controller.IDelivery
	mdw          []mux.MiddlewareFunc
}

func (r *appRouter) InitMainRoutes() {
	r.app.router.Use(r.logRequestMiddleware.Log)
	serviceManager := manager.NewServiceManager(r.connect)
	appRoutes := []appRoutes{
		{
			centerRoutes: handler.NewCustomerController(r.app.router, r.parse, r.responder, serviceManager.CustomerUsecase()),
			mdw: nil,
		},
		{
			centerRoutes: handler.NewMerchantController(r.app.router, r.parse, r.responder, serviceManager.MerchantUsecase()),
			mdw: nil,
		},
		{
			centerRoutes: handler.NewTrxController(r.app.router, r.parse, r.responder, serviceManager.TransactionUsecase()),
			mdw: nil,
		},
	}

	for _, r := range appRoutes {
		r.centerRoutes.InitRoute(r.mdw...)
	}
}

func NewAppRouter(app *superApp) *appRouter {
	return &appRouter{
		app,
		httpparse.NewJsonParse(),
		httpresponse.NewJSONResponder(),
		app.connect,
		middlewares.NewLogRequestMiddleware(),
	}
}
