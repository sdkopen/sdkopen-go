package webserver

import (
	"context"
	"mime/multipart"
	"net/http"

	commonhttp "github.com/sdkopen/sdkopen-go/common/http"
)

type WebContext interface {
	Context() context.Context
	Response() http.ResponseWriter
	Request() *http.Request
	RequestHeader(key string) []string
	RequestHeaders() map[string][]string
	PathParam(key string) string
	RawQuery() string
	QueryParam(key string) string
	QueryArrayParam(key string) []string
	DecodeQueryParams(object any) error
	DecodeBody(object any) error
	DecodeFormData(object any) error
	StringBody() (string, error)
	Path() string
	Method() string
	FormFile(key string) (multipart.File, *multipart.FileHeader, error)
	AddHeader(key string, value string)
	AddHeaders(headers map[string]string)
	Redirect(url string, statusCode commonhttp.HttpStatusCode)
	ServeFile(path string)
	JsonResponse(statusCode commonhttp.HttpStatusCode, body any)
	ErrorResponse(statusCode commonhttp.HttpStatusCode, err error)
	EmptyResponse(statusCode commonhttp.HttpStatusCode)
}
