package main

import (
	"github.com/s4kibs4mi/rapunzel-blog/cmd"
	"github.com/spf13/viper"
	"fmt"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func main() {
	viper.SetConfigName("config")
	viper.SetConfigType("json")
	viper.AddConfigPath("etc")
	viper.AddConfigPath(".")
	viper.AddConfigPath("/etc/config")
	viper.AddConfigPath("/app/config")
	err := viper.ReadInConfig()
	if err != nil {
		fmt.Println("Couldn't read config", err)
		return
	}
	cmd.Serve()
}
