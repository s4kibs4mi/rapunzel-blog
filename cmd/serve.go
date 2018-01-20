package cmd

import (
	"fmt"
	pb "github.com/s4kibs4mi/rapunzel-blog/protos"
	"google.golang.org/grpc"
	"net"
	"google.golang.org/grpc/reflection"
	"github.com/s4kibs4mi/rapunzel-blog/servers"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

func Serve() {
	lis, err := net.Listen("tcp", ":8090")
	if err != nil {
		fmt.Println(err)
		return
	}
	gRPCServer := grpc.NewServer()
	pb.RegisterUserServiceServer(gRPCServer, servers.NewUserServer())
	reflection.Register(gRPCServer)
	if err := gRPCServer.Serve(lis); err != nil {
		fmt.Println("Failed to run server, ", err)
		return
	}
}
