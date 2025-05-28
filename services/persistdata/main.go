package main

import (
	"database/sql"
	"log"
	"os"
	"sync"

	"github.com/JuanMartinCoder/LanChat/internal"
	"github.com/JuanMartinCoder/LanChat/internal/database"
	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	godotenv.Load()
	url_rabbit_mq := os.Getenv("URL_RMQ")

	conn, _, err := internal.CreateConnectionRabbitMQ(url_rabbit_mq)
	if err != nil {
		log.Fatal(err)
	}

	dbUrl := os.Getenv("DB_URL")
	if dbUrl == "" {
		log.Fatal("DB_URL is not set")
	}

	db, err := sql.Open("postgres", dbUrl)
	if err != nil {
		log.Fatalf("Error opening database: %s", err)
	}
	dbQueries := database.New(db)

	var wg sync.WaitGroup
	wg.Add(1)
	err = internal.SaveToDBMessages(conn, "group_chat", "message_to_db", "", dbQueries)
	if err != nil {
		log.Fatal(err)
	}
	wg.Wait()
}
