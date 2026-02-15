package messaging

import "time"

type Message struct {
	ID        string
	Topic     string
	Body      []byte
	Headers   map[string]string
	Timestamp time.Time
}

type PublishOption func(*publishConfig)

type publishConfig struct {
	Headers      map[string]string
	DelaySeconds int
}

func WithHeaders(headers map[string]string) PublishOption {
	return func(c *publishConfig) {
		c.Headers = headers
	}
}

func WithDelay(seconds int) PublishOption {
	return func(c *publishConfig) {
		c.DelaySeconds = seconds
	}
}

func applyOptions(opts []PublishOption) publishConfig {
	cfg := publishConfig{}
	for _, opt := range opts {
		opt(&cfg)
	}
	return cfg
}
