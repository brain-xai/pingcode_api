package auth

import (
	"context"
	"fmt"
	"net/url"
	"sync"

	apiauth "github.com/brain-xai/pingcode_api/internal/api/auth"
	"github.com/brain-xai/pingcode_api/internal/httpclient"
	modelauth "github.com/brain-xai/pingcode_api/sdk/model/auth"
)

// Service 认证服务
type Service struct {
	client *httpclient.Client
	baseURL string
	mu     sync.RWMutex
	token  *modelauth.Token
}

// NewService 创建认证服务
func NewService(client *httpclient.Client, baseURL string) *Service {
	return &Service{
		client: client,
		baseURL: baseURL,
	}
}

// GetToken 获取访问令牌（使用 client credentials）
func (s *Service) GetToken(ctx context.Context, clientID, clientSecret string) (*modelauth.Token, error) {
	// 构建请求参数
	query := url.Values{}
	query.Set("grant_type", "client_credentials")
	query.Set("client_id", clientID)
	query.Set("client_secret", clientSecret)

	// 发送请求
	var resp apiauth.GetTokenResponse
	if err := s.client.Get(ctx, "/v1/auth/token", query, &resp); err != nil {
		return nil, fmt.Errorf("failed to get token: %w", err)
	}

	token := resp.ToModel()

	// 缓存 token
	s.mu.Lock()
	s.token = token
	s.mu.Unlock()

	return token, nil
}

// GetCurrentToken 获取当前有效的令牌
func (s *Service) GetCurrentToken(ctx context.Context, clientID, clientSecret string) (string, error) {
	s.mu.RLock()
	token := s.token
	s.mu.RUnlock()

	// 如果没有 token 或已过期，重新获取
	if token == nil || token.IsExpired() {
		newToken, err := s.RefreshToken(ctx, clientID, clientSecret)
		if err != nil {
			return "", err
		}
		return newToken.AccessToken, nil
	}

	return token.AccessToken, nil
}

// RefreshToken 刷新访问令牌
func (s *Service) RefreshToken(ctx context.Context, clientID, clientSecret string) (*modelauth.Token, error) {
	return s.GetToken(ctx, clientID, clientSecret)
}

// GetCachedToken 获取缓存的令牌（不检查过期）
func (s *Service) GetCachedToken() *modelauth.Token {
	s.mu.RLock()
	defer s.mu.RUnlock()
	return s.token
}
