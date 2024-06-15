package workers

import (
	"bytes"
	"encoding/json"
	"github.com/ruskotwo/emotional-analyzer/internal/gorm/clients"
	"github.com/ruskotwo/emotional-analyzer/internal/queue/types"
	"log"
	"net/http"
)

type SendAnalysisResultWorker struct {
	clientsRepository *clients.Repository
}

func NewSendAnalysisResultWorker(
	clientsRepository *clients.Repository,
) *SendAnalysisResultWorker {
	return &SendAnalysisResultWorker{
		clientsRepository: clientsRepository,
	}
}

type SendAnalysisResultBody struct {
	Messages map[string]string `json:"messages"`
	Secret   string            `json:"secret"`
}

func (w SendAnalysisResultWorker) Handle(msg interface{}) {
	message, _ := msg.(types.QueueMessage)

	var task types.ToAnalysisTask
	err := json.Unmarshal(message.Body, &task)
	if err != nil {
		log.Printf("Invalid task %v\n", err)
		_ = message.Reject(true)
		return
	}

	client, err := w.clientsRepository.GetByID(task.ClientId)
	if err != nil {
		log.Printf("Not found client by ID %d: %v\n", task.ClientId, err)
		_ = message.Ack(true)
		return
	}

	body := SendAnalysisResultBody{
		Messages: task.Messages,
		Secret:   task.Secret,
	}
	jsonBody, err := json.Marshal(body)
	if err != nil {
		log.Fatalf("Cant marshal body %v\n", err)
	}

	r, err := http.NewRequest("POST", client.CallbackUrl, bytes.NewBuffer(jsonBody))
	if err != nil {
		log.Printf("Cant create request for %d: %v\n", task.ClientId, err)
		_ = message.Ack(true)
		return
	}

	r.Header.Add("Content-Type", "application/json")

	httpClient := http.Client{}
	res, err := httpClient.Do(r)
	if err != nil {
		// TODO: Retries
		log.Printf("Cant send result for %d: %v\n", task.ClientId, err)
		_ = message.Ack(true)
		return
	}

	log.Printf("Sent result for %d: %v\n", task.ClientId, res.StatusCode)
	_ = message.Ack(true)
}
