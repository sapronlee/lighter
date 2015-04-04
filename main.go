package main

import (
	"lighter/config"
)

func main() {
	app := config.ClassicApp()
	app.Run()
}
