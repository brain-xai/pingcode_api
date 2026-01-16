package sprint

import (
	"github.com/brain-xai/pingcode_api/internal/httpclient"
)

// Service 迭代管理服务
type Service struct {
	client  *httpclient.Client
	baseURL string
}

// NewService 创建迭代管理服务
func NewService(client *httpclient.Client, baseURL string) *Service {
	return &Service{
		client:  client,
		baseURL: baseURL,
	}
}
