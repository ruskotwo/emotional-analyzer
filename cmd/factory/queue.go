package factory

import (
	"github.com/google/wire"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	queue2 "github.com/ruskotwo/emotional-analyzer/internal/queue"
	"github.com/ruskotwo/emotional-analyzer/pkg/queue"
	"github.com/ruskotwo/emotional-analyzer/pkg/queue/rabbitmq"
)

var queueSet = wire.NewSet(
	provideRabbitMQ,
	wire.Bind(new(queue.Client), new(*rabbitmq.ClientImpl)),
	queue2.NewNames,
)

func provideRabbitMQ(cfg *config.Config) *rabbitmq.ClientImpl {
	client := rabbitmq.NewClient(cfg.Queue.RabbitMQ)

	client.CreateQueue(cfg.Queue.NameToAnalysis)
	client.CreateQueue(cfg.Queue.NameAnalysisResult)

	return client
}
