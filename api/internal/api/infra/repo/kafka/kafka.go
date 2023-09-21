package kafka

import (
	"fmt"
	"os"
	"os/signal"
	"sync"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/gommon/log"
)

func StartKafkaConsumer(wg *sync.WaitGroup, consumer *kafka.Consumer, topic string) {
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
		select {
		case ev := <-consumer.Events():
			switch e := ev.(type) {
			case *kafka.Message:
				fmt.Printf("Mensagem recebida em tópico %s:\n%s\n", e.TopicPartition, string(e.Value))
			case kafka.Error:
				fmt.Printf("Erro no consumidor: %v\n", e)
			}
		case <-sigchan:
			fmt.Println("Recebida interrupção do sistema. Fechando o consumidor...")
			run = false
		}
	}
}
