package initialize

import (
	"fmt"

	"github.com/marktran77/go-ecomerce-backend-api/global"
	"github.com/spf13/viper"
)

func LoadConfig() {
	viper := viper.New()
	viper.AddConfigPath("./config/")
	viper.SetConfigName("local")
	viper.SetConfigType("yaml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Failed to read configuration %w \n", err))
	}

	fmt.Println("Server Port::", viper.GetInt("server.port"))

	if err = viper.Unmarshal(&global.Config); err != nil {
		panic(fmt.Errorf("Unable to decode configuration %w\n", err))
	}
}
