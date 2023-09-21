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
	// // Configurar as opções do consumidor
	// config := &kafka.ConfigMap{
	// 	"bootstrap.servers": "localhost:9092", // Substitua pelo(s) endereço(s) do(s) broker(s) Kafka.
	// 	"group.id":          "de6dd398-fe30-4d47-a6d4-91116e80918b",
	// 	"auto.offset.reset": "earliest", // Pode ser "earliest" ou "latest" dependendo do comportamento desejado.
	// }

	// // Crie um consumidor Kafka
	// consumer, err := kafka.NewConsumer(config)
	// if err != nil {
	// 	fmt.Printf("Erro ao criar o consumidor: %v\n", err)
	// 	return
	// }
	defer consumer.Close()

	// Subscreva-se a um tópico
	err := consumer.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("Erro ao se inscrever no tópico: %v\n", err)
		return
	}

	defer wg.Done()

	// Capture interrupções do sistema
	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, os.Interrupt)

	// Aguarde mensagens

	log.Info("kafka consumer connected sucessfully")
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
