package config

import "github.com/spf13/viper"

type Env struct {
	Debug      bool
	Port       int
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
}

func GetEnv() Env {
	env := Env{
		Debug:      viper.GetBool("DEBUG"),
		Port:       viper.GetInt("PORT"),
		DBHost:     viper.GetString("NEO4J_HOST"),
		DBPort:     viper.GetInt("NEO4J_PORT"),
		DBUser:     viper.GetString("NEO4J_USER"),
		DBPassword: viper.GetString("NEO4J_PASSWORD"),
	}
	return env
}

func Init() {
	viper.AutomaticEnv()
	viper.SetConfigFile(".env")
	viper.ReadInConfig()
}
