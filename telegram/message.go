package telegram

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rlarkin212/anime-notifer/models"
	"github.com/rlarkin212/anime-notifer/util"
)

const (
	iso8601 = "2006-01-02"
)

func SendMessage(bot *tgbot.BotAPI, shows []models.ScheduleItem) {
	tgChatId, err := strconv.ParseInt(util.GetEnvVar("TG_CHAT_ID"), 10, 64)
	if err != nil {
		log.Fatal(fmt.Sprintf("parse int err: %s", err))
	}

	var b strings.Builder
	b.WriteString(fmt.Sprintf("AIRING TODAY : %s\n", time.Now().Format(iso8601)))

	for _, show := range shows {
		b.WriteString("-----------------\n")
		b.WriteString(fmt.Sprintf("%s\n", show.Title))
		b.WriteString(fmt.Sprintf("%s\n", show.Time))
		b.WriteString("-----------------\n")
	}

	msg := tgbot.NewMessage(tgChatId, b.String())
	bot.Send(msg)
}
