package config

import "github.com/spf13/viper"

type Config struct {
	Port string `mapstructure:"PORT"`
}

func LoadConfig() (c Config, err error) {
	viper.AddConfigPath("./pkg/config/dev")
	viper.SetConfigName("config")
	viper.SetConfigType("env")

	viper.AutomaticEnv()
	err = viper.ReadInConfig()
	if err != nil {
		return
	}
	err = viper.Unmarshal(&c)
	return

}
