package factory

import (
	"github.com/google/wire"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	queue2 "github.com/ruskotwo/emotional-analyzer/internal/queue"
	"github.com/ruskotwo/emotional-analyzer/internal/queue/workers"
	"github.com/ruskotwo/emotional-analyzer/pkg/queue"
	"github.com/ruskotwo/emotional-analyzer/pkg/queue/rabbitmq"
)

var queueSet = wire.NewSet(
	provideRabbitMQ,
	wire.Bind(new(queue.Client), new(*rabbitmq.ClientImpl)),
	workers.NewSendAnalysisResultWorker,
	queue2.NewWorkersManager,
)

func provideRabbitMQ(cfg *config.Config) *rabbitmq.ClientImpl {
	client := rabbitmq.NewClient(cfg.Queue.Clients.RabbitMQ)

	client.CreateQueue(cfg.Queue.List.ToAnalysis.Name)
	client.CreateQueue(cfg.Queue.List.AnalysisResult.Name)

	return client
}
