package main

import (
	"log"
	"os"

	accessbot "github.com/inclunet/accessibility/pkg/accessbot"
)

func main() {
	log.Println("Starting accessibility checker bot service...")
	bot, err := accessbot.New(os.Getenv("TELEGRAM_APITOKEN"))

	if err != nil {
		log.Println(err)
	}

	bot.Handle()
}
