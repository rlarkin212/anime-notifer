package entities

type Config struct {
	Env        string     `mapstructure:"env"`
	Shows      []string   `mapstructure:"shows"`
	SubsPlease subsPlease `mapstructure:"subsPlease"`
	Telegram   telegram   `mapstructure:"telegram"`
}

type subsPlease struct {
	BaseUrl  string `mapstructure:"baseUrl"`
	TimeZone string `mapstructure:"timeZone"`
}

type telegram struct {
	Token  string `mapstructure:"token"`
	ChatId string `mapstructure:"chatId"`
}
