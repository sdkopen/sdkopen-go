package restserver

type Route struct {
	Path       string
	HttpMethod HttpMethod
	Function   func(ctx WebContext)
}
