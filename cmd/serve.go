package cmd

import (
	"fmt"
	pb "github.com/s4kibs4mi/rapunzel-blog/protos"
	"github.com/s4kibs4mi/rapunzel-blog/api"
	"google.golang.org/grpc"
	"net"
	"google.golang.org/grpc/reflection"
)

func Serve() {
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	gRPCServer := grpc.NewServer()
	pb.RegisterUserServiceServer(gRPCServer, api.NewUserServer())
	reflection.Register(gRPCServer)
	if err := gRPCServer.Serve(lis); err != nil {
		fmt.Println("Failed to run server, ", err)
		return
	}
}
