package entities

type Config struct {
	Env         string       `mapstructure:"env"`
	Shows       []string     `mapstructure:"shows"`
	SubsPlease  subsPlease   `mapstructure:"subsPlease"`
	Telegram    telegram     `mapstructure:"telegram"`
	ManualShows []manualShow `mapstructure:"manualShows"`
}

type manualShow struct {
	Name string `mapstructure:"name"`
	Day  string `mapstructure:"day"`
	Time string `mapstructure:"time"`
}

type subsPlease struct {
	BaseUrl  string `mapstructure:"baseUrl"`
	TimeZone string `mapstructure:"timeZone"`
}

type telegram struct {
	Token  string `mapstructure:"token"`
	ChatId string `mapstructure:"chatId"`
}
