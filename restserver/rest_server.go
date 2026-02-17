package restserver

import (
	"log"

	"github.com/sdkopen/sdkopen-go/common/observer"
	"github.com/sdkopen/sdkopen-go/logging"
)

var (
	ServerRoutes      []Route
	ServerMiddlewares []IMiddleware
	ServerInstance    Server
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
	ServerInstance = server()
	ServerInstance.Initialize()
	ServerInstance.InjectMiddlewares()
	ServerInstance.InjectCustomMiddlewares()
	ServerInstance.InjectRoutes()

	observer.Attach(restObserver{})
	logging.Info("Service '%s' running in %d port", "REST-SERVER", 8080)
	log.Fatal(ServerInstance.ListenAndServe())
}
