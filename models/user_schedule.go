package models

type UserSchedule struct {
	TimeZone         string   `yaml:"time-zone"`
	NotificationType string   `yaml:"notification-type"`
	CronSchedule     string   `yaml:"cron-schedule"`
	Shows            []string `yaml:"shows"`
}
