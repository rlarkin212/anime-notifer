package models

type UserSchedule struct {
	TimeZone string   `yaml:"time-zone"`
	Shows    []string `yaml:"shows"`
}
