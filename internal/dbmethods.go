package internal

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/JuanMartinCoder/LanChat/internal/database"
	"github.com/jackc/pgx/v5/pgtype"
	amqp "github.com/rabbitmq/amqp091-go"
)

func SaveToDBMessages(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	db *database.Queries,
) error {
	ch, queue, err := DeclareAndBind(conn, exchange, queueName, key)
	if err != nil {
		return fmt.Errorf("could not declare and bind queue: %v", err)
	}

	msgs, err := ch.Consume(
		queue.Name, // queue
		"",         // consumer
		true,       // auto-ack
		false,      // exclusive
		false,      // no-local
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return fmt.Errorf("could not consume messages: %v", err)
	}

	MessageBulk := make([]database.InsertBulkMessagesParams, 0)

	go func() {
		defer ch.Close()
		for msg := range msgs {
			message, err := UnmarshallData(msg.Body)
			if err != nil {
				log.Fatalf("Unable to parse message:%v", err)
			}

			msqToDB := database.InsertBulkMessagesParams{
				NameFrom: message.FROM,
				NameTo:   message.TO,
				Message:  message.MESSAGE,
				CreatedAt: pgtype.Timestamp{
					Time:  time.Now().UTC(),
					Valid: true,
				},
			}

			MessageBulk = append(MessageBulk, msqToDB)

			if len(MessageBulk) > 50 {
				msgRes, err := db.InsertBulkMessages(context.Background(), MessageBulk)
				if err != nil {
					log.Fatal(err)
				}
				log.Printf("Data inserted/Received: %d", msgRes)
				MessageBulk = MessageBulk[:0]
			}
		}
	}()
	return nil
}
