package queue

type ToAnalysisTask struct {
	Messages map[string]string
	UserId   int
}

type AnalysisResultTask struct {
	Messages map[string]string
	UserId   int
}
