package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	// "golang.org/x/crypto/bcrypt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "Micro_Grpc/proto"
)

const (
	port = ":6000"
)

func SignUp(ctx *gin.Context) {
	var Body struct {
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		Email     string `json:"email"`
		Password  string `json:"password"`
	}

	if ctx.ShouldBindJSON(&Body) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "Please enter input correctly",
		})
		return
	}

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "could not connect server",
		})
		return
	}
	defer conn.Close()

	client := pb.NewAuthenticationServiceClient(conn)

	// hash, err := bcrypt.GenerateFromPassword([]byte(Body.Password), 10)
	// if err != nil {
	// 	ctx.JSON(http.StatusBadGateway, gin.H{
	// 		"error": "failed to hash",
	// 	})
	// 	ctx.Abort()
	// 	return
	// }
	// user := User{FirstName: Body.FirstName, LastName: Body.LastName, Email: Body.Email, Password: string(hash)}
	// result := InitDB().Create(&user)
	// if result.Error != nil {
	// 	ctx.JSON(http.StatusBadRequest, gin.H{
	// 		"error": "failed to create user",
	// 	})
	// 	ctx.Abort()
	// 	return
	// }

	req := &pb.SignupRequest{
		FirstName: Body.FirstName,
		LastName:  Body.LastName,
		Email:     Body.Email,
		Password:  Body.Password,
	}

	res,err:=client.SignUpService(context.Background(),req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"err": "Error from response",
		})
		return
	}

	fmt.Println(res)

	var Response struct{
		Name   string `json:"name"`
		Email  string `json:"email"`
		Result string `json:"result"`
	}

	Response.Result=res.Result
	Response.Name=res.Name
	Response.Email=res.Email

	ctx.JSON(http.StatusOK, gin.H{
		"success": "user signup Successfully completed..",
	})
}
