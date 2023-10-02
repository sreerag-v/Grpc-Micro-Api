package main

import (
	pb "Micro_Grpc/proto"
	"context"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (s *Server) SignUpService(ctx context.Context, req *pb.SignupRequest) (*pb.SignupRespone, error) {
	firstname := req.GetFirstName()
	lastname := req.GetLastName()
	email := req.GetEmail()
	password := req.GetPassword()

	hash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		log.Fatalf("Faild to hash password:%v", err)
	}
	users := &User{
		FirstName: firstname,
		LastName:  lastname,
		Email:     email,
		Password:  string(hash),
	}
	record := InitDB().Create(&users)
	if record.Error != nil {
		log.Fatalf("Faild to create database:%v", err)
	}
	log.Printf("User Signed With Cridentials << User >>:%v", firstname)

	return &pb.SignupRespone{
		Name:   firstname,
		Email:  email,
		Result: "Success",
	}, nil
}
