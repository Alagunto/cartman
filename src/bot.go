package src

import (
	"time"
	tb "gopkg.in/tucnak/telebot.v2"
)

func Bot(token string) *tb.Bot {
	b, err := tb.NewBot(tb.Settings{
		Token:  token,
		Poller: Poller(),
	})
	if err != nil {
		panic(err)
	}

	return b
}

func Poller() tb.Poller {
	return tb.Poller(&tb.LongPoller{Timeout: 15 * time.Second})
}
