package config

import (
	"fmt"

	"github.com/spf13/viper"
)

func initConfig() {
	viper.SetConfigType("yaml")
	viper.SetConfigName("config") // name of config file (without extension)
	viper.AddConfigPath("config/")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Config file not found...")
	}
}

/**
 * Get config by key
 *
 */
func Get(key string) string {
	initConfig()
	return viper.GetString(key)
}
