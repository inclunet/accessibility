package main

import (
	"log"
	"os"

	accessbot "github.com/inclunet/accessibility/pkg/accessbot"
)

func main() {
	log.Println("Starting accessibility checker bot service...")
	token := os.Getenv("TELEGRAM_TOKEN")

	bot, err := accessbot.New(token)

	if err != nil {
		log.Println(err)
	}

	bot.Handle()
}
