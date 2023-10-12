package main

import(
	"github.com/gin-gonic/gin"
   _"Micro_Grpc/client/docs"
   swaggerFiles "github.com/swaggo/files"
   ginSwagger "github.com/swaggo/gin-swagger"

)


func init(){
	InitDB()
}

// @title Documenting API (Micro-Grpc)
// @version 1.0
// @Description Sample Grpc Cilent-Server Interaction

// @contact.name Sreerag.V
// @contact.email sreeraghrg@gmail.com
// @contact.url https://github.com/sreerag-v

//host localhost:9000
func main(){

	router :=gin.Default()
	UserRoutes(router)

	router.GET("/swagger/*any",ginSwagger.WrapHandler(swaggerFiles.Handler))
	router.Run(":9000")
}