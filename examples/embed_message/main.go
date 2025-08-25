package main

import (
	"log"
	"time"

	"github.com/dozerokz/webhookergo"
)

func main() {
	webhookURL := "https://discord.com/api/webhooks/..."

	embed := webhookergo.NewEmbed().
		SetTitle("Server Alert").
		SetDescription("The server has restarted successfully").
		SetColorHex("#FF5733").
		SetTimestamp(time.Now()). // Same as .SetTimestampNow()
		AddField(webhookergo.NewField("Status", "Success", true)).
		AddField(webhookergo.NewField("Duration", "3m45s", true))

	err := webhookergo.SendEmbed(webhookURL, embed)
	if err != nil {
		log.Fatal(err)
	}
}
