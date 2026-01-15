package api

import (
	"github.com/brain-xai/pingcode_api/sdk/model/auth"
)

// GetTokenRequest 获取令牌请求
type GetTokenRequest struct {
	GrantType    string `url:"grant_type"`    // 固定为 "client_credentials"
	ClientID     string `url:"client_id"`     // PingCode 应用的 Client ID
	ClientSecret string `url:"client_secret"` // PingCode 应用的 Secret
}

// GetTokenResponse 获取令牌响应
type GetTokenResponse struct {
	AccessToken string `json:"access_token"`
	TokenType   string `json:"token_type"`
	ExpiresIn   int64  `json:"expires_in"`
}

// ToModel 将 API DTO 转换为对外模型
func (r *GetTokenResponse) ToModel() *auth.Token {
	return &auth.Token{
		AccessToken: r.AccessToken,
		TokenType:   r.TokenType,
		ExpiresIn:   r.ExpiresIn,
	}
}
