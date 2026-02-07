package restserver

type IMiddleware interface {
	Apply(ctx WebContext) error
}
