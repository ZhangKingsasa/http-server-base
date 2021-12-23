package options

import (
	log "github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type ServerRunOptions struct {
	Port           int      `mapstructure:"port"`
	DatabaseUrl    string   `mapstructure:"database_url"`
}

const configPath = "./conf/config.yaml"

func NewServerRunOptions() *ServerRunOptions {
	viper.SetConfigFile(configPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.WithFields(log.Fields{
			"configPath": configPath,
			"error": err.Error(),
		}).Fatal("Read config file failed")
	}
	conf := &ServerRunOptions{}
	if err := viper.Unmarshal(conf); err != nil {
		log.WithFields(log.Fields{
			"configPath": configPath,
			"error": err.Error(),
		}).Fatal("Unmarshal config failed")
	}
	return conf
}