package utils

import "github.com/spf13/viper"

func GetEnv(key string) (value string) {
	return viper.GetString(key)
}
