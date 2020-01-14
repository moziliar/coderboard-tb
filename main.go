package main

import (
	"fmt"
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
	"os"
)

var B *tb.Bot

func main() {
	var (
		port      = os.Getenv("PORT")
		publicURL = os.Getenv("PUBLIC_URL")
		token     = os.Getenv("TOKEN")
	)

	wh := &tb.Webhook{
		Listen:   ":" + port,
		Endpoint: &tb.WebhookEndpoint{PublicURL: publicURL},
	}

	pref := tb.Settings{
		Token:  token,
		Poller: wh,
	}

	B, err := tb.NewBot(pref)
	if err != nil {
		log.Fatalln("Fail to build", err)
	}

	// This button will be displayed in user's
	// reply keyboard.
	replyBtn := tb.ReplyButton{Text: "🌕 Button #1"}
	replyKeys := [][]tb.ReplyButton{
		[]tb.ReplyButton{replyBtn},
		// ...
	}

	// And this one — just under the message itself.
	// Pressing it will cause the client to send
	// the bot a callback.
	//
	// Make sure Unique stays unique as it has to be
	// for callback routing to work.
	inlineBtn := tb.InlineButton{
		Unique: "sad_moon",
		Text:   "🌚 Button #2",
	}
	inlineKeys := [][]tb.InlineButton{
		[]tb.InlineButton{inlineBtn},
		// ...
	}

	B.Handle(&replyBtn, func(m *tb.Message) {
		_, err := B.Send(m.Sender, "Reply pressed")
		if err != nil {
			fmt.Println(err)
		}
	})

	B.Handle(&inlineBtn, func(c *tb.Callback) {
		// on inline button pressed (callback!)

		// always respond!
		err := B.Respond(c, &tb.CallbackResponse{
			Text: "Wow",
		})
		if err != nil {
			fmt.Println(err)
		}
	})

	// Command: /start <PAYLOAD>
	B.Handle("/start", func(m *tb.Message) {
		if !m.Private() {
			return
		}

		_, err := B.Send(m.Sender, "Hello!", &tb.ReplyMarkup{
			ReplyKeyboard:  replyKeys,
			InlineKeyboard: inlineKeys,
		})
		if err != nil {
			fmt.Println(err)
		}
	})

	B.Handle("/hello", Greet)
	B.Handle("/menu", Menu)

	B.Start()
}
