package util

import (
	"io/ioutil"
	"log"

	"github.com/rlarkin212/anime-notifer/models"
	"gopkg.in/yaml.v2"
)

func ParseYaml(userSchedule *models.UserSchedule) *models.UserSchedule {
	yamlFile, err := ioutil.ReadFile("schedule.yaml")
	if err != nil {
		log.Fatal(err.Error())
	}

	err = yaml.Unmarshal(yamlFile, userSchedule)
	if err != nil {
		log.Fatal(err.Error())
	}

	return userSchedule
}
