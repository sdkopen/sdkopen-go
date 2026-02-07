package restserver

type HttpMethod int

const (
	MethodGet HttpMethod = iota
	MethodHead
	MethodPost
	MethodPut
	MethodPatch
	MethodDelete
	MethodConnect
	MethodOptions
	MethodTrace
)

func (httpMethod HttpMethod) String() string {
	return [...]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE"}[httpMethod]
}
