package queue

type ToAnalysisTask struct {
	Messages []string
	UserId   int
}

type AnalysisResultTask struct {
	Messages map[string]string
	Emotions map[string]string
	UserId   int
}
