package main

import (
	"context"
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/echoservice"
	"golang.conradwood.net/go-easyops/server"
	"golang.conradwood.net/go-easyops/utils"
	"google.golang.org/grpc"
	"os"
)

var (
	port = flag.Int("port", 4100, "The grpc server port")
)

type echoServer struct {
}

func main() {
	var err error
	flag.Parse()
	fmt.Printf("Starting EchoServiceServer...\n")
	server.SetHealth(common.Health_STARTING)

	sd := server.NewServerDef()
	sd.SetOnStartupCallback(server_started)
	sd.SetPort(*port)
	sd.SetRegister(server.Register(
		func(server *grpc.Server) error {
			e := new(echoServer)
			pb.RegisterEchoServiceServer(server, e)
			return nil
		},
	),
	)
	err = server.ServerStartup(sd)
	utils.Bail("Unable to start server", err)
	os.Exit(0)
}

// called after server was started and is listening
func server_started() {
	server.SetHealth(common.Health_STARTING)
	server.SetHealth(common.Health_READY)
}

/************************************
* grpc functions
************************************/

func (e *echoServer) Ping(ctx context.Context, req *common.Void) (*pb.PingResponse, error) {
	resp := &pb.PingResponse{Response: "pingresponse"}
	return resp, nil
}
