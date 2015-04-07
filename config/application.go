package config

import (
	"github.com/go-martini/martini"
	"github.com/martini-contrib/render"
	"github.com/robfig/config"
	"html/template"
	"lighter/helpers"
)

type App struct {
	Martini *martini.ClassicMartini
	Env     string
	config  *config.Config
}

func ClassicApp() *App {
	// 初始化martini
	mart := martini.Classic()
	mart.Use(render.Renderer(render.Options{
		Directory: "views",
		Layout:    "layout",
		Funcs:     []template.FuncMap{helpers.AssetHelpers()},
	}))

	// 加载配置文件
	conf, err := config.ReadDefault("config/config.ini")
	if err != nil {
		panic(err)
	}

	app := &App{Martini: mart, Env: martini.Env, config: conf}
	registerRoutes(app)
	return app
}

func (a *App) Run() {
	a.Martini.Run()
}

func (a *App) GetBoolConfig(name string) bool {
	b, _ := a.config.Bool(a.Env, name)
	return b
}

func (a *App) GetStringConfig(name string) string {
	s, _ := a.config.String(a.Env, name)
	return s
}

func (a *App) GetIntConfig(name string) int {
	i, _ := a.config.Int(a.Env, name)
	return i
}

func (a *App) GetFloatConfig(name string) float64 {
	f, _ := a.config.Float(a.Env, name)
	return f
}

func (a *App) GetRawStringConfig(name string) string {
	s, _ := a.config.RawString(a.Env, name)
	return s
}
