package main

import (
	"fmt"
	"log"
	"net"

	"github.com/athunlal/order-svc/pkg/client"
	"github.com/athunlal/order-svc/pkg/config"
	"github.com/athunlal/order-svc/pkg/db"
	"github.com/athunlal/order-svc/pkg/pb"
	"github.com/athunlal/order-svc/pkg/service"
	"google.golang.org/grpc"
)

func main() {
	c, err := config.LoadConfig()

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	handler := db.Init(c.DBUrl)

	lis, err := net.Listen("tcp", c.Port)

	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	productSvc := client.InitProductServiceClient(c.ProductSvcUrl)
	if err != nil {
		log.Fatalln("Failed at config", err)
	}

	s := service.Server{
		H:          handler,
		ProductSvc: productSvc,
	}

	fmt.Println("Order Svc on", c.Port)
	grpcServer := grpc.NewServer()

	pb.RegisterOrderServiceServer(grpcServer, &s)
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalln("Failed to serve:", err)
	}
}
