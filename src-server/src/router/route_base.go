package router

import "github.com/gin-gonic/gin"

func loadBase(router *gin.Engine) {
	router.GET("/", baseHomeland)
	router.GET("/index", baseIndexPage)
}

func baseHomeland(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": ctx.ClientIP(),
	})
}

func baseIndexPage(ctx *gin.Context) {
	ctx.Redirect(301, "/")
}
