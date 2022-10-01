package main

import (
	"github.com/spf13/viper"
	"mocae/internal/database"
	"mocae/internal/logger"
	"mocae/internal/srvhttp"
)

func main() {

	viper.AddConfigPath("./config/")
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.ReadInConfig()
	logger.InitLog()

	logger.LogMsg("Starting process", "info")
	logger.LogMsg("Read configfile config.json", "info")

	database.DatabaseConnect()
	database.CheckDbUpdate()
	srvhttp.Instance = database.DBInstance
	srvhttp.DBError = database.DBError
	srvhttp.HandleRequests()

}
