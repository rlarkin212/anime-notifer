package main

import (
	"fmt"
	"log"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rlarkin212/anime-notifer/helpers"
	"github.com/rlarkin212/anime-notifer/models"
	"github.com/rlarkin212/anime-notifer/subsplease"
	"github.com/rlarkin212/anime-notifer/telegram"
	"github.com/rlarkin212/anime-notifer/util"
)

const (
	baseUrl = "https://subsplease.org/api/?f=schedule&h=true"
)

var tgApi = util.GetEnvVar("TG_API_TOKEN")

func main() {
	userSchedule := models.UserSchedule{}
	_ = util.ParseYaml(&userSchedule)

	fmt.Println(userSchedule)

	execute(userSchedule)
}

func execute(userSchedule models.UserSchedule) {
	schedule := subsplease.FetchSchedule(baseUrl, userSchedule.TimeZone)
	fmt.Println(schedule.Schedule)

	inSchedule := []models.ScheduleItem{}
	usMap := helpers.SliceToStrMap(userSchedule.Shows)

	for _, item := range schedule.Schedule {
		if helpers.Contains(usMap, item.Title) {
			inSchedule = append(inSchedule, item)
		}
	}

	if len(inSchedule) > 0 {
		bot, err := tgbot.NewBotAPI(tgApi)
		if err != nil {
			log.Fatal(err.Error())
		}

		telegram.SendMessage(bot, inSchedule)
	} else {
		fmt.Println("no shows today")
	}
}
