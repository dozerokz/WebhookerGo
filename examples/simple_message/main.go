package main

import (
	"log"

	"github.com/dozerokz/webhookergo"
)

func main() {
	webhookURL := "https://discord.com/api/webhooks/..."

	err := webhookergo.SendSimple(webhookURL, "Hello, Discord!")
	if err != nil {
		log.Fatal(err)
	}
}
