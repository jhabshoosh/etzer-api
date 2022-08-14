package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
    Debug       bool
	Port		int
    DBHost   	string
    DBPort   	int
    DBUser   	string
    DBPassword  string
}

func GetEnv() Env {
	var env Env
    err := envconfig.Process("etzer", &env)
	if err != nil {
        log.Fatal(err.Error())
    }
	
    format := "Debug: %v\nPort: %d\nNeo4JHost: %s\nNeo4JPort: %d\nNeo4JUser: %s\nNeo4JPassword: %s\n"
    _, err = fmt.Printf(format, env.Debug, env.Port, env.DBHost, env.DBPort, env.DBUser, env.DBPassword)
    if err != nil {
        log.Fatal(err.Error())
    }

	return env
}