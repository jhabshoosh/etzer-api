package main

import (
	config "github.com/jhabshoosh/etzer-api/internal/config"
	server "github.com/jhabshoosh/etzer-api/internal/server"
)

func main() {
	config.Init()

	srv := server.Init()
	srv.InitRoutes()
	srv.Run()
}
