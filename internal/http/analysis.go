package http

import (
	"encoding/json"
	queue2 "github.com/ruskotwo/emotional-analyzer/internal/queue"
	"github.com/ruskotwo/emotional-analyzer/pkg/queue"
	"net/http"
)

type AnalysisHandler struct {
	queueClient queue.Client
	queueNames  *queue2.Names
}

func NewAnalysisHandler(
	queueClient queue.Client,
	queueNames *queue2.Names,
) *AnalysisHandler {
	return &AnalysisHandler{
		queueClient: queueClient,
		queueNames:  queueNames,
	}
}

type BodyAddToAnalysis struct {
	Messages map[string]string
}

func (h AnalysisHandler) handleAddToAnalysis(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	body := &BodyAddToAnalysis{}
	err := json.NewDecoder(r.Body).Decode(body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(body.Messages) < 0 {
		http.Error(w, "The request contains no messages", http.StatusBadRequest)
		return
	}

	task := queue2.ToAnalysisTask{
		Messages: body.Messages,
	}
	payload, err := json.Marshal(task)
	if err != nil {
		http.Error(w, "The request contains no messages", http.StatusInternalServerError)
		return
	}

	err = h.queueClient.Publish(
		h.queueNames.GetNameToAnalysis(),
		payload,
	)

	w.WriteHeader(http.StatusCreated)
}
