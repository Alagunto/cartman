package src

import (
	"github.com/labstack/echo/v4/middleware"
	"github.com/sirupsen/logrus"
	"gopkg.in/tucnak/telebot.v2"
	"net/http"
	"github.com/labstack/echo/v4"
)

func StartHttpServer(bot *telebot.Bot) {
	e := echo.New()

	// Middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	// Routes
	e.GET("/", hello)

	// Start server
	e.Logger.Fatal(e.Start(":1323"))
}
