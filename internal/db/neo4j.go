package db

import (
	"fmt"
	"log"

	"github.com/neo4j/neo4j-go-driver/neo4j"
)


type Neo4JDB struct {
	Connection neo4j.Driver
}

func NewNeo4jConnection(protocol string, host string, port int, user string, password string) (neo4j.Driver, error) {
	target := fmt.Sprintf("%s://%s:%d", protocol, host, port)

	log.Printf("Attempting to connect to neo 4j at %s with user %s password %s", target, user, password)

	driver, err := neo4j.NewDriver(
		target,
		neo4j.BasicAuth(user, password, ""),
		func(c *neo4j.Config) {
			c.Encrypted = false
		})

	if err != nil {
		log.Panic("Cannot connect to Neo4j Server", err)
		return nil, err
	}

	log.Println("Connected to Neo4j Server", "neo4j_server_uri", target)

	return driver, nil
}

