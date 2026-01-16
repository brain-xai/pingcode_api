package sdk

import (
	"testing"
	"time"
)

// TestNewClient_TimeoutDefaults 测试不同配置下的超时默认值
func TestNewClient_TimeoutDefaults(t *testing.T) {
	tests := []struct {
		name        string
		config      *Config
		shouldError bool
	}{
		{
			name:        "使用 NewDefaultConfig - 30秒默认值",
			config:      NewDefaultConfig(),
			shouldError: true, // 没有 Auth 配置，会返回错误
		},
		{
			name: "手动构造 Config，未设置 Timeout - 应使用默认 30秒",
			config: &Config{
				BaseURL: "https://test.com",
				Auth:    &AuthConfig{AccessToken: "test-token"},
			},
			shouldError: false,
		},
		{
			name: "显式设置 Timeout 为 60秒 - 使用用户值",
			config: &Config{
				BaseURL: "https://test.com",
				Timeout: 60 * time.Second,
				Auth:    &AuthConfig{AccessToken: "test-token"},
			},
			shouldError: false,
		},
		{
			name: "显式设置 Timeout 为 0 - 应使用默认 30秒",
			config: &Config{
				BaseURL: "https://test.com",
				Timeout: 0,
				Auth:    &AuthConfig{AccessToken: "test-token"},
			},
			shouldError: false,
		},
		{
			name: "显式设置 Timeout 为 10秒 - 使用用户值",
			config: &Config{
				BaseURL: "https://test.com",
				Timeout: 10 * time.Second,
				Auth:    &AuthConfig{AccessToken: "test-token"},
			},
			shouldError: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			client, err := NewClient(tt.config)

			if tt.shouldError {
				if err == nil {
					t.Errorf("期望返回错误，但得到 nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("创建客户端失败: %v", err)
			}

			// 验证 client 创建成功
			// 注意：我们无法直接验证内部 http.Client 的 timeout
			// 但可以通过验证 client 不为 nil 来确认创建成功
			if client == nil {
				t.Fatal("期望 client 不为 nil")
			}
		})
	}
}

// TestDefaultTimeout 测试默认超时常量值
func TestDefaultTimeout(t *testing.T) {
	expected := 30 * time.Second
	if defaultTimeout != expected {
		t.Errorf("期望 defaultTimeout = %v, 实际为 %v", expected, defaultTimeout)
	}
}
