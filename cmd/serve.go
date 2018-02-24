package cmd

import (
	"fmt"
	pb "github.com/s4kibs4mi/rapunzel-blog/proto/defs"
	"google.golang.org/grpc"
	"net"
	"google.golang.org/grpc/reflection"
	"github.com/s4kibs4mi/rapunzel-blog/servers"
	"github.com/spf13/viper"
	"github.com/spf13/cobra"
	"github.com/s4kibs4mi/rapunzel-blog/conn"
	"github.com/s4kibs4mi/rapunzel-blog/auth"
	"os"
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
	initDBs()
	runAndListen()
}

func initDBs() {
	ok := conn.NewMongodbConnection()
	if !ok {
		os.Exit(-1)
	}
}

func runAndListen() {
	lis, err := net.Listen("tcp", viper.GetString("app.address"))
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Listening on...", viper.GetString("app.address"))
	var serverOptions []grpc.ServerOption
	serverOptions = append(serverOptions, grpc.UnaryInterceptor(auth.UnaryAuthInterceptor))
	gRPCServer := grpc.NewServer(serverOptions...)
	pb.RegisterRapunzelBlogServiceServer(gRPCServer, servers.NewRapunzelBlogServer())
	reflection.Register(gRPCServer)
	if err := gRPCServer.Serve(lis); err != nil {
		fmt.Println("Failed to run server, ", err)
		return
	}
}
