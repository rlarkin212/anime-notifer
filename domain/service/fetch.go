package service

import (
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/rlarkin212/anime-notifer/domain/subsplease"
	"github.com/rlarkin212/anime-notifer/entities"
	sp "github.com/rlarkin212/anime-notifer/entities/subsplease"
	"github.com/rlarkin212/anime-notifer/helpers"
	"github.com/rlarkin212/anime-notifer/transport/telegram"
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
	m := helpers.SliceToStrMap(fs.config.Shows)

	for _, s := range schedule.Schedule {
		if _, ok := m[s.Title]; ok {
			items = append(items, s)
		}
	}

	log.Println(fmt.Printf("%d shows on your schedule today", len(items)))

	return items
}

func buildMessage(items []sp.Item) string {
	var b strings.Builder
	b.WriteString(fmt.Sprintf("AIRING TODAY : %s\n", time.Now().Format(iso8601)))

	for _, item := range items {
		b.WriteString("-----------------\n")
		b.WriteString(fmt.Sprintf("%s\n", item.Title))
		b.WriteString(fmt.Sprintf("%s\n", item.Time))
		b.WriteString("-----------------\n")
	}

	return b.String()
}
