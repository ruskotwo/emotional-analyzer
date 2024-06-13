package factory

import (
	"github.com/ruskotwo/emotional-analyzer/internal/http"
	"github.com/ruskotwo/emotional-analyzer/pkg/queue"
)

type Service struct {
	Server *http.Server
}

func NewService(
	_ queue.Client,
	server *http.Server,
) *Service {
	return &Service{server}
}
