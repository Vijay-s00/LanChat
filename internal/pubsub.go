package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	amqp "github.com/rabbitmq/amqp091-go"
)

func CreateConnectionRabbitMQ(url_rabbit_mq string) (*amqp.Connection, *amqp.Channel, error) {
	conn, err := amqp.Dial(url_rabbit_mq)
	if err != nil {
		return nil, nil, fmt.Errorf("error connecting to rabbitmq: %v", err)
	}
	log.Println("Connected to RabbitMQ:")
	ch, err := conn.Channel()
	if err != nil {
		return conn, nil, fmt.Errorf("error creating channel rabbitmq: %v", err)
	}
	return conn, ch, nil
}

func PublishJSON[T any](ch *amqp.Channel, exchange, key string, val T) error {
	b, err := json.Marshal(val)
	if err != nil {
		return err
	}

	publish := amqp.Publishing{
		ContentType: "application/json",
		Body:        b,
	}
	err = ch.PublishWithContext(context.Background(), exchange, key, false, false, publish)
	if err != nil {
		return err
	}

	return nil
}

func DeclareAndBind(conn *amqp.Connection, exchange, queueName, key string) (*amqp.Channel, amqp.Queue, error) {
	newChan, err := conn.Channel()
	if err != nil {
		return nil, amqp.Queue{}, err
	}
	queue, err := newChan.QueueDeclare(
		queueName, // name
		false,     // durable
		false,     // delete when unused
		true,      // exclusive
		false,     // no-wait
		nil,
	)
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("could not declare queue: %v", err)
	}

	err = newChan.QueueBind(
		queue.Name, // queue name
		key,        // routing key
		exchange,   // exchange
		false,      // no-wait
		nil,        // args
	)
	if err != nil {
		return nil, amqp.Queue{}, fmt.Errorf("could not bind queue: %v", err)
	}
	return newChan, queue, nil
}

func Subscribe(
	conn *amqp.Connection,
	exchange,
	queueName,
	key string,
	username string,
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

	go func() {
		defer ch.Close()
		for msg := range msgs {
			message, err := UnmarshallData(msg.Body)
			if err != nil {
				log.Fatalf("Unable to parse message:%v", err)
			}
			if message.FROM != username {
				fmt.Printf("%s: %s\n", message.FROM, message.MESSAGE)
			}
		}
	}()
	return nil
}
