package util

import (
	"os"
	"strings"

	"github.com/rlarkin212/anime-notifer/entities"
	"github.com/spf13/viper"
)

func LoadConfig(path, name, extension string) (*entities.Config, error) {
	config := &entities.Config{}

	viper.AddConfigPath(path)
	viper.SetConfigName(name)
	viper.SetConfigType(extension)
	viper.AutomaticEnv()

	err := viper.ReadInConfig()
	if err != nil {
		return nil, err
	}

	replaceEnvVars("${", "}")

	err = viper.Unmarshal(config)
	if err != nil {
		return nil, err
	}

	return config, nil
}

func replaceEnvVars(prefix, suffix string) {
	for _, k := range viper.AllKeys() {
		v := viper.GetString(k)
		if strings.HasPrefix(v, prefix) && strings.HasSuffix(v, suffix) {
			viper.Set(k, os.ExpandEnv(v))
		}
	}
}
