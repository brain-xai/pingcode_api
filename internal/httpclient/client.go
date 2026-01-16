package httpclient

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
	"sync"
	"time"

	"github.com/brain-xai/pingcode_api/internal/errors"
)

// Client HTTP 客户端封装
type Client struct {
	httpClient *http.Client
	baseURL    *url.URL
	token      string
	tokenMu    sync.RWMutex
}

// NewClient 创建新的 HTTP 客户端
func NewClient(baseURL string, timeout time.Duration) (*Client, error) {
	parsedURL, err := url.Parse(baseURL)
	if err != nil {
		return nil, fmt.Errorf("invalid base URL: %w", err)
	}

	return &Client{
		httpClient: &http.Client{
			Timeout: timeout,
		},
		baseURL: parsedURL,
	}, nil
}

// SetToken 设置访问令牌（线程安全）
func (c *Client) SetToken(token string) {
	c.tokenMu.Lock()
	defer c.tokenMu.Unlock()
	c.token = token
}

// GetToken 获取当前令牌（线程安全）
func (c *Client) GetToken() string {
	c.tokenMu.RLock()
	defer c.tokenMu.RUnlock()
	return c.token
}

// Get 发送 GET 请求
func (c *Client) Get(ctx context.Context, path string, query url.Values, result interface{}) error {
	return c.doRequest(ctx, http.MethodGet, path, query, nil, result)
}

// Post 发送 POST 请求
func (c *Client) Post(ctx context.Context, path string, query url.Values, body, result interface{}) error {
	return c.doRequest(ctx, http.MethodPost, path, query, body, result)
}

// Patch 发送 PATCH 请求
func (c *Client) Patch(ctx context.Context, path string, query url.Values, body, result interface{}) error {
	return c.doRequest(ctx, http.MethodPatch, path, query, body, result)
}

// Delete 发送 DELETE 请求
func (c *Client) Delete(ctx context.Context, path string, query url.Values) error {
	return c.doRequest(ctx, http.MethodDelete, path, query, nil, nil)
}

// buildPath 构建带路径参数的 URL（内部方法）
func (c *Client) buildPath(pathTemplate string, pathParams map[string]string) (string, error) {
	path := pathTemplate
	for key, value := range pathParams {
		placeholder := "{" + key + "}"
		if !strings.Contains(path, placeholder) {
			return "", fmt.Errorf("path placeholder %s not found in template %s", placeholder, pathTemplate)
		}
		path = strings.ReplaceAll(path, placeholder, url.PathEscape(value))
	}
	return path, nil
}

// GetWithPathParams 发送带路径参数的 GET 请求
func (c *Client) GetWithPathParams(ctx context.Context, pathTemplate string, pathParams map[string]string, query url.Values, result interface{}) error {
	path, err := c.buildPath(pathTemplate, pathParams)
	if err != nil {
		return err
	}
	return c.Get(ctx, path, query, result)
}

// PostWithPathParams 发送带路径参数的 POST 请求
func (c *Client) PostWithPathParams(ctx context.Context, pathTemplate string, pathParams map[string]string, query url.Values, body, result interface{}) error {
	path, err := c.buildPath(pathTemplate, pathParams)
	if err != nil {
		return err
	}
	return c.Post(ctx, path, query, body, result)
}

// PatchWithPathParams 发送带路径参数的 PATCH 请求
func (c *Client) PatchWithPathParams(ctx context.Context, pathTemplate string, pathParams map[string]string, query url.Values, body, result interface{}) error {
	path, err := c.buildPath(pathTemplate, pathParams)
	if err != nil {
		return err
	}
	return c.Patch(ctx, path, query, body, result)
}

// DeleteWithPathParams 发送带路径参数的 DELETE 请求
func (c *Client) DeleteWithPathParams(ctx context.Context, pathTemplate string, pathParams map[string]string, query url.Values) error {
	path, err := c.buildPath(pathTemplate, pathParams)
	if err != nil {
		return err
	}
	return c.Delete(ctx, path, query)
}

// doRequest 发送 HTTP 请求（内部方法）
func (c *Client) doRequest(ctx context.Context, method, path string, query url.Values, body, result interface{}) error {
	// 构建完整 URL
	relURL, err := url.Parse(path)
	if err != nil {
		return fmt.Errorf("invalid path: %w", err)
	}

	// 合并查询参数
	if query != nil {
		relURL.RawQuery = query.Encode()
	}

	fullURL := c.baseURL.ResolveReference(relURL)

	// 准备请求体
	var bodyReader io.Reader
	if body != nil {
		jsonData, err := json.Marshal(body)
		if err != nil {
			return fmt.Errorf("failed to marshal request body: %w", err)
		}
		bodyReader = bytes.NewReader(jsonData)
	}

	// 创建请求
	req, err := http.NewRequestWithContext(ctx, method, fullURL.String(), bodyReader)
	if err != nil {
		return fmt.Errorf("failed to create request: %w", err)
	}

	// 设置请求头
	if body != nil {
		req.Header.Set("Content-Type", "application/json")
	}

	// 添加 Authorization 头（除非是获取 token 的请求）
	if !strings.HasSuffix(path, "/v1/auth/token") {
		c.tokenMu.RLock()
		token := c.token
		c.tokenMu.RUnlock()

		if token != "" {
			req.Header.Set("Authorization", "Bearer "+token)
		}
	}

	// 发送请求
	resp, err := c.httpClient.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %w", err)
	}
	defer resp.Body.Close()

	// 读取响应体
	respBody, err := io.ReadAll(resp.Body)
	if err != nil {
		return fmt.Errorf("failed to read response body: %w", err)
	}

	// 处理错误状态码
	if resp.StatusCode >= 400 {
		return c.handleError(resp.StatusCode, respBody)
	}

	// 解析响应
	if result != nil && len(respBody) > 0 {
		if err := json.Unmarshal(respBody, result); err != nil {
			return fmt.Errorf("failed to unmarshal response: %w, body: %s", err, string(respBody))
		}
	}

	return nil
}

// handleError 处理 HTTP 错误响应
func (c *Client) handleError(statusCode int, body []byte) error {
	// 尝试解析为 API 错误格式
	var apiErr struct {
		Code    string `json:"code"`
		Message string `json:"message"`
	}

	if err := json.Unmarshal(body, &apiErr); err == nil && apiErr.Code != "" {
		return &errors.APIError{
			StatusCode: statusCode,
			Code:       apiErr.Code,
			Message:    apiErr.Message,
		}
	}

	// 如果无法解析，返回通用错误
	return &errors.APIError{
		StatusCode: statusCode,
		Code:       "http_error",
		Message:    string(body),
	}
}
