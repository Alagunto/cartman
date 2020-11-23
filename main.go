package main

import (
	"cartman/src"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
	"os"
)

func main() {
	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	b := src.Bot(os.Getenv("TELEGRAM_TOKEN"))

	b.Handle(tb.OnAddedToGroup, func (m *tb.Message) {
		b.Send(m.Chat, "Привет! Меня нужно запустить через /start")
	})

	b.Handle("/start@" + b.Me.Username, func(m *tb.Message) {
		b.Reply(m, src.GetChatToken(m.Chat.ID))
	})

	err := b.SetCommands([]tb.Command{{Text: "start", Description: "Активировать для чата и получить токен"}})
	if err != nil {
		panic(err)
	}

	b.Start()
}