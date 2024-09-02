package main

import (
	"flag"
	"log"

	clientTg "test/clients/telegram"
	eventConsumer "test/consumer/event-consumer"
	eventTg "test/events/telegram"
	"test/storage/files"
)

const (
	tgBotHost   = "api.telegram.org"
	storagePath = "storage"
	batchSize   = 100
)

func main() {

	eventsProcessor := eventTg.New(
		clientTg.New(tgBotHost, mustToken()),
		files.New(storagePath))

	log.Print("service started")

	consumer := eventConsumer.New(eventsProcessor, eventsProcessor, batchSize)
	if err := consumer.Start(); err != nil {
		log.Fatal()
	}
}

func mustToken() string {
	token := flag.String(
		"token-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
