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

// SignUp godoc
// @Summary Create a new user
// @Description Create a new user account
// @Accept json
// @Produce json
// @Param request body SignUpRequest true "User registration details"
// @Success 200 {object} ErrResponse
// @Failure 400 {object} ErrResponse
// @Failure 500 {object} SuccResponse
// @Router /signup [post]
func SignUp(ctx *gin.Context) {
	SignUpRequest := SignUpRequest{}
	if ctx.ShouldBindJSON(&SignUpRequest) != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"err": "Please enter input correctly",
		})
		return
	}

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrResponse{
			"could not connect server",
			err.Error(),
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
		FirstName: SignUpRequest.FirstName,
		LastName:  SignUpRequest.LastName,
		Email:     SignUpRequest.Email,
		Password:  SignUpRequest.Password,
	}

	res, err := client.SignUpService(context.Background(), req)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, ErrResponse{
			"Error from response",
			err.Error(),
		})
		return
	}

	fmt.Println(res)

	var Response struct {
		Name   string `json:"name"`
		Email  string `json:"email"`
		Result string `json:"result"`
	}

	Response.Result = res.Result
	Response.Name = res.Name
	Response.Email = res.Email

	ctx.JSON(http.StatusOK, SuccResponse{
		"user signup Successfully completed..",
	})
}
