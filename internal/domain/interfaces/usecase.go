package interfaces

import (
	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
)

type UseCase interface {
	Execute(context webserver.Context) (any, error)
}
