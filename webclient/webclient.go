package webclient

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	commonhttp "github.com/sdkopen/sdkopen-go/common/http"
)

type WebClient struct {
	baseURL string
	headers map[string]string
	client  *http.Client
}

func New(baseURL string) *WebClient {
	return &WebClient{
		baseURL: baseURL,
		headers: make(map[string]string),
		client:  &http.Client{},
	}
}

func (c *WebClient) WithHeader(key, value string) *WebClient {
	c.headers[key] = value
	return c
}

func (c *WebClient) WithTimeout(timeout time.Duration) *WebClient {
	c.client.Timeout = timeout
	return c
}

func (c *WebClient) Get(path string, result any) (*Response, error) {
	return c.doRequest(http.MethodGet, path, nil, result)
}

func (c *WebClient) Post(path string, body any, result any) (*Response, error) {
	return c.doRequest(http.MethodPost, path, body, result)
}

func (c *WebClient) Put(path string, body any, result any) (*Response, error) {
	return c.doRequest(http.MethodPut, path, body, result)
}

func (c *WebClient) Patch(path string, body any, result any) (*Response, error) {
	return c.doRequest(http.MethodPatch, path, body, result)
}

func (c *WebClient) Delete(path string, result any) (*Response, error) {
	return c.doRequest(http.MethodDelete, path, nil, result)
}

func (c *WebClient) doRequest(method, path string, body any, result any) (*Response, error) {
	url := fmt.Sprintf("%s%s", c.baseURL, path)

	var reqBody io.Reader
	if body != nil {
		jsonBytes, err := json.Marshal(body)
		if err != nil {
			return nil, fmt.Errorf("failed to marshal request body: %w", err)
		}
		reqBody = bytes.NewReader(jsonBytes)
	}

	req, err := http.NewRequest(method, url, reqBody)
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %w", err)
	}

	for key, value := range c.headers {
		req.Header.Set(key, value)
	}

	if body != nil {
		req.Header.Set("Content-Type", commonhttp.ContentTypeJSON.String())
	}

	resp, err := c.client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("request failed: %w", err)
	}
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("failed to read response body: %w", err)
	}

	response := &Response{
		StatusCode: commonhttp.HttpStatusCode(resp.StatusCode),
		Headers:    resp.Header,
		Body:       respBody,
	}

	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return response, fmt.Errorf("failed to decode response body: %w", err)
		}
	}

	return response, nil
}
