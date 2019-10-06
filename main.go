package main

import (
	"fmt"
	"github.com/Diode222/Mimiron/conf"
	"github.com/Diode222/Mimiron/manager"
	pb "github.com/Diode222/Mimiron/proto_gen"
	"github.com/Diode222/Mimiron/serviceServer"
	"google.golang.org/grpc"
	"log"
)

func main() {
	initService()
}

func initService() {
	grpcServer := grpc.NewServer()
	defer grpcServer.GracefulStop()

	pb.RegisterWordSplitServiceServer(grpcServer, serviceServer.NewSplittedWordServer())
	err := manager.GetServiceManger(conf.ETCD_ADDRESS).Register(conf.ETCD_ADDRESS, conf.LISTEN_IP, conf.SERVICE_IP, conf.SERVICE_PORT, grpcServer, conf.TTL)
	if err != nil {
		log.Printf(fmt.Sprintf("Register service to etcd failed, err: %s"))
		panic(err)
	}
}
