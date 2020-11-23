package src

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/tucnak/telebot.v2"
	"net/http"
)

var bot *telebot.Bot

func send(c echo.Context) error {
	token := c.Param("token")
	message := c.QueryParam("message")


	chat := &telebot.Chat{ID:GetChatByToken(token)}

	_, err := bot.Send(chat, message)
	if err != nil {
		return c.String(http.StatusInternalServerError, err.Error())
	}

	return c.String(http.StatusOK, "Ok bitch, paper in de flight")
}

func StartHttpServer(b *telebot.Bot) {
	bot = b

	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/:token/send", send)

	// Start server
	e.Logger.Fatal(e.Start(":1984"))
}
