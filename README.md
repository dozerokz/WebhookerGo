# WebhookerGo — Discord Webhook Library for Go

WebhookerGo is a lightweight and flexible Go library for sending messages and rich embeds to Discord via webhooks. It
allows developers to integrate Discord notifications, alerts, logs, or automated messages into their Go applications
quickly and efficiently.

![Discord Webhook](https://i.imgur.com/SuIIrhR.png)

## Features

- Send plain text messages or rich embeds to Discord webhooks.
- Chainable builder-style API for easy embed and message creation.
- Support for setting content, username, avatar URL, TTS, and multiple embeds.
- Convenience functions: SendSimple for text-only messages, SendEmbed for sending a single embed.
- Color support via Hex, RGB, or integers

## Installation

Use **go get** to install the package:
```
go get github.com/dozerokz/webhookergo
```

## Setup

You must first configure a Webhook on a Discord server before you can use this package. Instructions can be found
on [Discord's support website](https://support.discord.com/hc/en-us/articles/228383668).

You can read more about webhooks structure [here](https://discord.com/developers/docs/resources/webhook).

## Usage Examples

### Sending a Simple Message

```
err := webhookergo.SendSimple("https://discord.com/api/webhooks/...", "Hello, Discord!")
if err != nil {
    log.Fatal(err)
}
```

### Sending an Embed

```
embed := webhookergo.NewEmbed().
    SetTitle("Server Alert").
    SetDescription("The server has restarted successfully").
    SetColorHex("#FF5733").
    SetTimestampNow()

err := webhookergo.SendEmbed("https://discord.com/api/webhooks/...", embed)
if err != nil {
    log.Fatal(err)
}
```

### Creating webhook with Embed with Fields

```
webhook := webhookergo.NewWebhook().
	SetContent("Hello, Discord!")
		
embed := webhookergo.NewEmbed().
    SetTitle("Build Report").
    AddField(webhookergo.NewField("Status", "Success", true)).
    AddField(webhookergo.NewField("Duration", "3m45s", true))

webhook.AddEmbed(embed1)

err := webhook.Send("https://discord.com/api/webhooks/...")
if err != nil {
	log.Fatal(err)
}
```

For more detailed examples, check out the [examples](examples) folder.

## License

This project is open-source. You can use, modify, and distribute it under the [MIT License](LICENSE).
