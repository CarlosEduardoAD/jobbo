package main

import (
	"sync"

	"github.com/CarlosEduardoAD/jobbo-api/internal/api/app/routes"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/kafka"
	kafkaLib "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	var wg sync.WaitGroup

	config := &kafkaLib.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Substitua pelo(s) endereço(s) do(s) broker(s) Kafka.
		"group.id":          "de6dd398-fe30-4d47-a6d4-91116e80918b",
		"auto.offset.reset": "earliest", // Pode ser "earliest" ou "latest" dependendo do comportamento desejado.
	}

	// Crie um consumidor Kafka
	consumer, err := kafkaLib.NewConsumer(config)
	if err != nil {
		log.Error("Erro ao criar o consumidor: %v\n", err)
		return
	}

	wg.Add(1)
	go kafka.StartKafkaConsumer(&wg, consumer, "goalfy-mail")

	e := echo.New()
	routes.EmailRoutes(e)
	e.Logger.Fatal(e.Start(":9292"))

	wg.Wait()
}
