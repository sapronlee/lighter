package controllers

import (
	"github.com/martini-contrib/render"
)

func HomeAction(r render.Render) {
	r.HTML(200, "index", nil)
}

func LoginAction(r render.Render) {
	r.HTML(200, "index", nil)
}
