package config

import (
	"github.com/go-martini/martini"
	"github.com/robfig/config"
)

type App struct {
	Martini *martini.ClassicMartini
	Env     string
	Config  *config.Config
}

func ClassicApp() *App {
	// 初始化martini
	mart := martini.Classic()

	// 加载配置文件
	conf, err := config.ReadDefault("config/config.ini")
	if err != nil {
		panic(err)
	}

	app := &App{Martini: mart, Env: martini.Env, Config: conf}
	registerRoutes(app)
	return app
}

func (a *App) Run() {
	a.Martini.Run()
}
