package interfaces

import (
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

type Route struct {
	Path        string
	Method      string
	UseCase     UseCase
	Middlewares []webserver.Middleware
}
