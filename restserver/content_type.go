package restserver

type ContentType int

const (
	ContentTypeJSON ContentType = iota
	ContentTypeTextPlain
	ContentTypePDF
	ContentTypeOctetStream
)

func (ct ContentType) String() string {
	return [...]string{
		"application/json",
		"text/plain",
		"application/pdf",
		"application/octet-stream",
	}[ct]
}
