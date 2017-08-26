package controllers

import (
	"fmt"
	"github.com/spf13/viper"
)

func init() {
	fmt.Println("Reading in pallet.conf")

	viper.SetConfigName("pallet")
	viper.AddConfigPath("/etc/pallet/")
	viper.AddConfigPath("$HOME/.pallet")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}
}
