package handlers

import (
	"encoding/json"
	"github.com/go-oauth2/oauth2/v4/server"
	"github.com/ruskotwo/emotional-analyzer/internal/config"
	"github.com/ruskotwo/emotional-analyzer/internal/queue/types"
	"github.com/ruskotwo/emotional-analyzer/pkg/queue"
	"net/http"
	"strconv"
)

type AnalysisHandler struct {
	queueClient queue.Client
	config      *config.Config
	oAuth2      *server.Server
}

func NewAnalysisHandler(
	queueClient queue.Client,
	cfg *config.Config,
	oAuth2 *server.Server,
) *AnalysisHandler {
	return &AnalysisHandler{
		queueClient: queueClient,
		config:      cfg,
		oAuth2:      oAuth2,
	}
}

type BodyAddToAnalysis struct {
	Messages map[string]string `json:"messages"`
	Secret   string            `json:"secret"`
}

func (h AnalysisHandler) HandleAddToAnalysis(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	token, err := h.oAuth2.ValidationBearerToken(r)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusUnauthorized), http.StatusUnauthorized)
		return
	}

	body := &BodyAddToAnalysis{}
	err = json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(body.Messages) < 0 {
		http.Error(w, "The request contains no messages", http.StatusBadRequest)
		return
	}

	id, err := strconv.Atoi(token.GetUserID())
	if err != nil {
		http.Error(w, "Invalid token", http.StatusBadRequest)
		return
	}

	task := types.ToAnalysisTask{
		Messages: body.Messages,
		ClientId: uint(id),
		Secret:   body.Secret,
	}
	payload, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "The request contains no messages", http.StatusInternalServerError)
		return
	}

	err = h.queueClient.Publish(
		h.config.Queue.List.ToAnalysis.Name,
		payload,
	)

	w.WriteHeader(http.StatusCreated)
}
