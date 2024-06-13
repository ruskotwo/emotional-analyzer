package factory

import (
	"github.com/google/wire"
	"github.com/ruskotwo/emotional-analyzer/internal/http"
)

var httpSet = wire.NewSet(
	http.NewServer,
	http.NewAnalysisHandler,
)
