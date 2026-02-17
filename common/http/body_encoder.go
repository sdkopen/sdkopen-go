package commonhttp

import (
	"encoding/json"
	"fmt"
)

var contentTypeEncoder = map[ContentType]func(any) ([]byte, error){
	ContentTypeJSON:      JsonEncoder,
	ContentTypeTextPlain: StringEncoder,
}

func Encoder(contentType ContentType, response any) ([]byte, error) {
	return contentTypeEncoder[contentType](response)
}

func JsonEncoder(response any) ([]byte, error) {
	return json.Marshal(response)
}

func StringEncoder(response any) ([]byte, error) {
	return []byte(fmt.Sprintf("%v", response)), nil
}
