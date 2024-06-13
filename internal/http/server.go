package http

import (
	"fmt"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	"log"
	"net/http"
)

type Server struct {
	port int
}

func NewServer(
	cfg *config.Config,
	analysisHandler *AnalysisHandler,
) *Server {
	http.HandleFunc("/add", analysisHandler.handleAddToAnalysis)

	return &Server{
		port: cfg.HttpPort,
	}
}

func (s Server) Run() {
	log.Printf("Listening on :%d...", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil))
}
