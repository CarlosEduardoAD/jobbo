package main

import (
	"encoding/json"
	"sync"

	email_handler "github.com/CarlosEduardoAD/jobbo-api/internal/api/app/handlers"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/app/routes"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/domain/email"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/kafka"
	email_repo "github.com/CarlosEduardoAD/jobbo-api/internal/api/infra/repo/smtp"
	"github.com/CarlosEduardoAD/jobbo-api/internal/api/utils"
	kafkaLib "github.com/confluentinc/confluent-kafka-go/kafka"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
)

func main() {
	var wg sync.WaitGroup
	var msgChan = make(chan *kafkaLib.Message)
	dialer := utils.ConnectSMTP("smtp.gmail.com", 587, "karl.devcontato@gmail.com", "ehuf hvxx funu frov")

	config := &kafkaLib.ConfigMap{
		"bootstrap.servers": "localhost:9092", // Substitua pelo(s) endereço(s) do(s) broker(s) Kafka.
		"group.id":          "goalfy-mail",
		"auto.offset.reset": "latest", // Pode ser "earliest" ou "latest" dependendo do comportamento desejado.
	}

	// Crie um consumidor Kafka
	consumer, err := kafkaLib.NewConsumer(config)
	if err != nil {
		log.Error("Erro ao criar o consumidor: %v\n", err)
		return
	}

	wg.Add(1)
	go kafka.StartKafkaConsumer(&wg, consumer, "goalfy-mail", msgChan)

	for msg := range msgChan {
		var err error
		var emailInput *email.Email

		err = json.Unmarshal(msg.Value, emailInput)

		if err != nil {
			log.Error("error unmarshalling email from kafka", err)
			continue
		}

		err = emailInput.Validate()

		if err != nil {
			log.Error(`invalid email from Kafka => ` + string(msg.Value) + " err => " + err.Error())
			continue
		}

		emailToBeDelivered := utils.ConvertToMailMessage(email.NewEmail(emailInput.From, emailInput.To, emailInput.Subject, emailInput.Body))
		emailRepo := email_repo.NewEmailService(dialer)
		email_handler := email_handler.NewEmailHandler(emailRepo)

		err, _ = email_handler.DeliverEmail(emailToBeDelivered)

		if err != nil {
			log.Error("error delivering email from kafka: ", err)
			continue
		}

	}

	e := echo.New()
	routes.EmailRoutes(e)
	e.Logger.Fatal(e.Start(":9292"))

	wg.Wait()
}
