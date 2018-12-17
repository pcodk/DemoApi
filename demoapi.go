package demoapi

import (
	"os"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

)

//Initialize a new logger instance
func NewLogger() *logrus.Logger {
	logger := logrus.New()

	return logger
}

func LoadConfig() {
	configName := "dev"
	if os.Getenv("PCO_ENV") != "" {
		configName = os.Getenv("PCO_ENV")
	}

	viper.SetConfigName(configName)  // name of config file (without extension)'
	viper.AddConfigPath("./config/") // path to look for the config file in
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		logrus.Fatal("Unable to load config file")
	}

}
