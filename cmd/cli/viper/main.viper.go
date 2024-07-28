package main

import (
	"fmt"

	"github.com/spf13/viper"
)

type Server struct {
	Port int `mapstructure:"port"`
}

type Databases struct {
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	Host     string `mapstructure:"host"`
}

type Config struct {
	Server    Server
	Databases []Databases `mapstructure:"database"`
}

func main() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))

	var config Config
	if err = viper.Unmarshal(&config); err != nil {
		panic(fmt.Errorf("Unable to decode configuration %w\n", err))
	}

	fmt.Println("Config Port::", config.Server.Port)
	for _, value := range config.Databases {
		fmt.Printf("Databases user: %s, password: %s, host: %s \n", value.User, value.Password, value.Host)
	}

}
