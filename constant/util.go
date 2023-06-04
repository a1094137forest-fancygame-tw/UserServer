package constant

import (
	"log"

	"github.com/spf13/viper"

	"UserServer/config"
)

var ConnectServer = []string{
	config.UserServerUrl,
}

type Message struct {
	EventType int         `json:"event_type"`
	Data      interface{} `json:"data"`
}

func ReadConfig(configPath string) {
	viper.SetConfigFile(configPath)
	viper.AddConfigPath(".")

	viper.SetDefault("PORT", ":1002")
	viper.SetDefault("RUN_MODE", "debug")

	envs := []string{
		"PORT",
		"RUN_MODE",
	}

	for _, env := range envs {
		if err := viper.BindEnv(env); err != nil {
			log.Println(err)
		}
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}

	config.Initialize()
}
