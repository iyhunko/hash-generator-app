package main

import (
	"fmt"
	"github.com/iyhunko/hash-generation-app/config"
	grpc2 "github.com/iyhunko/hash-generation-app/grpc"
	pb "github.com/iyhunko/hash-generation-app/proto"
	"github.com/iyhunko/hash-generation-app/store"
	"google.golang.org/grpc"
	"log"
	"net"
)

func main() {
	log.Println("Starting GRPC api server")

	conf := config.InitConfig()
	cacheStorage := store.NewStore()

	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", conf.GRPCServerPort))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	hashServer := grpc2.NewHashServer(conf, cacheStorage)
	pb.RegisterHashServiceServer(grpcServer, &hashServer)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %s", err)
	}
}
