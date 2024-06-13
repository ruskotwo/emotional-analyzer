package config

import (
	"log"
	"os"
	"strconv"
)

type Queue struct {
	RabbitMQ           string
	NameToAnalysis     string
	NameAnalysisResult string
}

type Config struct {
	HttpPort int
	Queue    Queue
}

func NewConfig() *Config {
	var err error

	port := 3000
	if v, ok := os.LookupEnv("HTTP_PORT"); ok {
		if port, err = strconv.Atoi(v); err != nil {
			log.Panicf("incorrect port %s", v)
		}
	}

	return &Config{
		HttpPort: port,
		Queue: Queue{
			RabbitMQ:           os.Getenv("RABBIT_MQ_DSN"),
			NameToAnalysis:     os.Getenv("QUEUE_TO_ANALYSIS"),
			NameAnalysisResult: os.Getenv("QUEUE_ANALYSIS_RESULT"),
		},
	}
}
