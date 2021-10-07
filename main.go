package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"

	tgbot "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/rlarkin212/anime-notifer/helpers"
	"github.com/rlarkin212/anime-notifer/models"
	"github.com/rlarkin212/anime-notifer/subsplease"
	"github.com/rlarkin212/anime-notifer/telegram"
	"github.com/rlarkin212/anime-notifer/util"
	"github.com/robfig/cron/v3"
)

const (
	baseUrl = "https://subsplease.org/api/?f=schedule&h=true"
)

var tgApi = util.GetEnvVar("TG_API_TOKEN")

func main() {
	userSchedule := models.UserSchedule{}
	_ = util.ParseYaml(&userSchedule)

	fmt.Println(userSchedule.CronSchedule)

	cron := cron.New()
	cron.AddFunc(userSchedule.CronSchedule, func() {
		execute(userSchedule)
	})

	errs := make(chan error, 1)
	go cron.Start()

	execute(userSchedule)

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func execute(userSchedule models.UserSchedule) {
	schedule := subsplease.FetchSchedule(baseUrl, userSchedule.TimeZone)
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
	}
}
