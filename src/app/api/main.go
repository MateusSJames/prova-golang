package main

import (
	"Golang/src/app/api/endpoints/router"
)

func main() {
	startServer()
}

func startServer() {
	server := router.Start()
	server.Debug = true
	server.Logger.Info(server.Start(":8000"))
}
