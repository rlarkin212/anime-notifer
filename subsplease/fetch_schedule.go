package subsplease

import (
	"fmt"
	"log"
	"net/http"

	"github.com/rlarkin212/anime-notifer/models"
	"github.com/rlarkin212/anime-notifer/util"
)

func FetchSchedule(baseUrl string, timeZone string) models.ScheduleResponse {
	url := fmt.Sprintf("%s&tz=%s", baseUrl, timeZone)

	res, err := http.Get(url)
	if err != nil {
		log.Fatal(fmt.Sprintf("fetch schedule err : %s", err.Error()))
	}

	schedule := models.ScheduleResponse{}
	_ = util.UnmarshallResponseBody(res, &schedule)

	return schedule
}
