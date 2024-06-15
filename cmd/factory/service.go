package factory

import (
	"github.com/ruskotwo/emotional-analyzer/internal/http"
	"github.com/ruskotwo/emotional-analyzer/internal/queue"
)

type Service struct {
	WorkersManager *queue.WorkersManager
	Server         *http.Server
}

func NewService(
	workersManager *queue.WorkersManager,
	server *http.Server,
) *Service {
	return &Service{
		workersManager,
		server,
	}
}
