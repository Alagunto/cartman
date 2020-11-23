package src

import (
	"github.com/sirupsen/logrus"
	"strconv"
)

func GetTokenForChat(chatId int64) string {
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

	err = db.Write("chats_by_token", token, chat)
	if err != nil {
		logrus.Error(err)
	}


	return token
}

func GetChatByToken(token string) int64 {
	var chat string
	err := db.Read("chats_by_token", token, &chat)
	if err == nil {
		cid, err := strconv.Atoi(chat)
		if err == nil {
			return int64(cid)
		}
	}

	return -1
}
