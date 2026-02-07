package restserver

import (
	"log"

	"github.com/sdkopen/sdkopen-go/logging"
)

var (
	ServerRoutes      []Route
	ServerMiddlewares []IMiddleware
)

type Server interface {
	Initialize()
	Shutdown() error
	InjectMiddlewares()
	InjectCustomMiddlewares()
	InjectRoutes()
	ListenAndServe() error
}

func RegisterController(controller Controller) {
	ServerRoutes = append(ServerRoutes, controller.Routes()...)
}

func RegisterMiddleware(middleware IMiddleware) {
	ServerMiddlewares = append(ServerMiddlewares, middleware)
}

func ListenAndServe(server func() Server) {
	srv := server()
	srv.Initialize()
	srv.InjectMiddlewares()
	srv.InjectCustomMiddlewares()
	srv.InjectRoutes()

	logging.Info("Service '%s' running in %d port", "REST-SERVER", 8080)
	log.Fatal(srv.ListenAndServe())
}

func GET(route Route) {
	route.Path = "GET-" + route.Path
	ServerRoutes = append(ServerRoutes, route)
}
