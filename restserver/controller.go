package restserver

type Controller interface {
	Routes() (routes []Route)
}
