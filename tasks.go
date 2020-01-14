package main

import (
	tb "gopkg.in/tucnak/telebot.v2"
	"log"
)

func Greet(m *tb.Message) {
	_, err := B.Send(m.Sender, "Hi!")
	if err != nil {
		log.Println(err)
	}
}

func Menu(m *tb.Message) {
	_, err := B.Send(m.Sender, "Open Menu")
	if err != nil {
		log.Println(err)
	}
}
