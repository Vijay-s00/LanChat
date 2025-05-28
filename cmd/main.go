package main

import (
	"fmt"
	"log"
	"os"
	"sync"

	"github.com/JuanMartinCoder/LanChat/internal"
	"github.com/joho/godotenv"
)

func main() {
	fmt.Println("Welcome to Lan Chat")

	godotenv.Load()
	url_rabbit_mq := os.Getenv("URL_RMQ")

	conn, ch, err := internal.CreateConnectionRabbitMQ(url_rabbit_mq)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Enter a name:")
	username := internal.GetInput()
	fmt.Printf("Welcome %s to the Lan Chat\n", username)

	err = ch.ExchangeDeclare("group_chat", "fanout", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error declaring exchange: %v", err)
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

			err = internal.PublishJSON(ch, MessageToSend.Fan, MessageToSend.Key, MessageToSend)
			if err != nil {
				log.Fatalf("Message unable to send %v", err)
			}
		}
	}()

	err = internal.Subscribe(conn, "group_chat", username, "", username)
	if err != nil {
		log.Fatal(err)
	}

	wg.Wait()
}
