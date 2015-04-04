package main

import (
	"lighter/config"
)

func main() {
	app := config.InitApplication()
	app.Martini.Run()
}
