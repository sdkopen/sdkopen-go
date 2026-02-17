package webserver

import commonhttp "github.com/sdkopen/sdkopen-go/common/http"

type Route struct {
	Path       string
	HttpMethod commonhttp.HttpMethod
	Function   func(ctx WebContext)
}
