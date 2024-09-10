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
	storagePath = "files_storage"
	batchSize   = 100
)

/*

7430142777:AAE00lPQDOWUC4OmUTYedOLR-1HVLkclCus

go build in terminal

./test -tg-bot-token '7430142777:AAE00lPQDOWUC4OmUTYedOLR-1HVLkclCus'

*/

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
		"tg-bot-token",
		"",
		"token for access to telegram bot")

	flag.Parse()

	if *token == "" {
		log.Fatal("token is not specified")
	}

	return *token
}
