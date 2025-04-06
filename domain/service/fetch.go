package service

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/rlarkin212/anime-notifer/domain/subsplease"
	"github.com/rlarkin212/anime-notifer/entities"
	sp "github.com/rlarkin212/anime-notifer/entities/subsplease"
	"github.com/rlarkin212/anime-notifer/transport/telegram"
	"github.com/rlarkin212/glinq"
)

const (
	iso8601 = "2006-01-02"
)

type FetchService interface {
	FetchSchedule()
	notify(msg string)
}

type fetchService struct {
	subsPlease subsplease.Fetcher
	telegram   telegram.Sender
	config     *entities.Config
}

func NewFetchService(c *entities.Config) FetchService {
	return &fetchService{
		subsPlease: subsplease.NewSubsPlease(c),
		telegram:   telegram.NewTelegram(c),
		config:     c,
	}
}

func (fs *fetchService) FetchSchedule() {
	schedule := fs.subsPlease.Fetch()
	inSchedule := fs.checkSchedule(schedule)

	if len(inSchedule) > 0 {
		msg := buildMessage(inSchedule)
		fs.notify(msg)
	}
}

func (fs *fetchService) notify(msg string) {
	fs.telegram.Send(msg)
}

func (fs *fetchService) checkSchedule(schedule *sp.Response) []sp.Item {
	items := []sp.Item{}

	for _, s := range schedule.Schedule {
		for _, cs := range fs.config.Shows {
			if s.Title == cs.Title {
				s.Source = cs.Source

				items = append(items, s)
			}
		}
	}

	currentDay := time.Now().Weekday().String()
	for _, x := range fs.config.ManualShows {
		if x.Day == currentDay {
			items = append(items, sp.Item{
				Title:  x.Title,
				Time:   x.Time,
				Source: x.Source,
			})
		}
	}

	orderedItems := glinq.OrderBy(items, func(x sp.Item) string {
		return x.Time
	})

	log.Printf("%d shows on your schedule today ", len(orderedItems))

	return orderedItems
}

func buildMessage(items []sp.Item) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("AIRING TODAY : %s\n", time.Now().Format(iso8601)))
	b.WriteString("-----------------\n")

	for _, item := range items {
		b.WriteString(fmt.Sprintf("%s\n", item.Title))
		b.WriteString(fmt.Sprintf("%s\n", item.Time))
		b.WriteString(fmt.Sprintf("%s\n", item.Source))
		b.WriteString("-----------------\n")
	}

	return b.String()
}
