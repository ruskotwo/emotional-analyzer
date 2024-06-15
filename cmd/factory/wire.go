//go:build wireinject
// +build wireinject

package factory

import (
	"github.com/google/wire"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
)

func InitService() *Service {
	panic(
		wire.Build(
			config.NewConfig,
			httpSet,
			queueSet,
			gormSet,
			NewService,
		),
	)
}
