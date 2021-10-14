package util

import (
	"fmt"
	"log"

	"github.com/rlarkin212/anime-notifer/models"
	"gopkg.in/yaml.v2"
)

func ParseYaml(file []byte, userSchedule *models.UserSchedule) {
	err := yaml.Unmarshal(file, userSchedule)
	if err != nil {
		log.Fatal(fmt.Sprintf("yaml parse err: %s", err.Error()))
	}
}
