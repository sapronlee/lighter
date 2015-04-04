package config

import (
	"lighter/controllers"
)

func registerRoutes(app *App) {
	app.Martini.Get("/", controllers.HomeAction)
	app.Martini.Get("/login", controllers.LoginAction)
}
