package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/rlarkin212/anime-notifer/helpers"
	"github.com/rlarkin212/anime-notifer/models"
	"github.com/rlarkin212/anime-notifer/spservice"
	"github.com/rlarkin212/anime-notifer/util"
	"github.com/robfig/cron/v3"
)

const (
	baseUrl = "https://subsplease.org/api/?f=schedule&h=true"
)

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

	go func() {
		c := make(chan os.Signal, 1)
		signal.Notify(c, syscall.SIGINT)
		errs <- fmt.Errorf("%s", <-c)
	}()

	fmt.Printf("Terminated %s", <-errs)
}

func execute(userSchedule models.UserSchedule) {
	schedule := spservice.FetchSchedule(baseUrl, userSchedule.TimeZone)
	inSchedule := []models.ScheduleItem{}

	for _, item := range schedule.Schedule {
		if helpers.Contains(userSchedule.Shows, item.Title) {
			inSchedule = append(inSchedule, item)
		}
	}

	fmt.Println(inSchedule)
}
