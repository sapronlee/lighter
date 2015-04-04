package config

import (
	"github.com/go-martini/martini"
)

type Application struct {
	Martini *martini.ClassicMartini
}

func InitApplication() *Application {
	mart := martini.Classic()
	return &Application{Martini: mart}
}
