package main

import (
	"github.com/go-martini/martini"
	"lighter/web"
)

func main() {
	server := martini.Classic()
	web.RegiterRoutes(server)
	server.Run()
}
