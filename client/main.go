package main

import(
	"github.com/gin-gonic/gin"
)


func init(){
	InitDB()
}


func main(){
	router :=gin.Default()
	UserRoutes(router)
	router.Run(":9000")
}