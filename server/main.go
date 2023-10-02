package main

import (
	pb "Micro_Grpc/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

const (
	port = ":6000"
)

type Server struct {
	pb.AuthenticationServiceServer
}

func main() {
	InitDB()

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("Error while listening port:%s.........> %v", port, err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterAuthenticationServiceServer(grpcServer, &Server{})
	log.Printf("Server is Startig")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Faild to Serv %v", err)
	}
}
