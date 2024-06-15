package types

type ToAnalysisTask struct {
	Messages map[string]string `json:"messages"`
	ClientId uint              `json:"client_id"`
	Secret   string            `json:"secret"`
}

type AnalysisResultTask struct {
	Messages map[string]string `json:"messages"`
	ClientId uint              `json:"client_id"`
	Secret   string            `json:"secret"`
	Retries  uint              `json:"retries"`
}
