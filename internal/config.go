package config

import (
	"fmt"
	"log"

	"github.com/kelseyhightower/envconfig"
)

type Env struct {
    Debug       	bool
	Port			int
    Neo4JHost   	string
    Neo4JPort   	int
    Neo4JUser   	string
    Neo4JPassword   string
	Neo4JProtocol	string
}

func GetEnv() Env {
	var env Env
    err := envconfig.Process("etzer", &env)
	if err != nil {
        log.Fatal(err.Error())
    }
	
    format := "Debug: %v\nPort: %d\nNeo4JProtocol: %s\nNeo4JHost: %s\nNeo4JPort: %d\nNeo4JUser: %s\n"
    _, err = fmt.Printf(format, env.Debug, env.Neo4JProtocol, env.Neo4JHost, env.Neo4JPort, env.Neo4JUser)
    if err != nil {
        log.Fatal(err.Error())
    }

	return env
}