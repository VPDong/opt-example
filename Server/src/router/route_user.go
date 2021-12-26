package router

import (
	"server/src/service"

	"github.com/gin-gonic/gin"
)

func loadUser(router *gin.Engine) {
	userRouter := router.Group("/user")
	{
		userRouter.GET("/signup", userSignUp)
		userRouter.GET("/signdown", userSignDown)
		userRouter.GET("/signin", userSignIn)
		userRouter.GET("/signout", userSignOut)
		userRouter.GET("/details", userDetails)
	}
}

func userSignUp(ctx *gin.Context) {
	service.UserSignUp()
	ctx.JSON(200, gin.H{
		"msg": "signup",
	})
}

func userSignDown(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "signdown",
	})
}

func userSignIn(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "signin",
	})
}

func userSignOut(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "signout",
	})
}

func userDetails(ctx *gin.Context) {
	ctx.JSON(200, gin.H{
		"msg": "details",
	})
}
