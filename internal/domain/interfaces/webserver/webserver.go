package interfaces

type WebServer interface {
	Add(method string, path string, action func(context Context) (any, error), middlewares []Middleware)
	Run()
}
