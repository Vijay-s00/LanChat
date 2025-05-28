package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/JuanMartinCoder/LanChat/internal"
	"github.com/joho/godotenv"
	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	fmt.Println("Welcome to Lan Chat")

	godotenv.Load()
	url_rabbit_mq := os.Getenv("URL_RMQ")

	conn, err := amqp.Dial(url_rabbit_mq)
	if err != nil {
		log.Fatalf("Error connecting to RabbitMQ: %v", err)
	}
	log.Println("Connected to RabbitMQ:")
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Error getting Channel: %v", err)
	}
	defer ch.Close()

	fmt.Println("Enter a name:")
	username := internal.GetInput()
	fmt.Printf("Welcome %s to the Lan Chat\n", username)

	q, err := ch.QueueDeclare(username, false, false, true, false, nil)
	if err != nil {
		log.Fatalf("error declaring a queue: %v", err)
	}

	err = ch.ExchangeDeclare("group_chat", "fanout", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error declaring exchange: %v", err)
	}

	err = ch.QueueBind(q.Name, "", "group_chat", false, nil)
	if err != nil {
		log.Fatalf("Error binding a queue: %v", err)
	}

	var wg sync.WaitGroup
	wg.Add(1)
	var MessageToSend internal.Message
	go func() {
		defer wg.Done()
		for {
			MessageToSend.FROM = username
			MessageToSend.TO = "group_chat"
			MessageToSend.Fan = "group_chat"
			MessageToSend.Key = ""
			Input := internal.GetInput()
			MessageToSend.MESSAGE = Input

			MessageToJson, err := internal.MarshallData(MessageToSend)
			if err != nil {
				log.Fatalf("Error marshalling message %v", err)
			}

			err = ch.Publish(MessageToSend.Fan, MessageToSend.Key, false, false, amqp.Publishing{
				ContentType: "application/json",
				Body:        MessageToJson,
			})
			if err != nil {
				log.Fatalf("Message unable to send %v", err)
			}
		}
	}()

	msgs, err := ch.Consume(q.Name, "", true, false, false, false, nil)
	if err != nil {
		log.Fatalf("Unable to consume messages:%v", err)
	}
	wg.Add(1)
	go func() {
		defer wg.Done()
		for msg := range msgs {
			message, err := internal.UnmarshallData(msg.Body)
			if err != nil {
				log.Fatalf("Unable to parse message:%v", err)
			}
			if message.FROM != username {
				fmt.Printf("%s: %s\n", message.FROM, message.MESSAGE)
			}
		}
	}()
	wg.Wait()
}
