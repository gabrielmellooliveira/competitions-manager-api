package interfaces

type Middleware interface {
	Execute(context Context) error
}
