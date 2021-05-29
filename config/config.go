package config

import (
	"fmt"
	"github.com/spf13/viper"
)

var Config *viper.Viper

func init() {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./config")
	err := v.ReadInConfig() // Find and read the config file
	if err != nil {         // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
	fmt.Println("INIT CONFIG")
	Config = v
}
