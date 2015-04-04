package config

func registerRoutes(app *App) {
	app.Martini.Get("/", func() string {
		return "Hello world"
	})
}
