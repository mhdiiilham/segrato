package config

import "github.com/spf13/viper"

type Config struct {
	Production         bool   `mapstructure:"production"`
	Port               Port   `mapstructure:"port"`
	MongoDBURI         string `mapstructure:"mongoDBUri"`
	PasetoSymmetricKey string `mapstructure:"pasetoSymmentricKey"`
	Database           string `mapstructure:"database"`
}

type Port struct {
	Auth string `mapstructure:"auth"`
	Game string `mapstructure:"game"`
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
