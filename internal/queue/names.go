package queue

import "github.com/ruskotwo/emotional-analyzer/internal/config"

type Names struct {
	config *config.Config
}

func NewNames(cfg *config.Config) *Names {
	return &Names{
		cfg,
	}
}

func (n Names) GetNameToAnalysis() string {
	return n.config.Queue.NameToAnalysis
}

func (n Names) GetNameAnalysisResult() string {
	return n.config.Queue.NameAnalysisResult
}
