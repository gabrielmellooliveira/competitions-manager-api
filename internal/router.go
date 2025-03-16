package internal

import (
	"github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces"
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

func Router(webServer webserver.WebServer, routes []interfaces.Route) {
	for _, route := range routes {
		webServer.Add(route.Method, route.Path, route.UseCase.Execute, route.Middlewares)
	}
}
