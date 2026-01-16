package workitem

import (
	"github.com/brain-xai/pingcode_api/internal/httpclient"
)

// Service 工作项服务
type Service struct {
	client  *httpclient.Client
	baseURL string
}

// NewService 创建工作项服务
func NewService(client *httpclient.Client, baseURL string) *Service {
	return &Service{
		client:  client,
		baseURL: baseURL,
	}
}
