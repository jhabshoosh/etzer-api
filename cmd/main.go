package main

import (
	server "github.com/jhabshoosh/etzer-api/internal/server"
)

func main() {
	srv := server.Init()
	srv.InitRoutes()
	srv.Run()
}