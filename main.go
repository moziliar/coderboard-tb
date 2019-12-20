package main

import (
	"log"
	"os"

	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	webhook := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: webhook,
	}

	b, err := tb.NewBot(pref)
	if err != nil {
		log.Fatalln("Fail to build", err)
	}

	b.Handle("/hello", func(m *tb.Message) {
		_, err := b.Send(m.Sender, "Hi!")
		if err != nil {
			log.Println(err)
		}
	})

	b.Start()
}
