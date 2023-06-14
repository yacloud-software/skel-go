package main

import (
	"flag"
	"fmt"
	"golang.conradwood.net/apis/common"
	pb "golang.conradwood.net/apis/echoservice"
	"golang.conradwood.net/go-easyops/authremote"
	"golang.conradwood.net/go-easyops/utils"
	"os"
)

var (
	echoClient pb.EchoServiceClient
)

func main() {
	flag.Parse()

	echoClient = pb.GetEchoServiceClient()

	// a context with authentication
	ctx := authremote.Context()

	empty := &common.Void{}
	response, err := echoClient.Ping(ctx, empty)
	utils.Bail("Failed to ping server", err)
	fmt.Printf("Response to ping: %v\n", response)

	fmt.Printf("Done.\n")
	os.Exit(0)
}
