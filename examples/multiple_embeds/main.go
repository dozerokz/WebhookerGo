package main

import (
	"log"

	"github.com/dozerokz/webhookergo"
)

func main() {
	webhookURL := "https://discord.com/api/webhooks/..."
	avatarURL := "https://..."
	thumbnailURL := "https://..."
	footerURL := "https://..."
	imageURL := "http://..."

	webhook := webhookergo.NewWebhook().
		SetContent("Hello, Discord!").
		SetUsername("WebhooherGo").
		SetAvatarURL(avatarURL)

	embed1 := webhookergo.NewEmbed().
		SetTitle("Embed №1").
		SetDescription("Lorem ipsum dolor sit amet.").
		SetColorRGB(webhookergo.RGB{R: 100, G: 200, B: 300}).
		SetTimestampNow().
		SetImage(imageURL).
		AddField(webhookergo.NewField("Field №1",
			"Lorem ipsum dolor sit amet consectetur adipiscing elit",
			true)).
		AddField(webhookergo.NewField("Field №2",
			"Dolor sit amet consectetur adipiscing elit quisque faucibus.",
			true))

	field := webhookergo.NewEmptyField()
	field.SetName("Field №1").
		SetValue("Adipiscing elit quisque faucibus ex sapien vitae pellentesque.").
		SetInline(true)

	embed2 := webhookergo.NewEmbed().
		SetTitle("Embed №2").
		SetDescription("Lorem ipsum dolor sit amet.").
		SetColorInt(55555).
		SetThumbnail(thumbnailURL).
		SetFooter("Some Footer Text", footerURL).
		AddField(field)

	webhook.AddEmbed(embed1)
	webhook.AddEmbed(embed2)

	err := webhook.Send(webhookURL)
	if err != nil {
		log.Fatal(err)
	}
}
