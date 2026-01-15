package sdk

import (
	"context"
	"fmt"
	"sync"

	"github.com/brain-xai/pingcode_api/internal/httpclient"
	"github.com/brain-xai/pingcode_api/sdk/service/auth"
	"github.com/brain-xai/pingcode_api/sdk/service/global"
	"github.com/brain-xai/pingcode_api/sdk/service/project"
	"github.com/brain-xai/pingcode_api/sdk/service/ship"
)

// Client SDK 客户端
type Client struct {
	config    *Config
	http      *httpclient.Client
	auth      *auth.Service
	project   *project.Service
	global    *global.Service
	ship      *ship.Service
	projectMu sync.RWMutex
	globalMu  sync.RWMutex
	shipMu    sync.RWMutex
}

// NewClient 创建新的 SDK 客户端
// 如果配置中提供了 ClientID/ClientSecret，会自动获取 access_token
func NewClient(cfg *Config) (*Client, error) {
	// 验证配置
	if err := cfg.Validate(); err != nil {
		return nil, err
	}

	// 规范化 BaseURL
	baseURL := NormalizeBaseURL(cfg.BaseURL)

	// 创建 HTTP 客户端
	timeout := cfg.Timeout
	if timeout == 0 {
		timeout = 30 * defaultTimeout
	}

	httpClient, err := httpclient.NewClient(baseURL, timeout)
	if err != nil {
		return nil, fmt.Errorf("failed to create HTTP client: %w", err)
	}

	client := &Client{
		config: cfg,
		http:   httpClient,
	}

	// 初始化认证服务
	client.auth = auth.NewService(httpClient, baseURL)

	// 如果配置了 access_token，直接设置
	if cfg.HasAccessToken() {
		httpClient.SetToken(cfg.Auth.AccessToken)
	} else if cfg.HasClientCredentials() {
		// 使用 client_id/client_secret 获取 token
		token, err := client.auth.GetToken(context.Background(), cfg.Auth.ClientID, cfg.Auth.ClientSecret)
		if err != nil {
			return nil, fmt.Errorf("failed to get access token: %w", err)
		}
		httpClient.SetToken(token.AccessToken)
	}

	return client, nil
}

// Auth 返回认证服务（用于手动刷新 token）
func (c *Client) Auth() *auth.Service {
	return c.auth
}

// Project 返回项目服务（懒加载）
func (c *Client) Project() *project.Service {
	c.projectMu.Lock()
	defer c.projectMu.Unlock()

	if c.project == nil {
		c.project = project.NewService(c.http, c.config.BaseURL)
	}

	return c.project
}

// Global 返回全局服务（懒加载）
func (c *Client) Global() *global.Service {
	c.globalMu.Lock()
	defer c.globalMu.Unlock()

	if c.global == nil {
		c.global = global.NewService(c.http, c.config.BaseURL)
	}

	return c.global
}

// Ship 返回 Ship 服务（懒加载）
func (c *Client) Ship() *ship.Service {
	c.shipMu.Lock()
	defer c.shipMu.Unlock()

	if c.ship == nil {
		c.ship = ship.NewService(c.http, c.config.BaseURL)
	}

	return c.ship
}

const defaultTimeout = 30
