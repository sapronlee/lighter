package main

import (
	"fmt"
	"lighter/config"
)

func main() {
	app := config.ClassicApp()
	debug, _ := app.Config.Bool(app.Env, "debug")
	fmt.Printf("debug is %v", debug)
	app.Run()
}
