package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/gommon/log"
)

// StartKafkaConsumer starts a Kafka consumer.
//
// It takes the following parameters:
//   - wg: a pointer to a sync.WaitGroup object
//   - consumer: a pointer to a kafka.Consumer object
//   - topic: a string representing the Kafka topic
//   - eventsChan: a channel of type *kafka.Message
//
// It does not return any value.
func StartKafkaConsumer(wg *sync.WaitGroup, consumer *kafka.Consumer, topic string, eventsChan chan *kafka.Message) {

	defer consumer.Close()

	err := consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Erro ao se inscrever no tópico: %v\n", err)
		return
	}

	defer wg.Done()

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	log.Info("repository: kafka consumer connected sucessfully")
	run := true
	for run {
		msg, err := consumer.ReadMessage(-1)
		if err == nil {
			eventsChan <- msg
		}
	}
}
