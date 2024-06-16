package http

import (
	"fmt"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	"github.com/ruskotwo/emotional-analyzer/internal/http/handlers"
	"log"
	"net/http"
)

type Server struct {
	port int
}

func NewServer(
	cfg *config.Config,
	analysisHandler *handlers.AnalysisHandler,
	clientsHandler *handlers.ClientsHandler,
	oAuthHandler *handlers.OAuthHandler,
) *Server {
	http.HandleFunc("/register", clientsHandler.HandleRegister)
	http.HandleFunc("/oauth/token", oAuthHandler.HandleToken)

	http.HandleFunc("/analyze/task", analysisHandler.HandleAddToAnalysis)

	return &Server{
		port: cfg.HttpPort,
	}
}

func (s Server) Run() {
	log.Printf("Listening on :%d...", s.port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%d", s.port), nil))
}
