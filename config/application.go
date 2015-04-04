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

func (a *App) GetBoolConfig(name string) bool {
	b, _ := a.Config.Bool(a.Env, name)
	return b
}

func (a *App) GetStringConfig(name string) string {
	s, _ := a.Config.String(a.Env, name)
	return s
}

func (a *App) GetIntConfig(name string) int {
	i, _ := a.Config.Int(a.Env, name)
	return i
}

func (a *App) GetFloatConfig(name string) float64 {
	f, _ := a.Config.Float(a.Env, name)
	return f
}

func (a *App) GetRawStringConfig(name string) string {
	s, _ := a.Config.RawString(a.Env, name)
	return s
}
