package telegram

import (
	"fmt"
	"strconv"
	"strings"
	"time"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rlarkin212/anime-notifer/models"
	"github.com/rlarkin212/anime-notifer/util"
)

const (
	iso8601 = "2006-01-02T15:04:05-0700"
)

var tgChatId, _ = strconv.ParseInt(util.GetEnvVar("TG_CHAT_ID"), 10, 64)

func SendMessage(bot *tgbot.BotAPI, shows []models.ScheduleItem) {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("--- AIRING TODAY--- %s\n", time.Now().Format(iso8601)))

	for _, show := range shows {
		b.WriteString(fmt.Sprintf("%s\n", show.Title))
		b.WriteString(fmt.Sprintf("%s\n", show.Time))
		b.WriteString(fmt.Sprintf("https://subsplease.org%s\n", show.ImageURL))
		b.WriteString("-----------------")
	}

	msg := tgbot.NewMessage(tgChatId, b.String())
	bot.Send(msg)
}
