package config

import "github.com/spf13/viper"

type Config struct {
	Production         bool   `mapstructure:"production"`
	Port               string `mapstructure:"port"`
	MongoDBURI         string `mapstructure:"mongoDBUri"`
	PasetoSymmetricKey string `mapstructure:"pasetoSymmentricKey"`
	Database           string `mapstructure:"database"`
}

func ReadConfig() (configuration Config, err error) {
	viper.SetConfigType("yaml")
	viper.AddConfigPath("../")
	viper.SetConfigFile("config.yaml")

	err = viper.ReadInConfig()
	if err != nil {
		return
	}

	err = viper.Unmarshal(&configuration)
	return
}
