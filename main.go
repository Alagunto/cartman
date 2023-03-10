package main

import (
	"math/rand"
	"os"
	"time"

	"github.com/Alagunto/cartman/src"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	tb "gopkg.in/tucnak/telebot.v2"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	if err := godotenv.Load(); err != nil {
		logrus.Fatal(err)
	}

	b := src.Bot(os.Getenv("TELEGRAM_TOKEN"))

	b.Handle(tb.OnAddedToGroup, func(m *tb.Message) {
		b.Send(m.Chat, "Привет! Меня нужно запустить через /start")
	})

	b.Handle("/start@"+b.Me.Username, func(m *tb.Message) {
		b.Reply(m, src.GetTokenForChat(m.Chat.ID))
	})

	err := b.SetCommands([]tb.Command{{Text: "start", Description: "Активировать для чата и получить токен"}})
	if err != nil {
		panic(err)
	}

	go src.StartHttpServer(b)

	b.Start()
}
