package config

type Queue struct {
	Clients QueueClients
	List    QueueList
}

type QueueClients struct {
	RabbitMQ string
}

type QueueList struct {
	ToAnalysis     QueueListItem
	AnalysisResult QueueListItem
}

type QueueListItem struct {
	Name         string
	WorkersCount int
}
