package main

import (
	"github.com/gin-gonic/gin"
)

func UserRoutes(ctx *gin.Engine) {

	// c.GET("/",UserHome)
	// c.POST("/login",Login)
	ctx.POST("/signup", SignUp)

}
