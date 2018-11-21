package router

import (
	"github.com/julienschmidt/httprouter"
	"github.com/tribes/controller/home"
)

func InitRoutes(r *httprouter.Router) {
	setHomeRoutes(r)
}

func setHomeRoutes(r *httprouter.Router) {
	r.GET("/home", home.RequestHandler)
	r.POST("/login", home.RequestLoginHandler)
}
