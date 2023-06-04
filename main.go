package main

import (
	"log"
	"net"

	"google.golang.org/grpc"

	"UserServer/constant"
	v1 "UserServer/controllers/v1"
	"UserServer/mongodb"
	fancygame "UserServer/proto/fancygame"
)

func init() {
	//go proto.Initialize()
	constant.ReadConfig(".env")
	mongodb.Initialize()
}

func main() {
	log.Println("[info] create user Service :27031")
	lis, err := net.Listen("tcp", ":27031")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	fancygame.RegisterUserServer(s, &v1.UserServe{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
