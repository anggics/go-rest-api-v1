package main

import (
	"rest-api/interface/container"
	"rest-api/interface/server"
	"rest-api/shared/database"
)

func main() {
	server.StartApp(container.SetUpContainer(database.Getconnection().DB)).Run()
}
