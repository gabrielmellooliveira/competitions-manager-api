package webserver

import (
	"net/http"

	webserver "github.com/gabrielmellooliveira/competitions-manager-api/internal/domain/interfaces/webserver"
	"github.com/labstack/echo"
)

type WebServerHandler struct {
	Client *echo.Echo
}

func NewWebServerHandler() webserver.WebServer {
	return &WebServerHandler{
		Client: echo.New(),
	}
}

func (handler *WebServerHandler) Add(method string, path string, action func(context webserver.Context) (any, error), middlewares []webserver.Middleware) {
	var echoMiddlewares []echo.MiddlewareFunc
	for _, middleware := range middlewares {
		echoMiddlewares = append(
			echoMiddlewares,
			convertToEchoMiddleware(middleware),
		)
	}

	handler.Client.Add(method, path, func(c echo.Context) error {
		context := webserver.Context{
			Request:     *c.Request(),
			Response:    c.Response().Writer,
			QueryParams: c.QueryParams(),
			GetParam:    c.Param,
			Set:         c.Set,
		}

		result, err := action(context)
		if err != nil {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": err.Error()})
		}

		return c.JSON(http.StatusOK, result)
	}, echoMiddlewares...)
}

func convertToEchoMiddleware(middleware webserver.Middleware) func(next echo.HandlerFunc) echo.HandlerFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			context := webserver.Context{
				Request:     *c.Request(),
				Response:    c.Response().Writer,
				QueryParams: c.QueryParams(),
				GetParam:    c.Param,
				Set:         c.Set,
			}

			err := middleware.Execute(context)
			if err != nil {
				return c.JSON(http.StatusUnauthorized, map[string]string{"error": err.Error()})
			}

			return next(c)
		}
	}
}

func (handler *WebServerHandler) Run() {
	handler.Client.Start(":8080")
}
