package config

import (
	"log"
	"os"
	"strconv"
)

type Config struct {
	HttpPort int
	MysqlDSN string
	OAuth2   OAuth2
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
		MysqlDSN: os.Getenv("MYSQL_DSN"),
		OAuth2: OAuth2{
			Client: OAuth2Client{
				ID:     os.Getenv("OAUTH2_CLIENT_ID"),
				Secret: os.Getenv("OAUTH2_CLIENT_SECRET"),
				Domain: os.Getenv("OAUTH2_CLIENT_DOMAIN"),
			},
		},
		Queue: Queue{
			Clients: QueueClients{
				RabbitMQ: os.Getenv("RABBIT_MQ_DSN"),
			},
			List: QueueList{
				ToAnalysis: QueueListItem{
					Name: os.Getenv("QUEUE_TO_ANALYSIS"),
				},
				AnalysisResult: QueueListItem{
					Name:         os.Getenv("QUEUE_ANALYSIS_RESULT"),
					WorkersCount: 1,
				},
			},
		},
	}
}
