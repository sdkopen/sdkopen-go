package webclient

import (
	"net/http"

	commonhttp "github.com/sdkopen/sdkopen-go/common/http"
)

type Response struct {
	StatusCode commonhttp.HttpStatusCode
	Headers    http.Header
	Body       []byte
}
