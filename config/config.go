package config

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/viper"
)

type AppConfig struct {
	MayarToken    string
	Port          int
	WAServiceHost string
	WAServicePort string
	Admin         string
}

func InitConfig() *AppConfig {
	return ReadEnv()
}

func ReadEnv() *AppConfig {
	app := AppConfig{}

	isRead := true

	if val, found := os.LookupEnv("MAYARTOKEN"); found {
		app.MayarToken = val
		isRead = false
	}
	if val, found := os.LookupEnv("WASERVICEHOST"); found {
		app.WAServiceHost = val
		isRead = false
	}
	if val, found := os.LookupEnv("WASERVICEPORT"); found {
		app.WAServicePort = val
		isRead = false
	}
	if val, found := os.LookupEnv("ADMIN"); found {
		app.MayarToken = val
		isRead = false
	}
	if val, found := os.LookupEnv("PORT"); found {
		cnv, _ := strconv.Atoi(val)
		app.Port = cnv
		isRead = false
	}

	if isRead {
		viper.AddConfigPath(".")
		viper.SetConfigName("local")
		viper.SetConfigType("env")

		err := viper.ReadInConfig()
		if err != nil {
			log.Println("error read config : ", err.Error())
			return nil
		}
		err = viper.Unmarshal(&app)
		if err != nil {
			log.Println("error parse config : ", err.Error())
			return nil
		}
	}

	return &app
}
