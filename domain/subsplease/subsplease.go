package subsplease

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rlarkin212/anime-notifer/entities"
	"github.com/rlarkin212/anime-notifer/entities/subsplease"
	"github.com/rlarkin212/anime-notifer/util"
)

type Fetcher interface {
	Fetch() *subsplease.Response
}

type subspleaseFetcher struct {
	config *entities.Config
}

func NewSubsPlease(c *entities.Config) Fetcher {
	return &subspleaseFetcher{
		config: c,
	}
}

func (sp *subspleaseFetcher) Fetch() *subsplease.Response {
	url := fmt.Sprintf("%s&tz=%s", sp.config.SubsPlease.BaseUrl, sp.config.SubsPlease.TimeZone)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(fmt.Sprintf("fetch schedule err : %s", err.Error()))
	}

	response := &subsplease.Response{}
	_ = util.UnmarshallResponseBody(res, response)

	return response
}
