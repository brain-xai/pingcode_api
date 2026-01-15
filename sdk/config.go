package sdk

import (
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"
)

// AuthConfig 认证配置
type AuthConfig struct {
	ClientID     string // PingCode 应用的 Client ID
	ClientSecret string // PingCode 应用的 Secret
	// 可选：直接提供 access_token（跳过自动获取）
	AccessToken string
}

// Config SDK 配置
type Config struct {
	BaseURL    string         // API 基础 URL，如 https://open.pingcode.com
	HTTPClient *http.Client  // 可选的自定义 HTTP 客户端
	Timeout    time.Duration // 请求超时时间，默认 30s
	Auth       *AuthConfig   // 认证配置（必填）
}

// NewDefaultConfig 创建默认配置
func NewDefaultConfig() *Config {
	return &Config{
		BaseURL: "https://open.pingcode.com",
		Timeout: 30 * time.Second,
	}
}

// LoadConfigFromEnv 从环境变量加载配置
func LoadConfigFromEnv() (*Config, error) {
	cfg := &Config{
		BaseURL: getEnv("PINGCODE_BASE_URL", "https://open.pingcode.com"),
		Timeout: parseDuration(getEnv("PINGCODE_TIMEOUT", "30s")),
	}

	// 检查是否有直接提供的 access_token
	if accessToken := os.Getenv("PINGCODE_ACCESS_TOKEN"); accessToken != "" {
		cfg.Auth = &AuthConfig{
			AccessToken: accessToken,
		}
		return cfg, nil
	}

	// 否则使用 client_id 和 client_secret
	clientID := os.Getenv("PINGCODE_CLIENT_ID")
	clientSecret := os.Getenv("PINGCODE_CLIENT_SECRET")

	if clientID == "" || clientSecret == "" {
		return nil, ErrInvalidConfig
	}

	cfg.Auth = &AuthConfig{
		ClientID:     clientID,
		ClientSecret: clientSecret,
	}

	return cfg, nil
}

// Validate 验证配置有效性
func (c *Config) Validate() error {
	if c.BaseURL == "" {
		return fmt.Errorf("base URL is required: %w", ErrInvalidConfig)
	}

	if c.Auth == nil {
		return fmt.Errorf("auth config is required: %w", ErrInvalidConfig)
	}

	// 检查认证配置
	if c.Auth.AccessToken == "" {
		if c.Auth.ClientID == "" || c.Auth.ClientSecret == "" {
			return fmt.Errorf("either access_token or client_id/client_secret must be provided: %w", ErrInvalidConfig)
		}
	}

	return nil
}

// getEnv 获取环境变量，如果不存在则返回默认值
func getEnv(key, defaultValue string) string {
	if value := os.Getenv(key); value != "" {
		return value
	}
	return defaultValue
}

// parseDuration 解析持续时间字符串
func parseDuration(s string) time.Duration {
	d, err := time.ParseDuration(s)
	if err != nil {
		return 30 * time.Second
	}
	return d
}

// HasClientCredentials 检查是否配置了客户端凭据
func (c *Config) HasClientCredentials() bool {
	if c.Auth == nil {
		return false
	}
	return c.Auth.ClientID != "" && c.Auth.ClientSecret != ""
}

// HasAccessToken 检查是否配置了访问令牌
func (c *Config) HasAccessToken() bool {
	if c.Auth == nil {
		return false
	}
	return c.Auth.AccessToken != ""
}

// NormalizeBaseURL 规范化 Base URL，移除末尾斜杠
func NormalizeBaseURL(baseURL string) string {
	return strings.TrimSuffix(baseURL, "/")
}
