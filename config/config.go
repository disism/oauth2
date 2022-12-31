package config

import (
	"fmt"
	"github.com/spf13/viper"
	"time"
)

// InitConfig Initializing the configuration file.
func InitConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("./config")
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			return fmt.Errorf("CONFIG FILE NOT FOUND")
		} else {
			return err
		}
	}
	return nil
}

func GetAuthTokenSecret() string {
	return viper.GetString("authentication.token.secret")
}

func GetAuthTokenExpir() time.Duration {
	return time.Duration(viper.GetInt("authentication.token.expired")) * 24 * time.Hour
}

func GetDomain() string {
	return viper.GetString("domain")
}

func GetRedisAddr() string {
	return viper.GetString("redis.addr")
}

func GetRedisPassword() string {
	return viper.GetString("redis.password")
}
