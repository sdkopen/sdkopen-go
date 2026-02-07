package restserver

import (
	"encoding/json"
)

type DecoderConfig struct {
	ContentType ContentType
	Data        []byte
}

func Decoder[T any](config *DecoderConfig) (*T, error) {
	if config == nil {
		return nil, nil
	}

	switch config.ContentType {
	case ContentTypeJSON:
		return JsonDecoder[T](config.Data)
	default:
		return nil, nil
	}
}

func JsonDecoder[T any](data []byte) (*T, error) {
	var response T
	err := json.Unmarshal(data, &response)
	return &response, err
}
