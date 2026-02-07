package restserver

import (
	"context"
	"mime/multipart"
	"net/http"
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
	Redirect(url string, httpStatusCode HttpStatusCode)
	ServeFile(path string)
	JsonResponse(httpStatusCode HttpStatusCode, body any)
	ErrorResponse(httpStatusCode HttpStatusCode, err error)
	EmptyResponse(httpStatusCode HttpStatusCode)
}
