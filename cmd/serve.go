package cmd

import (
	"fmt"
	pb "github.com/s4kibs4mi/rapunzel-blog/protos"
	"google.golang.org/grpc"
	"net"
	"google.golang.org/grpc/reflection"
	"github.com/s4kibs4mi/rapunzel-blog/servers"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
)

/**
 * := Coded with love by Sakib Sami on 19/01/18.
 * := root@sakib.ninja
 * := www.sakib.ninja
 * := Coffee : Dream : Code
 */

var ServeCmd = cobra.Command{
	Use:   "serve",
	Run:   Serve,
	Short: "Use serve to start gRPC server",
}

func Serve(cmd *cobra.Command, args []string) {
	lis, err := net.Listen("tcp", viper.GetString("app.address"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on...", viper.GetString("app.address"))
	gRPCServer := grpc.NewServer()
	pb.RegisterUserServiceServer(gRPCServer, servers.NewUserServer())
	reflection.Register(gRPCServer)
	if err := gRPCServer.Serve(lis); err != nil {
		fmt.Println("Failed to run server, ", err)
		return
	}
}
