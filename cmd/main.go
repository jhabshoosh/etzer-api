package main

import (
	"os"

	server "github.com/jhabshoosh/etzer-api/internal/server"
)

func main() {

	os.Setenv("ETZER_DEBUG", "true")
	os.Setenv("ETZER_PORT", "8080")
	os.Setenv("ETZER_DBHOST", "localhost")
	os.Setenv("ETZER_DBPORT", "7687")
	os.Setenv("ETZER_DBUSER", "neo4j")
	os.Setenv("ETZER_DBPASSWORD", "test")


	srv := server.Init()
	srv.InitRoutes()
	srv.Run()
}