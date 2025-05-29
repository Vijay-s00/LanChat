package main

import (
	"context"
	"log"
	"os"
	"sync"

	"github.com/JuanMartinCoder/LanChat/internal"
	"github.com/JuanMartinCoder/LanChat/internal/database"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	url_rabbit_mq := os.Getenv("URL_RMQ")

	conn, ch, err := internal.CreateConnectionRabbitMQ(url_rabbit_mq)
	if err != nil {
		log.Fatal(err)
	}
	err = ch.ExchangeDeclare("group_chat", "fanout", false, false, false, false, nil)
	if err != nil {
		log.Fatalf("Error declaring exchange: %v", err)
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not set")
	}

	connDB, err := pgx.Connect(context.Background(), dbUrl)
	if err != nil {
		log.Fatal(err)
	}
	defer connDB.Close(context.Background())

	dbQueries := database.New(connDB)

	var wg sync.WaitGroup
	wg.Add(1)
	err = internal.SaveToDBMessages(conn, "group_chat", "message_to_db", "", dbQueries)
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
