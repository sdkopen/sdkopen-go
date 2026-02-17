package webserver

type Controller interface {
	Routes() (routes []Route)
}
