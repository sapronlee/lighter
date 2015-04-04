package web

import (
	"github.com/go-martini/martini"
	"lighter/web/controllers"
	"lighter/web/controllers/api/v1"
)

func RegiterRoutes(m *martini.ClassicMartini) {

	// GET /
	m.Get("/", controllers.HomeIndex)

	m.Group("/api/v1/users", func(r martini.Router) {
		r.Get("", apiv1.UserIndex)
	})
}
