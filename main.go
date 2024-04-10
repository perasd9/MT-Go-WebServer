package main

import "github/perasd9/MTWebServer/server"

func main() {

	var server *server.MTServer

	server = server.NewServer()

	server.Start()
}
