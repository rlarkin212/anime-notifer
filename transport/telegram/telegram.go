package telegram

import (
	"log"
	"strconv"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rlarkin212/anime-notifer/entities"
)

type Sender interface {
	Send(msg string)
}

type telegram struct {
	config *entities.Config
	bot    *tgbot.BotAPI
}

func NewTelegram(c *entities.Config) Sender {
	bot, err := tgbot.NewBotAPI(c.Telegram.Token)
	if err != nil {
		log.Printf("telegram bot err: %s", err.Error())
	}

	return &telegram{
		config: c,
		bot:    bot,
	}
}

func (t *telegram) Send(msg string) {
	chatId, err := strconv.ParseInt(t.config.Telegram.ChatId, 10, 64)
	if err != nil {
		log.Printf("parse int err: %s", err.Error())
	}

	message := tgbot.NewMessage(chatId, msg)
	t.bot.Send(message)
}
