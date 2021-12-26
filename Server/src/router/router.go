package router

import "github.com/gin-gonic/gin"

func Initialize(engine *gin.Engine) {
	staticRouter(engine)
	loadBase(engine)
	loadUser(engine)
}

func staticRouter(engine *gin.Engine) {
	engine.Static("/assets", "../res")
	engine.StaticFile("/favicon.ico", "../res/img/favicon.ico")
}
