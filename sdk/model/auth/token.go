package auth

import "time"

// Token 访问令牌
type Token struct {
	AccessToken string `json:"access_token"` // 访问令牌
	TokenType   string `json:"token_type"`   // 令牌类型，固定为 "Bearer"
	ExpiresIn   int64  `json:"expires_in"`   // 过期时间（10位时间戳）
}

// IsExpired 检查令牌是否已过期
// 提前5分钟判定为过期，留出安全边界
func (t *Token) IsExpired() bool {
	if t.ExpiresIn == 0 {
		return false
	}
	// 提前5分钟判定为过期
	expiryTime := t.ExpiresIn - 300
	return time.Now().Unix() > expiryTime
}
