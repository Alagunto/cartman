package src

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetChatToken(chatId int64) string {
	chat := strconv.Itoa(int(chatId))

	var token string
	err := db.Read("chats", chat, &token)
	if err == nil {
		return token
	}

	token = RandomToken(16)
	err = db.Write("chats", chat, token)
	if err != nil {
		logrus.Error(err)
	}

	return token
}
