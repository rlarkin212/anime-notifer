package main

import (
	"fmt"
	"log"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/rlarkin212/anime-notifer/domain/service"
	"github.com/rlarkin212/anime-notifer/entities"
	"github.com/rlarkin212/anime-notifer/util"
)

var c *entities.Config
var err error

func main() {
	c, err = util.LoadConfig(".", "config", "yaml")
	if err != nil {
		log.Fatal(fmt.Printf("config load err %s", err.Error()))
	}

	if c.Env == "prod" {
		lambda.Start(HandleRequest)
	} else {
		HandleRequest()
	}
}

func HandleRequest() {
	fetchService := service.NewFetchService(c)
	fetchService.FetchSchedule()
}
