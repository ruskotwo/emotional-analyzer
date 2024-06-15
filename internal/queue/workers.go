package queue

import (
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	"github.com/ruskotwo/emotional-analyzer/internal/queue/workers"
	"github.com/ruskotwo/emotional-analyzer/pkg/queue"
)

type WorkersManager struct {
	config *config.Config
}

func NewWorkersManager(
	queueClient queue.Client,
	cfg *config.Config,
	sendAnalysisResultWorker *workers.SendAnalysisResultWorker,
) *WorkersManager {
	m := &WorkersManager{
		config: cfg,
	}

	for i := 0; i < cfg.Queue.List.AnalysisResult.WorkersCount; i++ {
		queueClient.Consume(
			cfg.Queue.List.AnalysisResult.Name,
			sendAnalysisResultWorker.Handle,
		)
	}

	return m
}
