package restserver

type HttpMethod int

const (
	Get HttpMethod = iota
	Head
	Post
	Put
	Patch
	Delete
	Connect
	Options
	Trace
)

func (httpMethod HttpMethod) String() string {
	return [...]string{"GET", "HEAD", "POST", "PUT", "PATCH", "DELETE", "CONNECT", "OPTIONS", "TRACE"}[httpMethod]
}
